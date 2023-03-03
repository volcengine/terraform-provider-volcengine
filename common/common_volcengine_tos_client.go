package common

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
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
	var content *os.File
	defer func() {
		if content != nil {
			err = content.Close()
		}
	}()
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
		content, _ = os.Open(info.ContentPath)
		req.Body = content
		req.HTTPRequest.Header.Set("Content-Length", strconv.FormatInt(u.tryResolveLength(content), 10))
	}

	if len(info.Header) > 0 {
		for k, v := range info.Header {
			req.HTTPRequest.Header.Set(k, v)
		}
	}

	err = req.Send()
	return output, err
}

func (u *Tos) tryResolveLength(reader io.Reader) int64 {
	switch v := reader.(type) {
	case *bytes.Buffer:
		return int64(v.Len())
	case *bytes.Reader:
		return int64(v.Len())
	case *strings.Reader:
		return int64(v.Len())
	case *os.File:
		length, err := fileUnreadLength(v)
		if err != nil {
			return -1
		}
		return length
	case *io.LimitedReader:
		return v.N
	case *net.Buffers:
		if v != nil {
			length := int64(0)
			for _, p := range *v {
				length += int64(len(p))
			}
			return length
		}
		return 0
	default:
		return -1
	}
}

func fileUnreadLength(file *os.File) (int64, error) {
	offset, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	size := stat.Size()
	if offset > size || offset < 0 {
		return 0, fmt.Errorf("unexpected file size and(or) offset")
	}

	return size - offset, nil
}
