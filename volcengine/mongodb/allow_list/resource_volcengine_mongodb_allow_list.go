package allow_list

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongodb allow list can be imported using the allowListId, e.g.
```
$ terraform import volcengine_mongodb_allow_list.default acl-d1fd76693bd54e658912e7337d5b****
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
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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
				Description:  "The IP address type of allow list, valid value contains `IPv4`.",
			},
			"allow_list": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: mongoDBAllowListImportDiffSuppress,
				Description:      "IP address or IP address segment in CIDR format.",
			},
			"modify_mode": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				ValidateFunc: validation.StringInSlice([]string{"Cover"}, false),
				Description:  "The modify mode. Only support Cover mode.",
			},
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

func mongoDBAllowListImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if len(old) != len(new) {
		return false
	}
	oldAllowLists := strings.Split(old, ",")
	newAllowLists := strings.Split(new, ",")
	sort.Strings(oldAllowLists)
	sort.Strings(newAllowLists)
	// 根据前后allow list是否相等判断是否modify
	return reflect.DeepEqual(oldAllowLists, newAllowLists)
}
