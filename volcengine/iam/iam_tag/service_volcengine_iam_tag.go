package iam_tag

import (
	"errors"
	"fmt"
	"strconv"

	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamTagService struct {
	Client *ve.SdkClient
}

func NewIamTagService(c *ve.SdkClient) *VolcengineIamTagService {
	return &VolcengineIamTagService{
		Client: c,
	}
}

func (s *VolcengineIamTagService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamTagService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	maxResults := 100
	if val, ok := m["MaxResults"].(int); ok && val > 0 {
		maxResults = val
	}
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", maxResults, nil, func(condition map[string]interface{}) ([]interface{}, string, error) {
		universalClient := s.Client.UniversalClient
		action := "ListTagsForResources"

		// Convert MaxResults to string as required by IAM API
		if val, ok := condition["MaxResults"].(int); ok {
			condition["MaxResults"] = strconv.Itoa(val)
		}

		if _, ok := condition["NextToken"]; !ok {
			condition["NextToken"] = ""
		}

		logger.Debug(logger.ReqFormat, action, condition)

		resp, err := universalClient.DoCall(getUniversalInfo(action), &condition)
		if err != nil {
			return nil, "", err
		}

		logger.Debug(logger.RespFormat, action, resp)
		results, err := ve.ObtainSdkValue("Result.ResourceTags", *resp)
		if err != nil {
			return nil, "", err
		}
		if results == nil {
			results = []interface{}{}
		}
		rawTags, ok := results.([]interface{})
		if !ok {
			return nil, "", errors.New("Result.ResourceTags is not Slice")
		}

		var nextTokenStr string
		nextToken, _ := ve.ObtainSdkValue("Result.NextToken", *resp)
		if nextToken != nil {
			nextTokenStr = nextToken.(string)
		}

		return rawTags, nextTokenStr, nil
	})
}

func (s *VolcengineIamTagService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	// For a single tag resource, we might need to filter from ListTagsForResources
	return nil, errors.New("ReadResource not implemented for iam_tag")
}

func (s *VolcengineIamTagService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamTagService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineIamTagService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "TagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ResourceType"] = d.Get("resource_type")
				resourceNames := d.Get("resource_names").([]interface{})
				for i, name := range resourceNames {
					(*call.SdkParam)[fmt.Sprintf("ResourceNames.%d", i+1)] = name.(string)
				}
				tags := d.Get("tags").(*schema.Set).List()
				for i, tag := range tags {
					tm := tag.(map[string]interface{})
					(*call.SdkParam)[fmt.Sprintf("Tags.%d.Key", i+1)] = tm["key"]
					(*call.SdkParam)[fmt.Sprintf("Tags.%d.Value", i+1)] = tm["value"]
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s:%d", d.Get("resource_type"), time.Now().UnixNano()))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamTagService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// Tags are usually recreated if changed, or handled by adding/removing.
	// For this standalone resource, we'll just recreate it if key/value changes.
	return nil
}

func (s *VolcengineIamTagService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ResourceType"] = d.Get("resource_type")
				resourceNames := d.Get("resource_names").([]interface{})
				for i, name := range resourceNames {
					(*call.SdkParam)[fmt.Sprintf("ResourceNames.%d", i+1)] = name.(string)
				}
				tags := d.Get("tags").(*schema.Set).List()
				for i, tag := range tags {
					tm := tag.(map[string]interface{})
					(*call.SdkParam)[fmt.Sprintf("TagKeys.%d", i+1)] = tm["key"]
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamTagService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"resource_type": {
				TargetField: "ResourceType",
			},
			"resource_names": {
				TargetField: "ResourceNames",
				ConvertType: ve.ConvertListN,
			},
		},
		CollectField: "resource_tags",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ResourceType": {
				TargetField: "resource_type",
			},
			"ResourceName": {
				TargetField: "resource_name",
			},
			"TagKey": {
				TargetField: "tag_key",
			},
			"TagValue": {
				TargetField: "tag_value",
			},
			"NextToken": {
				TargetField: "next_token",
				Chain:       "resource_tags",
			},
		},
	}
}

func (s *VolcengineIamTagService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}
