package kms_mac_verification

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsMacVerifications() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsMacVerificationsRead,
		Schema: map[string]*schema.Schema{
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the keyring.",
			},
			"key_name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The name of the key.",
			},
			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.",
			},
			"message": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The message to verify, Base64 encoded.",
			},
			"mac_algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"HMAC_SM3", "HMAC_SHA_256"}, true),
				Description:  "The MAC algorithm. Valid values: `HMAC_SM3`, `HMAC_SHA_256`.",
			},
			"mac": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The MAC to verify, Base64 encoded. Verify the Hash-based Message Authentication Code (HMAC), HMAC KMS key, and MAC algorithm for the specified message.",
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
			"mac_verification_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The MAC verification info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The key id.",
						},
						"mac_valid": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the MAC is valid.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsMacVerificationsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsMacVerificationService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsMacVerifications())
}
