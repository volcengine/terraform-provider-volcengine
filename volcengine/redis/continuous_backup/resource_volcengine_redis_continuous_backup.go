package continuous_backup

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Redis Continuous Backup can be imported using the continuous:instanceId, e.g.
```
$ terraform import volcengine_redis_continuous_backup.default continuous:redis-asdljioeixxxx
```
*/

func ResourceVolcengineRedisContinuousBackup() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineRedisContinuousBackupRead,
		Create: resourceVolcengineRedisContinuousBackupCreate,
		Delete: resourceVolcengineRedisContinuousBackupDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Id of redis instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineRedisContinuousBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	redisContinuousBackupService := NewRedisContinuousBackupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(redisContinuousBackupService, d, ResourceVolcengineRedisContinuousBackup())
	if err != nil {
		return fmt.Errorf("error on read continuous backup %v, %v", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRedisContinuousBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	redisContinuousBackupService := NewRedisContinuousBackupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(redisContinuousBackupService, d, ResourceVolcengineRedisContinuousBackup())
	if err != nil {
		return fmt.Errorf("error on create continuous backup %v, %v", d.Id(), err)
	}
	return resourceVolcengineRedisContinuousBackupRead(d, meta)
}

func resourceVolcengineRedisContinuousBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	redisContinuousBackupService := NewRedisContinuousBackupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(redisContinuousBackupService, d, ResourceVolcengineRedisContinuousBackup())
	if err != nil {
		return fmt.Errorf("error on delete continuous backup %v, %v", d.Id(), err)
	}
	return nil
}
