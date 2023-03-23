package allow_list

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongosdb allow list can be imported using the allowListId, e.g.
```
$ terraform import volcengine_mongosdb_allow_list.default acl-d1fd76693bd54e658912e7337d5b****
```

*/

func ResourceVolcengineMongoDBAllowList() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineMongoDBAllowListCreate,
		Read:   resourceVolcengineMongoDBAllowListRead,
		Update: resourceVolcengineMongoDBAllowListUpdate,
		Delete: resourceVolcengineMongoDBAllowListDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"allow_list_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of allow list.",
			},
			"allow_list_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of allow list.",
			},
			"allow_list_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"IPv4"}, false),
				Description:  "The IP address type of allow list,valid value contains `IPv4`.",
			},
			"allow_list": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address or IP address segment in CIDR format.",
			},
			"modify_mode": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Default:      "Cover",
				ValidateFunc: validation.StringInSlice([]string{"Cover", "Append", "Delete"}, false),
				Description:  "The modify mode.",
			},
			//"apply_instance_num": {
			//	Type:     schema.TypeInt,
			//	Optional: true,
			//	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			//		return d.Id() == ""
			//	},
			//	Description: "The instance number bound to the allow list,this parameter is required if you need to modify `AllowList`.",
			//},
		},
	}
	return resource
}

func resourceVolcengineMongoDBAllowListCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBAllowListService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineMongoDBAllowList())
	if err != nil {
		return fmt.Errorf("Error on creating instance %q,%s", d.Id(), err)
	}
	return resourceVolcengineMongoDBAllowListRead(d, meta)
}

func resourceVolcengineMongoDBAllowListUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBAllowListService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineMongoDBAllowList())
	if err != nil {
		return fmt.Errorf("error on updating instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineMongoDBAllowListRead(d, meta)
}

func resourceVolcengineMongoDBAllowListDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBAllowListService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineMongoDBAllowList())
	if err != nil {
		return fmt.Errorf("error on deleting instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongoDBAllowListRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBAllowListService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineMongoDBAllowList())
	if err != nil {
		return fmt.Errorf("Error on reading instance %q,%s", d.Id(), err)
	}
	return err
}
