package download_task

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls download task can be imported using the id, e.g.
```
$ terraform import volcengine_tls_download_task.default task-1234567890
```

*/

func ResourceVolcengineTlsDownloadTask() *schema.Resource {
	return &schema.Resource{
		Read:   resourceVolcengineTlsDownloadTaskRead,
		Create: resourceVolcengineTlsDownloadTaskCreate,
		Delete: resourceVolcengineTlsDownloadTaskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the log topic to which the download task belongs.",
			},
			"task_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the download task.",
			},
			"query": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The query statement for the download task.",
			},
			"start_time": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The start time of the log data to download, in Unix timestamp format.",
			},
			"end_time": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The end time of the log data to download, in Unix timestamp format.",
			},
			"compression": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "gzip",
				ForceNew:    true,
				Description: "The compression format of the downloaded file. Valid values: gzip.",
			},
			"data_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "csv",
				ForceNew:    true,
				Description: "The data format of the downloaded file. Valid values: csv, json.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     20000000,
				ForceNew:    true,
				Description: "The maximum number of log entries to download.",
			},
			"sort": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "asc",
				ForceNew:    true,
				Description: "The sorting order of the log data. Valid values: asc, desc.",
			},
			"allow_incomplete": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether to allow incomplete download.",
			},
			"task_type": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the download task.",
			},
			"log_context_infos": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The info of the log context.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The source of the log.",
						},
						"context_flow": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The context flow of the log.",
						},
						"package_offset": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "The package offset of the log.",
						},
					},
				},
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the download task.",
			},
			"download_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The download URL for the completed task.",
			},
			"task_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the download task.",
			},
		},
	}
}

func resourceVolcengineTlsDownloadTaskCreate(d *schema.ResourceData, meta interface{}) error {
	downloadTaskService := NewTlsDownloadTaskService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(downloadTaskService, d, ResourceVolcengineTlsDownloadTask()); err != nil {
		return fmt.Errorf("error on creating tls download task %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsDownloadTaskRead(d, meta)
}

func resourceVolcengineTlsDownloadTaskRead(d *schema.ResourceData, meta interface{}) error {
	downloadTaskService := NewTlsDownloadTaskService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(downloadTaskService, d, ResourceVolcengineTlsDownloadTask()); err != nil {
		return fmt.Errorf("error on reading tls download task %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineTlsDownloadTaskDelete(d *schema.ResourceData, meta interface{}) error {
	downloadTaskService := NewTlsDownloadTaskService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(downloadTaskService, d, ResourceVolcengineTlsDownloadTask()); err != nil {
		return fmt.Errorf("error on deleting tls download task %q, %w", d.Id(), err)
	}
	return nil
}
