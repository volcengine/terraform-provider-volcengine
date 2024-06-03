package financial_relation

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
FinancialRelation can be imported using the sub_account_id:relation:relation_id, e.g.
```
$ terraform import volcengine_financial_relation.default resource_id
```

*/

func ResourceVolcengineFinancialRelation() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineFinancialRelationCreate,
		Read:   resourceVolcengineFinancialRelationRead,
		Update: resourceVolcengineFinancialRelationUpdate,
		Delete: resourceVolcengineFinancialRelationDelete,
		Importer: &schema.ResourceImporter{
			State: financialRelationImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"sub_account_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The sub account id.",
			},
			"account_alias": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The display name of the sub account.",
			},
			"relation": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 4}),
				Description:  "The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.",
			},
			"auth_list": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashInt,
				MinItems: 1,
				MaxItems: 5,
				Elem: &schema.Schema{
					Type:         schema.TypeInt,
					ValidateFunc: validation.IntBetween(1, 5),
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("relation").(int) != 4
				},
				Description: "The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.",
			},

			// computed fields
			"relation_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the financial relation.",
			},
			"status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The status of the financial relation.",
			},
		},
	}
	return resource
}

func resourceVolcengineFinancialRelationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFinancialRelationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineFinancialRelation())
	if err != nil {
		return fmt.Errorf("error on creating financial_relation %q, %s", d.Id(), err)
	}
	return resourceVolcengineFinancialRelationRead(d, meta)
}

func resourceVolcengineFinancialRelationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFinancialRelationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineFinancialRelation())
	if err != nil {
		return fmt.Errorf("error on reading financial_relation %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineFinancialRelationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFinancialRelationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineFinancialRelation())
	if err != nil {
		return fmt.Errorf("error on updating financial_relation %q, %s", d.Id(), err)
	}
	return resourceVolcengineFinancialRelationRead(d, meta)
}

func resourceVolcengineFinancialRelationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFinancialRelationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineFinancialRelation())
	if err != nil {
		return fmt.Errorf("error on deleting financial_relation %q, %s", d.Id(), err)
	}
	return err
}

var financialRelationImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	subAccountId, err := strconv.Atoi(items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	relation, err := strconv.Atoi(items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}

	if err := data.Set("sub_account_id", subAccountId); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("relation", relation); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("relation_id", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
