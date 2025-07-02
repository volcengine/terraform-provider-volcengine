package waf_prohibition

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafProhibitions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafProhibitionsRead,
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
			"start_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "starting time.",
			},
			"end_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "end time.",
			},
			"reason": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Attack type filtering.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain name of the website that needs to be queried.",
			},
			"letter_order_by": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The list shows the order.",
			},
			"ip_agg_group": {
				Description: "Details of the attack IP.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Attack source IP.",
						},
						"drop_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of attacks on the source IP of this attack.",
						},
						"reason": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "Reason for the ban.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"black": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of visits to the blacklist.",
									},
									"bot": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of Bot attacks.",
									},
									"geo_black": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of geographical location access control.",
									},
									"http_flood": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of CC attacks.",
									},
									"param_abnormal": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of API parameter exceptions.",
									},
									"route_abnormal": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of API routing exceptions.",
									},
									"sensitive_info": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of times sensitive information is leaked.",
									},
									"web_vulnerability": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of Web vulnerability attacks.",
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IP banned status.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status update time.",
						},
						"rule_tag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Ban rule ID.",
						},
						"rule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the ban rule.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineWafProhibitionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafProhibitionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafProhibitions())
}
