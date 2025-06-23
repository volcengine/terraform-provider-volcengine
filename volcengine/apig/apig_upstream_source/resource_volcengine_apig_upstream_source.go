package apig_upstream_source

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ApigUpstreamSource can be imported using the id, e.g.
```
$ terraform import volcengine_apig_upstream_source.default resource_id
```

*/

func ResourceVolcengineApigUpstreamSource() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineApigUpstreamSourceCreate,
		Read:   resourceVolcengineApigUpstreamSourceRead,
		Update: resourceVolcengineApigUpstreamSourceUpdate,
		Delete: resourceVolcengineApigUpstreamSourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The gateway id of the apig upstream source.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comments of the apig upstream source.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The source type of the apig upstream. Valid values: `K8S`, `Nacos`.",
			},
			"source_spec": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The source spec of apig upstream source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"k8s_source": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "The k8s source of apig upstream source.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_id": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "The cluster id of k8s source.",
									},
									"cluster_type": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Description: "The cluster type of k8s source.",
									},
								},
							},
						},
						"nacos_source": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "The nacos source of apig upstream source.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nacos_id": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "The nacos id of nacos source.",
									},
									"nacos_name": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Description: "The nacos name of nacos source.",
									},
									"address": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Description: "The address of nacos source.",
									},
									"http_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Description: "The http port of nacos source.",
									},
									"grpc_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Description: "The grpc port of nacos source.",
									},
									"context_path": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Description: "The context path of nacos source.",
									},
									"auth_config": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										MaxItems:    1,
										Description: "The auth config of nacos source.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"basic": {
													Type:        schema.TypeList,
													Optional:    true,
													ForceNew:    true,
													MaxItems:    1,
													Description: "The basic auth config of nacos source.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"username": {
																Type:        schema.TypeString,
																Required:    true,
																ForceNew:    true,
																Description: "The username of basic auth config of nacos source.",
															},
															"password": {
																Type:        schema.TypeString,
																Required:    true,
																ForceNew:    true,
																Description: "The password of basic auth config of nacos source.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ingress_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The ingress settings of apig upstream source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_ingress": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether to enable ingress.",
						},
						"enable_all_ingress_classes": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether to enable all ingress classes.",
						},
						"enable_ingress_without_ingress_class": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether to enable ingress without ingress class.",
						},
						"enable_all_namespaces": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether to enable all namespaces.",
						},
						"update_status": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "The update status of ingress settings.",
						},
						"ingress_classes": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The ingress classes of ingress settings.",
						},
						"watch_namespaces": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The watch namespaces of ingress settings.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineApigUpstreamSourceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamSourceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineApigUpstreamSource())
	if err != nil {
		return fmt.Errorf("error on creating apig_upstream_source %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigUpstreamSourceRead(d, meta)
}

func resourceVolcengineApigUpstreamSourceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamSourceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineApigUpstreamSource())
	if err != nil {
		return fmt.Errorf("error on reading apig_upstream_source %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineApigUpstreamSourceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamSourceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineApigUpstreamSource())
	if err != nil {
		return fmt.Errorf("error on updating apig_upstream_source %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigUpstreamSourceRead(d, meta)
}

func resourceVolcengineApigUpstreamSourceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamSourceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineApigUpstreamSource())
	if err != nil {
		return fmt.Errorf("error on deleting apig_upstream_source %q, %s", d.Id(), err)
	}
	return err
}
