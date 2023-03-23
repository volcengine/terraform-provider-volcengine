package instance_parameter

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongosdb parameter can be imported using the param:instanceId:parameterName, e.g.
```
$ terraform import volcengine_mongodb_instance_parameter.default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
```

*/

func mongoDBParameterImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 3 || items[0] != "param" {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'param:instanceId:parameterName'")
	}
	d.Set("instance_id", items[1])
	d.Set("parameter_name", items[2])
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
			"parameter_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of parameter.",
			},
			"parameter_role": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    true,
				Description: "The node type to which the parameter belongs.",
			},
			"parameter_value": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    true,
				Description: "The value of parameter.",
			},
		},
	}
	dataSource := DataSourceVolcengineMongoDBInstanceParameters().Schema["parameters"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineMongoDBInstanceParameterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	return fmt.Errorf("mongodb instance parameter not allow creating,please import first")
}

func resourceVolcengineMongoDBInstanceParameterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceParameterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineMongoDBInstanceParameter())
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
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineMongoDBInstanceParameter())
	if err != nil {
		return fmt.Errorf("Error on reading instance %q,%s", d.Id(), err)
	}
	return err
}