package access_key

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
AccessKey can be imported using the AccessKeyId:UserName,  e.g.
```
$ terraform import vestack_access_key.default AKLTYmQ2MGFmY2RjNzAxNDQ3NDhiMTZjZmE3MGUyZ****:Name
```

*/

func ResourceVestackAccessKey() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackAccessKeyCreate,
		Read:   resourceVestackAccessKeyRead,
		Update: resourceVestackAccessKeyUpdate,
		Delete: resourceVestackAccessKeyDelete,
		Importer: &schema.ResourceImporter{
			State: akSkImporter,
		},
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The user name.",
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "active",
				ValidateFunc: validation.StringInSlice([]string{"active", "inactive"}, false),
				Description:  "The status of the access key.",
			},
			"pgp_key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.",
			},
			"secret_file": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The file to save the access id and secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.",
			},
			"secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The secret of the access key.",
			},
			"encrypted_secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encrypted secret of the access key by pgp key, base64 encoded.",
			},
			"key_fingerprint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The key fingerprint of the encrypted secret.",
			},
			"create_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create date of the access key.",
			},
		},
	}
	return resource
}

func resourceVestackAccessKeyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAccessKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVestackAccessKey())
	if err != nil {
		return fmt.Errorf("error on creating access key  %q, %s", d.Id(), err)
	}
	return resourceVestackAccessKeyRead(d, meta)
}

func resourceVestackAccessKeyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAccessKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVestackAccessKey())
	if err != nil {
		return fmt.Errorf("error on reading access key %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackAccessKeyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAccessKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVestackAccessKey())
	if err != nil {
		return fmt.Errorf("error on updating access key %q, %s", d.Id(), err)
	}
	return resourceVestackAccessKeyRead(d, meta)
}

func resourceVestackAccessKeyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAccessKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVestackAccessKey())
	if err != nil {
		return fmt.Errorf("error on deleting access key %q, %s", d.Id(), err)
	}
	return err
}
