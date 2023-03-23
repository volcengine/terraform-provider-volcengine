package ssl_state

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineMongoDBSSLStates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineMongoDBSSLStatesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The mongodb instance ID to query.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of mongodb ssl state query.",
			},
			"ssl_state": {
				Description: "The collection of mongodb ssl state query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The mongodb instance id.",
						},
						"ssl_enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether SSL is enabled.",
						},
						"is_valid": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whetehr SSL is valid.",
						},
						"ssl_expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expire time of SSL.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineMongoDBSSLStatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewMongoDBSSLStateService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineMongoDBSSLStates())
}
