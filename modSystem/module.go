package modSystem

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type EnvInfo struct {
	OS          string // "windows", "linux", "darwin"
	IsCodespace bool
	MachineID   string
	MachineInfo string
}

func Detect() (*EnvInfo, error) {
	osName := runtime.GOOS
	isCodespace := os.Getenv("CODESPACES") == "true"

	info, err := getMachineInfo(osName)
	if err != nil {
		return nil, err
	}
	id, err := loadMachineID()
	if err != nil {
		return nil, err
	}

	return &EnvInfo{
		OS:          osName,
		IsCodespace: isCodespace,
		MachineID:   id,
		MachineInfo: info,
	}, nil
}

func getMachineInfo(goos string) (string, error) {
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

func loadMachineID() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	execDir := filepath.Dir(execPath)

	machineIDPath := filepath.Join(execDir, "machineid")

	data, err := os.ReadFile(machineIDPath)
	if err != nil {
		return "", err
	}
	strID := strings.TrimSpace(string(data))

	return strID, nil
}
