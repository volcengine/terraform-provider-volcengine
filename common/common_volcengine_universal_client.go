package common

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcenginequery"
)

type HttpMethod int

const (
	GET HttpMethod = iota
	HEAD
	POST
	PUT
	DELETE
)

type ContentType int

const (
	Default ContentType = iota
	ApplicationJSON
)

type Universal struct {
	Session *session.Session
}

type UniversalInfo struct {
	ServiceName string
	Action      string
	Version     string
	HttpMethod  HttpMethod
	ContentType ContentType
}

func NewUniversalClient(session *session.Session) *Universal {
	return &Universal{
		Session: session,
	}
}

func (u *Universal) newTargetClient(svc string, version string) *client.Client {
	config := u.Session.ClientConfig(svc)
	c := client.New(
		*config.Config,
		metadata.ClientInfo{
			SigningName:   config.SigningName,
			SigningRegion: config.SigningRegion,
			Endpoint:      config.Endpoint,
			APIVersion:    version,
			ServiceName:   svc,
			ServiceID:     svc,
		},
		config.Handlers,
	)
	c.Handlers.Build.PushBackNamed(corehandlers.SDKVersionUserAgentHandler)
	c.Handlers.Build.PushBackNamed(corehandlers.AddHostExecEnvUserAgentHandler)
	c.Handlers.Sign.PushBackNamed(volc.SignRequestHandler)
	c.Handlers.Build.PushBackNamed(volcenginequery.BuildHandler)
	c.Handlers.Unmarshal.PushBackNamed(volcenginequery.UnmarshalHandler)
	c.Handlers.UnmarshalMeta.PushBackNamed(volcenginequery.UnmarshalMetaHandler)
	c.Handlers.UnmarshalError.PushBackNamed(volcenginequery.UnmarshalErrorHandler)

	return c
}

func (u *Universal) getMethod(m HttpMethod) string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case HEAD:
		return "HEAD"
	default:
		return "GET"
	}
}

func getContentType(m ContentType) string {
	switch m {
	case ApplicationJSON:
		return "application/json"
	default:
		return ""
	}
}

func (u *Universal) DoCall(info UniversalInfo, input *map[string]interface{}) (output *map[string]interface{}, err error) {
	c := u.newTargetClient(info.ServiceName, info.Version)
	op := &request.Operation{
		HTTPMethod: u.getMethod(info.HttpMethod),
		HTTPPath:   "/",
		Name:       info.Action,
	}
	if input == nil {
		input = &map[string]interface{}{}
	}
	output = &map[string]interface{}{}
	req := c.NewRequest(op, input, output)

	if getContentType(info.ContentType) == "application/json" {
		req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	}
	err = req.Send()
	return output, err
}
