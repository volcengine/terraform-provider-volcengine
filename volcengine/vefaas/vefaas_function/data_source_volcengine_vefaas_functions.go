package vefaas_function

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVefaasFunctions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVefaasFunctionsRead,
		Schema: map[string]*schema.Schema{
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
			"items": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Function.",
						},
						"envs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Function environment variable.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Key of the environment variable.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Value of the environment variable.",
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of Function.",
						},
						"tags": ve.TagsSchemaComputed(),
						"owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner of Function.",
						},
						"runtime": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The runtime of Function.",
						},
						"code_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of code package.",
						},
						"memory_mb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum memory for a single instance.",
						},
						"tls_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Function log configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_log": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "TLS log function switch.",
									},
									"tls_topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic ID of TLS log topic.",
									},
									"tls_project_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The project ID of TLS log topic.",
									},
								},
							},
						},
						"vpc_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration of VPC.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of VPC.",
									},
									"enable_vpc": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the function enables private network access.",
									},
									"subnet_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The ID of subnet.",
									},
									"security_group_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The ID of security group.",
									},
									"enable_shared_internet_access": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Function access to the public network switch.",
									},
								},
							},
						},
						"nas_storage": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration of file storage NAS mount.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_nas": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable NAS storage mounting.",
									},
									"nas_configs": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The configuration of NAS.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"gid": {
													Type:     schema.TypeInt,
													Computed: true,
													Description: "User groups in the file system. " +
														"Customization is not supported yet. If this parameter is provided, " +
														"the parameter value is 1000 (consistent with the function run user gid).",
												},
												"uid": {
													Type:     schema.TypeInt,
													Computed: true,
													Description: "Users in the file system do not support customization yet. " +
														"If this parameter is provided, " +
														"its value can only be 1000 (consistent with the function run user uid).",
												},
												"remote_path": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Remote directory of the file system.",
												},
												"file_system_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The ID of NAS file system.",
												},
												"mount_point_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The ID of NAS mount point.",
												},
												"local_mount_path": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The directory of Function local mount.",
												},
											},
										},
									},
								},
							},
						},
						"source_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Code Source type, supports tos, zip, image (whitelist accounts support native/v1 custom images).",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of Function.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance type of the function instance.",
						},
						"code_size_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum code package size.",
						},
						"exclusive_mode": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Exclusive mode switch.",
						},
						"triggers_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of triggers for this Function.",
						},
						"initializer_sec": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Function to initialize timeout configuration.",
						},
						"last_update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
						},
						"max_concurrency": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum concurrency of a single instance.",
						},
						"request_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Request timeout (in seconds).",
						},
						"source_location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source address of the code/image.",
						},
						"tos_mount_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration of Object Storage TOS mount.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_tos": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable TOS storage mounting.",
									},
									"credentials": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "After enabling TOS, you need to provide an AKSK with access rights to the TOS domain name.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"access_key_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The AccessKey ID (AK) of the Volcano Engine account.",
												},
												"secret_access_key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Secret Access Key (SK) of the Volcano Engine account.",
												},
											},
										},
									},
									"mount_points": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "After enabling TOS, you need to provide a TOS storage configuration list, with a maximum of 5 items.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"endpoint": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "TOS Access domain name.",
												},
												"read_only": {
													Type:     schema.TypeBool,
													Computed: true,
													Description: "Function local directory access permissions. " +
														"After mounting the TOS Bucket, whether the function local " +
														"mount directory has read-only permissions.",
												},
												"bucket_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "TOS bucket.",
												},
												"bucket_path": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The mounted TOS Bucket path.",
												},
												"local_mount_path": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Function local mount directory.",
												},
											},
										},
									},
								},
							},
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Custom listening port for the instance.",
						},
						"command": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The custom startup command for the instance.",
						},
						"cpu_strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Function CPU charging policy.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVefaasFunctionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVefaasFunctionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVefaasFunctions())
}
