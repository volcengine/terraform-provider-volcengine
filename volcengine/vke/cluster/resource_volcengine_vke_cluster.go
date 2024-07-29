package cluster

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VkeCluster can be imported using the id, e.g.
```
$ terraform import volcengine_vke_cluster.default cc9l74mvqtofjnoj5****
```

*/

func ResourceVolcengineVkeCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineVkeClusterCreate,
		Read:   resourceVolcengineVkeClusterRead,
		Update: resourceVolcengineVkeClusterUpdate,
		Delete: resourceVolcengineVkeClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ClientToken is a case-sensitive string of no more than 64 ASCII characters passed in by the caller.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the cluster.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the cluster.",
			},
			"delete_protection_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The delete protection of the cluster, the value is `true` or `false`.",
			},
			"kubernetes_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if k == "kubernetes_version" && strings.Contains(old, new) {
						return true
					}
					return false
				},
				Description: "The version of Kubernetes specified when creating a VKE cluster (specified to patch version), with an example value of `1.24`. If not specified, the latest Kubernetes version supported by VKE is used by default, which is a 3-segment version format starting with a lowercase v, that is, KubernetesVersion with IsLatestVersion=True in the return value of ListSupportedVersions.",
			},
			"tags": ve.TagsSchema(),
			"cluster_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The config of the cluster.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ids": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 6,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
							Description: "The subnet ID for the cluster control plane to communicate within the private network.\n" +
								"Up to 3 subnets can be selected from each available zone, and a maximum of 2 subnets can be added to each available zone\n" +
								"Cannot support deleting configured subnets.",
						},
						"api_server_public_access_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Cluster API Server public network access configuration, the value is `true` or `false`.",
						},
						"api_server_public_access_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							Optional:         true,
							DiffSuppressFunc: ApiServerPublicAccessConfigFieldDiffSuppress,
							Description:      "Cluster API Server public network access configuration.",
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
													Type:         schema.TypeString,
													Optional:     true,
													Description:  "Billing type of public IP, the value is `PostPaidByBandwidth` or `PostPaidByTraffic`.",
													ValidateFunc: validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
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
						},
						"resource_public_access_default_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Node public network access configuration, the value is `true` or `false`.",
						},
						"ip_family": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "[SkipDoc]The IpFamily configuration,the value is `Ipv4` or `DualStack`.",
						},
					},
				},
			},
			"pods_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The config of the pods.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pod_network_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							Description:  "The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.",
							ValidateFunc: validation.StringInSlice([]string{"Flannel", "VpcCniShared"}, false),
						},
						"flannel_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							ForceNew:         true,
							Optional:         true,
							Description:      "Flannel network configuration.",
							DiffSuppressFunc: FlannelFieldDiffSuppress,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pod_cidrs": {
										Type:     schema.TypeSet,
										Optional: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "Pod CIDR for the Flannel container network.",
									},
									"max_pods_per_node": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Description: "The maximum number of single-node Pod instances for a Flannel container network, the value can be `16` or `32` or `64` or `128` or `256`.",
									},
								},
							},
						},
						"vpc_cni_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							Optional:         true,
							Description:      "VPC-CNI network configuration.",
							DiffSuppressFunc: VpcCniConfigFieldDiffSuppress,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											return true
										},
										Description: "The private network where the cluster control plane network resides.",
									},
									"subnet_ids": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "A list of Pod subnet IDs for the VPC-CNI container network.",
									},
								},
							},
						},
					},
				},
			},
			"services_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "The config of the services.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_cidrsv4": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The IPv4 private network address exposed by the service.",
						},
					},
				},
			},
			"kubeconfig_public": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kubeconfig data with public network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.",
			},
			"kubeconfig_private": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kubeconfig data with private network access, returned in BASE64 encoding, it is suggested to use vke_kubeconfig instead.",
			},
			"eip_allocation_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Eip allocation Id.",
			},
			"logging_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Cluster log configuration information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_project_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The TLS log item ID of the collection target.",
						},
						"log_setups": {
							Type:        schema.TypeSet,
							Optional:    true,
							Set:         logSetupsHash,
							Description: "Cluster logging options. This structure can only be modified and added, and cannot be deleted. When encountering a `cannot be deleted` error, please query the log setups of the current cluster and fill in the current `tf` file.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_type": {
										Type:     schema.TypeString,
										Required: true,
										Description: "The current types of logs that can be enabled are:\n" +
											"Audit: Cluster audit logs.\n" +
											"KubeApiServer: kube-apiserver component logs.\n" +
											"KubeScheduler: kube-scheduler component logs.\n" +
											"KubeControllerManager: kube-controller-manager component logs.",
									},
									"log_ttl": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      30,
										ValidateFunc: validation.IntBetween(1, 3650),
										Description:  "The storage time of logs in Log Service. After the specified log storage time is exceeded, the expired logs in this log topic will be automatically cleared. The unit is days, and the default is 30 days. The value range is 1 to 3650, specifying 3650 days means permanent storage.",
									},
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
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
}

func resourceVolcengineVkeClusterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(clusterService, d, ResourceVolcengineVkeCluster())
	if err != nil {
		return fmt.Errorf("error on creating cluster  %q, %w", d.Id(), err)
	}
	return resourceVolcengineVkeClusterRead(d, meta)
}

func resourceVolcengineVkeClusterRead(d *schema.ResourceData, meta interface{}) (err error) {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(clusterService, d, ResourceVolcengineVkeCluster())
	if err != nil {
		return fmt.Errorf("error on reading cluster %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineVkeClusterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(clusterService, d, ResourceVolcengineVkeCluster())
	if err != nil {
		return fmt.Errorf("error on updating cluster  %q, %w", d.Id(), err)
	}
	return resourceVolcengineVkeClusterRead(d, meta)
}

func resourceVolcengineVkeClusterDelete(d *schema.ResourceData, meta interface{}) (err error) {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(clusterService, d, ResourceVolcengineVkeCluster())
	if err != nil {
		return fmt.Errorf("error on deleting cluster %q, %w", d.Id(), err)
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
