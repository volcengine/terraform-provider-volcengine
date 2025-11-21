package nat_gateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNatGateways() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNatGatewaysRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of NatGateway IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "The Name Regex of NatGateway.",
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
				Description: "The total count of NatGateway query.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the VPC.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the Subnet.",
			},
			"spec": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The specification of the NatGateway.",
			},
			"nat_gateway_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the NatGateway.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the NatGateway.",
			},
			"nat_gateways": {
				Description: "The collection of NatGateway query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the NatGateway.",
						},
						"nat_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the NatGateway.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network type of the NatGateway.",
						},
						"nat_gateway_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the NatGateway.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPC.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Subnet.",
						},
						"network_interface_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the network interface.",
						},
						"spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The specification of the NatGateway.",
						},
						"eip_addresses": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allocation_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of Eip.",
									},
									"eip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The address of Eip.",
									},
									"using_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The using status of Eip.",
									},
								},
							},
							Description: "The eip addresses of the NatGateway.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the NatGateway.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the NatGateway.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the NatGateway is locked.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason why locking NatGateway.",
						},
						"billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing type of the NatGateway.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the NatGateway.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deleted time of the NatGateway.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the NatGateway.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the NatGateway.",
						},
						"tags": ve.TagsSchemaComputed(),
						"snat_entry_ids": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: "A list of snat entry ids.",
						},
						"dnat_entry_ids": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: "A list of dnat entry ids.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNatGatewaysRead(d *schema.ResourceData, meta interface{}) error {
	natGatewayService := NewNatGatewayService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(natGatewayService, d, DataSourceVolcengineNatGateways())
}
