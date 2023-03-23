package cluster_bind

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineBioosClusterBindService struct {
	Client *ve.SdkClient
}

func (s *VolcengineBioosClusterBindService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineBioosClusterBindService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListClustersOfWorkspace"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		}
		if err != nil {
			return data, err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.ReqFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineBioosClusterBindService) ReadResource(resData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		res     map[string]interface{}
		ok      bool
	)
	ids := strings.Split(resData.Id(), ":")
	if len(ids) != 2 {
		return nil, nil
	}
	req := map[string]interface{}{
		"ID":   ids[0],
		"Type": resData.Get("type").(string),
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, nil
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	for _, cluster := range data {
		clusterMap, _ := cluster.(map[string]interface{})
		if clusterMap["ID"] == ids[1] {
			res = clusterMap
		}
	}
	if len(res) == 0 {
		return res, fmt.Errorf("Bioos Cluster Bind %s not exist ", id)
	}
	return res, nil
}

func (s *VolcengineBioosClusterBindService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineBioosClusterBindService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineBioosClusterBindService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "BindClusterToWorkspace",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"workspace_id": {
					TargetField: "ID",
				},
				"cluster_id": {
					TargetField: "ClusterID",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["ID"], ":", (*call.SdkParam)["ClusterID"]))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineBioosClusterBindService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineBioosClusterBindService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UnbindClusterAndWorkspace",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"ID":        data.Get("workspace_id").(string),
				"ClusterID": data.Get("cluster_id").(string),
				"Type":      data.Get("type").(string),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineBioosClusterBindService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"workspace_id": {
				TargetField: "ID",
			},
		},
	}
}

func (s *VolcengineBioosClusterBindService) ReadResourceId(id string) string {
	return id
}

func NewVolcengineBioosClusterBindService(c *ve.SdkClient) *VolcengineBioosClusterBindService {
	return &VolcengineBioosClusterBindService{
		Client: c,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "bio",
		Action:      actionName,
		Version:     "2021-03-04",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
