package cr_repository

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CR Repository can be imported using the registry:namespace:name, e.g.
```
$ terraform import volcengine_cr_repository.default cr-basic:namespace-1:repo-1
```

*/

func crRepositoryImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'registry:namespace:name'")
	}
	if err := d.Set("registry", items[0]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	if err := d.Set("namespace", items[1]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	if err := d.Set("name", items[2]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	return []*schema.ResourceData{d}, nil
}

func ResourceVolcengineCrRepository() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCrRepositoryCreate,
		Read:   resourceVolcengineCrRepositoryRead,
		Update: resourceVolcengineCrRepositoryUpdate,
		Delete: resourceVolcengineCrRepositoryDelete,
		Importer: &schema.ResourceImporter{
			State: crRepositoryImporter,
		},
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The CrRegistry name.",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target namespace name.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of CrRepository.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of CrRepository.",
			},
			"access_level": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Private",
				ValidateFunc: validation.StringInSlice([]string{"Private", "Public"}, false),
				Description:  "The access level of CrRepository.",
			},
		},
	}
	dataSource := DataSourceVolcengineCrRepositories().Schema["repositories"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineCrRepositoryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRepositoryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCrRepository())
	if err != nil {
		return fmt.Errorf("error on creating CrRepository %q,%s", d.Id(), err)
	}
	return resourceVolcengineCrRepositoryRead(d, meta)
}

func resourceVolcengineCrRepositoryUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRepositoryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCrRepository())
	if err != nil {
		return fmt.Errorf("error on updating CrRepository  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrRepositoryRead(d, meta)
}

func resourceVolcengineCrRepositoryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRepositoryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCrRepository())
	if err != nil {
		return fmt.Errorf("error on deleting CrRepository %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCrRepositoryRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrRepositoryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCrRepository())
	if err != nil {
		return fmt.Errorf("Error on reading CrRepository %q,%s", d.Id(), err)
	}
	return err
}
