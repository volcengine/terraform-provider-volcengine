package vpc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVpcs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVpcsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of VPC IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Vpc.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the VPC.",
			},
			"vpc_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc name to query.",
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
				Description: "The total count of Vpc query.",
			},
			"vpcs": {
				Description: "The collection of Vpc query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of VPC.",
						},

						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of VPC.",
						},

						"vpc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of VPC.",
						},

						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of VPC.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of VPC.",
						},
						"cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cidr block of VPC.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of VPC.",
						},
						"subnet_ids": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The subnet ID list of VPC.",
						},
						"nat_gateway_ids": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The nat gateway ID list of VPC.",
						},
						"route_table_ids": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The route table ID list of VPC.",
						},
						"security_group_ids": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The security group ID list of VPC.",
						},
						"dns_servers": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The dns server list of VPC.",
						},
						"associate_cens": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cen_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of CEN.",
									},
									"cen_owner_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The owner ID of CEN.",
									},
									"cen_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of CEN.",
									},
								},
							},
							Description: "The associate cen list of VPC.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of VPC.",
						},
						"auxiliary_cidr_blocks": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The auxiliary cidr block list of VPC.",
						},
						"ipv6_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPv6 CIDR block of the VPC.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the VPC.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVpcsRead(d *schema.ResourceData, meta interface{}) error {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(vpcService, d, DataSourceVolcengineVpcs())
}
