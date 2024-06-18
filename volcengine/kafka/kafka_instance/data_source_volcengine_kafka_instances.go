package kafka_instance

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

var TagsHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["key"], m["value"]))
	return hashcode.String(buf.String())
}

func DataSourceVolcengineKafkaInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKafkaInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of instance.",
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of instance.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone id of instance.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The tags of instance.",
				Set:         TagsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of tag.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of tag.",
						},
					},
				},
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
			"instances": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of account.",
						},
						"compute_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The compute spec of instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of instance.",
						},
						"eip_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of eip.",
						},
						"instance_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of instance.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of instance.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of instance.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of region.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of zone.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of vpc.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of subnet.",
						},
						"storage_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The storage type of instance.",
						},
						"storage_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The storage space of instance.",
						},
						"usable_partition_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The usable partition number of instance.",
						},
						"used_group_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used group number of instance.",
						},
						"used_partition_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used partition number of instance.",
						},
						"used_topic_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used topic number of instance.",
						},
						"used_storage_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The used storage space of instance.",
						},
						"private_domain_on_public": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether enable private domain on public.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of project.",
						},
						"tags": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The Tags of instance.",
							Set:         TagsHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of tags.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of tags.",
									},
								},
							},
						},
						// charge detail info
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of instance.",
						},
						"charge_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge status of instance.",
						},
						"charge_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge start time of instance.",
						},
						"charge_expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge expire time of instance.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of instance.",
						},
						"overdue_reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue reclaim time of instance.",
						},
						"period_unit": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The period unit of instance.",
						},
						"auto_renew": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The auto renew status of instance.",
						},
						"connection_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Connection info of the instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The endpoint type of instance.",
									},
									"network_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The network type of instance.",
									},
									"internal_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The internal endpoint of instance.",
									},
									"public_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public endpoint of instance.",
									},
								},
							},
						},
						// parameters
						"parameters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameters of the instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter name.",
									},
									"parameter_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter value.",
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

func dataSourceVolcengineKafkaInstancesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKafkaInstanceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKafkaInstances())
}
