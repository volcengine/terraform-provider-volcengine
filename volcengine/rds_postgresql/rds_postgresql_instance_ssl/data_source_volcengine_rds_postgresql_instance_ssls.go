package rds_postgresql_instance_ssl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceSsls() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceSslsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of the PostgreSQL instance IDs.",
			},
			"download_certificate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to include SSL certificate raw bytes for each instance.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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
			"ssls": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the postgresql Instance.",
						},
						"ssl_enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable SSL.",
						},
						"force_encryption": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to force encryption.",
						},
						"is_valid": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the SSL certificate is valid.",
						},
						"ssl_expire_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The expiration time of the SSL certificate. " +
								"The format is: yyyy-MM-ddTHH:mm:ss(UTC time).",
						},
						"tls_version": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The supported TLS versions.",
						},
						"address": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The protected addresses.",
						},
						"certificate": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "Raw byte stream array of certificate zip.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstanceSslsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceSslService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceSsls())
}
