package veecp_cluster

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpCluster can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_cluster.default resource_id
```

*/

func ResourceVolcengineVeecpCluster() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpClusterCreate,
		Read:   resourceVolcengineVeecpClusterRead,
		Update: resourceVolcengineVeecpClusterUpdate,
		Delete: resourceVolcengineVeecpClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ClientToken is a case-sensitive string of no more than 64 ASCII characters passed in by the caller.",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "Cluster name. " +
					"Under the same region, the name must be unique. " +
					"Supports upper and lower case English letters, Chinese characters, numbers, and hyphens (-). " +
					"Numbers cannot be at the first position, and hyphens (-) cannot be at the first or last position. " +
					"The length is limited to 2 to 64 characters.",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				//Computed: true,
				ForceNew:    true,
				Description: "Cluster description. Length is limited to within 300 characters.",
			},
			"profile": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Edge cluster: Edge. Non-edge cluster: Cloud. When using edge hosting, set this item to Edge.",
			},
			"delete_protection_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Description: "Cluster deletion protection. " +
					"Values: false: (default value) Deletion protection is off. " +
					"true: Enable deletion protection. The cluster cannot be directly deleted. " +
					"After creating a cluster, when calling Delete edge cluster, " +
					"configure the Force parameter and choose to forcibly delete the cluster.",
			},
			"edge_tunnel_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "Whether to enable the edge tunnel. " +
					"Values: false: (default value) Edge tunnel is off. " +
					"true: Enable edge tunnel. " +
					"Note: This parameter is not supported to be modified after the cluster is created.",
			},
			"kubernetes_version": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
				Description: "Specify the Kubernetes version when creating a cluster." +
					" The format is x.xx. The default value is the latest version in the supported Kubernetes version list (currently 1.20).",
			},
			"cluster_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "Network configuration of cluster control plane and nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ids": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The subnet ID for communication within the private network (VPC) of the cluster control plane. " +
								"You can call the private network API to obtain the subnet ID. " +
								"Note: When creating a cluster, please ensure that all specified SubnetIds (including but not limited to this parameter) belong to the same private network." +
								" It is recommended that you choose subnets in different availability zones as much as possible to improve the high availability of the cluster control plane." +
								" Please note that this parameter is not supported to be modified after the cluster is created. " +
								"Please configure it reasonably.",
						},
						"api_server_public_access_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Cluster API Server public network access configuration, values:\nfalse: (default value). closed\ntrue: opened.",
						},
						"api_server_public_access_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: ApiServerPublicAccessConfigFieldDiffSuppress,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"public_access_network_config": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										ForceNew:    true,
										Description: "Public network access network configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"billing_type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Billing type of public IP, the value is `PostPaidByBandwidth` or `PostPaidByTraffic`.",
												},
												"bandwidth": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The peak bandwidth of the public IP, unit: Mbps.",
												},
											},
										},
									},
								},
							},
							Description: "Cluster API Server public network access configuration information. It takes effect only when ApiServerPublicAccessEnabled=true.",
						},
						"resource_public_access_default_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
							Description: "Node public network access configuration, " +
								"values:\nfalse: (default value). Do not enable public network access." +
								" Existing NAT gateways and rules are not affected. " +
								"true: Enable public network access. After enabling, " +
								"a NAT gateway is automatically created for the cluster's private network and corresponding rules are configured." +
								" Note: This parameter cannot be modified after the cluster is created. " +
								"Please configure it reasonably.",
						},
					},
				},
			},
			"pods_config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "Container (Pod) network configuration of the cluster.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pod_network_mode": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: "Container network model, values: " +
								"Flannel: Flannel network model, an independent Underlay container network solution. " +
								"Combined with the global routing capability of a private network (VPC), " +
								"it realizes a high-performance network experience for the cluster. " +
								"VpcCniShared: VPC-CNI network model, " +
								"an Underlay container network solution implemented based on the elastic network interface (ENI) of a private network," +
								" with high network communication performance. " +
								"Description: After the cluster is created, this parameter is not supported to be modified temporarily." +
								" Please configure it reasonably.",
						},
						"flannel_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							ForceNew:         true,
							Optional:         true,
							DiffSuppressFunc: FlannelFieldDiffSuppress,
							Description: "Flannel network configuration." +
								" It can be configured only when PodNetworkMode=Flannel, but it is not mandatory.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pod_cidrs": {
										Type:     schema.TypeSet,
										Required: true,
										ForceNew: true,
										Set:      schema.HashString,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "Pod CIDR of Flannel model container network." +
											" Only configurable when PodNetworkMode=Flannel, but not mandatory." +
											" Note: The number of Pods in the cluster is limited by the number of IPs in this CIDR." +
											" This parameter cannot be modified after cluster creation. Please plan the Pod CIDR reasonably." +
											" Cannot conflict with the following network segments: private network network segments corresponding to ClusterConfig.SubnetIds." +
											" All clusters within the same private network's FlannelConfig.PodCidrs. " +
											"All clusters within the same private network's ServiceConfig.ServiceCidrsv4." +
											" Different clusters within the same private network's FlannelConfig.PodCidrs cannot conflict.",
									},
									"max_pods_per_node": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Description: "Upper limit of the number of single-node Pod instances in the Flannel model container network. " +
											"Values: 64(default value), 16, 32, 128, 256.",
									},
								},
							},
						},
						"vpc_cni_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							ForceNew:         true,
							Optional:         true,
							DiffSuppressFunc: VpcCniConfigFieldDiffSuppress,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subnet_ids": {
										Type:     schema.TypeSet,
										Required: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "A list of Pod subnet IDs for the VPC-CNI container network.",
									},
								},
							},
							Description: "VPC-CNI network configuration. PodNetworkMode=VpcCniShared, but it is not mandatory.",
						},
					},
				},
			},
			"services_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster service (Service) network configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_cidrsv4": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "CIDR used by services within the cluster. " +
								"It cannot conflict with the following network segments: " +
								"FlannelConfig.PodCidrs. " +
								"SubnetIds of all clusters within the same private network or FlannelConfig.VpcConfig.SubnetIds. " +
								"ServiceConfig.ServiceCidrsv4 of all clusters within the same private network (this parameter)." +
								"It is stated that currently only one array element is supported. " +
								"When multiple values are specified, only the first value takes effect.",
						},
					},
				},
			},
			//"tags": ve.TagsSchema(),
			"logging_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Cluster log configuration information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_project_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The TLS log item ID of the collection target.",
						},
						"log_setups": {
							Type:        schema.TypeSet,
							Optional:    true,
							ForceNew:    true,
							Set:         logSetupsHash,
							Description: "Cluster logging options. This structure can only be modified and added, and cannot be deleted. When encountering a `cannot be deleted` error, please query the log setups of the current cluster and fill in the current `tf` file.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										Description: "The current types of logs that can be enabled are:\n" +
											"Audit: Cluster audit logs.\n" +
											"KubeApiServer: kube-apiserver component logs.\n" +
											"KubeScheduler: kube-scheduler component logs.\n" +
											"KubeControllerManager: kube-controller-manager component logs.",
									},
									"log_ttl": {
										Type:         schema.TypeInt,
										Optional:     true,
										ForceNew:     true,
										Default:      30,
										ValidateFunc: validation.IntBetween(1, 3650),
										Description:  "The storage time of logs in Log Service. After the specified log storage time is exceeded, the expired logs in this log topic will be automatically cleared. The unit is days, and the default is 30 days. The value range is 1 to 3650, specifying 3650 days means permanent storage.",
									},
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										ForceNew:    true,
										Default:     false,
										Description: "Whether to enable the log option, true means enable, false means not enable, the default is false. When Enabled is changed from false to true, a new Topic will be created.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVeecpClusterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpClusterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpCluster())
	if err != nil {
		return fmt.Errorf("error on creating veecp_cluster %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpClusterRead(d, meta)
}

func resourceVolcengineVeecpClusterRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpClusterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpCluster())
	if err != nil {
		return fmt.Errorf("error on reading veecp_cluster %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpClusterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpClusterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVeecpCluster())
	if err != nil {
		return fmt.Errorf("error on updating veecp_cluster %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpClusterRead(d, meta)
}

func resourceVolcengineVeecpClusterDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpClusterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpCluster())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_cluster %q, %s", d.Id(), err)
	}
	return err
}

func logSetupsHash(i interface{}) int {
	if i == nil {
		return hashcode.String("")
	}
	m := i.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v#%v", m["log_type"], m["log_ttl"], m["enabled"]))
	return hashcode.String(buf.String())
}

func ApiServerPublicAccessConfigFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	apiServerPublicAccessEnabled := d.Get("cluster_config").([]interface{})[0].(map[string]interface{})["api_server_public_access_enabled"].(bool)
	return !apiServerPublicAccessEnabled
}

func FlannelFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	podNetworkMode := d.Get("pods_config").([]interface{})[0].(map[string]interface{})["pod_network_mode"].(string)
	return podNetworkMode != "Flannel"
}

func VpcCniConfigFieldDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	podNetworkMode := d.Get("pods_config").([]interface{})[0].(map[string]interface{})["pod_network_mode"].(string)
	return podNetworkMode != "VpcCniShared"
}
