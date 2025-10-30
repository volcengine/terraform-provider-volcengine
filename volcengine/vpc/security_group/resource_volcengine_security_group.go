package security_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
SecurityGroup can be imported using the id, e.g.
```
$ terraform import volcengine_security_group.default sg-273ycgql3ig3k7fap8t3dyvqx
```

*/

func ResourceVolcengineSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineSecurityGroupCreate,
		Read:   resourceVolcengineSecurityGroupRead,
		Update: resourceVolcengineSecurityGroupUpdate,
		Delete: resourceVolcengineSecurityGroupDelete,
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
				Description: "Id of the VPC.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of SecurityGroup.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of SecurityGroup.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of SecurityGroup.",
			},
			"security_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of SecurityGroup.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Type of SecurityGroup. Valid values: `cidr_only`. If this parameter is not specified, it is a normal security group.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of SecurityGroup.",
			},
			"tags": ve.TagsSchema(),
		},
	}
}

func resourceVolcengineSecurityGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	securityGroupService := NewSecurityGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(securityGroupService, d, ResourceVolcengineSecurityGroup())
	if err != nil {
		return fmt.Errorf("error on creating securityGroupService  %q, %w", d.Id(), err)
	}
	return resourceVolcengineSecurityGroupRead(d, meta)
}

func resourceVolcengineSecurityGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	securityGroupService := NewSecurityGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(securityGroupService, d, ResourceVolcengineSecurityGroup())
	if err != nil {
		return fmt.Errorf("error on reading securityGroupService %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineSecurityGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	securityGroupService := NewSecurityGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(securityGroupService, d, ResourceVolcengineSecurityGroup())
	if err != nil {
		return fmt.Errorf("error on updating securityGroupService  %q, %w", d.Id(), err)
	}
	return resourceVolcengineSecurityGroupRead(d, meta)
}

func resourceVolcengineSecurityGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	securityGroupService := NewSecurityGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(securityGroupService, d, ResourceVolcengineSecurityGroup())
	if err != nil {
		return fmt.Errorf("error on deleting securityGroupService %q, %w", d.Id(), err)
	}
	return err
}
