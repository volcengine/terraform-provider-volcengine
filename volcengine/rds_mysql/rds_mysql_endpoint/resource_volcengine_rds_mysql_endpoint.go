package rds_mysql_endpoint

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsMysqlEndpoint can be imported using the instance id and endpoint id, e.g.
```
$ terraform import volcengine_rds_mysql_endpoint.default mysql-3c25f219***:mysql-3c25f219****-custom-eeb5
```

*/

func ResourceVolcengineRdsMysqlEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsMysqlEndpointCreate,
		Read:   resourceVolcengineRdsMysqlEndpointRead,
		Update: resourceVolcengineRdsMysqlEndpointUpdate,
		Delete: resourceVolcengineRdsMysqlEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: endpointImporter,
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
				Description: "The id of the mysql instance.",
			},
			"endpoint_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The id of the endpoint. Import an exist endpoint, usually for import a default endpoint generated with instance creating.",
			},
			"read_write_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ReadOnly",
				Description: "Reading and writing mode: ReadWrite, ReadOnly(Default).",
			},
			"endpoint_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the endpoint.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the endpoint.",
			},
			"nodes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:      schema.HashString,
				Required: true,
				Description: "List of node IDs configured for the endpoint. Required when EndpointType is Custom. " +
					"To add a master node to the terminal, there is no need to fill in the master node ID, just fill in `Primary`.",
			},
			"auto_add_new_nodes": {
				Type:     schema.TypeBool,
				Computed: false,
				Optional: true,
				Description: "When the terminal type is a read-write terminal or a read-only terminal, " +
					"support is provided for setting whether new nodes are automatically added." +
					" The values are:\ntrue: Automatically add.\nfalse: Do not automatically add (default).",
			},
			"read_write_spliting": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable read-write splitting. Values: true: Yes. Default value. false: No.",
			},
			"read_only_node_max_delay_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					//读写类型的终端，且开通读写分离后支持设置此参数。
					return d.Get("read_write_mode").(string) == "ReadOnly" ||
						!d.Get("read_write_spliting").(bool)
				},
				Description: "The maximum delay threshold for read-only nodes, when the delay time of a read-only node exceeds this value, " +
					"the read traffic will not be sent to that node, unit: seconds. " +
					"Value range: 0~3600. Default value: 30.",
			},
			"read_only_node_distribution_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !d.Get("read_write_spliting").(bool)
				},
				Description: "Read weight distribution mode. " +
					"This parameter needs to be passed in when the read-write separation setting is true. " +
					"When used as a request parameter in the CreateDBEndpoint and ModifyDBEndpoint interfaces, the value range is as follows: " +
					"LoadSchedule: Load scheduling. " +
					"RoundRobinCustom: Polling scheduling with custom weights. " +
					"RoundRobinAuto: Polling scheduling with automatically allocated weights.",
			},
			"read_only_node_weight": {
				Type: schema.TypeSet,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					//当 ReadOnlyNodeDistributionType 取值为 Custom 时，需要传入此参数。
					return d.Get("read_only_node_distribution_type").(string) != "RoundRobinCustom"
				},
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Read-only nodes require NodeId to be passed, while primary nodes do not require it.",
						},
						"node_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The primary node needs to pass in the NodeType as Primary, while the read-only node does not need to pass it in.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The read weight of the node increases by 100, with a maximum value of 10000.",
						},
					},
				},
				Optional: true,
				Description: "Customize read weight distribution, that is, pass in the read request weight of the master node and read-only nodes. " +
					"It increases by 100 and the maximum value is 10000. " +
					"When the ReadOnlyNodeDistributionType value is Custom, " +
					"this parameter needs to be passed in.",
			},
			"dns_visibility": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Values:\nfalse: Volcano Engine private network resolution (default).\ntrue: Volcano Engine private and public network resolution.",
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Connection address, Please note that the connection address can only modify the prefix." +
					" In one call, it is not possible to modify both the connection address prefix and the port at the same time.",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The port. Cannot modify public network port. In one call, it is not possible to modify both the connection address prefix and the port at the same time.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsMysqlEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMysqlEndpoint())
	if err != nil {
		return fmt.Errorf("error on creating rds_mysql_endpoint %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlEndpointRead(d, meta)
}

func resourceVolcengineRdsMysqlEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMysqlEndpoint())
	if err != nil {
		return fmt.Errorf("error on reading rds_mysql_endpoint %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsMysqlEndpoint())
	if err != nil {
		return fmt.Errorf("error on updating rds_mysql_endpoint %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlEndpointRead(d, meta)
}

func resourceVolcengineRdsMysqlEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsMysqlEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting rds_mysql_endpoint %q, %s", d.Id(), err)
	}
	return err
}

func endpointImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'instanceId:endpointId'")
	}
	instanceId := items[0]
	endpointId := items[1]
	_ = d.Set("instance_id", instanceId)
	_ = d.Set("endpoint_id", endpointId)

	return []*schema.ResourceData{d}, nil
}
