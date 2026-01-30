package log_context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

// DataSourceVolcengineTlsLogContexts 定义
func DataSourceVolcengineTlsLogContexts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsLogContextRead,
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the topic.",
			},
			"context_flow": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The context flow of the log.",
			},
			"package_offset": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The package offset of the log.",
			},
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source of the log.",
			},
			"prev_logs": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
				Description: "The number of previous logs.",
			},
			"next_logs": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
				Description: "The number of next logs.",
			},
			"describe_log_context": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to describe log context.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"log_contexts": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of log contexts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prev_over": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the previous logs are over.",
						},
						"next_over": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the next logs are over.",
						},
						"log_context_infos": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The infos of context log.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
					},
				},
			},
		},
	}
}

// dataSourceVolcengineTlsLogContextRead 读取日志上下文
func dataSourceVolcengineTlsLogContextRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	req := map[string]interface{}{
		"TopicId":       d.Get("topic_id"),
		"ContextFlow":   d.Get("context_flow"),
		"PackageOffset": d.Get("package_offset"),
		"Source":        d.Get("source"),
		"PrevLogs":      d.Get("prev_logs"),
		"NextLogs":      d.Get("next_logs"),
	}

	results, err := service.DescribeLogContext(req)
	if err != nil {
		return fmt.Errorf("Error reading tls log context: %s", err)
	}

	if len(results) == 0 {
		return fmt.Errorf("tls log context not found")
	}

	// Process result and set ID
	d.SetId(fmt.Sprintf("%s-%s-%d-%s", d.Get("topic_id"), d.Get("context_flow"), d.Get("package_offset"), d.Get("source")))

	// Set log_contexts list
	d.Set("log_contexts", results)

	// Handle output_file
	if outputFile, ok := d.GetOk("output_file"); ok && outputFile != "" {
		s, _ := json.MarshalIndent(results, "", "\t")
		if err := ioutil.WriteFile(outputFile.(string), s, 0644); err != nil {
			return fmt.Errorf("Error saving output file: %s", err)
		}
	}

	return nil
}
