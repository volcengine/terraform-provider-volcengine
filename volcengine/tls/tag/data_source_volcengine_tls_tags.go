package tag

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsTags() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsTagsRead,
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the resource. Valid values: project, topic, shipper, host_group, host, consumer_group, rule, alarm, alarm_notify_group, etl_task, import_task, schedule_sql_task, download_task, trace_instance.",
			},
			"resource_ids": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The IDs of the resources.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"max_results": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of results returned per page. Default value: 20. Maximum value: 100.",
			},
			"next_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The token to get the next page of results. If this parameter is left empty, it means to get the first page of results.",
			},
			"tag_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The tag filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the tag filter.",
						},
						"values": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The values of the tag filter.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The key of the tag.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value of the tag.",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the resource.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the resource.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsTagsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTagService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsTags())
}
