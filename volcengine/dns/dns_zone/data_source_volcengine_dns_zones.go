package dns_zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineZonesRead,
		Schema: map[string]*schema.Schema{
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
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the domain.",
			},
			"tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"values": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The Value of Tags.",
						},
					},
				},
			},
			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The keyword included in domains.",
			},
			"order_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The key for sorting the results.",
			},
			"search_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The matching mode for the Key parameter.",
			},
			"search_order": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The sorting order of the results.",
			},
			"stage": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the domain.",
			},
			"trade_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The edition of the domain.",
			},
			"zones": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the zone.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the domain.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The most recent update time of the domain.",
						},
						"cache_stage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The most recent update time of the domain.",
						},
						"dns_security": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of DNS DDoS protection service.",
						},
						"expired_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The expiration time of the domain.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the instance.",
						},
						"last_operator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the account that last updated this domain.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The remarks for the domain.",
						},
						"trade_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The edition of the domain.",
						},
						"zid": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the domain.",
						},
						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain name.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the domain.",
						},
						"tags": ve.TagsSchemaComputed(),
						"is_sub_domain": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the domain is a subdomain.",
						},
						"record_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of DNS records contained in the domain.",
						},
						"allocate_dns_server_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The list of DNS servers allocated to the domain by BytePlus DNS.",
						},
						"auto_renew": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether automatic domain renewal is enabled.",
						},
						"instance_no": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the instance. For free edition, the value of this field is null.",
						},
						"is_ns_correct": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the configuration of NS servers is correct. If the configuration is correct, the status of the domain in BytePlus DNS is Active.",
						},
						"real_dns_server_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The list of DNS servers actually used by the domain.",
						},
						"stage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of the domain.",
						},
						"sub_domain_host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain prefix of the subdomain. If the domain is not a subdomain, this parameter is null.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineZonesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewZoneService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineZones())
}
