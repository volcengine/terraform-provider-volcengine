package vmp_silence_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

// DataSourceVolcengineVmpSilencePolicies 数据源：查询静默策略列表
func DataSourceVolcengineVmpSilencePolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpSilencePoliciesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of silence policy ids.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of silence policy.",
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Active", "Disabled", "Expired"}, false),
				Description:  "The status of silence policy: Active/Disabled/Expired.",
			},
			"sources": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The sources of silence policy: General/LarkBot.",
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
			"silence_policies": {
				Description: "The list of silence policies.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the silence policy.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the silence policy.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the silence policy.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the silence policy.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the silence policy, in RFC3339 format.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the silence policy, in RFC3339 format.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the silence policy.",
						},
						"auto_delete_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The auto delete time of the silence policy.",
						},
						"time_range_matchers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The matching time in the alert silence policy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"date": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The time period for alarm silence.",
									},
									"location": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Time zone.",
									},
									"periodic_date": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The cycle of alarm silence.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"time": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Time periods, e.g. 20:00~21:12,22:00~23:12.",
												},
												"weekday": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Weekdays, e.g. 1,3,5.",
												},
												"day_of_month": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Days of the month, e.g. 1,15,30.",
												},
											},
										},
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

// dataSourceVolcengineVmpSilencePoliciesRead 读取静默策略数据源
func dataSourceVolcengineVmpSilencePoliciesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVmpSilencePolicyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVmpSilencePolicies())
}
