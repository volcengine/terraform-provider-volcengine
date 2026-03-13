package download_url

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsDownloadUrls() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsDownloadUrlsRead,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the download task.",
			},
			"download_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The download URL of the download task.",
			},
		},
	}
}

func dataSourceVolcengineTlsDownloadUrlsRead(d *schema.ResourceData, meta interface{}) error {
	downloadTaskService := NewTlsDownloadTaskService(meta.(*ve.SdkClient))
	req := map[string]interface{}{
		"task_id": d.Get("task_id"),
	}

	results, err := downloadTaskService.ReadDownloadUrl(req)
	if err != nil {
		return fmt.Errorf("error on reading tls download url %q, %w", d.Get("task_id"), err)
	}

	if len(results) == 0 {
		return fmt.Errorf("download url not found for task %s", d.Get("task_id"))
	}

	result, ok := results[0].(map[string]interface{})
	if !ok {
		return fmt.Errorf("download url result is not map[string]interface{}")
	}
	if taskId, ok := result["task_id"].(string); ok {
		d.SetId(taskId)
	} else {
		return fmt.Errorf("task_id in result is not string")
	}
	d.Set("download_url", result["download_url"])

	return nil
}
