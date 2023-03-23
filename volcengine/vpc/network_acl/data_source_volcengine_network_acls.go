package network_acl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNetworkAcls() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNetworkAclsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Network Acl IDs.",
			},
			"network_acl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of Network Acl.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id of Network Acl.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The subnet id of Network Acl.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Network Acl.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Network Acl query.",
			},
			"network_acls": {
				Description: "The collection of Network Acl query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Network Acl.",
						},
						"network_acl_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Network Acl.",
						},
						"network_acl_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Name of Network Acl.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of Network Acl.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of Network Acl.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Status of Network Acl.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of Network Acl.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of Network Acl.",
						},
						"acl_entry_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of Network acl entry.",
						},
						"ingress_acl_entries": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ingress entries info of Network Acl.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_acl_entry_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of entry.",
									},
									"network_acl_entry_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of entry.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of entry.",
									},
									"policy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The policy of entry.",
									},
									"source_cidr_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SourceCidrIp of entry.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol of entry.",
									},
									"priority": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The priority of entry.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The port of entry.",
									},
								},
							},
						},
						"egress_acl_entries": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The egress entries info of Network Acl.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_acl_entry_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of entry.",
									},
									"network_acl_entry_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of entry.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of entry.",
									},
									"policy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The policy of entry.",
									},
									"destination_cidr_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The DestinationCidrIp of entry.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol of entry.",
									},
									"priority": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The priority of entry.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The port of entry.",
									},
								},
							},
						},
						"resources": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resources info of Network Acl.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource id of Network Acl.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource status of Network Acl.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNetworkAclsRead(d *schema.ResourceData, meta interface{}) error {
	aclService := NewNetworkAclService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(aclService, d, DataSourceVolcengineNetworkAcls())
}
