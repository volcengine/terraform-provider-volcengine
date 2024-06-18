package kafka_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKafkaGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKafkaGroupsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance id of kafka group.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of kafka group, support fuzzy matching.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of kafka group.",
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
			"groups": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of kafka group.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of kafka group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKafkaGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKafkaGroupService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineKafkaGroups())
}
