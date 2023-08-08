package iam_access_key

import (
	"fmt"
	"os"
	"strings"

	"encoding/json"
	"os/user"
)

func writeToFile(filePath string, data interface{}) error {
	var out string
	switch _data := data.(type) {
	case string:
		out = _data
	case nil:
		return nil
	default:
		bs, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return fmt.Errorf("MarshalIndent data %#v got an error: %#v", data, err)
		}
		out = string(bs)
	}

	if strings.HasPrefix(filePath, "~") {
		home, err := getUserHomeDir()
		if err != nil {
			return err
		}
		if home != "" {
			filePath = strings.Replace(filePath, "~", home, 1)
		}
	}

	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil { // ignore_security_alert
			return err
		}
	}

	return os.WriteFile(filePath, []byte(out), 0422) // ignore_security_alert
}

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("get current user got an error: %#v", err)
	}
	return usr.HomeDir, nil
}
