package clb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineClbs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineClbsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Clb IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Clb.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of Clb.",
			},
			"tags": ve.TagsSchema(),

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Clb query.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the VPC.",
			},
			"eni_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The private ip address of the Clb.",
			},
			"load_balancer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Clb.",
			},
			"clbs": {
				Description: "The collection of Clb query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the Clb.",
						},
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Clb.",
						},
						"load_balancer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Clb.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Clb.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Clb.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Clb.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the Clb.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Clb.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc ID of the Clb.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet ID of the Clb.",
						},
						"modification_protection_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification protection status of the Clb.",
						},
						"modification_protection_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification protection reason of the Clb.",
						},
						"eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Eip ID of the Clb.",
						},
						"eip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Eip address of the Clb.",
						},
						"eni_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Eni ID of the Clb.",
						},
						"eni_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Eni address of the Clb.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the Clb.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason why Clb is locked.",
						},
						"load_balancer_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The specifications of the Clb.",
						},
						"load_balancer_billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing type of the Clb.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the Clb.",
						},
						"tags": ve.TagsSchemaComputed(),
						"master_zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The master zone ID of the CLB.",
						},
						"slave_zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The slave zone ID of the CLB.",
						},
						"renew_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The renew type of the CLB. When the value of the load_balancer_billing_type is `PrePaid`, the query returns this field.",
						},
						"renew_period_times": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The renew period times of the CLB. When the value of the renew_type is `AutoRenew`, the query returns this field.",
						},
						"remain_renew_times": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The remain renew times of the CLB. When the value of the renew_type is `AutoRenew`, the query returns this field.",
						},
						"instance_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The billing status of the CLB.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expired time of the CLB.",
						},
						"reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reclaim time of the CLB.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the Clb.",
						},
						"overdue_reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The over reclaim time of the CLB.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected recycle time of the Clb.",
						},
						"eip_billing_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"isp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ISP of the EIP assigned to CLB, the value can be `BGP`.",
									},
									"eip_billing_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The billing type of the EIP assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.",
									},
									"bandwidth": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The peek bandwidth of the EIP assigned to CLB. The value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.",
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

func dataSourceVolcengineClbsRead(d *schema.ResourceData, meta interface{}) error {
	clbService := NewClbService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(clbService, d, DataSourceVolcengineClbs())
}
