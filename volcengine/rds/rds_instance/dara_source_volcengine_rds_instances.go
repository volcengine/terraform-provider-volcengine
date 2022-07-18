package rds_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsInstancesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of RDS instance IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of RDS instance.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of RDS instance query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the RDS instance.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the RDS instance.",
			},
			"instance_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the RDS instance.",
			},
			"create_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of creating RDS instance.",
			},
			"create_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of creating RDS instance.",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region of the RDS instance.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The available zone of the RDS instance.",
			},
			"rds_instances": {
				Description: "The collection of RDS instance query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the RDS instance.",
						},
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the RDS instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the RDS instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the RDS instance.",
						},
						"db_engine": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The engine of the RDS instance.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The engine version of the RDS instance.",
						},
						"storage_space_gb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total storage GB of the RDS instance.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region of the RDS instance.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The available zone of the RDS instance.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc ID of the RDS instance.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the RDS instance.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of the RDS instance.",
						},
						"charge_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge status of the RDS instance.",
						},
						"read_only_instance_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The ID list of read only instance.",
						},
						"instance_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							MinItems:    1,
							Description: "The spec type detail of RDS instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_num": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The cpu core count of spec type.",
									},
									"mem_in_gb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The memory size(GB) of spec type.",
									},
									"spec_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of spec type.",
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

func dataSourceVolcengineRdsInstancesRead(d *schema.ResourceData, meta interface{}) error {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	return rdsInstanceService.Dispatcher.Data(rdsInstanceService, d, DataSourceVolcengineRdsInstances())
}
