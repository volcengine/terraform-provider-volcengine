package ha_vip

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineHaVips() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineHaVipsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Ha Vip IDs.",
			},
			"ha_vip_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of Ha Vip.",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ip address of Ha Vip.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of Ha Vip.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of Ha Vip.",
			},
			"tags": ve.TagsSchema(),
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of vpc.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of subnet.",
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
			"ha_vips": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the Ha Vip.",
						},
						"ha_vip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the Ha Vip.",
						},
						"ha_vip_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Ha Vip.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ip address of the Ha Vip.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Ha Vip.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Ha Vip.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the Ha Vip.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the Ha Vip.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of the Ha Vip.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Ha Vip.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the Ha Vip.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the Ha Vip.",
						},
						"tags": ve.TagsSchemaComputed(),
						"master_instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The master instance id of the Ha Vip.",
						},
						"associated_eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The associated eip id of the Ha Vip.",
						},
						"associated_eip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The associated eip address of the Ha Vip.",
						},
						"associated_instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The associated instance type of the Ha Vip.",
						},
						"associated_instance_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The associated instance ids of the Ha Vip.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineHaVipsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewHaVipService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineHaVips())
}
