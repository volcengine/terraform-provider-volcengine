package iam_allowed_ip_address

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*
Import
Iam AllowedIpAddress key don't support import
*/
func ResourceVolcengineIamAllowedIpAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineIamAllowedIpAddressCreate,
		Read:   resourceVolcengineIamAllowedIpAddressRead,
		Update: resourceVolcengineIamAllowedIpAddressUpdate,
		Delete: resourceVolcengineIamAllowedIpAddressDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"enable_ip_list": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to enable the IP whitelist.",
			},
			"ip_list": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The IP whitelist list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The IP address.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The description of the IP address.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineIamAllowedIpAddressCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewIamAllowedIpAddressService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineIamAllowedIpAddress())
	if err != nil {
		return fmt.Errorf("error on creating allowed ip addresses: %s", err)
	}
	return resourceVolcengineIamAllowedIpAddressRead(d, meta)
}

func resourceVolcengineIamAllowedIpAddressRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamAllowedIpAddressService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Read(service, d, ResourceVolcengineIamAllowedIpAddress())
	if err != nil {
		return fmt.Errorf("error on reading allowed ip addresses: %s", err)
	}
	return nil
}

func resourceVolcengineIamAllowedIpAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewIamAllowedIpAddressService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Update(service, d, ResourceVolcengineIamAllowedIpAddress())
	if err != nil {
		return fmt.Errorf("error on updating allowed ip addresses: %s", err)
	}
	return resourceVolcengineIamAllowedIpAddressRead(d, meta)
}

func resourceVolcengineIamAllowedIpAddressDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewIamAllowedIpAddressService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineIamAllowedIpAddress())
	if err != nil {
		return fmt.Errorf("error on deleting allowed ip addresses: %s", err)
	}
	return nil
}
