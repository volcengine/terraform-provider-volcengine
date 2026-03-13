package iam_tag

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Iam tag can be imported using the ResourceType, ResourceName and TagKey, e.g.
```
$ terraform import volcengine_iam_tag.default User:jonny:key1
```

*/

func ResourceVolcengineIamTag() *schema.Resource {
	tagsSchema := ve.TagsSchema()
	tagsSchema.ForceNew = true
	return &schema.Resource{
		Create: resourceVolcengineIamTagCreate,
		Read:   resourceVolcengineIamTagRead,
		Delete: resourceVolcengineIamTagDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				id := d.Id()
				parts := strings.Split(id, ":")
				if len(parts) != 3 {
					return nil, fmt.Errorf("invalid import id %q, expected format: ResourceType:ResourceName:TagKey (e.g., User:jonny:key1)", id)
				}

				resourceType := parts[0]
				resourceName := parts[1]
				tagKey := parts[2]

				service := NewIamTagService(meta.(*ve.SdkClient))
				m := map[string]interface{}{
					"ResourceType":    resourceType,
					"ResourceNames.1": resourceName,
				}
				cloudResources, err := service.ReadResources(m)
				if err != nil {
					return nil, fmt.Errorf("error querying tags for import: %s", err)
				}

				var tagValue string
				found := false
				for _, cr := range cloudResources {
					crm, ok := cr.(map[string]interface{})
					if !ok {
						continue
					}
					if crm["TagKey"] == tagKey {
						if v, ok := crm["TagValue"].(string); ok {
							tagValue = v
						}
						found = true
						break
					}
				}

				if !found {
					return nil, fmt.Errorf("tag %q not found for %s %q", tagKey, resourceType, resourceName)
				}

				_ = d.Set("resource_type", resourceType)
				_ = d.Set("resource_names", []interface{}{resourceName})
				_ = d.Set("tags", schema.NewSet(ve.TagsHash, []interface{}{
					map[string]interface{}{
						"key":   tagKey,
						"value": tagValue,
					},
				}))

				// Set a stable ID
				d.SetId(fmt.Sprintf("%s:%s:%s", resourceType, resourceName, tagKey))

				return []*schema.ResourceData{d}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the resource. Valid values: User, Role.",
			},
			"resource_names": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "The names of the resource.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": tagsSchema,
		},
	}
}

func resourceVolcengineIamTagCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewIamTagService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineIamTag())
	if err != nil {
		return fmt.Errorf("error on creating iam tag %q, %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineIamTagRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamTagService(meta.(*ve.SdkClient))
	resourceType, ok := d.Get("resource_type").(string)
	if !ok {
		return fmt.Errorf("resource_type is not string")
	}
	vNames := d.Get("resource_names")
	resourceNames, ok := vNames.([]interface{})
	if !ok {
		return fmt.Errorf("resource_names is not []interface{}")
	}
	vTags := d.Get("tags")
	tagsSet, ok := vTags.(*schema.Set)
	if !ok {
		return fmt.Errorf("tags is not *schema.Set")
	}

	m := map[string]interface{}{
		"ResourceType": resourceType,
	}
	for i, name := range resourceNames {
		nameStr, ok := name.(string)
		if !ok {
			continue
		}
		m[fmt.Sprintf("ResourceNames.%d", i+1)] = nameStr
	}

	cloudResources, err := service.ReadResources(m)
	if err != nil {
		return fmt.Errorf("error on reading iam tags for %v, %s", resourceNames, err)
	}

	// Build a map for quick lookup: map[ResourceName]map[TagKey]TagValue
	cloudTagMap := make(map[string]map[string]string)
	for _, cr := range cloudResources {
		crm, ok := cr.(map[string]interface{})
		if !ok {
			return fmt.Errorf("cloud tag item is not map")
		}
		rName, ok := crm["ResourceName"].(string)
		if !ok {
			return fmt.Errorf("cloud tag ResourceName is not string")
		}
		tKey, ok := crm["TagKey"].(string)
		if !ok {
			return fmt.Errorf("cloud tag TagKey is not string")
		}
		tValue, ok := crm["TagValue"].(string)
		if !ok {
			return fmt.Errorf("cloud tag TagValue is not string")
		}
		if rName == "" || tKey == "" {
			continue
		}
		if _, ok := cloudTagMap[rName]; !ok {
			cloudTagMap[rName] = make(map[string]string)
		}
		cloudTagMap[rName][tKey] = tValue
	}

	// For each tag in our set, check if it exists on all resources in our list
	tagsToFind := tagsSet.List()
	for _, tag := range tagsToFind {
		tm, ok := tag.(map[string]interface{})
		if !ok {
			return fmt.Errorf("tag item is not map")
		}
		tagKey, ok := tm["key"].(string)
		if !ok {
			return fmt.Errorf("tag key is not string")
		}
		tagValue, ok := tm["value"].(string)
		if !ok {
			return fmt.Errorf("tag value is not string")
		}

		for _, name := range resourceNames {
			nameStr, ok := name.(string)
			if !ok {
				return fmt.Errorf("resource name item is not string")
			}
			found := false
			if resTags, ok := cloudTagMap[nameStr]; ok {
				if val, ok2 := resTags[tagKey]; ok2 && val == tagValue {
					found = true
				}
			}

			if !found {
				d.SetId("")
				return nil
			}
		}
	}

	return nil
}

func resourceVolcengineIamTagDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewIamTagService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineIamTag())
	if err != nil {
		return fmt.Errorf("error on deleting iam tag %q, %s", d.Id(), err)
	}
	return nil
}
