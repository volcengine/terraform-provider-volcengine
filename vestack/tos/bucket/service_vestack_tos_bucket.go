package bucket

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackTosBucketService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketService(c *ve.SdkClient) *VestackTosBucketService {
	return &VestackTosBucketService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackTosBucketService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackTosBucketService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	tos := s.Client.TosClient
	var (
		action  string
		resp    *map[string]interface{}
		results interface{}
	)
	action = "ListBuckets"
	logger.Debug(logger.ReqFormat, action, nil)
	resp, err = tos.DoTosCall(ve.TosInfo{
		HttpMethod: ve.GET,
	}, nil)
	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue("Buckets", *resp)
	if err != nil {
		return data, err
	}
	data = results.([]interface{})
	return data, err
}

func (s *VestackTosBucketService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	results, err = s.ReadResources(nil)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, fmt.Errorf("Value is not map ")
		}
	}

	if len(data) == 0 {
		return data, fmt.Errorf("bucket %s not exist ", instanceId)
	}
	return data, nil
}

func (VestackTosBucketService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return nil
}

func (VestackTosBucketService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VestackTosBucketService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	//create bucket
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceTos,
			Action:          "CreateBucket",
			ConvertMode:     ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
				},
				"tos_acl": {
					ConvertType: ve.ConvertDefault,
					TargetField: "x-tos-acl",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
				"tos_storage_class": {
					ConvertType: ve.ConvertDefault,
					TargetField: "tos-storage-class",
					SpecialParam: &ve.SpecialParam{
						Type: ve.HeaderParam,
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建Bucket
				return s.Client.TosClient.DoTosCall(ve.TosInfo{
					HttpMethod: ve.PUT,
					Domain:     (*call.SdkParam)[ve.TosDomain].(string),
					Header:     (*call.SdkParam)[ve.TosHeader].(map[string]string),
				}, nil)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(s.Client.Region + ":" + (*call.SdkParam)[ve.TosDomain].(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VestackTosBucketService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VestackTosBucketService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VestackTosBucketService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {

	name, ok := data.GetOk("bucket_name")
	return ve.DataSourceInfo{
		ServiceCategory: ve.ServiceTos,
		RequestConverts: map[string]ve.RequestConvert{
			"bucket_name": {
				Ignore: true,
			},
		},
		NameField:        "Name",
		IdField:          "BucketId",
		CollectField:     "buckets",
		ResponseConverts: map[string]ve.ResponseConvert{},
		ExtraData: func(sourceData []interface{}) (extraData []interface{}, err error) {
			for _, v := range sourceData {
				if v.(map[string]interface{})["Location"].(string) != s.Client.Region {
					continue
				}
				if ok {
					if name.(string) == v.(map[string]interface{})["Name"].(string) {
						v.(map[string]interface{})["BucketId"] = s.Client.Region + ":" + v.(map[string]interface{})["Name"].(string)
						extraData = append(extraData, v)
						break
					} else {
						continue
					}
				} else {
					v.(map[string]interface{})["BucketId"] = s.Client.Region + ":" + v.(map[string]interface{})["Name"].(string)
					extraData = append(extraData, v)
				}

			}
			return extraData, err
		},
	}
}

func (VestackTosBucketService) ReadResourceId(s string) string {
	return s[strings.Index(s, ":")+1:]
}
