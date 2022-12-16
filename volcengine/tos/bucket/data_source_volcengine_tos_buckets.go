package bucket

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTosBuckets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTosBucketsRead,
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name the TOS bucket.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of TOS bucket.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of TOS bucket query.",
			},
			"buckets": {
				Description: "The collection of TOS bucket query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name the TOS bucket.",
						},
						"is_truncated": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The truncated the TOS bucket.",
							Deprecated:  "The Field is Deprecated.",
						},
						"marker": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The marker the TOS bucket.",
							Deprecated:  "The Field is Deprecated.",
						},
						"max_keys": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The max keys the TOS bucket.",
							Deprecated:  "The Field is Deprecated.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The prefix the TOS bucket.",
							Deprecated:  "The Field is Deprecated.",
						},
						"creation_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create date of the TOS bucket.",
						},
						"location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The location of the TOS bucket.",
						},
						"extranet_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The extranet endpoint of the TOS bucket.",
						},
						"intranet_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The intranet endpoint the TOS bucket.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTosBucketsRead(d *schema.ResourceData, meta interface{}) error {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	return tosBucketService.Dispatcher.Data(tosBucketService, d, DataSourceVolcengineTosBuckets())
}
