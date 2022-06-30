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
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the Clb.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected recycle time of the Clb.",
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
					},
				},
			},
		},
	}
}

func dataSourceVolcengineClbsRead(d *schema.ResourceData, meta interface{}) error {
	clbService := NewClbService(meta.(*ve.SdkClient))
	return clbService.Dispatcher.Data(clbService, d, DataSourceVolcengineClbs())
}
