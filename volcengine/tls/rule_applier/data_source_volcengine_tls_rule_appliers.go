package rule_applier

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsRuleAppliers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsRuleAppliersRead,
		Schema: map[string]*schema.Schema{
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The rule id.",
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
			"host_group_infos": {
				Description: "The host group info list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host group id.",
						},
						"host_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host group name.",
						},
						"host_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host group type.",
						},
						"host_identifier": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host identifier.",
						},
						"host_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The host count.",
						},
						"rule_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The rule count.",
						},
						"iam_project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The iam project name.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modify time.",
						},
						"auto_update": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to auto update.",
						},
						"update_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update start time.",
						},
						"update_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update end time.",
						},
						"service_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to service logging.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsRuleAppliersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsRuleApplierService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsRuleAppliers())
}
