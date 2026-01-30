package rule_bound_host_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func ResourceVolcengineTlsRuleBoundHostGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineTlsRuleBoundHostGroupCreate,
		Read:   resourceVolcengineTlsRuleBoundHostGroupRead,
		Delete: resourceVolcengineTlsRuleBoundHostGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the rule.",
			},
			"host_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the host group.",
			},
		},
	}
}

func resourceVolcengineTlsRuleBoundHostGroupCreate(d *schema.ResourceData, meta interface{}) error {
	TlsRuleBoundHostGroupService := NewTlsRuleBoundHostGroupService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(TlsRuleBoundHostGroupService, d, ResourceVolcengineTlsRuleBoundHostGroup()); err != nil {
		return fmt.Errorf("error on creating tls rule bond host group  %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsRuleBoundHostGroupRead(d, meta)
}

func resourceVolcengineTlsRuleBoundHostGroupRead(d *schema.ResourceData, meta interface{}) error {
	TlsRuleBoundHostGroupService := NewTlsRuleBoundHostGroupService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(TlsRuleBoundHostGroupService, d, ResourceVolcengineTlsRuleBoundHostGroup()); err != nil {
		return fmt.Errorf("error on reading tls rule bond host group %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineTlsRuleBoundHostGroupDelete(d *schema.ResourceData, meta interface{}) error {
	TlsRuleBoundHostGroupService := NewTlsRuleBoundHostGroupService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(TlsRuleBoundHostGroupService, d, ResourceVolcengineTlsRuleBoundHostGroup()); err != nil {
		return fmt.Errorf("error on deleting tls rule bond host group %q, %w", d.Id(), err)
	}
	return nil
}
