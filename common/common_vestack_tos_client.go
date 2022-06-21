package common

import (
	"fmt"
	"io/ioutil"

	"github.com/volcengine/volcstack-go-sdk/volcstack/client"
	"github.com/volcengine/volcstack-go-sdk/volcstack/client/metadata"
	"github.com/volcengine/volcstack-go-sdk/volcstack/corehandlers"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
)

const (
	TosInfoUrlParam = "TOS_URL_PARAM"
	TosInfoInput    = "TOS_INPUT"
)

type Tos struct {
	Session *session.Session
}

type TosInfo struct {
	ContentType ContentType
	HttpMethod  HttpMethod
	Path        []string
	UrlParam    map[string]string
	Header      map[string]string
	Domain      string
	ContentPath string
}

func NewTosClient(session *session.Session) *Tos {
	return &Tos{
		Session: session,
	}
}

func (u *Tos) getMethod(m HttpMethod) string {
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

func (u *Tos) newTosClient(domain string) *client.Client {
	svc := "tos"
	config := u.Session.ClientConfig(svc)
	var (
		endpoint string
	)
	if domain == "" {
		if config.Config.DisableSSL != nil && *config.Config.DisableSSL {
			endpoint = fmt.Sprintf("%s://tos-%s.volces.com", "http", config.SigningRegion)
		} else {
			endpoint = fmt.Sprintf("%s://tos-%s.volces.com", "https", config.SigningRegion)
		}
	} else {
		if config.Config.DisableSSL != nil && *config.Config.DisableSSL {
			endpoint = fmt.Sprintf("%s://%s.tos-%s.volces.com", "http", domain, config.SigningRegion)
		} else {
			endpoint = fmt.Sprintf("%s://%s.tos-%s.volces.com", "https", domain, config.SigningRegion)
		}

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
	c.Handlers.Sign.PushBackNamed(tosSignRequestHandler)
	c.Handlers.Build.PushBackNamed(tosBuildHandler)
	c.Handlers.Unmarshal.PushBackNamed(tosUnmarshalHandler)
	c.Handlers.UnmarshalError.PushBackNamed(tosUnmarshalErrorHandler)

	return c
}

func (u *Tos) DoTosCall(info TosInfo, input *map[string]interface{}) (output *map[string]interface{}, err error) {
	c := u.newTosClient(info.Domain)
	trueInput := make(map[string]interface{})
	var httpPath string

	if len(info.Path) > 0 {
		for _, v := range info.Path {
			httpPath = httpPath + "/" + v
		}
	}

	op := &request.Operation{
		HTTPMethod: u.getMethod(info.HttpMethod),
		HTTPPath:   httpPath,
	}
	if input == nil {
		input = &map[string]interface{}{}
	}
	trueInput[TosInfoInput] = input
	if len(info.UrlParam) > 0 {
		trueInput[TosInfoUrlParam] = info.UrlParam
	}
	output = &map[string]interface{}{}
	req := c.NewRequest(op, &trueInput, output)

	if getContentType(info.ContentType) == "application/json" {
		req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if info.ContentPath != "" && (op.HTTPMethod == "PUT" || op.HTTPMethod == "POST") {
		content, _ := ioutil.ReadFile(info.ContentPath)
		req.SetBufferBody(content)
	}

	if len(info.Header) > 0 {
		for k, v := range info.Header {
			req.HTTPRequest.Header.Set(k, v)
		}
	}

	err = req.Send()
	return output, err
}
