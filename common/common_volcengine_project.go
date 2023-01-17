package common

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type Project struct {
	ProjectSchemeKey string
	ProjectResultKey string
	ResourceId       string
	ResourceType     string
	Service          string
}

func projectActionInfo(actionName string) UniversalInfo {
	return UniversalInfo{
		ServiceName: "iam",
		Version:     "2021-08-01",
		HttpMethod:  GET,
		ContentType: Default,
		Action:      actionName,
	}
}

func MergeProjectInfo(client *SdkClient, resourceData *schema.ResourceData, project *Project, data *map[string]interface{}) (err error) {
	limit := 1000
	actionName := "ListProjectResources"
	var (
		projectName      interface{}
		output           *map[string]interface{}
		offset           int
		total            interface{}
		projectResources interface{}
		resourceId       interface{}
	)

	projectName = resourceData.Get(project.ProjectSchemeKey)

	if projectName.(string) == "" {
		projectName = (*data)[project.ProjectResultKey]
	}

	if projectName.(string) == "" {
		return nil
	}

	input := map[string]interface{}{
		"ProjectName":    projectName.(string),
		"ResourceRegion": client.Region,
		"ServiceName":    project.Service,
		"Limit":          limit,
		"Offset":         offset,
	}
	for {
		logger.Debug(logger.RespFormat, actionName, input)
		output, err = client.UniversalClient.DoCall(projectActionInfo(actionName), &input)
		if err != nil {
			return err
		}
		total, err = ObtainSdkValue("Result.Total", *output)
		if err != nil {
			return err
		}
		projectResources, err = ObtainSdkValue("Result.ProjectResources", *output)
		if err != nil {
			return err
		}
		for index, _ := range projectResources.([]interface{}) {
			resourceId, err = ObtainSdkValue("Result.ProjectResources."+strconv.Itoa(index)+".ResourceID", *output)
			if err != nil {
				return err
			}
			if resourceId.(string) == project.ResourceId {
				(*data)[project.ProjectResultKey] = projectName
				return nil
			}
		}
		if int(total.(float64)) <= offset+limit {
			break
		}
		offset = (offset + 1) * limit
		input["Offset"] = offset
	}

	return err
}

func UpdateProjectInfo(project *Project) Callback {
	var (
		err       error
		resp      *map[string]interface{}
		accountId interface{}
	)

	return Callback{
		Call: SdkCall{
			Action:      "MoveProjectResource",
			ConvertMode: RequestConvertInConvert,
			Convert: map[string]RequestConvert{
				project.ProjectSchemeKey: {
					TargetField: "TargetProjectName",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (bool, error) {
				if d.HasChange(project.ProjectSchemeKey) {
					_, n := d.GetChange(project.ProjectSchemeKey)
					if n == "" {
						return false, fmt.Errorf("Can not set resource project nil ")
					}
				}

				if _, ok := (*call.SdkParam)["TargetProjectName"]; !ok {
					return false, nil
				}
				actionName := "GetProject"
				input := map[string]interface{}{
					"ProjectName": (*call.SdkParam)["TargetProjectName"].(string),
				}
				resp, err = client.UniversalClient.DoCall(projectActionInfo(actionName), &input)
				if err != nil {
					return false, err
				}

				accountId, err = ObtainSdkValue("Result.AccountID", *resp)
				if err != nil {
					return false, err
				}
				(*call.SdkParam)["ResourceTrn.1"] = "trn:" + project.Service + ":" + client.Region + ":" + strconv.Itoa(int(accountId.(float64))) + ":" + project.ResourceType + "/" + project.ResourceId
				return true, err
			},
			ExecuteCall: func(resourceData *schema.ResourceData, client *SdkClient, call SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam)
				resp, err = client.UniversalClient.DoCall(projectActionInfo(call.Action), call.SdkParam)
				if err != nil {
					return resp, err
				}
				return resp, err
			},
		},
	}
}
