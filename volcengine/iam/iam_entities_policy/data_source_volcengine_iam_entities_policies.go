package iam_entities_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamEntitiesPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamEntitiesPoliciesRead,
		Schema: map[string]*schema.Schema{
			"policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the policy.",
			},
			"policy_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the policy.",
			},
			"entity_filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The entity filter.",
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
			"users": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of users.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the user.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the user.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the user.",
						},
						"attach_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The attach date of the user.",
						},
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the user.",
						},
						"policy_scope": policyScopeSchema(),
					},
				},
			},
			"roles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of roles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the role.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the role.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the role.",
						},
						"attach_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The attach date of the role.",
						},
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the role.",
						},
						"policy_scope": policyScopeSchema(),
					},
				},
			},
			"user_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of user groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the user group.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the user group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the user group.",
						},
						"attach_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The attach date of the user group.",
						},
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the user group.",
						},
						"policy_scope": policyScopeSchema(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamEntitiesPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamEntitiesPolicyService(meta.(*ve.SdkClient))

	// Prepare request parameters manually to match service expectations
	m := map[string]interface{}{
		"PolicyName": d.Get("policy_name"),
		"PolicyType": d.Get("policy_type"),
	}
	if v, ok := d.GetOk("entity_filter"); ok {
		m["EntityFilter"] = v
	}

	// Call ReadResources which now returns fully mapped snake_case data
	data, err := service.ReadResources(m)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		res := data[0].(map[string]interface{})
		// Set ID from the result
		if id, ok := res["id"].(string); ok {
			d.SetId(id)
		} else {
			// Fallback ID
			d.SetId(d.Get("policy_name").(string))
		}

		// Set top-level fields
		d.Set("users", res["users"])
		d.Set("roles", res["roles"])
		d.Set("user_groups", res["user_groups"])
		d.Set("total_count", res["total_count"])
	} else {
		// If no data, we should still set an ID to avoid "not found" errors in some cases
		// But for data sources, empty results are often okay.
		d.SetId(d.Get("policy_name").(string))
	}

	return nil
}

func policyScopeSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: "The scope of the policy.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"policy_scope_type": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The type of the policy scope.",
				},
				"project_name": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The name of the project.",
				},
				"project_display_name": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The display name of the project.",
				},
				"attach_date": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The attach date of the policy scope.",
				},
			},
		},
	}
}
