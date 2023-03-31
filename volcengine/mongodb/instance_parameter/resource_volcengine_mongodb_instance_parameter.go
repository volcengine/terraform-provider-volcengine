package instance_parameter

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongodb parameter can be imported using the param:instanceId:parameterName, e.g.
```
$ terraform import volcengine_mongodb_instance_parameter.default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
```
Note: This resource must be imported before it can be used.
Please note that instance_id and parameter_name must correspond to the ID of the above import.

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
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
				Description: "The instance ID. This field cannot be modified after the resource is imported.",
			},
			"parameter_name": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
				Description: "The name of parameter. This field cannot be modified after the resource is imported.",
			},
			"parameter_role": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Node",
					"Shard",
					"ConfigServer",
					"Mongos",
				}, false),
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
	return fmt.Errorf("This resource must be imported before it can be used. ")
}

func resourceVolcengineMongoDBInstanceParameterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongoDBInstanceParameterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineMongoDBInstanceParameter())
	if err != nil {
		return fmt.Errorf("Error on updating instance %q, %s ", d.Id(), err)
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
		return fmt.Errorf("Error on reading instance %q, %s ", d.Id(), err)
	}
	return err
}

func mongoDBParameterImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 3 || items[0] != "param" {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'param:instanceId:parameterName'")
	}
	d.Set("instance_id", items[1])
	d.Set("parameter_name", items[2])
	return []*schema.ResourceData{d}, nil
}
