package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

func getTerraformConfigFile() (string, error) {
	var tfConfigDir string
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		tfConfigDir = filepath.Join(homeDir, "terraform.d")
	} else {
		tfConfigDir = filepath.Join(homeDir, ".terraform.d")
	}
	return filepath.Join(tfConfigDir, "credentials.tfrc.json"), nil
}

func readTerraformConfigCredentials(file string) (map[string]interface{}, error) {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var raw map[string]interface{}
	if err = json.Unmarshal(src, &raw); err != nil {
		return nil, err
	}
	creds, ok := raw["credentials"].(map[string]interface{})
	if !ok {
		return nil, errors.New("malformed JSON file: couldn't find credentials key")
	}
	return creds, nil
}

func GetTFCloudToken() (string, error) {
	tfConfigFile, err := getTerraformConfigFile()
	if err != nil {
		return "", err
	}
	tfConfigCreds, err := readTerraformConfigCredentials(tfConfigFile)
	if err != nil {
		return "", err
	}
	tfcCreds, ok := tfConfigCreds[TFCloudHostname].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("malformed JSON file: couldn't find %s key", TFCloudHostname)
	}
	tfcToken, ok := tfcCreds["token"].(string)
	if !ok {
		return "", errors.New("malformed JSON file: couldn't find token key")
	}
	return tfcToken, nil
}
