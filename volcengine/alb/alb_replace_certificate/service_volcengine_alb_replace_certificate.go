package alb_replace_certificate

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

type VolcengineAlbReplaceCertificateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewAlbReplaceCertificateService(c *ve.SdkClient) *VolcengineAlbReplaceCertificateService {
	return &VolcengineAlbReplaceCertificateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineAlbReplaceCertificateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineAlbReplaceCertificateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		return []interface{}{}, nil
	})
}

func (s *VolcengineAlbReplaceCertificateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	// 获取证书类型
	certificateType, _ := resourceData.Get("certificate_type").(string)

	// 根据证书类型调用不同的查询API
	var action string
	var condition map[string]interface{}

	if certificateType == "server" {
		action = "DescribeCertificates"
		condition = map[string]interface{}{
			"ServerCertificateIds.1": id,
		}
	} else {
		action = "DescribeCACertificates"
		condition = map[string]interface{}{
			"CACertificateIds.1": id,
		}
	}

	results, err = s.readResourcesWithAction(action, condition)
	if err != nil {
		return data, err
	}

	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("alb_replace_certificate %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineAlbReplaceCertificateService) readResourcesWithAction(action string, condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	bytes, _ := json.Marshal(condition)
	logger.Debug(logger.ReqFormat, action, string(bytes))

	// 统一处理API调用
	var callCondition *map[string]interface{}
	if condition == nil {
		callCondition = nil
	} else {
		callCondition = &condition
	}

	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), callCondition)
	if err != nil {
		return nil, err
	}
	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, string(respBytes))

	// 根据不同的action处理不同的响应字段
	if action == "DescribeCertificates" {
		results, err = ve.ObtainSdkValue("Result.ServerCertificates", *resp)
	} else {
		results, err = ve.ObtainSdkValue("Result.CACertificates", *resp)
	}
	if err != nil {
		return nil, err
	}

	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return nil, errors.New("Result is not Slice")
	}
	return data, nil
}

func (s *VolcengineAlbReplaceCertificateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("alb_replace_certificate status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineAlbReplaceCertificateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// 获取证书类型
	certificateType, _ := resourceData.Get("certificate_type").(string)
	var callback ve.Callback
	if certificateType == "server" {
		callback = s.createServerCertificateResource(resourceData, resource)
	} else {
		callback = s.createCACertificateResource(resourceData, resource)
	}

	return []ve.Callback{callback}
}

func (s *VolcengineAlbReplaceCertificateService) createServerCertificateResource(resourceData *schema.ResourceData, resource *schema.Resource) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ReplaceCertificate",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"certificate_name": {
					TargetField: "CertificateName",
				},
				"description": {
					TargetField: "Description",
				},
				"project_name": {
					TargetField: "ProjectName",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// 设置必需的参数
				(*call.SdkParam)["OldCertificateId"] = d.Get("old_certificate_id").(string)

				// 根据update_mode处理不同的参数
				updateMode := d.Get("update_mode").(string)
				(*call.SdkParam)["UpdateMode"] = updateMode
				if updateMode == "new" {
					// new模式需要PublicKey和PrivateKey
					if publicKey, ok := d.GetOk("public_key"); ok {
						(*call.SdkParam)["PublicKey"] = publicKey
					}
					if privateKey, ok := d.GetOk("private_key"); ok {
						(*call.SdkParam)["PrivateKey"] = privateKey
					}
				} else if updateMode == "stock" {
					// stock模式需要ServerCertificateId
					certificateSource, ok := d.GetOk("certificate_source")
					if !ok {
						return false, fmt.Errorf("certificate_source is required when update_mode is 'stock'")
					}
					(*call.SdkParam)["CertificateSource"] = certificateSource.(string)
					if certificateSource.(string) == "alb" {
						(*call.SdkParam)["CertificateId"] = d.Get("certificate_id").(string)
					} else if certificateSource.(string) == "cert_center" {
						(*call.SdkParam)["CertCenterCertificateId"] = d.Get("cert_center_certificate_id").(string)
					}
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				certificateType := d.Get("certificate_type").(string)
				oldId := d.Get("old_certificate_id").(string)
				d.SetId(fmt.Sprintf("replace:%s:%s", certificateType, oldId))
				return nil
			},
		},
	}
	return callback
}

func (s *VolcengineAlbReplaceCertificateService) createCACertificateResource(resourceData *schema.ResourceData, resource *schema.Resource) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ReplaceCACertificate",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"certificate_name": {
					TargetField: "CACertificateName",
				},
				"description": {
					TargetField: "Description",
				},
				"project_name": {
					TargetField: "ProjectName",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// 设置必需的参数
				(*call.SdkParam)["OldCACertificateId"] = d.Get("old_certificate_id").(string)
				(*call.SdkParam)["UpdateMode"] = d.Get("update_mode").(string)

				// 根据update_mode处理不同的参数
				updateMode := d.Get("update_mode").(string)
				if updateMode == "new" {
					// new模式需要CACertificate
					if caCert, ok := d.GetOk("ca_certificate"); ok {
						(*call.SdkParam)["CACertificate"] = caCert
					}
				} else if updateMode == "stock" {
					// stock模式需要CACertificateId
					if certificateId, ok := d.GetOk("certificate_id"); ok {
						(*call.SdkParam)["CACertificateId"] = certificateId
					}
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				certificateType := d.Get("certificate_type").(string)
				oldId := d.Get("old_certificate_id").(string)
				d.SetId(fmt.Sprintf("replace:%s:%s", certificateType, oldId))
				return nil
			},
		},
	}
	return callback
}

func (VolcengineAlbReplaceCertificateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineAlbReplaceCertificateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineAlbReplaceCertificateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineAlbReplaceCertificateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineAlbReplaceCertificateService) ReadResourceId(id string) string {
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
