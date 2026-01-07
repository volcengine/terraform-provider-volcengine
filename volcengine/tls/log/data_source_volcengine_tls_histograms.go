package log

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsHistograms() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsHistogramsRead,
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The topic id.",
			},
			"start_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The start time.",
			},
			"end_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The end time.",
			},
			"query": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The query statement.",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The interval.",
			},
			"result_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The result status.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count.",
			},
			"histogram_infos": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The histogram info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count.",
						},
						"start_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start time.",
						},
						"end_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The end time.",
						},
						"result_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The result status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsHistogramsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	req := map[string]interface{}{
		"TopicId":   d.Get("topic_id"),
		"StartTime": d.Get("start_time"),
		"EndTime":   d.Get("end_time"),
		"Query":     d.Get("query"),
	}

	if v, ok := d.GetOk("interval"); ok {
		req["Interval"] = v
	}

	results, err := service.DescribeHistogramV1(req)
	if err != nil {
		return fmt.Errorf("Error reading tls histogram: %s", err)
	}

	if len(results) == 0 {
		return fmt.Errorf("tls histogram not found")
	}

	result := results[0].(map[string]interface{})
	d.SetId(fmt.Sprintf("%s-%d-%d", result["topic_id"], result["start_time"], result["end_time"]))
	d.Set("result_status", result["result_status"])
	d.Set("total_count", result["total_count"])

	if v, ok := result["histogram_infos"].([]interface{}); ok {
		histograms := make([]map[string]interface{}, len(v))
		for i, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				histograms[i] = map[string]interface{}{
					"count":         m["Count"],
					"start_time":    m["StartTime"],
					"end_time":      m["EndTime"],
					"result_status": m["ResultStatus"],
				}
			}
		}
		d.Set("histogram_infos", histograms)
	}

	return nil
}
