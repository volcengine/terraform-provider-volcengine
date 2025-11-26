package rds_postgresql_database_endpoint

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
RdsPostgresqlDatabaseEndpoint can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_database_endpoint.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlDatabaseEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlDatabaseEndpointCreate,
		Read:   resourceVolcengineRdsPostgresqlDatabaseEndpointRead,
		Update: resourceVolcengineRdsPostgresqlDatabaseEndpointUpdate,
		Delete: resourceVolcengineRdsPostgresqlDatabaseEndpointDelete,
		Importer: &schema.ResourceImporter{State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
			parts := strings.Split(d.Id(), ":")
			if len(parts) < 2 {
				return []*schema.ResourceData{d}, fmt.Errorf("import id must be 'instance_id:endpoint_id'")
			}
			_ = d.Set("instance_id", parts[0])
			return []*schema.ResourceData{d}, nil
		}},
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
				Description: "The ID of the RDS PostgreSQL instance.",
			},
			"endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "The ID of the connection endpoint. " +
					"The ID of the default endpoint is in the form of instance_id-cluster.",
			},
			"endpoint_name": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The name of the connection endpoint. " +
					"If not provided, the connection endpoint will be automatically named Custom Endpoint.",
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Custom", "Cluster"}, false),
				Description: "Type of the connection endpoint. Valid values: `Custom`(custom endpoint), `Cluster`(default endpoint). " +
					"When create a new endpoint, the value must be `Custom`. " +
					"The default cluster endpoint does not support creation; you can use import to bring it under Terraform management.",
			},
			"read_write_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ReadOnly", "ReadWrite"}, false),
				Description:  "ReadWrite or ReadOnly.",
			},
			"nodes": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "List of nodes configured for the connection endpoint. Required when EndpointType is Custom. " +
					"The primary node does not need to pass the node ID; it is sufficient to pass the string \"Primary\".",
			},
			// 注意：ModifyDBEndpointAddress 和 ModifyDBEndpointDNS  要求 NetworkType 必传，且值为 Private，在 service 层进行设置
			"domain_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Private address domain prefix to modify. " +
					"Do not set this field when creating a endpoint.",
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Private address port to modify. The value range is 1000~65534. " +
					"Do not set this field when creating a endpoint.",
			},
			"dns_visibility": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "Whether to enable public network resolution. " +
					"false: Default value, Volcano Engine private network resolution. " +
					"true: Volcano Engine private network and public network resolution. " +
					"Do not set this field when creating a endpoint.",
			},
			"global_read_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Whether to enable the global read-only mode for the instance. " +
					"There is no default value. If no value is passed, the request will be ignored. " +
					"Do not set this field when creating a endpoint.",
			},
			// 读写分离，只有默认终端支持该功能
			"read_write_splitting": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				// 要求endpoint_id 以 -cluster 结尾，在service层进行校验
				Description: "Whether to enable read-write separation. Only default endpoint supports this feature.",
			},
			"read_only_node_max_delay_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "The maximum delay threshold for read-only nodes. " +
					"When the delay time of a read-only node exceeds this value, read traffic will not be sent to that node. " +
					"The value range is 0~3600. Default value is 30 seconds.",
			},
			"read_only_node_distribution_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"Default", "Custom"}, false),
				Description: "Read-only weight distribution mode, Default or Custom. " +
					"Default: Standard weight allocation. Custom: Custom weight allocation.",
			},
			"read_only_node_weight": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A read-only node requires passing in the NodeId. A primary node does not need to pass in the NodeId.",
						},
						"node_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Node type. Primary or ReadOnly.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "Custom read weight allocation. Increases by 100, with a maximum value of 40000. Weights cannot all be set to 0.",
						},
					},
				},
				Description: "Custom read weight allocation. This parameter needs to be set when the value of read_only_node_distribution_type is Custom.",
			},
			"write_node_halt_writing": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "Whether to prohibit the terminal from sending write requests to the write node. " +
					"To avoid having no available connection endpoints to carry write operations, this configuration can only be enabled when the instance has other read-write endpoints.",
			},
			"read_write_proxy_connection": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: "The number of proxy connections set for the terminal after enabling read-write separation. " +
					"The minimum value of the proxy connection count is 20.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlDatabaseEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlDatabaseEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlDatabaseEndpoint())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_database_endpoint %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlDatabaseEndpointRead(d, meta)
}

func resourceVolcengineRdsPostgresqlDatabaseEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlDatabaseEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlDatabaseEndpoint())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_database_endpoint %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlDatabaseEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlDatabaseEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlDatabaseEndpoint())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_database_endpoint %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlDatabaseEndpointRead(d, meta)
}

func resourceVolcengineRdsPostgresqlDatabaseEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlDatabaseEndpointService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlDatabaseEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_database_endpoint %q, %s", d.Id(), err)
	}
	return err
}
