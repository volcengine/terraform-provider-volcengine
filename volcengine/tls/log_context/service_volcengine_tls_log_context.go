package log_context

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

// Service 日志服务结构体
type Service struct {
	Client *ve.SdkClient
}

// NewService 创建日志服务实例
func NewService(c *ve.SdkClient) *Service {
	return &Service{
		Client: c,
	}
}

// GetClient 获取客户端
func (s *Service) GetClient() *ve.SdkClient {
	return s.Client
}

// ReadResources 读取资源列表 - 实现SearchLogs批量查询日志
func (s *Service) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	logger.Debug(logger.ReqFormat, "ReadResources", m)
	var (
		resp *map[string]interface{}
	)

	// Check if it's a SearchLogs request
	if _, hasQuery := m["Query"]; hasQuery {
		action := "SearchLogs"
		req := map[string]interface{}{
			"TopicId":       m["TopicId"],
			"Query":         m["Query"],
			"StartTime":     m["StartTime"],
			"EndTime":       m["EndTime"],
			"Limit":         m["Limit"],
			"Context":       m["Context"],
			"Sort":          m["Sort"],
			"HighLight":     m["Highlight"],
			"AccurateQuery": m["AccurateQuery"],
		}
		logger.Debug(logger.ReqFormat, action, req)

		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.ApplicationJSON,
			HttpMethod:  ve.POST,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &req)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, action, resp)

		// Extract full response
		response, err := ve.ObtainSdkValue("RESPONSE", *resp)
		if err != nil {
			return data, err
		}

		respMap, ok := response.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("Response is not map")
		}

		// Construct the result map
		result := map[string]interface{}{
			"result_status":       respMap["ResultStatus"],
			"hit_count":           respMap["Count"], // Map Count to hit_count
			"list_over":           respMap["ListOver"],
			"analysis":            respMap["Analysis"],
			"limit":               respMap["Limit"],
			"context":             respMap["Context"],
			"elapsed_millisecond": respMap["ElapsedMillisecond"],
		}

		// Process AnalysisResult
		if ar, ok := respMap["AnalysisResult"].(map[string]interface{}); ok {
			arMap := make(map[string]string)
			for k, v := range ar {
				switch val := v.(type) {
				case string:
					arMap[k] = val
				case []interface{}, map[string]interface{}:
					b, _ := json.Marshal(val)
					arMap[k] = string(b)
				default:
					arMap[k] = fmt.Sprintf("%v", val)
				}
			}
			result["analysis_result"] = arMap
		}

		// Process HighLight
		if highlightRaw, ok := respMap["HighLight"]; ok && highlightRaw != nil {
			if highlightSlice, ok := highlightRaw.([]interface{}); ok {
				var processedHighlights []interface{}
				for _, hlRaw := range highlightSlice {
					if hlMap, ok := hlRaw.(map[string]interface{}); ok {
						newHl := map[string]interface{}{}
						if k, ok := hlMap["Key"]; ok {
							newHl["key"] = k
						}
						if v, ok := hlMap["Value"]; ok {
							newHl["value"] = v
						}
						processedHighlights = append(processedHighlights, newHl)
					}
				}
				result["highlight"] = processedHighlights
			}
		}

		// Process logs
		if logsRaw, ok := respMap["Logs"]; ok && logsRaw != nil {
			if logsSlice, ok := logsRaw.([]interface{}); ok {
				var processedLogs []interface{}
				for _, logRaw := range logsSlice {
					if logMap, ok := logRaw.(map[string]interface{}); ok {
						newLog := map[string]interface{}{}
						// Extract known fields
						if t, ok := logMap["__time__"]; ok {
							newLog["timestamp"] = t
						}
						if s, ok := logMap["__source__"]; ok {
							newLog["source"] = s
						}
						if f, ok := logMap["__path__"]; ok {
							newLog["filename"] = f
						}
						// log_id might be in __context_flow__ or similar? Doc doesn't say.
						// We can try to use a unique ID if needed, or leave empty.

						// Convert all content values to string to satisfy TypeMap schema
						contentMap := make(map[string]string)
						for k, v := range logMap {
							switch val := v.(type) {
							case string:
								contentMap[k] = val
							case float64:
								if val == float64(int64(val)) {
									contentMap[k] = fmt.Sprintf("%.0f", val)
								} else {
									contentMap[k] = fmt.Sprintf("%v", val)
								}
							case int, int32, int64:
								contentMap[k] = fmt.Sprintf("%d", val)
							case bool:
								contentMap[k] = fmt.Sprintf("%t", val)
							default:
								contentMap[k] = fmt.Sprintf("%v", val)
							}
						}
						newLog["content"] = contentMap
						processedLogs = append(processedLogs, newLog)
					}
				}
				result["logs"] = processedLogs
			}
		}

		// Add identifying info for the data source itself
		result["topic_id"] = req["TopicId"]

		// Return as single item list
		return []interface{}{result}, nil
	}

	// Default: return empty slice for other operations
	return []interface{}{}, nil
}

// ReadResource 读取单个资源 - 实现DescribeLogContext获取日志上下文
func (s *Service) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	// Check if it's a DescribeLogContext request
	logID := resourceData.Get("log_id").(string)
	if logID != "" {
		action := "DescribeLogContext"
		req := map[string]interface{}{
			"TopicId":  resourceData.Get("topic_id"),
			"LogId":    logID,
			"PrevLogs": resourceData.Get("prev_logs"),
			"NextLogs": resourceData.Get("next_logs"),
		}
		logger.Debug(logger.ReqFormat, action, req)

		resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.ApplicationJSON,
			HttpMethod:  ve.POST,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &req)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, action, resp)

		// Extract context logs from response
		result, err := ve.ObtainSdkValue("RESPONSE", *resp)
		if err != nil {
			return data, err
		}
		if result == nil {
			return map[string]interface{}{}, nil
		}
		if data, ok := result.(map[string]interface{}); ok {
			return data, nil
		}
		return map[string]interface{}{}, fmt.Errorf("Response is not map[string]interface{}")
	}

	// Default: return empty map for other operations
	return map[string]interface{}{},
		nil
}

