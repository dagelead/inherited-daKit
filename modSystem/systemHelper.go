package modSystem

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func readFirstAvailable(paths []string) (string, error) {
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err == nil {
			return strings.TrimSpace(string(data)), nil
		}
	}
	return "", errors.New("no machine ID file found")
}

func readWindowsMachineID() (string, error) {
	out, err := exec.Command("reg", "query",
		`HKLM\SOFTWARE\Microsoft\Cryptography`,
		"/v", "MachineGuid").Output()
	if err != nil {
		return "", err
	}

	// Example output:
	//    MachineGuid    REG_SZ    9f3c...etc
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "MachineGuid") {
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				return strings.TrimSpace(parts[2]), nil
			}
		}
	}

	return "", errors.New("MachineGuid not found")
}

func getIOPlatformUUID() (string, error) {
	out, err := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "IOPlatformUUID") {
			parts := strings.Split(line, "\"")
			if len(parts) >= 4 {
				return parts[3], nil
			}
		}
	}

	return "", errors.New("IOPlatformUUID not found")
}
