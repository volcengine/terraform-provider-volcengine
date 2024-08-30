package vepfs_file_system

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"strings"
)

func DataSourceVolcengineVepfsFileSystems() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVepfsFileSystemsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Vepfs File System IDs.",
			},
			"store_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Store Type of Vepfs File System.",
			},
			"file_system_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of Vepfs File System. This field support fuzzy query.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The query status list of Vepfs File System.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone id of File System.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project of Vepfs File System.",
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

			"file_systems": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vepfs file system.",
						},
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
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the region.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the zone.",
						},
						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the zone.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth info of the vepfs file system.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version info of the vepfs file system.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the account.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the vepfs file system.",
						},
						"file_system_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the vepfs file system.",
						},
						"store_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The store type of the vepfs file system.",
						},
						"store_type_cn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The store type cn name of the vepfs file system.",
						},
						"protocol_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol type of the vepfs file system.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of the vepfs file system.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the vepfs file system.",
						},
						"charge_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge status of the vepfs file system.",
						},
						"project": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the vepfs file system.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the vepfs file system.",
						},
						"last_modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last modify time of the vepfs file system.",
						},
						"expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expire time of the vepfs file system.",
						},
						"stop_service_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The stop service time of the vepfs file system.",
						},
						"free_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The free time of the vepfs file system.",
						},
						"capacity_info": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The capacity info of the vepfs file system.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total_tib": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The total size. Unit: TiB.",
									},
									"used_gib": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The used size. Unit: GiB.",
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The tags of the vepfs file system.",
							Set:         vepfsTagsResponseHash,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Key of Tags.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Value of Tags.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Type of Tags.",
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

func dataSourceVolcengineVepfsFileSystemsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVepfsFileSystemService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVepfsFileSystems())
}

var vepfsTagsResponseHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["key"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["value"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["type"].(string))))
	return hashcode.String(buf.String())
}
