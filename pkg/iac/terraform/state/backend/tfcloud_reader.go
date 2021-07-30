package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	pkghttp "github.com/cloudskiff/driftctl/pkg/http"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const BackendKeyTFCloud = "tfcloud"
const TFCloudAPI = "https://app.terraform.io/api/v2"

type TFCloudAttributes struct {
	HostedStateDownloadUrl string `json:"hosted-state-download-url"`
}

type TFCloudData struct {
	Attributes TFCloudAttributes `json:"attributes"`
}

type TFCloudBody struct {
	Data TFCloudData `json:"data"`
}

type TFCloudBackend struct {
	request *http.Request
	client  pkghttp.HTTPClient
	reader  io.ReadCloser
}

func NewTFCloudReader(client pkghttp.HTTPClient, workspaceId string, opts *Options) (*TFCloudBackend, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/workspaces/%s/current-state-version", TFCloudAPI, workspaceId), nil)
	if err != nil {
		return nil, err
	}

	token := opts.TFCloudToken
	if token == "" {
		tfConfigFile, err := getTerraformConfigFile()
		if err != nil {
			return nil, err
		}
		file, err := os.Open(tfConfigFile)
		if err != nil {
			return nil, err
		}
		reader := NewTFCloudConfigReader(file)
		token, err = reader.GetToken()
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Content-Type", "application/vnd.api+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return &TFCloudBackend{req, client, nil}, nil
}

func (t *TFCloudBackend) Read(p []byte) (n int, err error) {
	if t.reader == nil {
		res, err := t.client.Do(t.request)
		if err != nil {
			return 0, err
		}

		if res.StatusCode < 200 || res.StatusCode >= 400 {
			return 0, errors.Errorf("error requesting terraform cloud backend state: status code: %d", res.StatusCode)
		}

		body := TFCloudBody{}
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(bodyBytes, &body)
		if err != nil {
			return 0, err
		}

		rawURL := body.Data.Attributes.HostedStateDownloadUrl
		logrus.WithFields(logrus.Fields{"hosted-state-download-url": rawURL}).Trace("Terraform Cloud backend response")

		h, err := NewHTTPReader(t.client, rawURL, &Options{})
		if err != nil {
			return 0, err
		}
		n, err = h.Read(p)
		if err != nil {
			return 0, err
		}
		t.reader = h.reader
		return n, nil
	}
	return t.reader.Read(p)
}

func (t *TFCloudBackend) Close() error {
	if t.reader != nil {
		return t.reader.Close()
	}
	return errors.New("Unable to close reader as nothing was opened")
}
