package cr_vpc_endpoint

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CR Vpc endpoint can be imported using the crVpcEndpoint:registry, e.g.
```
$ terraform import volcengine_cr_vpc_endpoint.default crVpcEndpoint:cr-basic
```

*/

func ResourceVolcengineCrVpcEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineCrVpcEndpointRead,
		Create: resourceVolcengineCrVpcEndpointCreate,
		Update: resourceVolcengineCrVpcEndpointUpdate,
		Delete: resourceVolcengineCrVpcEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: crVpcEndpointImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Cr Registry name.",
			},
			"vpcs": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of vpc meta.",
				Set:         vpcHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the vpc.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The id of the subnet. If not specified, the subnet with the most remaining IPs under the VPC will be automatically selected.",
						},
						"account_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The id of the account. When you need to expose the Enterprise Edition instance to a VPC under another primary account, you need to specify the ID of the primary account to which the VPC belongs.",
						},
					},
				},
			},
		},
	}
	dataSource := DataSourceVolcengineCrVpcEndpoints().Schema["endpoints"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineCrVpcEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrVpcEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCrVpcEndpoint())
	if err != nil {
		return fmt.Errorf("error on creating CrVpcEndpoint %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrVpcEndpointRead(d, meta)
}

func resourceVolcengineCrVpcEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrVpcEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCrVpcEndpoint())
	if err != nil {
		return fmt.Errorf("error on updating CrVpcEndpoint  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrVpcEndpointRead(d, meta)
}

func resourceVolcengineCrVpcEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrVpcEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCrVpcEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting CrVpcEndpoint %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCrVpcEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrVpcEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCrVpcEndpoint())
	if err != nil {
		return fmt.Errorf("error on reading CrVpcEndpoint %q, %s", d.Id(), err)
	}
	return err
}

func crVpcEndpointImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must start with 'crVpcEndpoint:',eg: 'crVpcEndpoint:[registry-1]'")
	}
	if err := d.Set("registry", items[1]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	return []*schema.ResourceData{d}, nil
}

func vpcHash(v interface{}) int {
	var buf bytes.Buffer
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%v:", m["vpc_id"]))
	return hashcode.String(buf.String())
}
