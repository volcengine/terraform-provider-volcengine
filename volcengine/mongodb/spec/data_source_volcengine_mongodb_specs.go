package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineMongoDBSpecs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineSpecsRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of region query.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region ID to query.",
			},
			"specs": {
				Description: "A list of supported node specification information for MongoDB instances.",
				Type:        schema.TypeList,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mongos_node_specs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The collection of mongos node specs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_num": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The max cpu cores.",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The max connections.",
									},
									"mem_in_gb": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The memory in GB.",
									},
									"spec_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The mongos node spec name.",
									},
								},
							},
						},

						"node_specs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The collection of node specs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_num": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The cpu cores.",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The max connections.",
									},
									"max_storage": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The max storage.",
									},
									"mem_in_db": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The memory in GB.",
									},
									"spec_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node spec name.",
									},
								},
							},
						},
						"shard_node_specs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The collection of shard node specs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_num": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The cpu cores.",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The max connections.",
									},
									"max_storage": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The max storage.",
									},
									"mem_in_gb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The memory in GB.",
									},
									"spec_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The shard node spec name.",
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

func dataSourceVolcengineSpecsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewSpecService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineMongoDBSpecs())
}
