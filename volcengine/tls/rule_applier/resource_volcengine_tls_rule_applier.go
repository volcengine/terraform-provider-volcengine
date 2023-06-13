package rule_applier

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls rule applier can be imported using the rule id and host group id, e.g.
```
$ terraform import volcengine_tls_rule_applier.default fa************:bcb*******
```

*/

func ResourceVolcengineTlsRuleApplier() *schema.Resource {
	return &schema.Resource{
		Read:   resourceVolcengineTlsRuleApplierRead,
		Create: resourceVolcengineTlsRuleApplierCreate,
		Delete: resourceVolcengineTlsRuleApplierDelete,
		Importer: &schema.ResourceImporter{
			State: importTlsRuleApply,
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
				Description: "The id of the rule.",
			},
			"host_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the host group.",
			},
		},
	}
}

func resourceVolcengineTlsRuleApplierCreate(d *schema.ResourceData, meta interface{}) error {
	TlsRuleService := NewTlsRuleApplierService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(TlsRuleService, d, ResourceVolcengineTlsRuleApplier()); err != nil {
		return fmt.Errorf("error on creating tls rule Applier %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsRuleApplierRead(d, meta)
}

func resourceVolcengineTlsRuleApplierRead(d *schema.ResourceData, meta interface{}) error {
	TlsRuleService := NewTlsRuleApplierService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(TlsRuleService, d, ResourceVolcengineTlsRuleApplier()); err != nil {
		return fmt.Errorf("error on reading tls rule Applier %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineTlsRuleApplierDelete(d *schema.ResourceData, meta interface{}) error {
	TlsRuleService := NewTlsRuleApplierService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(TlsRuleService, d, ResourceVolcengineTlsRuleApplier()); err != nil {
		return fmt.Errorf("error on deleting tls rule Applier %q, %w", d.Id(), err)
	}
	return nil
}

func importTlsRuleApply(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form ruleId:hostGroupId")
	}
	err = data.Set("rule_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("host_group_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
