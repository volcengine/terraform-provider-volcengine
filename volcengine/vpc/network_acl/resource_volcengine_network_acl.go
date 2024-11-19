package network_acl

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Network Acl can be imported using the id, e.g.
```
$ terraform import volcengine_network_acl.default nacl-172leak37mi9s4d1w33pswqkh
```

*/

func ResourceVolcengineNetworkAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNetworkAclCreate,
		Read:   resourceVolcengineNetworkAclRead,
		Update: resourceVolcengineNetworkAclUpdate,
		Delete: resourceVolcengineNetworkAclDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The vpc id of Network Acl.",
			},
			"network_acl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of Network Acl.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Network Acl.",
			},
			"ingress_acl_entries": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The ingress entries of Network Acl.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_acl_entry_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of entry.",
						},
						"network_acl_entry_name": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Description: "The name of entry.",
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Description: "The description of entry.",
						},
						"policy": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Default:     "accept",
							Description: "The policy of entry, default is `accept`. The value can be `accept` or `drop`.",
						},
						"source_cidr_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The SourceCidrIp of entry.",
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Default: "all",
							Description: "The protocol of entry, default is `all`. " +
								"The value can be `icmp` or `gre` or `tcp` or `udp` or `all`.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of entry.",
						},
						"port": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Default: "-1/-1",
							Description: "The port of entry. Default is `-1/-1`. When Protocol is `all`, `icmp` or `gre`, " +
								"the port range is `-1/-1`, which means no port restriction. " +
								"When the Protocol is `tcp` or `udp`, the port range is `1~65535`, and the format is `1/200`, `80/80`, " +
								"which means port 1 to port 200, port 80.",
						},
					},
				},
			},
			"egress_acl_entries": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The egress entries of Network Acl.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_acl_entry_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of entry.",
						},
						"network_acl_entry_name": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Description: "The name of entry.",
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Description: "The description of entry.",
						},
						"policy": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Default:     "accept",
							Description: "The policy of entry. Default is `accept`. The value can be `accept` or `drop`.",
						},
						"destination_cidr_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The DestinationCidrIp of entry.",
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Default: "all",
							Description: "The protocol of entry. " +
								"The value can be `icmp` or `gre` or `tcp` or `udp` or `all`. Default is `all`.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of entry.",
						},
						"port": {
							Type:     schema.TypeString,
							Optional: true,
							//Computed:    true,
							Default: "-1/-1",
							Description: "The port of entry. Default is `-1/-1`. " +
								"When Protocol is `all`, `icmp` or `gre`, the port range is `-1/-1`, " +
								"which means no port restriction." +
								"When the Protocol is `tcp` or `udp`, the port range is `1~65535`, and the format is `1/200`, `80/80`," +
								"which means port 1 to port 200, port 80.",
						},
					},
				},
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the network acl.",
			},
			"tags": ve.TagsSchema(),
		},
	}
}

func resourceVolcengineNetworkAclCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewNetworkAclService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(aclService, d, ResourceVolcengineNetworkAcl())
	if err != nil {
		return fmt.Errorf("error on creating network acl %q, %w", d.Id(), err)
	}
	return resourceVolcengineNetworkAclRead(d, meta)
}

func resourceVolcengineNetworkAclRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewNetworkAclService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(aclService, d, ResourceVolcengineNetworkAcl())
	if err != nil {
		return fmt.Errorf("error on reading network acl %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineNetworkAclUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewNetworkAclService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(aclService, d, ResourceVolcengineNetworkAcl())
	if err != nil {
		return fmt.Errorf("error on updating network acl %q, %w", d.Id(), err)
	}
	return resourceVolcengineNetworkAclRead(d, meta)
}

func resourceVolcengineNetworkAclDelete(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewNetworkAclService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(aclService, d, ResourceVolcengineNetworkAcl())
	if err != nil {
		return fmt.Errorf("error on deleting network acl %q, %w", d.Id(), err)
	}
	return err
}
