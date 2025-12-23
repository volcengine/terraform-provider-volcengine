package rds_postgresql_replication_slot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlReplicationSlots() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlReplicationSlotsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the PostgreSQL instance",
			},
			"slot_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the slot.",
			},
			"slot_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the slot: physical or logical.",
			},
			"plugin": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the plugin used by the logical replication slot to parse WAL logs.",
			},
			"data_base": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The database where the replication slot is located.",
			},
			"temporary": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the slot is temporary.",
			},
			"slot_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the replication slot: ACTIVE or INACTIVE.",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ip address.",
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
			"replication_slots": {
				Description: "Replication slots under the specified query conditions in the instance.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"slot_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the slot.",
						},
						"slot_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the slot: physical or logical.",
						},
						"plugin": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the plugin used by the logical replication slot to parse WAL logs.",
						},
						"data_base": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The database where the replication slot is located.",
						},
						"temporary": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether the slot is temporary.",
						},
						"slot_status": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The status of the replication slot: ACTIVE or INACTIVE.",
						},
						"wal_delay": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The cumulative WAL log volume corresponding to this replication slot. The unit is Byte.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ip address.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlReplicationSlotsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlReplicationSlotService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlReplicationSlots())
}
