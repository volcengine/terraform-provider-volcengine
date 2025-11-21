package nat_ip

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNatIps() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNatIpsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Nat IP ids.",
			},
			"nat_ip_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Nat IP.",
			},
			"nat_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the Nat gateway.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "The Name Regex of Nat ip.",
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
			"nat_ips": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the Nat Ip.",
						},
						"nat_ip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the Nat Ip.",
						},
						"nat_ip_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Nat Ip.",
						},
						"nat_ip_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Nat Ip.",
						},
						"nat_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ip address of the Nat Ip.",
						},
						"nat_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the Nat gateway.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Nat Ip.",
						},
						"is_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the Ip is the default Nat Ip.",
						},
						"using_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The using status of the Nat Ip.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNatIpsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNatIpService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineNatIps())
}
