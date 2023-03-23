package vpc

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VPC can be imported using the id, e.g.
```
$ terraform import volcengine_veenedge_vpc.default vpc-mizl7m1k
```

If you need to create a VPC, you need to apply for permission from the administrator in advance.
You can only delete the vpc from web consul
*/

func ResourceVolcengineVpc() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpcCreate,
		Read:   resourceVolcengineVpcRead,
		Update: resourceVolcengineVpcUpdate,
		Delete: resourceVolcengineVpcDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsCIDR,
				ForceNew:     true,
				Description:  "The cidr info.",
			},
			"vpc_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the VPC.",
			},
			"cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the cluster.",
			},
			"desc": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of the VPC.",
			},
		},
	}
	return resource
}

func resourceVolcengineVpcCreate(d *schema.ResourceData, meta interface{}) (err error) {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(vpcService, d, ResourceVolcengineVpc())
	if err != nil {
		return fmt.Errorf("error on creating vpc  %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcRead(d, meta)
}

func resourceVolcengineVpcRead(d *schema.ResourceData, meta interface{}) (err error) {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(vpcService, d, ResourceVolcengineVpc())
	if err != nil {
		return fmt.Errorf("error on reading vpc %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpcUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(vpcService, d, ResourceVolcengineVpc())
	if err != nil {
		return fmt.Errorf("error on updating vpc  %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcRead(d, meta)
}

func resourceVolcengineVpcDelete(d *schema.ResourceData, meta interface{}) (err error) {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(vpcService, d, ResourceVolcengineVpc())
	if err != nil {
		return fmt.Errorf("error on deleting vpc %q, %s", d.Id(), err)
	}
	return err
}