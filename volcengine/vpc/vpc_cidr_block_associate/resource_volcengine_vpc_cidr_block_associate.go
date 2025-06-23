package vpc_cidr_block_associate

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The VpcCidrBlockAssociate is not support import.

*/

func ResourceVolcengineVpcCidrBlockAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpcCidrBlockAssociateCreate,
		Read:   resourceVolcengineVpcCidrBlockAssociateRead,
		Delete: resourceVolcengineVpcCidrBlockAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the VPC.",
			},
			"secondary_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The secondary cidr block of the VPC.",
			},
		},
	}
	return resource
}

func resourceVolcengineVpcCidrBlockAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcCidrBlockAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVpcCidrBlockAssociate())
	if err != nil {
		return fmt.Errorf("error on creating vpc_cidr_block_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcCidrBlockAssociateRead(d, meta)
}

func resourceVolcengineVpcCidrBlockAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcCidrBlockAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVpcCidrBlockAssociate())
	if err != nil {
		return fmt.Errorf("error on reading vpc_cidr_block_associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpcCidrBlockAssociateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcCidrBlockAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVpcCidrBlockAssociate())
	if err != nil {
		return fmt.Errorf("error on updating vpc_cidr_block_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcCidrBlockAssociateRead(d, meta)
}

func resourceVolcengineVpcCidrBlockAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcCidrBlockAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVpcCidrBlockAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting vpc_cidr_block_associate %q, %s", d.Id(), err)
	}
	return err
}
