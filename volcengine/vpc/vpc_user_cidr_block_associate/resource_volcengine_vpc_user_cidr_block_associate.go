package vpc_user_cidr_block_associate

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

func ResourceVolcengineVpcUserCidrBlockAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpcUserCidrBlockAssociateCreate,
		Read:   resourceVolcengineVpcUserCidrBlockAssociateRead,
		Delete: resourceVolcengineVpcUserCidrBlockAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the VPC.",
			},
			"user_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The user cidr block of the VPC.",
			},
		},
	}
	return resource
}

func resourceVolcengineVpcUserCidrBlockAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcUserCidrBlockAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVpcUserCidrBlockAssociate())
	if err != nil {
		return fmt.Errorf("error on creating vpc_user_cidr_block_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcUserCidrBlockAssociateRead(d, meta)
}

func resourceVolcengineVpcUserCidrBlockAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcUserCidrBlockAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVpcUserCidrBlockAssociate())
	if err != nil {
		return fmt.Errorf("error on reading vpc_user_cidr_block_associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpcUserCidrBlockAssociateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcUserCidrBlockAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVpcUserCidrBlockAssociate())
	if err != nil {
		return fmt.Errorf("error on updating vpc_user_cidr_block_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcUserCidrBlockAssociateRead(d, meta)
}

func resourceVolcengineVpcUserCidrBlockAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcUserCidrBlockAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVpcUserCidrBlockAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting vpc_user_cidr_block_associate %q, %s", d.Id(), err)
	}
	return err
}
