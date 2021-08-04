package backend

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestTFCloudConfigReader_GetToken(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		want    string
		wantErr error
	}{
		{
			name:    "get terraform cloud creds with config file",
			src:     `{"credentials": {"app.terraform.io": {"token": "token.creds.test"}}}`,
			want:    "token.creds.test",
			wantErr: nil,
		},
		{
			name:    "test with wrong credentials key in config file",
			src:     `{"test": {"app.terraform.io": {"token": "token.creds.test"}}}`,
			want:    "",
			wantErr: fmt.Errorf("malformed JSON file: token not found"),
		},
		{
			name:    "test with wrong terraform cloud hostname key in config file",
			src:     `{"credentials": {"test": {"token": "token.creds.test"}}}`,
			want:    "",
			wantErr: fmt.Errorf("malformed JSON file: token not found"),
		},
		{
			name:    "test with wrong terraform cloud token key in config file",
			src:     `{"credentials": {"app.terraform.io": {"test": "token.creds.test"}}}`,
			want:    "",
			wantErr: fmt.Errorf("malformed JSON file: token not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readerCloser := ioutil.NopCloser(strings.NewReader(tt.src))
			defer readerCloser.Close()
			r := NewTFCloudConfigReader(readerCloser)
			got, err := r.GetToken()
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