// RefreshResourceState 刷新资源状态
func (s *Service) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

// WithResourceResponseHandlers 接口结果 -> terraform 映射
func (s *Service) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

// CreateResource 创建资源
func (s *Service) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

// ModifyResource 修改资源
func (s *Service) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

// RemoveResource 删除资源
func (s *Service) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	// Logs API doesn't support delete operation, so we return empty callbacks
	return []ve.Callback{}
}

// DatasourceResources data_source读取资源
func (s *Service) DatasourceResources(d *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	if _, ok := resource.Schema["log_contexts"]; ok {
		return ve.DataSourceInfo{
			IdField:      "topic_id",
			CollectField: "log_contexts",
		}
	}

	if _, ok := resource.Schema["histograms"]; ok {
		return ve.DataSourceInfo{
			IdField:      "HistogramId",
			CollectField: "histograms",
			ExtraData: func(i []interface{}) ([]interface{}, error) {
				for index, ele := range i {
					element := ele.(map[string]interface{})
					i[index].(map[string]interface{})["HistogramId"] = fmt.Sprintf("%s-%d-%d", element["topic_id"], element["start_time"], element["end_time"])
				}
				return i, nil
			},
		}
	}

	return ve.DataSourceInfo{
		IdField:      "topic_id", // Use topic_id as ID for the search result resource
		CollectField: "logs",
	}
}

// DescribeHistogramV1 获取日志直方图
func (s *Service) DescribeHistogramV1(m map[string]interface{}) (data []interface{}, err error) {
	action := "DescribeHistogramV1"
	req := map[string]interface{}{
		"TopicId":   m["TopicId"],
		"StartTime": m["StartTime"],
		"EndTime":   m["EndTime"],
		"Query":     m["Query"],
		"Interval":  m["Interval"],
	}
	logger.Debug(logger.ReqFormat, action, req)

	resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.ApplicationJSON,
		HttpMethod:  ve.POST,
		Path:        []string{action},
		Client:      s.Client.BypassSvcClient.NewTlsClient(),
	}, &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	// Extract histogram from response
	histogram, err := ve.ObtainSdkValue("RESPONSE.Histogram", *resp)
	if err != nil {
		return data, err
	}
	resultStatus, _ := ve.ObtainSdkValue("RESPONSE.ResultStatus", *resp)
	totalCount, _ := ve.ObtainSdkValue("RESPONSE.TotalCount", *resp)

	// Return as a list of one item
	data = []interface{}{
		map[string]interface{}{
			"histogram_infos": histogram,
			"result_status":   resultStatus,
			"total_count":     totalCount,
			"topic_id":        req["TopicId"],
			"start_time":      req["StartTime"],
			"end_time":        req["EndTime"],
			"query":           req["Query"],
			"interval":        req["Interval"],
		},
	}
	return data, nil
}

// DescribeLogContext 获取日志上下文
func (s *Service) DescribeLogContext(m map[string]interface{}) (data []interface{}, err error) {
	action := "DescribeLogContext"
	req := map[string]interface{}{
		"TopicId":       m["TopicId"],
		"ContextFlow":   m["ContextFlow"],
		"PackageOffset": m["PackageOffset"],
		"Source":        m["Source"],
		"PrevLogs":      m["PrevLogs"],
		"NextLogs":      m["NextLogs"],
	}
	logger.Debug(logger.ReqFormat, action, req)

	resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.ApplicationJSON,
		HttpMethod:  ve.POST,
		Path:        []string{action},
		Client:      s.Client.BypassSvcClient.NewTlsClient(),
	}, &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	// Extract context logs from response
	infosRaw, err := ve.ObtainSdkValue("RESPONSE.LogContextInfos", *resp)
	if err != nil {
		return data, err
	}
	prevOver, _ := ve.ObtainSdkValue("RESPONSE.PrevOver", *resp)
	nextOver, _ := ve.ObtainSdkValue("RESPONSE.NextOver", *resp)

	var infos []interface{}
	if infosRaw != nil {
		if rawList, ok := infosRaw.([]interface{}); ok {
			for _, item := range rawList {
				if logMap, ok := item.(map[string]interface{}); ok {
					newMap := make(map[string]string)
					for k, v := range logMap {
						switch val := v.(type) {
						case string:
							newMap[k] = val
						case float64:
							if val == float64(int64(val)) {
								newMap[k] = fmt.Sprintf("%.0f", val)
							} else {
								newMap[k] = fmt.Sprintf("%v", val)
							}
						case int, int32, int64:
							newMap[k] = fmt.Sprintf("%d", val)
						case bool:
							newMap[k] = fmt.Sprintf("%t", val)
						default:
							newMap[k] = fmt.Sprintf("%v", val)
						}
					}
					infos = append(infos, newMap)
				}
			}
		}
	}

	result := map[string]interface{}{
		"prev_over":         prevOver,
		"next_over":         nextOver,
		"log_context_infos": infos,
	}

	return []interface{}{result}, nil
}

// ReadResourceId 获取资源ID
func (s *Service) ReadResourceId(id string) string {
	return id
}
