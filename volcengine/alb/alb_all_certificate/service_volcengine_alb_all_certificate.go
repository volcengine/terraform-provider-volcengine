package alb_all_certificate

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineAlbAllCertificateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewAlbAllCertificateService(c *ve.SdkClient) *VolcengineAlbAllCertificateService {
	return &VolcengineAlbAllCertificateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineAlbAllCertificateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineAlbAllCertificateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAllCertificates"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))

		// Extract Certificates from the response
		results, err = ve.ObtainSdkValue("Result.Certificates", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Certificates is not Slice")
		}
		// 过滤系统标签，目前证书不会自动生成系统标签，兜底一下
		data, err = removeSystemTags(data)
		if err != nil {
			return data, err
		}
		return data, err
	})
}

func (s *VolcengineAlbAllCertificateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"Id": id,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("alb_all_certificate %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineAlbAllCertificateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("alb_all_certificate status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineAlbAllCertificateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (VolcengineAlbAllCertificateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineAlbAllCertificateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineAlbAllCertificateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineAlbAllCertificateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "CertificateIds",
				ConvertType: ve.ConvertWithN,
			},
			"certificate_name": {
				TargetField: "CertificateName",
			},
			"certificate_type": {
				TargetField: "CertificateType",
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"value": {
						TargetField: "Values.1",
					},
				},
			},
			"project_name": {
				TargetField: "ProjectName",
			},
		},
		NameField:    "CertificateName",
		IdField:      "CertificateId",
		CollectField: "certificates",
		ResponseConverts: map[string]ve.ResponseConvert{
			"CertificateId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func removeSystemTags(data []interface{}) ([]interface{}, error) {
	var (
		ok      bool
		result  map[string]interface{}
		results []interface{}
		tags    []interface{}
	)
	for _, d := range data {
		if result, ok = d.(map[string]interface{}); !ok {
			return results, errors.New("The elements in data are not map ")
		}
		tags, ok = result["Tags"].([]interface{})
		if ok {
			tags = ve.FilterSystemTags(tags)
			result["Tags"] = tags
		}
		results = append(results, result)
	}
	return results, nil
}

func (s *VolcengineAlbAllCertificateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "alb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
