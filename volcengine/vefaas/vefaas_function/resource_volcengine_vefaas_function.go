package vefaas_function

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VefaasFunction can be imported using the id, e.g.
```
$ terraform import volcengine_vefaas_function.default resource_id
```

*/

func ResourceVolcengineVefaasFunction() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVefaasFunctionCreate,
		Read:   resourceVolcengineVefaasFunctionRead,
		Update: resourceVolcengineVefaasFunctionUpdate,
		Delete: resourceVolcengineVefaasFunctionDelete,
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
				Description: "The name of Function.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of Function.",
			},
			"runtime": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The runtime of Function.",
			},
			"exclusive_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Exclusive mode switch.",
			},
			"request_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Request timeout (in seconds).",
			},
			"max_concurrency": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Maximum concurrency of a single instance.",
			},
			"memory_mb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory for a single instance.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Code Source type, supports tos, zip, image (whitelist accounts support native/v1 custom images).",
			},
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Code source.",
			},
			"envs": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Function environment variable.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of the environment variable.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Value of the environment variable.",
						},
					},
				},
			},
			"vpc_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The configuration of VPC.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of VPC.",
						},
						"enable_vpc": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether the function enables private network access.",
						},
						"subnet_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The ID of subnet.",
						},
						"security_group_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The ID of security group.",
						},
						"enable_shared_internet_access": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Function access to the public network switch.",
						},
					},
				},
			},
			"tls_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Function log configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_log": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "TLS log function switch.",
						},
						"tls_topic_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The topic ID of TLS log topic.",
						},
						"tls_project_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The project ID of TLS log topic.",
						},
					},
				},
			},
			"source_access_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Access configuration for the image repository.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "The image repository password.",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Mirror repository username.",
						},
					},
				},
			},
			"nas_storage": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The configuration of file storage NAS mount.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_nas": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to enable NAS storage mounting.",
						},
						"nas_configs": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The configuration of NAS.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gid": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: "User groups in the file system. " +
											"Customization is not supported yet. If this parameter is provided, " +
											"the parameter value is 1000 (consistent with the function run user gid).",
									},
									"uid": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: "Users in the file system do not support customization yet. " +
											"If this parameter is provided, " +
											"its value can only be 1000 (consistent with the function run user uid).",
									},
									"remote_path": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Remote directory of the file system.",
									},
									"file_system_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The ID of NAS file system.",
									},
									"mount_point_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The ID of NAS mount point.",
									},
									"local_mount_path": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The directory of Function local mount.",
									},
								},
							},
						},
					},
				},
			},
			"tos_mount_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The configuration of Object Storage TOS mount.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_tos": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to enable TOS storage mounting.",
						},
						"credentials": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "After enabling TOS, you need to provide an AKSK with access rights to the TOS domain name.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_key_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The AccessKey ID (AK) of the Volcano Engine account.",
									},
									"secret_access_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The Secret Access Key (SK) of the Volcano Engine account.",
									},
								},
							},
						},
						"mount_points": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "After enabling TOS, you need to provide a TOS storage configuration list, with a maximum of 5 items.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "TOS Access domain name.",
									},
									"read_only": {
										Type:     schema.TypeBool,
										Optional: true,
										Description: "Function local directory access permissions. " +
											"After mounting the TOS Bucket, whether the function local " +
											"mount directory has read-only permissions.",
									},
									"bucket_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "TOS bucket.",
									},
									"bucket_path": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The mounted TOS Bucket path.",
									},
									"local_mount_path": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Function local mount directory.",
									},
								},
							},
						},
					},
				},
			},
			"initializer_sec": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Function to initialize timeout configuration.",
			},
			"command": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The custom startup command for the instance.",
			},
			"cpu_strategy": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Function CPU charging policy.",
			},
			"code_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The size of code package.",
			},
			"code_size_limit": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum code package size.",
			},
			"source_location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Maximum code package size.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the function.",
			},
			"last_update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last update time of the function.",
			},
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The owner of Function.",
			},
			"triggers_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of triggers for this Function.",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Custom listening port for the instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineVefaasFunctionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasFunctionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVefaasFunction())
	if err != nil {
		return fmt.Errorf("error on creating vefaas_function %q, %s", d.Id(), err)
	}
	return resourceVolcengineVefaasFunctionRead(d, meta)
}

func resourceVolcengineVefaasFunctionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasFunctionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVefaasFunction())
	if err != nil {
		return fmt.Errorf("error on reading vefaas_function %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVefaasFunctionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasFunctionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVefaasFunction())
	if err != nil {
		return fmt.Errorf("error on updating vefaas_function %q, %s", d.Id(), err)
	}
	return resourceVolcengineVefaasFunctionRead(d, meta)
}

func resourceVolcengineVefaasFunctionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasFunctionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVefaasFunction())
	if err != nil {
		return fmt.Errorf("error on deleting vefaas_function %q, %s", d.Id(), err)
	}
	return err
}
