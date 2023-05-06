package project

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsProjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsProjectsRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The id of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.",
				ConflictsWith: []string{"project_name"},
			},
			"project_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The name of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.",
				ConflictsWith: []string{"project_id"},
			},
			"is_full_name": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to match accurately when filtering based on ProjectName.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("project_name").(string) == ""
				},
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IAM project name of the tls project.",
			},
			"tags": ve.TagsSchema(),
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of tls project.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of tls project query.",
			},
			"tls_projects": {
				Description: "The collection of tls project query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the tls project.",
						},
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the tls project.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the tls project.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the tls project.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the tls project.",
						},
						"inner_net_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The inner net domain of the tls project.",
						},
						"topic_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of topics in the tls project.",
						},
						"iam_project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IAM project name of the tls project.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsProjectsRead(d *schema.ResourceData, meta interface{}) error {
	tlsProjectService := NewTlsProjectService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(tlsProjectService, d, DataSourceVolcengineTlsProjects())
}
