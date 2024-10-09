package sign

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/stretchr/testify/require"

	"github.com/MicroOps-cn/fuck-web/config"
)

const cfg = `
storage:
  default: 
    sqlite: 
      path: ":memory:"
global: {
    
}
`

func TestSignHttpRequest(t *testing.T) {
	logger := logs.New()
	logs.SetDefaultLogger(logger)
	err := config.ReloadConfigFromYamlReader(logs.NewTraceLogger(), config.NewConverter("fuck-web.yaml", strings.NewReader(cfg)))
	require.NoError(t, err)
	type args struct {
		method      string
		contentType string
		body        string
		url         string
		header      url.Values
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    bool
	}{{
		name: "Test Get - OK",
		args: args{
			method: "GET",
			url:    "http://localhost:8080/api/v1/users?pageSize=100",
			header: map[string][]string{
				"date": {time.Now().Format(time.RFC1123)},
			},
		},
		wantErr: false,
		want:    true,
	}, {
		name: "Test POST JSON - OK",
		args: args{
			method:      "POST",
			contentType: MimeJSON,
			body:        `{"id":"123124343"}`,
			url:         "http://localhost:8080/api/v1/users?pageSize=100",
			header: map[string][]string{
				"date": {time.Now().Format(time.RFC1123)},
			},
		},
		wantErr: false,
		want:    true,
	}, {
		name: "Test PUT XML - OK",
		args: args{
			method:      "PUT",
			contentType: MimeXML,
			body:        `<xml><Meta>12312</Meta></xml>`,
			url:         "http://localhost:8080/api/v1/users?pageSize=100",
			header: map[string][]string{
				"date": {time.Now().Format(time.RFC1123)},
			},
		},
		wantErr: false,
		want:    true,
	}, {
		name: "Test PUT XML - Fail",
		args: args{
			method:      "PUT",
			contentType: MimeXML,
			body:        `<xml><Meta>12312</Meta></xml>`,
			url:         "http://localhost:8080/api/v1/users?pageSize=100",
			header: map[string][]string{
				"date": {time.Now().Format(time.Stamp)},
			},
		},
		wantErr: true,
		want:    false,
	}, {
		name: "Test POST Unknown - Fail",
		args: args{
			method:      "POST",
			contentType: "application/stream",
			body:        `{"id":"123124343"}`,
			url:         "http://localhost:8080/api/v1/users?pageSize=100",
			header: map[string][]string{
				"date": {time.Now().Format(time.RFC1123)},
			},
		},
		wantErr: true,
		want:    false,
	}}
	pub1, pub2, private, err := GenerateECDSAKeyPair()
	require.NoError(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, bytes.NewBuffer([]byte(tt.args.body)))
			require.NoError(t, err)
			for name, vals := range tt.args.header {
				for _, val := range vals {
					req.Header.Add(name, val)
				}
			}
			if len(tt.args.contentType) > 0 {
				req.Header.Set("Content-Type", tt.args.contentType)
			}
			sign, err := GetSignFromHTTPRequest(req, pub1, pub2, private, ECDSA)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignHttpRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			payload, err := GetPayloadFromHTTPRequest(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignHttpRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if ret := ECDSAVerify(pub1, pub2, payload, sign); ret != tt.want {
				t.Errorf("ECDSAVerify() ret = %v, want %v", ret, tt.want)
			}
			if tt.args.body != "" {
				all, err := ioutil.ReadAll(req.Body)
				require.NoError(t, err)
				require.Equal(t, string(all), tt.args.body)
			}
		})
	}
}
