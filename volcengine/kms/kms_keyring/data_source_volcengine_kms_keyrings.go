package kms_keyring

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsKeyrings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsKeyringsRead,
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
			"keyring_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The name of the keyring.",
			},
			"keyring_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The type of the keyring.",
			},
			"description": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The description of the keyring.",
			},
			"creation_date_range": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The creation date of the keyring.",
			},
			"update_date_range": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The update date of the keyring.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project.",
			},
			"keyrings": {
				Description: "The information about the keyring.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of the keyring. The value is in the UUID format.",
						},
						"creation_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The date when the keyring was created.",
						},
						"update_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The date when the keyring was updated.",
						},
						"keyring_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the keyring.",
						},
						"keyring_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the keyring.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the keyring.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The tenant ID of the keyring.",
						},
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The information about the tenant resource name (TRN).",
						},
						"key_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Key ring key count.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsKeyringsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsKeyringService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsKeyrings())
}
