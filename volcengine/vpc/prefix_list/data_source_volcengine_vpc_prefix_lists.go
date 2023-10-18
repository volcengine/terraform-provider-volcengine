package prefix_list

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVpcPrefixLists() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVpcPrefixListsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of prefix list ids.",
			},
			"prefix_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A Name of prefix list.",
			},
			"ip_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP version of prefix list.",
			},
			"tag_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of tag filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The key of the tag.",
						},
						"values": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The values of the tag.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
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
			"prefix_lists": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the prefix list.",
						},
						"prefix_list_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The prefix list id.",
						},
						"prefix_list_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The prefix list name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description.",
						},
						"ip_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ip version of the prefix list.",
						},
						"max_entries": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of entries, which is the maximum number of items that can be added to the prefix list.",
						},
						"association_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of associated resources for prefix list.",
						},
						"cidrs": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "CIDR address block information for prefix list.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the prefix list.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the prefix list.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the prefix list.",
						},
						"prefix_list_entries": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The prefix list entries.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix_list_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The prefix list id.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Description.",
									},
									"cidr": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CIDR address blocks for prefix list entries.",
									},
								},
							},
						},
						"prefix_list_associations": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Collection of resources associated with VPC prefix list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Associated resource ID.",
									},
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Related resource types.",
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

func dataSourceVolcengineVpcPrefixListsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVpcPrefixListService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVpcPrefixLists())
}
