package address_book

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
AddressBook can be imported using the id, e.g.
```
$ terraform import volcengine_address_book.default resource_id
```

*/

func ResourceVolcengineAddressBook() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAddressBookCreate,
		Read:   resourceVolcengineAddressBookRead,
		Update: resourceVolcengineAddressBookUpdate,
		Delete: resourceVolcengineAddressBookDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the address book.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the address book.",
			},
			"group_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the address book. Valid values: `ip`, `port`, `domain`.",
			},
			"address_list": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The address list of the address book.\n " +
					"When group_type is `ip`, fill in IPv4/CIDRV4 addresses in the address list.\n " +
					"When group_type is `port`, fill in the port information in the address list, supporting two formats: 22 and 100/200.\n " +
					"When group_type is `domain`, fill in the domain name information in the address list.",
			},

			// computed fields
			"ref_cnt": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The reference count of the address book.",
			},
		},
	}
	return resource
}

func resourceVolcengineAddressBookCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAddressBookService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAddressBook())
	if err != nil {
		return fmt.Errorf("error on creating address_book %q, %s", d.Id(), err)
	}
	return resourceVolcengineAddressBookRead(d, meta)
}

func resourceVolcengineAddressBookRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAddressBookService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAddressBook())
	if err != nil {
		return fmt.Errorf("error on reading address_book %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAddressBookUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAddressBookService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAddressBook())
	if err != nil {
		return fmt.Errorf("error on updating address_book %q, %s", d.Id(), err)
	}
	return resourceVolcengineAddressBookRead(d, meta)
}

func resourceVolcengineAddressBookDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAddressBookService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAddressBook())
	if err != nil {
		return fmt.Errorf("error on deleting address_book %q, %s", d.Id(), err)
	}
	return err
}
