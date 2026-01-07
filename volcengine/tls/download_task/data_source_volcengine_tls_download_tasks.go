package download_task

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsDownloadTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsDownloadTasksRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "A list of download task IDs.",
			},
			"task_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the download task.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the log topic to which the download tasks belong.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results (after running `terraform plan`).",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of download tasks queried.",
			},
			"download_tasks": {
				Description: "The collection of download task results.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the download task.",
						},
						"task_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the download task.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the log topic to which the download task belongs.",
						},
						"query": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The query statement for the download task.",
						},
						"start_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start time of the log data to download, in Unix timestamp format.",
						},
						"end_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The end time of the log data to download, in Unix timestamp format.",
						},
						"compression": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The compression format of the downloaded file.",
						},
						"data_format": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The data format of the downloaded file.",
						},
						"limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of log entries to download.",
						},
						"sort": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The sorting order of the log data.",
						},
						"task_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the download task.",
						},
						"download_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The download URL for the completed task.",
						},
						"log_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the downloaded log data.",
						},
						"log_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of the downloaded logs.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the download task.",
						},
						"task_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The type of the download task.",
						},
						"allow_incomplete": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to allow incomplete download.",
						},
						"log_context_infos": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The info of the log context.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of the log.",
									},
									"context_flow": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The context flow of the log.",
									},
									"package_offset": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The package offset of the log.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsDownloadTasksRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsDownloadTaskService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsDownloadTasks())
}
