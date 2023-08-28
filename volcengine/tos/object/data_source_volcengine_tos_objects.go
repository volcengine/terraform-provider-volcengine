package object

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTosObjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTosObjectRead,
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name the TOS bucket.",
			},
			"object_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name the TOS Object.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of TOS Object.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of TOS Object query.",
			},
			"objects": {
				Description: "The collection of TOS Object query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name the TOS Object.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The name the TOS Object size.",
						},
						"storage_class": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name the TOS Object storage class.",
						},
						"content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The content the TOS Object when content type is json or text and xml.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTosObjectRead(d *schema.ResourceData, meta interface{}) error {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(tosBucketService, d, DataSourceVolcengineTosObjects())
}
