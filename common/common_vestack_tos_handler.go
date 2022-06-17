package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"
)

var tosSignRequestHandler = request.NamedHandler{Name: "TosSignRequestHandler", Fn: tosSign}
var tosBuildHandler = request.NamedHandler{Name: "TosBuildHandler", Fn: tosBuild}
var tosUnmarshalHandler = request.NamedHandler{Name: "TosUnmarshalHandler", Fn: tosUnmarshal}
var tosUnmarshalErrorHandler = request.NamedHandler{Name: "TosUnmarshalErrorHandler", Fn: tosUnmarshalError}

func tosSign(req *request.Request) {
	region := req.ClientInfo.SigningRegion

	var (
		c Credentials
	)

	region = volcstack.StringValue(req.Config.Region)

	name := req.ClientInfo.SigningName
	if name == "" {
		name = req.ClientInfo.ServiceID
	}

	value, _ := req.Config.Credentials.Get()

	c = Credentials{
		AccessKeyID:     value.AccessKeyID,
		SecretAccessKey: value.SecretAccessKey,
		SessionToken:    value.SessionToken,
		Region:          region,
		Service:         "tos",
	}
	r := sign(req.HTTPRequest, c)
	req.HTTPRequest.Header = r.Header
}

func tosBuild(r *request.Request) {
	body := url.Values{}

	if r.Config.ExtraUserAgent != nil && *r.Config.ExtraUserAgent != "" {
		if strings.HasPrefix(*r.Config.ExtraUserAgent, "/") {
			request.AddToUserAgent(r, *r.Config.ExtraUserAgent)
		} else {
			request.AddToUserAgent(r, "/"+*r.Config.ExtraUserAgent)
		}

	}

	r.HTTPRequest.URL.RawQuery = body.Encode()
	r.HTTPRequest.Host = r.HTTPRequest.URL.Host

	if reflect.TypeOf(r.Params) == reflect.TypeOf(&map[string]interface{}{}) {
		m := *(r.Params).(*map[string]interface{})
		for k, v := range m {
			if reflect.TypeOf(v).String() == "string" {
				body.Add(k, v.(string))
			} else {
				body.Add(k, fmt.Sprintf("%v", v))
			}
		}
	}

	r.HTTPRequest.URL.Query()
	if r.Config.ExtraHttpParameters != nil {
		extra := r.Config.ExtraHttpParameters(r.Context())
		if extra != nil {
			for k, value := range extra {
				body.Add(k, value)
			}
		}
	}
	r.HTTPRequest.URL.RawQuery = body.Encode()
}

type tosMetadata struct {
	algorithm       string
	credentialScope string
	signedHeaders   string
	date            string
	region          string
	service         string
}

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	Service         string
	Region          string
	SessionToken    string
}

type tosError struct {
	Code      string
	RequestId string
	HostId    string
	Message   string
}

func sign(request *http.Request, c Credentials) *http.Request {
	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return sign4(request, c)
}

// Sign4 signs a request with Signed Signature Version 4.
func sign4(request *http.Request, credential Credentials) *http.Request {
	keys := credential

	prepareRequestV4(request)
	meta := new(tosMetadata)
	meta.service, meta.region = keys.Service, keys.Region

	// Task 0 设置SessionToken的header
	if credential.SessionToken != "" {
		request.Header.Set("X-Tos-Security-Token", credential.SessionToken)
	}

	// Task 1
	hashedCanonReq := hashedCanonicalRequestV4(request, meta)

	// Task 2
	stringToSign := stringToSignV4(request, hashedCanonReq, meta)

	// Task 3
	signingKey := signingKeyV4(keys.SecretAccessKey, meta.date, meta.region, meta.service)
	signature := signatureV4(signingKey, stringToSign)

	request.Header.Set("Authorization", buildAuthHeaderV4(signature, meta, keys))

	return request
}

func hashedCanonicalRequestV4(request *http.Request, meta *tosMetadata) string {
	payload := readAndReplaceBody(request)
	payloadHash := hashSHA256(payload)
	request.Header.Set("X-Tos-Content-Sha256", payloadHash)

	request.Header.Set("Host", request.Host)

	var sortedHeaderKeys []string
	for key := range request.Header {
		switch key {
		case "Content-Type", "Content-Md5", "Host", "X-Tos-Security-Token":
		default:
			if !strings.HasPrefix(key, "X-") {
				continue
			}
		}
		sortedHeaderKeys = append(sortedHeaderKeys, strings.ToLower(key))
	}
	sort.Strings(sortedHeaderKeys)

	var headersToSign string
	for _, key := range sortedHeaderKeys {
		value := strings.TrimSpace(request.Header.Get(key))
		if key == "host" {
			if strings.Contains(value, ":") {
				split := strings.Split(value, ":")
				port := split[1]
				if port == "80" || port == "443" {
					value = split[0]
				}
			}
		}
		headersToSign += key + ":" + value + "\n"
	}
	meta.signedHeaders = concat(";", sortedHeaderKeys...)
	canonicalRequest := concat("\n", request.Method, normuri(request.URL.Path), normquery(request.URL.Query()), headersToSign, meta.signedHeaders, payloadHash)

	return hashSHA256([]byte(canonicalRequest))
}

