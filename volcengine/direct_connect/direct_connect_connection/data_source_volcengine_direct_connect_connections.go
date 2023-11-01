package direct_connect_connection

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDirectConnectConnections() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDirectConnectConnectionsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
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
			"direct_connect_connection_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of directi connect connection.",
			},
			"line_operator": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The operator of the physical leased line,valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.",
			},
			"direct_connect_access_point_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the physical leased line access point.",
			},
			"peer_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The peer access point of the physical leased line.",
			},
			"connection_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The connection type of physical leased line,valid value contains `SharedConnection`,`DedicatedConnection`.",
			},
			"tag_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The filter tag of direct connect.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The tag key of cloud resource instance.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The tag value of cloud resource instance.",
						},
					},
				},
			},
			"direct_connect_connections": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The connection type of direct connect.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID which the physical leased line belongs.",
						},
						"direct_connect_connection_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of direct connect connection.",
						},
						"direct_connect_connection_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of direct connect connection.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of direct connect connection.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of direct connect.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of direct connect.",
						},
						"port_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The port type of direct connect.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth of direct connect.",
						},
						"expect_bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The expect bandwidth of direct connect.",
						},
						"line_operator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The operator of physical leased line.",
						},
						"direct_connect_access_point_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The access point id of direct connect.",
						},
						"peer_location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The peer access point of the physical leased line.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of physical leased line.",
						},
						"vlan_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The vlan ID of shared connection,if `connection_type` is `DedicatedConnection`,this parameter returns 0.",
						},
						"parent_connection_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the physical leased line to which the shared leased line belongs. If the physical leased line type is an exclusive leased line, this parameter returns empty.",
						},
						"parent_connection_account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of physical leased line to which the shared leased line belongs.If the physical leased line type is an exclusive leased line,this parameter returns empty.",
						},
						"customer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dedicated line contact name.",
						},
						"customer_contact_phone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dedicated line contact phone.",
						},
						"customer_contact_email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dedicated line contact email.",
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "All tags that physical leased line added.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The tag key.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The tag value.",
									},
								},
							},
						},
						"port_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dedicated line port spec.",
						},
						"billing_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The dedicated line billing type,only support `1` for yearly and monthly billing currently.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dedicated line billing status.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected resource force collection time.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expired time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDirectConnectConnectionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDirectConnectConnectionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDirectConnectConnections())
}
