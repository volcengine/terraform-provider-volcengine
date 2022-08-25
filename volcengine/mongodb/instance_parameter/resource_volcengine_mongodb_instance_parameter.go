package instance_parameter

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongosdb parameter can be imported using the param:instanceId, e.g.
```
$ terraform import volcengine_mongosdb_instance_parameter.default param:mongo-replica-e405f8e2****
```

*/

func mongoDBParameterImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 || items[0] != "param" {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'param:instanceId'")
	}
	return []*schema.ResourceData{d}, nil
}

func ResourceVolcengineMongoDBInstanceParameter() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineMongoDBInstanceParameterCreate,
		Read:   resourceVolcengineMongoDBInstanceParameterRead,
		Update: resourceVolcengineMongoDBInstanceParameterUpdate,
		Delete: resourceVolcengineMongoDBInstanceParameterDelete,
		Importer: &schema.ResourceImporter{
			State: mongoDBParameterImporter,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance ID.",
			},
			"parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The parameters to modify.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parameter_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of parameter.",
						},
						"parameter_role": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The node type to which the parameter belongs.",
						},
						"parameter_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of parameter.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineMongoDBInstanceParameterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
}

func resourceVolcengineMongoDBInstanceParameterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceParameterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineMongoDBInstanceParameter())
	if err != nil {
		return fmt.Errorf("error on updating instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineMongoDBInstanceParameterRead(d, meta)
}

func resourceVolcengineMongoDBInstanceParameterDelete(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
}

func resourceVolcengineMongoDBInstanceParameterRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceParameterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineMongoDBInstanceParameter())
	if err != nil {
		return fmt.Errorf("Error on reading instance %q,%s", d.Id(), err)
	}
	return err
}
