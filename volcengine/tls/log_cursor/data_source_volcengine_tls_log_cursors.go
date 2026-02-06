package log_cursor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsLogCursors() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsLogCursorRead,
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the topic.",
			},
			"shard_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the shard.",
			},
			"from": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The time point of the cursor. The value is a Unix timestamp in seconds, or \"begin\" or \"end\".",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"log_cursors": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of log cursors.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cursor": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cursor value.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the topic.",
						},
						"shard_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the shard.",
						},
						"from": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time point of the cursor.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsLogCursorRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	req := map[string]interface{}{
		"TopicId": d.Get("topic_id"),
		"ShardId": d.Get("shard_id"),
		"From":    d.Get("from"),
	}

	results, err := service.DescribeCursor(req)
	if err != nil {
		return fmt.Errorf("Error reading tls log cursor: %s", err)
	}

	if len(results) == 0 {
		return fmt.Errorf("tls log cursor not found")
	}

	d.SetId(fmt.Sprintf("%s-%d-%s", d.Get("topic_id"), d.Get("shard_id"), d.Get("from")))
	d.Set("log_cursors", results)

	if outputFile, ok := d.GetOk("output_file"); ok && outputFile != "" {
		s, _ := json.MarshalIndent(results, "", "\t")
		if err := ioutil.WriteFile(outputFile.(string), s, 0644); err != nil {
			return fmt.Errorf("Error saving output file: %s", err)
		}
	}

	return nil
}
