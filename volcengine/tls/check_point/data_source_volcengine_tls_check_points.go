package check_point

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsCheckPoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsCheckPointRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the project.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the topic.",
			},
			"shard_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the shard.",
			},
			"consumer_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the consumer group.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"check_points": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of checkpoints.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"checkpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The checkpoint value.",
						},
						"shard_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the shard.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsCheckPointRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	req := map[string]interface{}{
		"ProjectId":         d.Get("project_id"),
		"TopicId":           d.Get("topic_id"),
		"ShardId":           d.Get("shard_id"),
		"ConsumerGroupName": d.Get("consumer_group_name"),
	}

	results, err := service.DescribeCheckPoint(req)
	if err != nil {
		return fmt.Errorf("Error reading tls check point: %s", err)
	}

	if len(results) == 0 {
		return fmt.Errorf("tls check point not found")
	}

	d.SetId(fmt.Sprintf("%s-%s-%s-%s", d.Get("project_id"), d.Get("topic_id"), d.Get("shard_id"), d.Get("consumer_group_name")))
	d.Set("check_points", results)

	if outputFile, ok := d.GetOk("output_file"); ok && outputFile != "" {
		s, _ := json.MarshalIndent(results, "", "\t")
		if err := ioutil.WriteFile(outputFile.(string), s, 0644); err != nil {
			return fmt.Errorf("Error saving output file: %s", err)
		}
	}

	return nil
}
