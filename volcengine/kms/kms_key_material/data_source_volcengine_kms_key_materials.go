package kms_key_material

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsKeyMaterials() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsKeyMaterialsRead,
		Schema: map[string]*schema.Schema{
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of keyring.",
			},
			"key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of key.",
			},
			"key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of key. When key_id is not specified, both keyring_name and key_name must be specified.",
			},
			"wrapping_key_spec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "RSA_2048",
				ValidateFunc: validation.StringInSlice([]string{"RSA_2048", "EC_SM2"}, true),
				Description: "The wrapping key spec. Valid values: `RSA_2048`, `EC_SM2`. Default value: `RSA_2048`. " +
					"When the user's master key protection level is SOFTWARE, selecting EC_SM2 is prohibited.",
			},
			"wrapping_algorithm": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "RSAES_OAEP_SHA_256",
				ValidateFunc: validation.StringInSlice([]string{"RSAES_OAEP_SHA_256", "RSAES_OAEP_SHA_1", "RSAES_PKCS1_V1_5", "SM2PKE"}, true),
				Description: "The wrapping algorithm. Valid values: `RSAES_OAEP_SHA_256`, `RSAES_OAEP_SHA_1`, `RSAES_PKCS1_V1_5`, `SM2PKE`. Default value: `RSAES_OAEP_SHA_256`. " +
					"When the wrapping_key_spec is EC_SM2, only SM2PKE is supported.",
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
			"import_parameters": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The import parameters info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keyring_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of keyring.",
						},
						"key_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of key.",
						},
						"public_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public key used to encrypt key materials, Base64 encoded.",
						},
						"import_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The import token, Base64 encoded.",
						},
						"token_expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The token expire time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsKeyMaterialsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsKeyMaterialService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsKeyMaterials())
}
