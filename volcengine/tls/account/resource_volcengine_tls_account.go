package account

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The TlsAccount is not support import.

*/

func ResourceVolcengineTlsAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineTlsAccountCreate,
		Read:   resourceVolcengineTlsAccountRead,
		Delete: resourceVolcengineTlsAccountDelete,
		Schema: map[string]*schema.Schema{
			"arch_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the log service architecture. Valid values: 2.0 (new architecture), 1.0 (old architecture).",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the log service. Valid values: Activated (already activated), NonActivated (not activated).",
			},
		},
	}
}

func resourceVolcengineTlsAccountCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAccountService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsAccount())
}

func resourceVolcengineTlsAccountRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAccountService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsAccount())
}

func resourceVolcengineTlsAccountDelete(d *schema.ResourceData, meta interface{}) error {
	// No delete API available for tls account
	// The resource is just a representation of the account status, deletion doesn't make sense
	return nil
}
