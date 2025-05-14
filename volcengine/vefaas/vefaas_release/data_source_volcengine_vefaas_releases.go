package vefaas_release

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVefaasReleases() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVefaasReleasesRead,
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
			"function_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of Function.",
			},
			"filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Query the filtering conditions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Filter key enumeration.",
						},
						"values": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The filtering value of the query.",
						},
					},
				},
			},
			"order_by": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Query the sorting parameters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Key names used for sorting.",
						},
						"ascend": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether the sorting result is sorted in ascending order.",
						},
					},
				},
			},
			"items": {
				Description: "The list of function publication records.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of function release.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of function release.",
						},
						"function_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Function.",
						},
						"finish_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Finish time.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the published information.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the published information.",
						},
						"last_update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of the published information.",
						},
						"source_revision_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The historical version numbers released.",
						},
						"target_revision_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The target version number released.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVefaasReleasesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVefaasReleaseService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVefaasReleases())
}
