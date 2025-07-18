package ecs_instance

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
ECS Instance can be imported using the id, e.g.
If Import,The data_volumes is sort by volume name
```
$ terraform import volcengine_ecs_instance.default i-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineEcsInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsInstanceCreate,
		Read:   resourceVolcengineEcsInstanceRead,
		Update: resourceVolcengineEcsInstanceUpdate,
		Delete: resourceVolcengineEcsInstanceDelete,
		Exists: resourceVolcengineEcsInstanceExist,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The available zone ID of ECS instance.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Image ID of ECS instance.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance type of ECS instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of ECS instance.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of ECS instance.",
			},
			"host_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The host name of ECS instance.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password of ECS instance.",
			},
			"key_pair_name": {
				Type:     schema.TypeString,
				Optional: true,
				//ForceNew: true,
				//Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !d.HasChange("image_id")
				},
				Description: "The ssh key name of ECS instance. This field can be modified only when the `image_id` is modified.",
			},
			"keep_image_credential": {
				Type:     schema.TypeBool,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !d.HasChange("image_id")
				},
				Description: "Whether to keep the mirror settings. Only custom images and shared images support this field.\n When the value of this field is true, the Password and KeyPairName cannot be specified.\n When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"instance_charge_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"PostPaid",
					"PrePaid",
				}, false),
				Description: "The charge type of ECS instance, the value can be `PrePaid` or `PostPaid`.",
			},
			"spot_strategy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"NoSpot",
					"SpotAsPriceGo",
					"SpotWithPriceLimit",
				}, false),
				Description: "The spot strategy will auto" +
					"remove instance in some conditions.Please make sure you can maintain instance lifecycle before " +
					"auto remove.The spot strategy of ECS instance, values:\n NoSpot (default): indicates creating a normal pay-as-you-go instance." +
					"\nSpotAsPriceGo: spot instance with system automatically bidding and following the current market price." +
					"\nSpotWithPriceLimit: spot instance with a set upper limit for bidding price.",
			},
			"spot_price_limit": {
				Type:     schema.TypeFloat,
				Optional: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("instance_charge_type").(string) == "PostPaid" && d.Get("spot_strategy").(string) == "SpotWithPriceLimit" {
						return false
					}
					return true
				},
				Description: "The maximum hourly price for spot instances supports up to three decimal places. " +
					"This parameter only takes effect when SpotStrategy=SpotWithPriceLimit.",
			},
			"user_data": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: UserDateImportDiffSuppress,
				Description:      "The user data of ECS instance, this field must be encrypted with base64.",
			},
			"security_enhancement_strategy": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Active",
					"InActive",
				}, false),
				Default:     "Active",
				Description: "The security enhancement strategy of ECS instance. The value can be Active or InActive. Default is Active.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"hpc_cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The hpc cluster ID of ECS instance.",
			},
			"period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          12,
				DiffSuppressFunc: EcsInstanceImportDiffSuppress,
				Description:      "The period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 12. Unit is Month.",
			},
			//"period_unit": {
			//	Type:     schema.TypeString,
			//	Optional: true,
			//	Default:  "Month",
			//	ValidateFunc: validation.StringInSlice([]string{
			//		"Month",
			//	}, false),
			//	DiffSuppressFunc: ve.EcsInstanceImportDiffSuppress,
			//	Description:      "The period unit of ECS instance.Only effective when instance_charge_type is PrePaid. Default is Month.",
			//},
			"auto_renew": {
				Type:     schema.TypeBool,
				Optional: true,
				//ForceNew: true,
				Default:          true,
				DiffSuppressFunc: AutoRenewDiffSuppress,
				Description:      "The auto renew flag of ECS instance.Only effective when instance_charge_type is PrePaid. Default is true.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"auto_renew_period": {
				Type:     schema.TypeInt,
				Optional: true,
				//ForceNew: true,
				Default:          1,
				DiffSuppressFunc: AutoRenewDiffSuppress,
				Description:      "The auto renew period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 1.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},

			"include_data_volumes": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				DiffSuppressFunc: EcsInstanceImportDiffSuppress,
				Description:      "The include data volumes flag of ECS instance.Only effective when change instance charge type.include_data_volumes.",
			},

			//"vpc_id": {
			//	Type:     schema.TypeString,
			//	Optional: true,
			//	Computed: true,
			//	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			//		return d.Id() == ""
			//	},
			//	Description: "The vpc ID of primary networkInterface. This field is only effective when modifying the instance.",
			//},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet ID of primary networkInterface.",
			},

			"security_group_ids": {
				Type:        schema.TypeSet,
				Required:    true,
				MaxItems:    5,
				MinItems:    1,
				Description: "The security group ID set of primary networkInterface.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},

			"network_interface_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of primary networkInterface.",
			},

			"primary_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				//DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				//	return !d.HasChange("subnet_id")
				//},
				Description: "The private ip address of primary networkInterface.",
			},

			"eip_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"eip_address"},
				Description: "The id of an existing Available EIP which will be automatically assigned to this instance. \n" +
					"It is not recommended to use this field, it is recommended to use `volcengine_eip_associate` resource to bind EIP.",
			},
			"eip_address": {
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				ConflictsWith: []string{"eip_id"},
				Description: "The config of the eip which will be automatically created and assigned to this instance. `Prepaid` type eip cannot be created in this way, please use `volcengine_eip_address`.\n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"charge_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "PayByBandwidth",
							ForceNew:    true,
							Description: "The billing type of the EIP Address. Valid values: `PayByBandwidth`, `PayByTraffic`. Default is `PayByBandwidth`.",
						},
						"isp": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "BGP",
							ForceNew:    true,
							Description: "The ISP of the EIP. Valid values: `BGP`, `ChinaMobile`, `ChinaUnicom`, `ChinaTelecom`, `SingleLine_BGP`, `Static_BGP`.",
						},
						"bandwidth_package_id": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The id of the bandwidth package, indicates that the public IP address will be added to the bandwidth package.",
						},
						"bandwidth_mbps": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
							ForceNew:    true,
							Description: "The peek bandwidth of the EIP. The value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic. Default is 1.",
						},
					},
				},
			},

			"system_volume_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of system volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.",
			},

			"system_volume_size": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "The size of system volume. " +
					"The value range of the system volume size is ESSD_PL0: 20~2048, ESSD_FlexPL: 20~2048, PTSSD: 10~500.",
			},

			"system_volume_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of system volume.",
			},

			"deployment_set_id": {
				Type:     schema.TypeString,
				Optional: true,
				//Computed:    true,
				Description: "The ID of Ecs Deployment Set. This field only used to associate a deployment set to the ECS instance. Setting this field to null means disassociating the instance from the deployment set. \n" +
					"The current deployment set id of the ECS instance is the `deployment_set_id_computed` field.",
			},
			"deployment_set_id_computed": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of Ecs Deployment Set. Computed field.",
			},

			"ipv6_address_count": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				Description:   "The number of IPv6 addresses to be automatically assigned from within the CIDR block of the subnet that hosts the ENI. Valid values: 1 to 10.",
				ValidateFunc:  validation.IntBetween(1, 10),
				ConflictsWith: []string{"ipv6_addresses"},
			},

			"ipv6_addresses": {
				Type:        schema.TypeSet,
				MaxItems:    10,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Set:         schema.HashString,
				Description: "One or more IPv6 addresses selected from within the CIDR block of the subnet that hosts the ENI. Support up to 10.\n You cannot specify both the ipv6_addresses and ipv6_address_count parameters.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ConflictsWith: []string{"ipv6_address_count"},
			},

			"cpu_options": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				MinItems:    1,
				Description: "The option of cpu,only support for ebm.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threads_per_core": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The per core of threads, only support for ebm. `1` indicates disabling hyper threading function.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								//暂时增加这个逻辑 在不包含ebm的情况下 忽略掉这个变化 目前这个方式比较hack 后续接口能力完善后改变一下
								if it, ok := d.Get("instance_type").(string); ok {
									its := strings.Split(it, ".")
									if len(its) == 3 && !strings.Contains(strings.ToLower(its[1]), "ebm") {
										return true
									} else {
										return false
									}
								} else {
									return true
								}
							},
						},
						"numa_per_socket": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "The number of subnuma in socket, only support for ebm. `1` indicates disabling SNC/NPS function. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								//暂时增加这个逻辑 在不包含ebm的情况下 忽略掉这个变化 目前这个方式比较hack 后续接口能力完善后改变一下
								if it, ok := d.Get("instance_type").(string); ok {
									its := strings.Split(it, ".")
									if len(its) == 3 && !strings.Contains(strings.ToLower(its[1]), "ebm") {
										return true
									} else {
										return false
									}
								} else {
									return true
								}
							},
						},
					},
				},
			},

			"data_volumes": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    15,
				MinItems:    1,
				Computed:    true,
				Description: "The data volumes collection of  ECS instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_type": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The type of volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.",
						},
						"size": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
							Description: "The size of volume. " +
								"The value range of the data volume size is ESSD_PL0: 10~32768, ESSD_FlexPL: 10~32768, PTSSD: 20~8192.",
						},
						"delete_with_instance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							ForceNew:    true,
							Description: "The delete with instance flag of volume.",
						},
					},
				},
			},

			"secondary_network_interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				MinItems:    1,
				Description: "The secondary networkInterface detail collection of ECS instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The subnet ID of secondary networkInterface.",
						},
						"security_group_ids": {
							Type:        schema.TypeSet,
							Required:    true,
							ForceNew:    true,
							MaxItems:    5,
							MinItems:    1,
							Description: "The security group ID set of secondary networkInterface.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"primary_ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The private ip address of secondary networkInterface.",
						},
					},
				},
			},
			"install_run_command_agent": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Description: "Whether to install the Run Command Agent. Default is false. This field is only effective when creating a new instance.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of the ecs instance.",
			},
			"tags": ve.TagsSchema(),
		},
	}
	dataSource := DataSourceVolcengineEcsInstances().Schema["instances"].Elem.(*schema.Resource).Schema
	delete(dataSource, "network_interfaces")
	delete(dataSource, "volumes")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineEcsInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(instanceService, d, ResourceVolcengineEcsInstance())
	if err != nil {
		return fmt.Errorf("error on creating ecs instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsInstanceRead(d, meta)
}

func resourceVolcengineEcsInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(instanceService, d, ResourceVolcengineEcsInstance())
	if err != nil {
		return fmt.Errorf("error on reading ecs instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(instanceService, d, ResourceVolcengineEcsInstance())
	if err != nil {
		return fmt.Errorf("error on updating ecs instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsInstanceRead(d, meta)
}

func resourceVolcengineEcsInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(instanceService, d, ResourceVolcengineEcsInstance())
	if err != nil {
		return fmt.Errorf("error on deleting ecs instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsInstanceExist(d *schema.ResourceData, meta interface{}) (flag bool, err error) {
	err = resourceVolcengineEcsInstanceRead(d, meta)
	if err != nil {
		if strings.Contains(err.Error(), "notfound") || strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not associate") ||
			strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "not_found") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
