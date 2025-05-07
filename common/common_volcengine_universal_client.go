package common

import (
	"context"
	"fmt"
	"strings"

	"github.com/volcengine/terraform-provider-volcengine/logger"
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

type RegionType int

const (
	Regional RegionType = iota
	Global
)

var tobRegion = map[string]bool{
	"cn-beijing-autodriving":  true,
	"ap-southeast-3":          true,
	"cn-shanghai-autodriving": true,
}

type Universal struct {
	Session   *session.Session
	endpoints map[string]string
}

type UniversalInfo struct {
	ServiceName string
	Action      string
	Version     string
	HttpMethod  HttpMethod
	ContentType ContentType
	RegionType  RegionType
}

func NewUniversalClient(session *session.Session, endpoints map[string]string) *Universal {
	return &Universal{
		Session:   session,
		endpoints: endpoints,
	}
}

func (u *Universal) loadEndpoint(info UniversalInfo, defaultEndpoint, region string) string {
	var endpoint string
	// firstly, load endpoint from customer_endpoints
	if len(u.endpoints) > 0 {
		if value, ok := u.endpoints[info.ServiceName]; ok && value != "" {
			endpoint = defaultEndpoint[0:strings.Index(defaultEndpoint, "//")] + "//" + value
		}
	}

	// todo: secondly, query endpoint by location DescribeOpenAPIEndpoints

	// thirdly, combine standard endpoint for target region
	if v, exist := tobRegion[region]; exist && v {
		serviceName := strings.ReplaceAll(strings.ToLower(info.ServiceName), "_", "-")
		regionType := getRegionType(info.RegionType)
		var standardEndpoint string
		if regionType == RegionalService {
			standardEndpoint = fmt.Sprintf("%s.%s.%s", serviceName, region, VolcengineIpv4EndpointSuffix)
		} else if regionType == GlobalService {
			standardEndpoint = fmt.Sprintf("%s.%s", serviceName, VolcengineIpv4EndpointSuffix)
		}
		endpoint = defaultEndpoint[0:strings.Index(defaultEndpoint, "//")] + "//" + standardEndpoint
	}

	// lastly, use defaultEndpoint
	if endpoint == "" {
		endpoint = defaultEndpoint
	}
	logger.DebugInfo("service: %s, endpoint: %s", info.ServiceName, endpoint)
	return endpoint
}

func (u *Universal) newTargetClient(info UniversalInfo) *client.Client {
	config := u.Session.ClientConfig(info.ServiceName)
	endpoint := u.loadEndpoint(info, config.Endpoint, config.SigningRegion)

	c := client.New(
		*config.Config,
		metadata.ClientInfo{
			SigningName:   config.SigningName,
			SigningRegion: config.SigningRegion,
			Endpoint:      endpoint,
			APIVersion:    info.Version,
			ServiceName:   info.ServiceName,
			ServiceID:     info.ServiceName,
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

func getRegionType(m RegionType) string {
	switch m {
	case Global:
		return "Global"
	default:
		return "Regional"
	}
}

func (u *Universal) DoCall(info UniversalInfo, input *map[string]interface{}) (output *map[string]interface{}, err error) {
	rate := GetRateInfoMap(info.ServiceName, info.Action, info.Version)
	if rate == nil {
		return u.doCall(info, input)
	}

	// 开始限流
	ctx := context.Background()
	if err = rate.Limiter.Wait(ctx); err != nil {
		return nil, err
	}
	if err = rate.Semaphore.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer func() {
		rate.Semaphore.Release(1)
	}()

	return u.doCall(info, input)
}

func (u *Universal) doCall(info UniversalInfo, input *map[string]interface{}) (output *map[string]interface{}, err error) {
	c := u.newTargetClient(info)
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
