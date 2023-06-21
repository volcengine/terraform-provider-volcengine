package common

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
)

type UnsubscribeEnabled interface {
	// UnsubscribeInfo 判断是否需要退订
	UnsubscribeInfo(*schema.ResourceData, *schema.Resource) *UnsubscribeInfo
}

type UnsubscribeInfo struct {
	Product         string
	InstanceId      string
	NeedUnsubscribe bool
}

var unsubscribeRate *Rate

func init() {
	unsubscribeRate = &Rate{
		Limiter:   rate.NewLimiter(10, 10),
		Semaphore: semaphore.NewWeighted(20),
	}
}

type UnsubscribeService struct {
	Client *SdkClient
}

func NewUnsubscribeService(c *SdkClient) *UnsubscribeService {
	return &UnsubscribeService{
		Client: c,
	}
}

func (u *UnsubscribeService) UnsubscribeInstance(info *UnsubscribeInfo) []Callback {
	var call []Callback
	unsubscribe := Callback{
		Call: SdkCall{
			Action:      "UnsubscribeInstance",
			ConvertMode: RequestConvertIgnore,
			ContentType: ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceID"] = info.InstanceId
				(*call.SdkParam)["Product"] = info.Product
				(*call.SdkParam)["UnsubscribeRelatedInstance"] = false
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (*map[string]interface{}, error) {
				defer func() {
					unsubscribeRate.Semaphore.Release(1)
				}()
				var err error
				ctx := context.Background()
				err = unsubscribeRate.Limiter.Wait(ctx)
				if err != nil {
					return nil, err
				}
				err = unsubscribeRate.Semaphore.Acquire(ctx, 1)
				if err != nil {
					return nil, err
				}

				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return u.Client.UniversalClient.DoCall(u.getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *SdkClient, resp *map[string]interface{}, call SdkCall) error {
				return nil
			},
		},
	}
	call = append(call, unsubscribe)
	return call
}

func (u *UnsubscribeService) getUniversalInfo(actionName string) UniversalInfo {
	return UniversalInfo{
		ServiceName: "billing",
		Version:     "2022-01-01",
		HttpMethod:  POST,
		Action:      actionName,
	}
}
