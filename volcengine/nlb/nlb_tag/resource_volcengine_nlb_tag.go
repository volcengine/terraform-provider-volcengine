package nlb_tag

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Nlb tag can be imported using the ResourceType, ResourceId and TagKey, e.g.
```
$ terraform import volcengine_nlb_tag.foo nlb:nlb-29mxxx:key1
```

*/

func ResourceVolcengineNlbTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNlbTagCreate,
		Read:   resourceVolcengineNlbTagRead,
		Update: resourceVolcengineNlbTagUpdate,
		Delete: resourceVolcengineNlbTagDelete,
		CustomizeDiff: func(diff *schema.ResourceDiff, meta interface{}) error {
			if v, ok := diff.GetOk("tags"); ok {
				tagsSet, ok := v.(*schema.Set)
				if !ok {
					return fmt.Errorf("tags is not *schema.Set")
				}
				tags := tagsSet.List()
				keys := make(map[string]bool)
				for _, t := range tags {
					tag, ok := t.(map[string]interface{})
					if !ok {
						return fmt.Errorf("tag item is not map")
					}
					key, ok := tag["key"].(string)
					if !ok {
						return fmt.Errorf("tag key is not string")
					}
					if keys[key] {
						return fmt.Errorf("duplicate tag key %q found in tags blocks", key)
					}
					keys[key] = true
				}
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				id := d.Id()
				parts := strings.Split(id, ":")
				if len(parts) < 2 || len(parts) > 3 {
					return nil, fmt.Errorf("invalid import id %q, expected format: ResourceType:ResourceId or ResourceType:ResourceId:TagKey (e.g., nlb:nlb-29ml5i7992g3k1e1hgiptpis0:key1)", id)
				}

				resourceType := parts[0]
				resourceId := parts[1]

				service := NewNlbTagService(meta.(*ve.SdkClient))
				m := map[string]interface{}{
					"ResourceType":  resourceType,
					"ResourceIds.1": resourceId,
				}
				cloudResources, err := service.ReadResources(m)
				if err != nil {
					return nil, fmt.Errorf("error querying tags for import: %s", err)
				}

				var tags []interface{}
				foundKey := false
				var targetKey string
				if len(parts) == 3 {
					targetKey = parts[2]
				}

				for _, cr := range cloudResources {
					crm, ok := cr.(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("cloud tag item is not map")
					}
					tagKey, ok := crm["TagKey"].(string)
					if !ok {
						return nil, fmt.Errorf("cloud tag key is not string")
					}
					if strings.HasPrefix(tagKey, "sys:") {
						continue
					}
					if targetKey != "" && tagKey == targetKey {
						foundKey = true
					}
					tagValue, ok := crm["TagValue"].(string)
					if !ok {
						return nil, fmt.Errorf("cloud tag value is not string")
					}
					tags = append(tags, map[string]interface{}{
						"key":   tagKey,
						"value": tagValue,
					})
				}

				if targetKey != "" && !foundKey {
					return nil, fmt.Errorf("tag key %q not found for %s %q", targetKey, resourceType, resourceId)
				}

				_ = d.Set("resource_type", resourceType)
				_ = d.Set("resource_id", resourceId)
				_ = d.Set("tags", schema.NewSet(ve.TagsHash, tags))

				// Set a stable ID
				d.SetId(fmt.Sprintf("%s:%s", resourceType, resourceId))

				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the resource.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the resource. Valid values: `nlb`, `nlb_listener`, `nlb_servergroup`, `nlb_security_policy`.",
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      ve.TagsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineNlbTagCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbTagService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineNlbTag())
	if err != nil {
		return fmt.Errorf("error on creating nlb tag %q, %w", d.Id(), err)
	}
	return resourceVolcengineNlbTagRead(d, meta)
}

func resourceVolcengineNlbTagRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbTagService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Read(service, d, ResourceVolcengineNlbTag())
	if err != nil {
		return fmt.Errorf("error on reading nlb tag %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineNlbTagUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbTagService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Update(service, d, ResourceVolcengineNlbTag())
	if err != nil {
		return fmt.Errorf("error on updating nlb tag %q, %w", d.Id(), err)
	}
	return resourceVolcengineNlbTagRead(d, meta)
}

func resourceVolcengineNlbTagDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbTagService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineNlbTag())
	if err != nil {
		return fmt.Errorf("error on deleting nlb tag %q, %w", d.Id(), err)
	}
	return err
}
