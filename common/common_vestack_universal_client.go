package common

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack/client"
	"github.com/volcengine/volcstack-go-sdk/volcstack/client/metadata"
	"github.com/volcengine/volcstack-go-sdk/volcstack/corehandlers"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
	"github.com/volcengine/volcstack-go-sdk/volcstack/signer/volc"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackquery"
)

type HttpMethod int

const (
	GET HttpMethod = iota
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
	c.Handlers.Build.PushBackNamed(volcstackquery.BuildHandler)
	c.Handlers.Unmarshal.PushBackNamed(volcstackquery.UnmarshalHandler)
	c.Handlers.UnmarshalMeta.PushBackNamed(volcstackquery.UnmarshalMetaHandler)
	c.Handlers.UnmarshalError.PushBackNamed(volcstackquery.UnmarshalErrorHandler)

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
	default:
		return "GET"
	}
}

func (u *Universal) getContentType(m ContentType) string {
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

	if u.getContentType(info.ContentType) == "application/json" {
		req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	}
	err = req.Send()
	return output, err
}
