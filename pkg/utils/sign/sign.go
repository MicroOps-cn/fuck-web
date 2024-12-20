package sign

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/MicroOps-cn/fuck/sets"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/errors"
)

const (
	MimeJSON       = "application/json"
	MimeXML        = "application/xml"
	MimeUrlencoded = "application/x-www-form-urlencoded"
)

var innerKey = sets.New[string]("authKey", "authSign", "authSecret", "authMethod", "authAlgorithm")

func GetPayloadFromHTTPRequest(r *http.Request) (string, error) {
	if date := r.Header.Get("date"); date != "" {
		requestTime, err := time.Parse(time.RFC1123, date)
		if err != nil {
			return "", err
		} else if time.Since(requestTime) > time.Minute*10 {
			return "", errors.ParameterError("request has expired")
		}
	}
	var bodyHash string
	if r.ContentLength > 0 {
		contentType, _, _ := strings.Cut(r.Header.Get("content-type"), ";")
		if len(contentType) > 0 {
			switch contentType {
			case MimeJSON, MimeXML, MimeUrlencoded:
				if r.ContentLength < int64(config.Get().Global.MaxBodySize) {
					body, err := ioutil.ReadAll(r.Body)
					_ = r.Body.Close()
					if err != nil {
						return "", err
					}
					bodyHash = fmt.Sprintf("%x", md5.Sum(body))
					r.Body = io.NopCloser(bytes.NewBuffer(body))
				}
			}
		}
	}
	if len(bodyHash) == 0 {
		bodyHash = r.Header.Get("x-body-hash")
	}
	if len(bodyHash) == 0 && r.ContentLength > 0 {
		return "", fmt.Errorf("failed to get hash value of body")
	}
	payload := strings.Builder{}
	payload.WriteString(r.Method + "\n")
	payload.WriteString(bodyHash + "\n")
	payload.WriteString(r.Header.Get("content-type") + "\n")
	payload.WriteString(r.Header.Get("date") + "\n")
	urlQuery := url.Values{}
	for key, value := range r.URL.Query() {
		if !innerKey.Has(key) {
			for _, v := range value {
				urlQuery.Add(key, v)
			}
		}
	}
	payload.WriteString(r.URL.RawPath + "?" + urlQuery.Encode())
	return payload.String(), nil
}

type AuthAlgorithm string

const (
	HmacSha1   AuthAlgorithm = "HMAC-SHA1"
	HmacSha256 AuthAlgorithm = "HMAC-SHA256"
	HmacSha512 AuthAlgorithm = "HMAC-SHA512"
	ECDSA      AuthAlgorithm = "ECDSA"
)

func GetSignFromHTTPRequest(r *http.Request, key, secret, private string, algorithm AuthAlgorithm) (string, error) {
	payload, err := GetPayloadFromHTTPRequest(r)
	if err != nil {
		return "", err
	}
	switch algorithm {
	case "", HmacSha1:
		return SumSha1Hmac(secret, payload), nil
	case HmacSha256:
		return SumSha256Hmac(secret, payload), nil
	case ECDSA:
		return ECDSASign(private, payload)
	}
	return "", errors.ParameterError("unknown sign algorithm")
}

func Verify(key, secret, private string, algorithm AuthAlgorithm, signStr, payload string) bool {
	if len(signStr) == 0 {
		return false
	}
	switch algorithm {
	case "", HmacSha1:
		return SumSha1Hmac(secret, payload) == signStr
	case HmacSha256:
		return SumSha256Hmac(secret, payload) == signStr
	case "ECDSA":
		return ECDSAVerify(key, secret, payload, signStr)
	}
	return false
}
