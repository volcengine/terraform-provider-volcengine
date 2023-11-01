package direct_connect_gateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDirectConnectGateways() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDirectConnectGatewaysRead,
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
			"direct_connect_gateway_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The direst connect gateway name.",
			},
			"cen_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CEN ID which direct connect gateway belongs.",
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
			"direct_connect_gateways": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID that direct connect gateway belongs.",
						},
						"direct_connect_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The direct connect gateway ID.",
						},
						"direct_connect_gateway_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The direct connect gateway name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of direct connect gateway.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of direct connect gateway.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of direct connect gateway.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of direct connect gateway.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of direct connect gateway.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason of the direct connect gateway locked.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource freeze time. Only when the resource is frozen due to arrears, this parameter will have a return value, otherwise it will return a null value.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected resource force collection time. Only when the resource is frozen due to arrears, this parameter will have a return value, otherwise it will return a null value.",
						},
						"associate_cens": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The CEN information associated with the direct connect gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cen_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The cen ID.",
									},
									"cen_owner_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CEN owner's ID.",
									},
									"cen_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CEN status.",
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The tags that direct connect gateway added.",
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
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDirectConnectGatewaysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDirectConnectGatewayService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDirectConnectGateways())
}
