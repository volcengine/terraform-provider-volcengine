package vpc

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VPC can be imported using the id, e.g.
```
$ terraform import volcengine_vpc.default vpc-mizl7m1kqccg5smt1bdpijuj
```

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
		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsCIDR,
				Description:  "A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).",
			},
			"vpc_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the VPC.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the VPC.",
			},
			"dns_servers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsIPAddress,
				},
				Set:         schema.HashString,
				Description: "The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of VPC.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of VPC.",
			},
		},
	}
	ve.MergeDateSourceToResource(DataSourceVolcengineVpcs().Schema["vpcs"].Elem.(*schema.Resource).Schema, &resource.Schema)
	return resource
}

func resourceVolcengineVpcCreate(d *schema.ResourceData, meta interface{}) (err error) {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	err = vpcService.Dispatcher.Create(vpcService, d, ResourceVolcengineVpc())
	if err != nil {
		return fmt.Errorf("error on creating vpc  %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcRead(d, meta)
}

func resourceVolcengineVpcRead(d *schema.ResourceData, meta interface{}) (err error) {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	err = vpcService.Dispatcher.Read(vpcService, d, ResourceVolcengineVpc())
	if err != nil {
		return fmt.Errorf("error on reading vpc %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpcUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	err = vpcService.Dispatcher.Update(vpcService, d, ResourceVolcengineVpc())
	if err != nil {
		return fmt.Errorf("error on updating vpc  %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcRead(d, meta)
}

func resourceVolcengineVpcDelete(d *schema.ResourceData, meta interface{}) (err error) {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	err = vpcService.Dispatcher.Delete(vpcService, d, ResourceVolcengineVpc())
	if err != nil {
		return fmt.Errorf("error on deleting vpc %q, %s", d.Id(), err)
	}
	return err
}
