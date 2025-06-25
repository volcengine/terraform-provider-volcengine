package instance_spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineInstanceSpecs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineInstanceSpecsRead,
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
			"arch_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The architecture type of the Redis instance.",
			},
			"instance_class": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of Redis instance.",
			},
			"instance_specs": {
				Description: "The List of Redis instance specifications.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arch_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The architecture type of the Redis instance.",
						},
						"instance_class": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of Redis instance.",
						},
						"node_numbers": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Description: "The list of the number of nodes allowed to be used per shard. " +
								"The number of nodes allowed for different instance types varies.",
						},
						"shard_numbers": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Description: "The list of shards that the instance is allowed to use. " +
								"The number of shards allowed for use varies among different instance architecture types.",
						},
						"shard_capacity_specs": {
							Description: "The List of capacity specifications for a single shard.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"shard_capacity": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Single-shard memory capacity.",
									},
									"default_bandwidth_per_shard": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default bandwidth of the instance under the current memory capacity.",
									},
									"max_additional_bandwidth_per_shard": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The upper limit of bandwidth that an instance is allowed to modify under the current memory capacity.",
									},
									"max_connections_per_shard": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The default maximum number of connections for a single shard.",
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

func dataSourceVolcengineInstanceSpecsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewInstanceSpecService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineInstanceSpecs())
}
