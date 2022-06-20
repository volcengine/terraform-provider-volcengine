package cluster

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
VkeCluster can be imported using the id, e.g.
```
$ terraform import vestack_vke_cluster.default cc9l74mvqtofjnoj5****
```

*/

func ResourceVestackVkeCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackVkeClusterCreate,
		Read:   resourceVestackVkeClusterRead,
		Update: resourceVestackVkeClusterUpdate,
		Delete: resourceVestackVkeClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Description: "The delete protection of the cluster.",
			},
			"kubernetes_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The version of Kubernetes specified when creating a VKE cluster (specified to patch version), if not specified, the latest Kubernetes version supported by VKE is used by default, which is a 3-segment version format starting with a lowercase v, that is, KubernetesVersion with IsLatestVersion=True in the return value of ListSupportedVersions.",
			},
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
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The list of Subnet IDs.",
						},
						"api_server_public_access_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Cluster API Server public network access configuration.",
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
													Description:  "Billing type of public IP.",
													ValidateFunc: validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
												},
												"bandwidth": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Peak bandwidth of public IP.",
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
							Description: "Node public network access configuration.",
						},
					},
				},
			},
			"pods_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "The config of the pods.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pod_network_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							Description:  "Container Pod Network Type (CNI).",
							ValidateFunc: validation.StringInSlice([]string{"Flannel", "VpcCniShared"}, false),
						},
						"flannel_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							Optional:         true,
							ForceNew:         true,
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
										Description: "Container Pod Network CIDR.",
									},
									"max_pods_per_node": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Description: "Maximum number of Pod instances on a single node.",
									},
								},
							},
						},
						"vpc_cni_config": {
							Type:             schema.TypeList,
							MaxItems:         1,
							Optional:         true,
							ForceNew:         true,
							Description:      "VPC-CNI network configuration.",
							DiffSuppressFunc: VpcCniConfigFieldDiffSuppress,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Maximum number of Pod instances on a single node.",
									},
									"subnet_ids": {
										Type:     schema.TypeSet,
										Optional: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set:         schema.HashString,
										Description: "List of subnets corresponding to the container Pod network.",
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
				Description: "Kubeconfig data with public network access, returned in BASE64 encoding.",
			},
			"kubeconfig_private": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kubeconfig data with private network access, returned in BASE64 encoding.",
			},
			"eip_allocation_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Eip allocation Id.",
			},
		},
	}
}

func resourceVestackVkeClusterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	err = clusterService.Dispatcher.Create(clusterService, d, ResourceVestackVkeCluster())
	if err != nil {
		return fmt.Errorf("error on creating cluster  %q, %w", d.Id(), err)
	}
	return resourceVestackVkeClusterRead(d, meta)
}

func resourceVestackVkeClusterRead(d *schema.ResourceData, meta interface{}) (err error) {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	err = clusterService.Dispatcher.Read(clusterService, d, ResourceVestackVkeCluster())
	if err != nil {
		return fmt.Errorf("error on reading cluster %q, %w", d.Id(), err)
	}
	return err
}

func resourceVestackVkeClusterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	err = clusterService.Dispatcher.Update(clusterService, d, ResourceVestackVkeCluster())
	if err != nil {
		return fmt.Errorf("error on updating cluster  %q, %w", d.Id(), err)
	}
	return resourceVestackVkeClusterRead(d, meta)
}

func resourceVestackVkeClusterDelete(d *schema.ResourceData, meta interface{}) (err error) {
	clusterService := NewVkeClusterService(meta.(*ve.SdkClient))
	err = clusterService.Dispatcher.Delete(clusterService, d, ResourceVestackVkeCluster())
	if err != nil {
		return fmt.Errorf("error on deleting cluster %q, %w", d.Id(), err)
	}
	return err
}
