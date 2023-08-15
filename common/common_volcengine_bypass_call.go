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
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

const (
	BypassInfoUrlParam = "BYPASS_URL_PARAM"
	BypassInfoInput    = "BYPASS_INPUT"
)

type BypassSvc struct {
	Session *session.Session
	info    *BypassSvcInfo
}

type BypassSvcInfo struct {
	ContentType ContentType
	HttpMethod  HttpMethod
	Path        []string
	UrlParam    map[string]string
	Header      map[string]string
	Domain      string
	ContentPath string
	Client      *client.Client
}

func NewBypassClient(session *session.Session) *BypassSvc {
	return &BypassSvc{
		Session: session,
	}
}

func (u *BypassSvc) getMethod(m HttpMethod) string {
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

func (u *BypassSvc) DoBypassSvcCall(info BypassSvcInfo, input *map[string]interface{}) (output *map[string]interface{}, err error) {
	var content *os.File
	defer func() {
		if content != nil {
			err = content.Close()
		}
	}()
	u.info = &info
	var c *client.Client
	if info.Client == nil {
		c = u.NewTosClient()
	} else {
		c = info.Client
	}
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
	trueInput[BypassInfoInput] = input
	if len(info.UrlParam) > 0 {
		trueInput[BypassInfoUrlParam] = info.UrlParam
	}
	output = &map[string]interface{}{}
	req := c.NewRequest(op, &trueInput, output)

	if getContentType(info.ContentType) == "application/json" {
		req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if info.ContentPath != "" && (op.HTTPMethod == "PUT" || op.HTTPMethod == "POST") {
		content, _ = os.Open(info.ContentPath)
		req.Body = content
		req.ResetBody()
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

func (u *BypassSvc) tryResolveLength(reader io.Reader) int64 {
	switch v := reader.(type) {
	case *bytes.Buffer:
		return int64(v.Len())
	case *bytes.Reader:
		return int64(v.Len())
	case *strings.Reader:
		return int64(v.Len())
	case *os.File:
		length, err := u.fileUnreadLength(v)
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

func (u *BypassSvc) fileUnreadLength(file *os.File) (int64, error) {
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
