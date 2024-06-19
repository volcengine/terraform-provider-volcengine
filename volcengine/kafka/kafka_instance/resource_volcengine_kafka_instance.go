package kafka_instance

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
KafkaInstance can be imported using the id, e.g.
```
$ terraform import volcengine_kafka_instance.default kafka-insbjwbbwb
```

*/

func ResourceVolcengineKafkaInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKafkaInstanceCreate,
		Read:   resourceVolcengineKafkaInstanceRead,
		Update: resourceVolcengineKafkaInstanceUpdate,
		Delete: resourceVolcengineKafkaInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The version of instance, the value can be `2.2.2` or `2.8.2`.",
			},
			"compute_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The compute spec of instance.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id of instance.",
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The user name of instance. " +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"user_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				Description: "The user password of instance. " +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"storage_space": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The storage space of instance.",
			},
			"partition_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The partition number of instance.",
			},
			"storage_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ESSD_FlexPL",
				ForceNew:    true,
				Description: "The storage type of instance. The value can be ESSD_FlexPL or ESSD_PL0.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of instance.",
			},
			"need_rebalance": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Id() == "" {
						return true
					}
					if !d.HasChange("compute_spec") {
						return true
					}
					return false
				},
				Description: "Whether enable rebalance. Only effected in modify when compute_spec field is changed.",
			},
			"rebalance_time": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Id() == "" {
						return true
					}
					// 此参数仅在NeedRebalance为true时生效
					if !d.HasChange("compute_spec") || !d.Get("need_rebalance").(bool) {
						return true
					}
					return false
				},
				Description: "The rebalance time.",
			},
			"instance_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of instance.",
			},
			"charge_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostPaid", "PrePaid"}, false),
				Description:  "The charge type of instance, the value can be `PrePaid` or `PostPaid`.",
			},
			"auto_renew": {
				Type:     schema.TypeBool,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					//在计费方式没有发生变化的时候 auto_renew 的变化会被忽略
					if !d.HasChange("charge_type") {
						return true
					}
					if d.Get("charge_type").(string) == "PostPaid" {
						return true
					}
					return false
				},
				Description: "The auto renew flag of instance. Only effective when instance_charge_type is PrePaid. Default is false.",
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					//在计费方式没有发生变化的时候 period的变化会被忽略
					if !d.HasChange("charge_type") {
						return true
					}
					if d.Get("charge_type").(string) == "PostPaid" {
						return true
					}
					return false
				},
				Description: "The period of instance. Only effective when instance_charge_type is PrePaid. Unit is Month.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The tags of instance.",
				Set:         TagsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Value of Tags.",
						},
					},
				},
			},
			"parameters": {
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         parameterHash,
				Description: "Parameter of the instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parameter_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Parameter name.",
						},
						"parameter_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Parameter value.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineKafkaInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKafkaInstance())
	if err != nil {
		return fmt.Errorf("error on creating kafka_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaInstanceRead(d, meta)
}

func resourceVolcengineKafkaInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKafkaInstance())
	if err != nil {
		return fmt.Errorf("error on reading kafka_instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKafkaInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKafkaInstance())
	if err != nil {
		return fmt.Errorf("error on updating kafka_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaInstanceRead(d, meta)
}

func resourceVolcengineKafkaInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKafkaInstance())
	if err != nil {
		return fmt.Errorf("error on deleting kafka_instance %q, %s", d.Id(), err)
	}
	return err
}

var parameterHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["parameter_name"], m["parameter_value"]))
	return hashcode.String(buf.String())
}
