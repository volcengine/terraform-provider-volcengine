package kms_secret

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsSecrets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsSecretsRead,
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
			"filters": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The filters you can specify to query all the secrets that meet the criteria. This parameter is composed of key/value pairs.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project to which the secret belongs.",
			},
			"secrets": {
				Description: "The information about the secret.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of the secret. The value is in the UUID format.",
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
						"secret_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the secret.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The tenant ID of the secret.",
						},
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The information about the tenant resource name (TRN).",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the secret.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the secret.",
						},
						"secret_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the secret.",
						},
						"encryption_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The TRN of the KMS key used to encrypt the secret value.",
						},
						"managed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the secret is hosted.",
						},
						"extended_config": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The extended configurations of the secret.",
						},
						"rotation_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rotation state of the secret.",
						},
						"rotation_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The interval at which automatic rotation is performed.",
						},
						"last_rotation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rotation state of the secret.",
						},
						"schedule_rotation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The next time the secret will be rotated.",
						},
						"secret_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of secret.",
						},
						"schedule_delete_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the secret will be deleted.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsSecretsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsSecretService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsSecrets())
}