func stringToSignV4(request *http.Request, hashedCanonReq string, meta *tosMetadata) string {
	requestTs := request.Header.Get("X-Tos-Date")

	meta.algorithm = "TOS4-HMAC-SHA256"
	meta.date = tsDateV4(requestTs)
	meta.credentialScope = concat("/", meta.date, meta.region, meta.service, "request")

	return concat("\n", meta.algorithm, requestTs, meta.credentialScope, hashedCanonReq)
}

func signatureV4(signingKey []byte, stringToSign string) string {
	return hex.EncodeToString(hmacSHA256(signingKey, stringToSign))
}

func prepareRequestV4(request *http.Request) *http.Request {
	necessaryDefaults := map[string]string{
		"X-Tos-Date": timestampV4(),
	}

	for header, value := range necessaryDefaults {
		if request.Header.Get(header) == "" {
			request.Header.Set(header, value)
		}
	}

	if request.URL.Path == "" {
		request.URL.Path += "/"
	}

	return request
}

func signingKeyV4(secretKey, date, region, service string) []byte {
	kDate := hmacSHA256([]byte(secretKey), date)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "request")
	return kSigning
}

func buildAuthHeaderV4(signature string, meta *tosMetadata, keys Credentials) string {
	credential := keys.AccessKeyID + "/" + meta.credentialScope

	return meta.algorithm +
		" Credential=" + credential +
		", SignedHeaders=" + meta.signedHeaders +
		", Signature=" + signature
}

func timestampV4() string {
	return now().Format("20060102T150405Z")
}

func tsDateV4(timestamp string) string {
	return timestamp[:8]
}

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func hashSHA256(content []byte) string {
	h := sha256.New()
	h.Write(content)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashMD5(content []byte) string {
	h := md5.New()
	h.Write(content)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func readAndReplaceBody(request *http.Request) []byte {
	if request.Body == nil {
		return []byte{}
	}
	payload, _ := ioutil.ReadAll(request.Body)
	request.Body = ioutil.NopCloser(bytes.NewReader(payload))
	return payload
}

func concat(delim string, str ...string) string {
	return strings.Join(str, delim)
}

var now = func() time.Time {
	return time.Now().UTC()
}

func normuri(uri string) string {
	parts := strings.Split(uri, "/")
	for i := range parts {
		parts[i] = encodePathFrag(parts[i])
	}
	return strings.Join(parts, "/")
}

func encodePathFrag(s string) string {
	hexCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			hexCount++
		}
	}
	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		} else {
			t[j] = c
			j++
		}
	}
	return string(t)
}

func shouldEscape(c byte) bool {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
		return false
	}
	if '0' <= c && c <= '9' {
		return false
	}
	if c == '-' || c == '_' || c == '.' || c == '~' {
		return false
	}
	return true
}

func normquery(v url.Values) string {
	queryString := v.Encode()

	return strings.Replace(queryString, "+", "%20", -1)
}

func tosUnmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read volcstackbody err, %v\n", err)
			r.Error = err
			return
		}

		if reflect.TypeOf(r.Data) == reflect.TypeOf(&map[string]interface{}{}) {
			if len(body) == 0 {
				return
			}
			if err = json.Unmarshal(body, &r.Data); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				r.Error = err
				return
			}
		}

	}
}

func tosUnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read volcstackbody err, %v\n", err)
			r.Error = err
			return
		}
		tos := tosError{}
		if err = json.Unmarshal(body, &tos); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			r.Error = err
			return
		}
		r.Error = volcstackerr.NewRequestFailure(
			volcstackerr.New(tos.Code, tos.Message, nil),
			r.HTTPResponse.StatusCode,
			tos.RequestId,
		)

		return
	} else {
		r.Error = volcstackerr.NewRequestFailure(
			volcstackerr.New("ServiceUnavailableException", "service is unavailable", nil),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
		return
	}
}
