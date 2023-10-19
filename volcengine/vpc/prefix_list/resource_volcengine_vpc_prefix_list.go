package prefix_list

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcPrefixList can be imported using the id, e.g.
```
$ terraform import volcengine_vpc_prefix_list.default resource_id
```

*/

func ResourceVolcengineVpcPrefixList() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpcPrefixListCreate,
		Read:   resourceVolcengineVpcPrefixListRead,
		Update: resourceVolcengineVpcPrefixListUpdate,
		Delete: resourceVolcengineVpcPrefixListDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"prefix_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the prefix list.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the prefix list.",
			},
			"ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"IPv4", "IPv6"}, false),
				Description:  "IP version type. Possible values:\nIPv4 (default): IPv4 type.\nIPv6: IPv6 type.",
			},
			"max_entries": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Maximum number of entries, which is the maximum number of entries that can be added to the prefix list. The value range is 1 to 200.",
			},
			"prefix_list_entries": {
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         entriesHash,
				Description: "Prefix list entry list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "CIDR of prefix list entries.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of prefix list entries.",
						},
					},
				},
			},
			"prefix_list_associations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of resources associated with VPC prefix list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Associated resource ID.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Related resource types.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Tags.",
				Set:         tagsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The Value of Tags.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVpcPrefixListCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcPrefixListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVpcPrefixList())
	if err != nil {
		return fmt.Errorf("error on creating vpc_prefix_list %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcPrefixListRead(d, meta)
}

func resourceVolcengineVpcPrefixListRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcPrefixListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVpcPrefixList())
	if err != nil {
		return fmt.Errorf("error on reading vpc_prefix_list %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpcPrefixListUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcPrefixListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVpcPrefixList())
	if err != nil {
		return fmt.Errorf("error on updating vpc_prefix_list %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcPrefixListRead(d, meta)
}

func resourceVolcengineVpcPrefixListDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcPrefixListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVpcPrefixList())
	if err != nil {
		return fmt.Errorf("error on deleting vpc_prefix_list %q, %s", d.Id(), err)
	}
	return err
}

var entriesHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["cidr"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["description"].(string))))
	return hashcode.String(buf.String())
}

var tagsHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["key"], m["value"]))
	return hashcode.String(buf.String())
}
