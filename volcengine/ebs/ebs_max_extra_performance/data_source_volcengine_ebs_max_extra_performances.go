package ebs_max_extra_performance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEbsMaxExtraPerformances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEbsMaxExtraPerformancesRead,
		Schema: map[string]*schema.Schema{
			"volume_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the volume.",
			},
			"volume_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the volume. Valid values: `TSSD_TL0`.",
			},
			"size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The size of the volume. Unit: GiB.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},

			"performances": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"baseline": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The baseline of the performance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iops": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The baseline of the iops.",
									},
									"throughput": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The baseline of the throughput.",
									},
								},
							},
						},
						"limit": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The limit of the performance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iops": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The limit of the iops.",
									},
									"throughput": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The limit of the throughput.",
									},
								},
							},
						},
						"max_extra_performance_can_purchase": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The max extra performance can purchase.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"extra_performance_type_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the extra performance.",
									},
									"limit": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The limit of the extra performance.",
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

func dataSourceVolcengineEbsMaxExtraPerformancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEbsMaxExtraPerformanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineEbsMaxExtraPerformances())
}
