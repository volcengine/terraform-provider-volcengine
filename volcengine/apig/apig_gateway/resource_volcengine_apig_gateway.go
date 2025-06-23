package apig_gateway

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ApigGateway can be imported using the id, e.g.
```
$ terraform import volcengine_apig_gateway.default resource_id
```

*/

func ResourceVolcengineApigGateway() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineApigGatewayCreate,
		Read:   resourceVolcengineApigGatewayRead,
		Update: resourceVolcengineApigGatewayUpdate,
		Delete: resourceVolcengineApigGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the api gateway.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comments of the api gateway.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The type of the api gateway. Valid values: `standard`, `serverless`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The project name of the api gateway.",
			},
			"tags": ve.TagsSchema(),
			"network_spec": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The network spec of the api gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The vpc id of the network spec.",
						},
						"subnet_ids": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The subnet ids of the network spec.",
						},
					},
				},
			},
			"backend_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The backend spec of the api gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_vke_with_flannel_cni_supported": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Whether the api gateway support vke flannel cni.",
						},
						"vke_pod_cidr": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The vke pod cidr of the api gateway.",
						},
					},
				},
			},
			"monitor_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The monitor spec of the api gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Whether the api gateway enable monitor.",
						},
						"workspace_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return !d.Get("monitor_spec.0.enable").(bool)
							},
							Description: "The workspace id of the monitor. This field is required when `enable` is true.",
						},
					},
				},
			},
			"log_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The log spec of the api gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether the api gateway enable tls log.",
						},
						"project_id": {
							Type:     schema.TypeString,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return !d.Get("log_spec.0.enable").(bool)
							},
							Description: "The project id of the tls. This field is required when `enable` is true.",
						},
						"topic_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return !d.Get("log_spec.0.enable").(bool)
							},
							Description: "The topic id of the tls.",
						},
					},
				},
			},
			"resource_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The resource spec of the api gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"replicas": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The replicas of the resource spec.",
						},
						"instance_spec_code": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The instance spec code of the resource spec. Valid values: `1c2g`, `2c4g`, `4c8g`, `8c16g`.",
						},
						"clb_spec_code": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "The clb spec code of the resource spec. Valid values: `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.",
						},
						"public_network_billing_type": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "The public network billing type of the resource spec. Valid values: `traffic`, `bandwidth`.",
						},
						"public_network_bandwidth": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "The public network bandwidth of the resource spec.",
						},
						"network_type": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "The network type of the resource spec. The default values for both `enable_public_network` and `enable_private_network` are true.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_public_network": {
										Type:        schema.TypeBool,
										Required:    true,
										ForceNew:    true,
										Description: "Whether the api gateway enable public network.",
									},
									"enable_private_network": {
										Type:        schema.TypeBool,
										Required:    true,
										ForceNew:    true,
										Description: "Whether the api gateway enable private network.",
									},
								},
							},
						},
					},
				},
			},

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the api gateway.",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The error message of the api gateway.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the api gateway.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the api gateway.",
			},
		},
	}
	return resource
}

func resourceVolcengineApigGatewayCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigGatewayService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineApigGateway())
	if err != nil {
		return fmt.Errorf("error on creating apig_gateway %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigGatewayRead(d, meta)
}

func resourceVolcengineApigGatewayRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigGatewayService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineApigGateway())
	if err != nil {
		return fmt.Errorf("error on reading apig_gateway %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineApigGatewayUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigGatewayService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineApigGateway())
	if err != nil {
		return fmt.Errorf("error on updating apig_gateway %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigGatewayRead(d, meta)
}

func resourceVolcengineApigGatewayDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigGatewayService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineApigGateway())
	if err != nil {
		return fmt.Errorf("error on deleting apig_gateway %q, %s", d.Id(), err)
	}
	return err
}
