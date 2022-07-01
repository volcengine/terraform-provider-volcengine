package iam_access_key

import (
	"fmt"
	"os"
	"strings"

	"encoding/json"
	"io/ioutil"
	"os/user"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func writeToFile(filePath string, data interface{}) error {
	var out string
	switch data.(type) {
	case string:
		out = data.(string)
		break
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

	return ioutil.WriteFile(filePath, []byte(out), 422) // ignore_security_alert
}

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("get current user got an error: %#v", err)
	}
	return usr.HomeDir, nil
}

var akSkImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	tmpId := data.Id()
	items := strings.Split(tmpId, ":")
	if len(items) < 1 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	data.SetId(items[0])
	if len(items) > 1 {
		if err := data.Set("user_name", tmpId[len(items[0])+1:]); err != nil {
			return []*schema.ResourceData{data}, err
		}
	}
	return []*schema.ResourceData{data}, nil
}
