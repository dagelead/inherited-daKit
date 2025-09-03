package modSystem

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

var envUserid string = "codespace"

// getEnvAsCodespaceUser runs the 'codespace-id' CLI as the 'codespace' user
// to fetch an environment variable value from that user's environment.
func getEnvAsNormalUser(envName string) (string, error) {
	// Lookup the "codespace" user
	usr, err := user.Lookup(envUserid)
	if err != nil {
		return "", fmt.Errorf("failed to lookup '%s' user: %w", envUserid, err)
	}

	uid, err := strconv.Atoi(usr.Uid)
	if err != nil {
		return "", fmt.Errorf("invalid UID: %w", err)
	}
	gid, err := strconv.Atoi(usr.Gid)
	if err != nil {
		return "", fmt.Errorf("invalid GID: %w", err)
	}

	// Prepare the command
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get current executable path: %w", err)
	}
	exeDir := filepath.Dir(exePath)
	cliPath := filepath.Join(exeDir, "userEnv")
	//cmd := exec.Command(cliPath, envName)
	cmd := exec.Command("bash", "-l", "-c", fmt.Sprintf("%s %s", cliPath, envName))

	// Drop to codespace user
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: uint32(uid),
			Gid: uint32(gid),
		},
	}
	cmd.Env = []string{
		fmt.Sprintf("HOME=%s", usr.HomeDir),
		"PATH=/usr/local/bin:/usr/bin:/bin",
	}

	// Capture output
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run codespace-id: %v, output: %s", err, out.String())
	}

	strValue := out.String()
	err = nil
	if len(strValue) > 0 {
		if strValue[0] == '-' { // means error;
			// print the error;
			fmt.Println(strValue)
			err = errors.New(strValue)
			strValue = ""
		}
	}

	return strValue, err
}
