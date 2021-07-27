package backend

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestGetTFCloudToken(t *testing.T) {
	originalEnv := os.Getenv("HOME")
	os.Setenv("HOME", "/tmp")
	defer func() { os.Setenv("HOME", originalEnv) }()
	homedir.DisableCache = true
	defer func() { homedir.DisableCache = false }()

	tests := []struct {
		name    string
		file    string
		want    string
		wantErr error
	}{
		{
			name:    "get terraform cloud creds with config file",
			file:    "./testdata/tfc_creds_valid.json",
			want:    "token.creds.test",
			wantErr: nil,
		},
		{
			name:    "test with wrong credentials key in config file",
			file:    "./testdata/tfc_creds_invalid_credentials.json",
			want:    "",
			wantErr: fmt.Errorf("malformed JSON file: couldn't find credentials key"),
		},
		{
			name:    "test with wrong terraform cloud hostname key in config file",
			file:    "./testdata/tfc_creds_invalid_tfc_hostname.json",
			want:    "",
			wantErr: fmt.Errorf("malformed JSON file: couldn't find app.terraform.io key"),
		},
		{
			name:    "test with wrong terraform cloud hostname key in config file",
			file:    "./testdata/tfc_creds_invalid_tfc_token.json",
			want:    "",
			wantErr: fmt.Errorf("malformed JSON file: couldn't find token key"),
		},
		{
			name:    "test with unreadable config file",
			file:    "",
			want:    "",
			wantErr: fmt.Errorf("open /tmp/.terraform.d/credentials.tfrc.json: no such file or directory"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.file != "" {
				input, err := ioutil.ReadFile(tt.file)
				if err != nil {
					panic(err)
				}
				tfConfigDir := filepath.Join("/tmp", ".terraform.d")
				err = os.MkdirAll(tfConfigDir, fs.ModePerm)
				if err != nil {
					panic(err)
				}
				err = ioutil.WriteFile(filepath.Join(tfConfigDir, "credentials.tfrc.json"), input, os.ModePerm)
				if err != nil {
					panic(err)
				}
				defer os.RemoveAll(tfConfigDir)
			}

			got, err := GetTFCloudToken()
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("GetTFCloudToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTFCloudToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
