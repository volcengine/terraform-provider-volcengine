package instance_parameter

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongodb parameter can be imported using the param:instanceId:parameterName:parameterRole, e.g.
```
$ terraform import volcengine_mongodb_instance_parameter.default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
```

*/

func ResourceVolcengineMongoDBInstanceParameter() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineMongoDBInstanceParameterCreate,
		Read:   resourceVolcengineMongoDBInstanceParameterRead,
		Update: resourceVolcengineMongoDBInstanceParameterUpdate,
		Delete: resourceVolcengineMongoDBInstanceParameterDelete,
		Importer: &schema.ResourceImporter{
			State: mongoDBParameterImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance ID.",
			},
			"parameter_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.",
			},
			"parameter_role": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.",
			},
			"parameter_value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of parameter.",
			},
		},
	}

	return resource
}

func resourceVolcengineMongoDBInstanceParameterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceParameterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineMongoDBInstanceParameter())
	if err != nil {
		return fmt.Errorf("Error on creating instance parameters %q, %s ", d.Id(), err)
	}
	return nil
}

func resourceVolcengineMongoDBInstanceParameterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceParameterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineMongoDBInstanceParameter())
	if err != nil {
		return fmt.Errorf("Error on updating instance parameters %q, %s ", d.Id(), err)
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
		return fmt.Errorf("Error on reading instance parameters %q, %s ", d.Id(), err)
	}
	return err
}

func mongoDBParameterImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 4 || items[0] != "param" {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'param:instanceId:parameterName'")
	}
	_ = d.Set("instance_id", items[1])
	_ = d.Set("parameter_name", items[2])
	_ = d.Set("parameter_role", items[3])
	return []*schema.ResourceData{d}, nil
}
