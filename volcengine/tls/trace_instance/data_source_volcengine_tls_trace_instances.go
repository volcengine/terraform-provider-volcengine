package trace_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsTraceInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsTraceInstancesRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the project.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project.",
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IAM project name.",
			},
			"trace_instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the trace instance.",
			},
			"trace_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the trace instance.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the trace instance.",
			},
			"cs_account_channel": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "CS account channel identifier.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of trace instances.",
			},
			"trace_instances": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of trace instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the project.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the trace instance.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the trace instance.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the trace instance.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the project.",
						},
						"trace_topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the trace topic.",
						},
						"backend_config": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The backend config of the trace instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ttl": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Total log retention time in days.",
									},
									"enable_hot_ttl": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable tiered storage.",
									},
									"hot_ttl": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Standard storage duration in days.",
									},
									"cold_ttl": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Infrequent storage duration in days.",
									},
									"archive_ttl": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Archive storage duration in days.",
									},
									"auto_split": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable auto split.",
									},
									"max_split_partitions": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Max split partitions.",
									},
								},
							},
						},
						"trace_topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the trace topic.",
						},
						"trace_instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the trace instance.",
						},
						"cs_account_channel": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CS account channel identifier.",
						},
						"dependency_topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the dependency topic.",
						},
						"trace_instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the trace instance.",
						},
						"trace_instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the trace instance.",
						},
						"dependency_topic_topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the dependency topic.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsTraceInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTraceInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsTraceInstances())
}
