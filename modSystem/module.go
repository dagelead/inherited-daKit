package modSystem

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

type EnvInfo struct {
	OS          string // "windows", "linux", "darwin"
	IsCodespace bool
	MachineID   string
}

func Detect() (*EnvInfo, error) {
	osName := runtime.GOOS
	isCodespace := os.Getenv("CODESPACES") == "true"

	var id string = ""
	var err error = nil

	if isCodespace {
		id, err = getGithubMachineID()
	} else {
		id, err = getMachineID(osName)
	}

	if err != nil {
		return nil, err
	}

	return &EnvInfo{
		OS:          osName,
		IsCodespace: isCodespace,
		MachineID:   id,
	}, nil
}

func getMachineID(goos string) (string, error) {
	switch goos {
	case "linux":
		return readFirstAvailable([]string{
			"/etc/machine-id",
			"/var/lib/dbus/machine-id",
		})
	case "windows":
		return readWindowsMachineID()
	case "darwin":
		return getIOPlatformUUID()
	default:
		return "", errors.New("unsupported OS for machine ID")
	}
}
func getGithubMachineID() (string, error) {
	repo := getEnv("GITHUB_REPOSITORY") //os.Getenv("GITHUB_REPOSITORY")
	space := getEnv("CODESPACE_NAME")   //os.Getenv("CODESPACE_NAME")

	if space == "" {
		return "", errors.New("CODESPACE_NAME not set")
	}

	return repo + "." + space, nil
}
func getEnv(key string) string {
	strValue, err := getEnvAsNormalUser(key)
	if err != nil {
		fmt.Println("fetch env error: ", err)
		return ""
	}
	return strValue
}
