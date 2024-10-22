package vedb_mysql_endpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVedbMysqlEndpoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVedbMysqlEndpointsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the instance.",
			},
			"endpoint_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the endpoint.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"endpoints": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the endpoint.",
						},
						"endpoint_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the endpoint.",
						},
						"endpoint_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Connect terminal type. The value is fixed as Custom, indicating a custom terminal.",
						},
						"read_write_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Endpoint read-write mode. Values:\n ReadWrite: Read and write terminal.\n ReadOnly: Read-only terminal (default).",
						},
						"endpoint_name": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Connect the endpoint name. The setting rules are as follows:\n " +
								"It cannot start with a number or a hyphen (-).\n " +
								"It can only contain Chinese characters, letters, numbers, underscores (_), and hyphens (-).\n " +
								"The length is 1 to 64 characters.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description information for connecting endpoint. The length cannot exceed 200 characters.",
						},
						"node_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Connect the node IDs associated with the endpoint." +
								"The filling rules are as follows:\n" +
								"When the value of ReadWriteMode is ReadWrite, at least two nodes must be passed in, and the master node must be passed in.\n" +
								"When the value of ReadWriteMode is ReadOnly, one or more read-only nodes can be passed in.",
						},
						"auto_add_new_nodes": {
							Type:     schema.TypeBool,
							Computed: true,
							Description: "Set whether newly created read-only nodes will automatically join this connection endpoint." +
								" Values:\ntrue: Automatically join.\nfalse: Do not automatically join (default).",
						},
						"master_accept_read_requests": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
							Description: "Set whether to enable transaction splitting. " +
								"For detailed introduction to transaction splitting, please refer to transaction splitting." +
								" Value range:\ntrue: Enabled (default).\nfalse: Disabled.\n" +
								"Description\nOnly when the value of ReadWriteMode is ReadWrite, is enabling transaction splitting supported.",
						},
						"consist_level": {
							Type:     schema.TypeString,
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
							Computed: true,
							Description: "When there is a large delay, " +
								"the timeout period for read-only nodes to synchronize the latest data, in us. " +
								"The value range is from 1us to 100000000us, and the default value is 10000us.\n" +
								"Explanation\n This parameter takes effect only when the value of ConsistLevel is Global or Session.",
						},
						"consist_timeout_action": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Timeout policy after data synchronization timeout of read-only nodes supports the following two policies:" +
								"\nReturnError: Return SQL error (wait replication complete timeout, please retry)." +
								"\nReadMaster: Send a request to the master node (default).\n" +
								"Description\n This parameter takes effect only when the value of ConsistLevel is Global or Session.",
						},
						"addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The address information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_visibility": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Parsing method. Currently, the return value can only be false (Volcengine private network parsing).",
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Instance intranet access domain name.",
									},
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP address.",
									},
									"network_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network type:\nPrivate: Private network VPC.\nPublic: Public network access.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Instance intranet access port.",
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Subnet ID. " +
											"The subnet must belong to the selected availability zone.\n" +
											"Description\n " +
											"A subnet is an IP address block within a private network. " +
											"All cloud resources in a private network must be deployed within a subnet. " +
											"The subnet assigns private IP addresses to cloud resources.",
									},
									"eip_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The EIP id.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVedbMysqlEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVedbMysqlEndpointService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVedbMysqlEndpoints())
}
