package bandwidth_package

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
BandwidthPackage can be imported using the id, e.g.
```
$ terraform import volcengine_bandwidth_package.default bwp-2zeo05qre24nhrqpy****
```

*/

func ResourceVolcengineBandwidthPackage() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineBandwidthPackageCreate,
		Read:   resourceVolcengineBandwidthPackageRead,
		Update: resourceVolcengineBandwidthPackageUpdate,
		Delete: resourceVolcengineBandwidthPackageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bandwidth_package_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the bandwidth package.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the bandwidth package.",
			},
			"isp": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "BGP",
				ForceNew:    true,
				Description: "Route type, default to BGP.",
			},
			"billing_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "PostPaidByBandwidth",
				Description: "BillingType of the Ipv6 bandwidth. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`.",
			},
			"bandwidth": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Bandwidth upper limit of shared bandwidth package, unit: Mbps. Valid values: 2 to 5000.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.",
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("billing_type").(string) != "PrePaid"
				},
				Description: "Duration of purchasing shared bandwidth package on an annual or monthly basis. " +
					"The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.",
			},
			"security_protection_types": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Security protection types for shared bandwidth packages. " +
					"Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank." +
					"If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection," +
					" which supports adding basic protection type public IP addresses." +
					"If left blank, it indicates a shared bandwidth package with basic protection, " +
					"which supports the addition of public IP addresses with enhanced protection.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the bandwidth package.",
			},
			"tags": ve.TagsSchema(),
		},
	}
	return resource
}

func resourceVolcengineBandwidthPackageCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on creating bandwidth_package %q, %s", d.Id(), err)
	}
	return resourceVolcengineBandwidthPackageRead(d, meta)
}

func resourceVolcengineBandwidthPackageRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on reading bandwidth_package %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineBandwidthPackageUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on updating bandwidth_package %q, %s", d.Id(), err)
	}
	return resourceVolcengineBandwidthPackageRead(d, meta)
}

func resourceVolcengineBandwidthPackageDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on deleting bandwidth_package %q, %s", d.Id(), err)
	}
	return err
}
