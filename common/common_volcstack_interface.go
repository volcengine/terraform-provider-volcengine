package common

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type ResourceService interface {
	//GetClient 获取客户端
	GetClient() *SdkClient
	// ReadResources 读取资源列表
	ReadResources(map[string]interface{}) ([]interface{}, error)
	// ReadResource 读取单个资源
	ReadResource(*schema.ResourceData, string) (map[string]interface{}, error)
	// RefreshResourceState 刷新资源状态
	RefreshResourceState(*schema.ResourceData, []string, time.Duration, string) *resource.StateChangeConf
	// WithResourceResponseHandlers 接口结果 -> terraform 映射
	WithResourceResponseHandlers(map[string]interface{}) []ResourceResponseHandler
	// CreateResource 创建资源
	CreateResource(*schema.ResourceData, *schema.Resource) []Callback
	// ModifyResource 修改资源
	ModifyResource(*schema.ResourceData, *schema.Resource) []Callback
	// RemoveResource 删除资源
	RemoveResource(*schema.ResourceData, *schema.Resource) []Callback
	// DatasourceResources data_source读取资源
	DatasourceResources(*schema.ResourceData, *schema.Resource) DataSourceInfo
	// ReadResourceId 获取资源ID
	ReadResourceId(string) string
}
