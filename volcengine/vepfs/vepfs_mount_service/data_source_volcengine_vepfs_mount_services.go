package vepfs_mount_service

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVepfsMountServices() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVepfsMountServicesRead,
		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of Vepfs File System.",
			},
			"mount_service_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of mount service.",
			},
			"mount_service_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of mount service. This field support fuzzy query.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The query status list of mount service.",
			},
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

			"mount_services": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the mount service.",
						},
						"mount_service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the mount service.",
						},
						"mount_service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the mount service.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the mount service.",
						},
						"project": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project of the mount service.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the mount service.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id of the mount service.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone id of the mount service.",
						},
						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone name of the mount service.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the mount service.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of the mount service.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the mount service.",
						},
						"nodes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The nodes info of the mount service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of ecs instance.",
									},
									"default_password": {
										Type:        schema.TypeString,
										Computed:    true,
										Sensitive:   true,
										Description: "The default password of ecs instance.",
									},
								},
							},
						},
						"attach_file_systems": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The attached file system info of the mount service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file_system_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the vepfs file system.",
									},
									"file_system_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the vepfs file system.",
									},
									"account_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The account id of the vepfs file system.",
									},
									"customer_path": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the vepfs file system.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the vepfs file system.",
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

func dataSourceVolcengineVepfsMountServicesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVepfsMountServiceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVepfsMountServices())
}
