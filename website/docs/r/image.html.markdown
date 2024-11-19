---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_image"
sidebar_current: "docs-volcengine-resource-image"
description: |-
  Provides a resource to manage image
---
# volcengine_image
Provides a resource to manage image
## Example Usage
```hcl
resource "volcengine_image" "foo" {
  image_name         = "acc-test-image"
  description        = "acc-test"
  instance_id        = "i-ydi2q1s7wgqc6ild****"
  create_whole_image = false
  project_name       = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `image_name` - (Required) The name of the custom image.
* `boot_mode` - (Optional) The boot mode of the custom image. Valid values: `BIOS`, `UEFI`. This field is only effective when modifying the image.
* `create_whole_image` - (Optional) Whether to create whole image. Default is false. This field is only effective when creating a new custom image.
* `description` - (Optional) The description of the custom image.
* `instance_id` - (Optional, ForceNew) The instance id of the custom image. Only one of `instance_id, snapshot_id, snapshot_group_id` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `project_name` - (Optional) The project name of the custom image.
* `snapshot_group_id` - (Optional, ForceNew) The snapshot group id of the custom image. Only one of `instance_id, snapshot_id, snapshot_group_id` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `snapshot_id` - (Optional, ForceNew) The snapshot id of the custom image. Only one of `instance_id, snapshot_id, snapshot_group_id` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `architecture` - The architecture of Image.
* `created_at` - The create time of Image.
* `is_support_cloud_init` - Whether the Image support cloud-init.
* `os_name` - The name of Image operating system.
* `os_type` - The operating system type of Image.
* `platform_version` - The platform version of Image.
* `platform` - The platform of Image.
* `share_status` - The share mode of Image.
* `size` - The size(GiB) of Image.
* `status` - The status of Image.
* `updated_at` - The update time of Image.
* `visibility` - The visibility of Image.


## Import
Image can be imported using the id, e.g.
```
$ terraform import volcengine_image.default resource_id
```

