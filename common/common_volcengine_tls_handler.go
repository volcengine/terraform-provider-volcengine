package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

var tlsUnmarshalErrorHandler = request.NamedHandler{Name: "TlsUnmarshalErrorHandler", Fn: tlsUnmarshalError}

func (u *BypassSvc) NewTlsClient() *client.Client {
	svc := "tls"
	config := u.Session.ClientConfig(svc)
	var (
		endpoint string
	)
	if config.Config.DisableSSL != nil && *config.Config.DisableSSL {
		endpoint = fmt.Sprintf("%s://tls-%s.volces.com", "http", config.SigningRegion)
	} else {
		endpoint = fmt.Sprintf("%s://tls-%s.volces.com", "https", config.SigningRegion)
	}

	c := client.New(
		*config.Config,
		metadata.ClientInfo{
			SigningName:   config.SigningName,
			SigningRegion: config.SigningRegion,
			Endpoint:      endpoint,
			ServiceName:   svc,
			ServiceID:     svc,
		},
		config.Handlers,
	)
	c.Handlers.Build.PushBackNamed(corehandlers.SDKVersionUserAgentHandler)
	c.Handlers.Build.PushBackNamed(corehandlers.AddHostExecEnvUserAgentHandler)
	c.Handlers.Sign.PushBackNamed(volc.SignRequestHandler)
	c.Handlers.Build.PushBackNamed(bypassBuildHandler)
	c.Handlers.Unmarshal.PushBackNamed(bypassUnmarshalHandler)
	c.Handlers.UnmarshalError.PushBackNamed(tlsUnmarshalErrorHandler)

	return c
}

type tlsError struct {
	ErrorCode    string
	ErrorMessage string
	RequestId    string
}

func tlsUnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read volcenginebody err, %v\n", err)
			r.Error = err
			return
		}
		tos := tlsError{}
		if err = json.Unmarshal(body, &tos); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			r.Error = err
			return
		}
		r.Error = volcengineerr.NewRequestFailure(
			volcengineerr.New(tos.ErrorCode, tos.ErrorMessage, nil),
			r.HTTPResponse.StatusCode,
			r.HTTPResponse.Header.Get("X-Tls-Requestid"),
		)

		return
	} else {
		r.Error = volcengineerr.NewRequestFailure(
			volcengineerr.New("ServiceUnavailableException", "service is unavailable", nil),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
		return
	}
}
