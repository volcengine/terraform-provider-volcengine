package host_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsHostGroupRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsHostGroupRulesRead,
		Schema: map[string]*schema.Schema{
			"host_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of host group.",
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
			"rule_infos": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of rule info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of rule.",
						},
						"rule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of rule.",
						},
						"paths": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The paths of rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"pause": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The pause status of rule.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of topic.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of topic.",
						},
						"log_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of log.",
						},
						"input_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The type of input.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of rule.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modify time of rule.",
						},
						"log_sample": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The sample of the log.",
						},
						"extract_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The extract rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delimiter":   {Type: schema.TypeString, Computed: true},
									"begin_regex": {Type: schema.TypeString, Computed: true},
									"log_regex":   {Type: schema.TypeString, Computed: true},
									"keys":        {Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"time_key":    {Type: schema.TypeString, Computed: true},
									"time_format": {Type: schema.TypeString, Computed: true},
									"filter_key_regex": {
										Type: schema.TypeList, Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key":   {Type: schema.TypeString, Computed: true},
												"regex": {Type: schema.TypeString, Computed: true},
											},
										},
									},
									"un_match_up_load_switch": {Type: schema.TypeBool, Computed: true},
									"un_match_log_key":        {Type: schema.TypeString, Computed: true},
									"log_template": {
										Type: schema.TypeList, Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type":   {Type: schema.TypeString, Computed: true},
												"format": {Type: schema.TypeString, Computed: true},
											},
										},
									},
								},
							},
						},
						"exclude_paths": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Collect the blacklist list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type":  {Type: schema.TypeString, Computed: true},
									"value": {Type: schema.TypeString, Computed: true},
								},
							},
						},
						"user_define_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "User-defined collection rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_raw_log": {Type: schema.TypeBool, Computed: true},
									"fields":         {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"tail_files":     {Type: schema.TypeBool, Computed: true},
									"parse_path_rule": {
										Type: schema.TypeList, Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"path_sample": {Type: schema.TypeString, Computed: true},
												"regex":       {Type: schema.TypeString, Computed: true},
												"keys":        {Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
											},
										},
									},
									"shard_hash_key": {
										Type: schema.TypeList, Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hash_key": {Type: schema.TypeString, Computed: true},
											},
										},
									},
									"plugin": {
										Type: schema.TypeList, Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"processors": {Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
											},
										},
									},
									"advanced": {
										Type: schema.TypeList, Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"close_inactive": {Type: schema.TypeInt, Computed: true},
												"close_removed":  {Type: schema.TypeBool, Computed: true},
												"close_renamed":  {Type: schema.TypeBool, Computed: true},
												"close_eof":      {Type: schema.TypeBool, Computed: true},
												"close_timeout":  {Type: schema.TypeInt, Computed: true},
											},
										},
									},
								},
							},
						},
						"container_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Container collection rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stream":                        {Type: schema.TypeString, Computed: true},
									"container_name_regex":          {Type: schema.TypeString, Computed: true},
									"include_container_label_regex": {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"exclude_container_label_regex": {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"include_container_env_regex":   {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"exclude_container_env_regex":   {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"env_tag":                       {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"kubernetes_rule": {
										Type: schema.TypeList, Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"namespace_name_regex":    {Type: schema.TypeString, Computed: true},
												"workload_type":           {Type: schema.TypeString, Computed: true},
												"workload_name_regex":     {Type: schema.TypeString, Computed: true},
												"include_pod_label_regex": {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
												"exclude_pod_label_regex": {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
												"pod_name_regex":          {Type: schema.TypeString, Computed: true},
												"label_tag":               {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
												"annotation_tag":          {Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
											},
										},
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

func dataSourceVolcengineTlsHostGroupRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	m := map[string]interface{}{
		"host_group_id": d.Get("host_group_id"),
	}
	data, err := service.ReadRules(m)
	if err != nil {
		return err
	}
	d.SetId(d.Get("host_group_id").(string))
	if len(data) > 0 {
		if err := d.Set("rule_infos", data); err != nil {
			return err
		}
	}
	d.Set("total_count", len(data))
	return nil
}
