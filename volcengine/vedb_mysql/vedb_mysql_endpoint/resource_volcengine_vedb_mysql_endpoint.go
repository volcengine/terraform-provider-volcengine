package vedb_mysql_endpoint

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VedbMysqlEndpoint can be imported using the instance id:endpoint id, e.g.
```
$ terraform import volcengine_vedb_mysql_endpoint.default vedbm-iqnh3a7z****:vedbm-2pf2xk5v****-Custom-50yv
```
Note: The master node endpoint only supports modifying the EndpointName and Description parameters. If values are passed in for other parameters, these values will be ignored without generating an error.
The default endpoint does not support modifying the ReadWriteMode, AutoAddNewNodes, and Nodes parameters. If values are passed in for these parameters, these values will be ignored without generating an error.
*/

func ResourceVolcengineVedbMysqlEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlEndpointCreate,
		Read:   resourceVolcengineVedbMysqlEndpointRead,
		Update: resourceVolcengineVedbMysqlEndpointUpdate,
		Delete: resourceVolcengineVedbMysqlEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: vedbMysqlEndpointImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"endpoint_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the endpoint.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the instance.",
			},
			"endpoint_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Connect endpoint type. The value is fixed as Custom, indicating a custom endpoint.",
			},
			"read_write_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Endpoint read-write mode. Values:\n ReadWrite: Read and write endpoint.\n ReadOnly: Read-only endpoint (default).",
			},
			"endpoint_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Connect the endpoint name. The setting rules are as follows:\n " +
					"It cannot start with a number or a hyphen (-).\n " +
					"It can only contain Chinese characters, letters, numbers, underscores (_), and hyphens (-).\n " +
					"The length is 1 to 64 characters.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description information for connecting endpoint. The length cannot exceed 200 characters.",
			},
			"node_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Connect the node IDs associated with the endpoint." +
					"The filling rules are as follows:\n" +
					"When the value of ReadWriteMode is ReadWrite, at least two nodes must be passed in, and the master node must be passed in.\n" +
					"When the value of ReadWriteMode is ReadOnly, one or more read-only nodes can be passed in.",
			},
			// 这个不能暴露出去，要不会触发node ids变更
			//"auto_add_new_nodes": {
			//	Type:     schema.TypeBool,
			//	Optional: true,
			//	Default:  false,
			//	Description: "Set whether newly created read-only nodes will automatically join this connection endpoint." +
			//		" Values:\ntrue: Automatically join.\nfalse: Do not automatically join (default).",
			//},
			"master_accept_read_requests": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					//仅当 ReadWriteMode 取值为ReadWrite 时，支持开启主节点接受读。
					if d.Get("read_write_mode").(string) != "ReadWrite" {
						return true
					}
					return false
				},
				Description: "The master node accepts read requests. " +
					"Value range:\ntrue: (default) After enabling the master node to accept read functions, " +
					"non-transactional read requests will be sent to the master node or read-only nodes in a load-balanced mode according to the number of active requests." +
					"\nfalse: After disabling the master node from accepting read requests, at this time," +
					" the master node only accepts transactional read requests, " +
					"and non-transactional read requests will not be sent to the master node.\n" +
					"Description\nOnly when the value of ReadWriteMode is ReadWrite, " +
					"enabling the master node to accept reads is supported.",
			},
			"distributed_transaction": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					//仅当 ReadWriteMode 取值为ReadWrite 时，支持开启事务拆分。
					if d.Get("read_write_mode").(string) != "ReadWrite" {
						return true
					}
					return false
				},
				Description: "Set whether to enable transaction splitting. " +
					"For detailed introduction to transaction splitting, please refer to transaction splitting." +
					" Value range:\ntrue: Enabled (default).\nfalse: Disabled.\n" +
					"Description\nOnly when the value of ReadWriteMode is ReadWrite, is enabling transaction splitting supported.",
			},
			"consist_level": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Consistency level. " +
					"For detailed introduction of consistency level, " +
					"please refer to consistency level. " +
					"Value range:\nEventual: eventual consistency.\n" +
					"Session: session consistency.\nGlobal: global consistency.\n" +
					"Description\nWhen the value of ReadWriteMode is ReadWrite, " +
					"the selectable consistency levels are Eventual, Session (default), and Global." +
					"\nWhen the value of ReadWriteMode is ReadOnly, the consistency level is Eventual by default and cannot be changed.",
			},
			"consist_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					//当 ConsistLevel 取值为 Global 或 Session 时，该参数才生效。
					if d.Get("consist_level").(string) != "Global" && d.Get("consist_level").(string) != "Session" {
						return true
					}
					return false
				},
				Description: "When there is a large delay, " +
					"the timeout period for read-only nodes to synchronize the latest data, in us. " +
					"The value range is from 1us to 100000000us, and the default value is 10000us.\n" +
					"Explanation\n This parameter takes effect only when the value of ConsistLevel is Global or Session.",
			},
			"consist_timeout_action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					//当 ConsistLevel 取值为 Global 或 Session 时，该参数才生效。
					if d.Get("consist_level").(string) != "Global" && d.Get("consist_level").(string) != "Session" {
						return true
					}
					return false
				},
				Description: "Timeout policy after data synchronization timeout of read-only nodes supports the following two policies:" +
					"\nReturnError: Return SQL error (wait replication complete timeout, please retry)." +
					"\nReadMaster: Send a request to the master node (default).\n" +
					"Description\n This parameter takes effect only when the value of ConsistLevel is Global or Session.",
			},
		},
	}
	return resource
}

func resourceVolcengineVedbMysqlEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVedbMysqlEndpoint())
	if err != nil {
		return fmt.Errorf("error on creating vedb_mysql_endpoint %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlEndpointRead(d, meta)
}

func resourceVolcengineVedbMysqlEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVedbMysqlEndpoint())
	if err != nil {
		return fmt.Errorf("error on reading vedb_mysql_endpoint %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVedbMysqlEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVedbMysqlEndpoint())
	if err != nil {
		return fmt.Errorf("error on updating vedb_mysql_endpoint %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlEndpointRead(d, meta)
}

func resourceVolcengineVedbMysqlEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVedbMysqlEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting vedb_mysql_endpoint %q, %s", d.Id(), err)
	}
	return err
}

var vedbMysqlEndpointImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("instance_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("endpoint_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
