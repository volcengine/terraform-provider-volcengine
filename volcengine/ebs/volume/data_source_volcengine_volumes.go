package volume

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVolumes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVolumesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Volume IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Volume.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Volume query.",
			},
			"volume_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of Volume.",
			},
			"volume_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of Volume.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Id of Zone.",
			},
			"volume_status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{"available", "attaching", "attached",
					"detaching", "creating", "deleting", "error", "extending"}, false),
				Description: "The Status of Volume, the value can be `available` or `attaching` or `attached` or `detaching` or `creating` or `deleting` or `error` or `extending`.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Id of instance.",
			},
			"kind": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"system", "data"}, false),
				Description:  "The Kind of Volume.",
			},
			"tags": ve.TagsSchema(),
			"volumes": {
				Description: "The collection of Volume query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"billing_type": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"pay_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trade_status": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"expired_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"renew_type": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"delete_with_instance": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"baseline_performance": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The baseline performance of the volume.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iops": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The baseline IOPS performance size for volume.",
									},
									"throughput": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The baseline Throughput performance size for volume. Unit: MB/s.",
									},
								},
							},
						},
						"total_performance": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The baseline performance of the volume.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iops": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The total IOPS performance size for volume.",
									},
									"throughput": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The total Throughput performance size for volume. Unit: MB/s.",
									},
								},
							},
						},
						"extra_performance": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The extra performance of the volume.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"extra_performance_type_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of extra performance for volume.",
									},
									"iops": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The extra IOPS performance size for volume.",
									},
									"throughput": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The extra Throughput performance size for volume. Unit: MB/s.",
									},
								},
							},
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVolumesRead(d *schema.ResourceData, meta interface{}) error {
	volumeService := NewVolumeService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(volumeService, d, DataSourceVolcengineVolumes())
}
