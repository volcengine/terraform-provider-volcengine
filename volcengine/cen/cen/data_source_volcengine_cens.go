package cen

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCens() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCensRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of cen IDs.",
			},
			"cen_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of cen names.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of cen.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of cen query.",
			},
			"tags": ve.TagsSchema(),
			"cens": {
				Description: "The collection of cen query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cen.",
						},
						"cen_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cen.",
						},
						"cen_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cen.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of the cen.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cen.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the cen.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the cen.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the cen.",
						},
						"cen_bandwidth_package_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "A list of bandwidth package IDs of the cen.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCensRead(d *schema.ResourceData, meta interface{}) error {
	cenService := NewCenService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(cenService, d, DataSourceVolcengineCens())
}