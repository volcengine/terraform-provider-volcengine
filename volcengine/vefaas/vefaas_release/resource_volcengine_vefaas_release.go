package vefaas_release

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VefaasRelease can be imported using the id, e.g.
```
$ terraform import volcengine_vefaas_release.default FunctionId:ReleaseRecordId
```

*/

func ResourceVolcengineVefaasRelease() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVefaasReleaseCreate,
		Read:   resourceVolcengineVefaasReleaseRead,
		Update: resourceVolcengineVefaasReleaseUpdate,
		Delete: resourceVolcengineVefaasReleaseDelete,
		Importer: &schema.ResourceImporter{
			State: vefaasReleaseImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"function_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of Function.",
			},
			"revision_number": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "When the RevisionNumber to be released is 0, " +
					"the Latest code (Latest) will be released and a new version will be created. " +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"target_traffic_weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Target percentage of published traffic.",
			},
			"rolling_step": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Percentage of grayscale step size.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of released this time.",
			},
			"max_instance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Upper limit of the number of function instances.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of function release.",
			},
			"status_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Detailed information of the function release status.",
			},
			"stable_revision_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current version number that is stably running online.",
			},
			"new_revision_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version number of the newly released version.",
			},
			"old_revision_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version number of the old version.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current release start time.",
			},
			"current_traffic_weight": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current percentage of current published traffic.",
			},
			"error_code": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Error code when the release fails.",
			},
			"failed_instance_logs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Download link for the failed instance log.",
			},
			"release_record_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of Release record.",
			},
		},
	}
	return resource
}

func resourceVolcengineVefaasReleaseCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasReleaseService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVefaasRelease())
	if err != nil {
		return fmt.Errorf("error on creating vefaas_release %q, %s", d.Id(), err)
	}
	return resourceVolcengineVefaasReleaseRead(d, meta)
}

func resourceVolcengineVefaasReleaseRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasReleaseService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVefaasRelease())
	if err != nil {
		return fmt.Errorf("error on reading vefaas_release %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVefaasReleaseUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasReleaseService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVefaasRelease())
	if err != nil {
		return fmt.Errorf("error on updating vefaas_release %q, %s", d.Id(), err)
	}
	return resourceVolcengineVefaasReleaseRead(d, meta)
}

func resourceVolcengineVefaasReleaseDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasReleaseService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVefaasRelease())
	if err != nil {
		return fmt.Errorf("error on deleting vefaas_release %q, %s", d.Id(), err)
	}
	return err
}
