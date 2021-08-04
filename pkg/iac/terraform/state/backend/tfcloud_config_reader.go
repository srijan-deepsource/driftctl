package backend

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

type Container struct {
	Credentials struct {
		TerraformCloud struct {
			Token string
		} `json:"app.terraform.io"`
	}
}

type TFCloudConfigReader struct {
	reader io.ReadCloser
}

func NewTFCloudConfigReader(reader io.ReadCloser) *TFCloudConfigReader {
	return &TFCloudConfigReader{reader}
}

func (r *TFCloudConfigReader) GetToken() (string, error) {
	defer r.reader.Close()

	b, err := ioutil.ReadAll(r.reader)
	if err != nil {
		return "", errors.New("unable to read file")
	}

	var container Container
	if err := json.Unmarshal(b, &container); err != nil {
		return "", err
	}
	if container.Credentials.TerraformCloud.Token == "" {
		return "", errors.New("malformed JSON file: token not found")
	}
	return container.Credentials.TerraformCloud.Token, nil
}

func getTerraformConfigFile() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	tfConfigDir := ".terraform.d"
	if runtime.GOOS == "windows" {
		tfConfigDir = "terraform.d"
	}
	return filepath.Join(homeDir, tfConfigDir, "credentials.tfrc.json"), nil
}
