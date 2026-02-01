---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_images"
sidebar_current: "docs-volcengine-datasource-images"
description: |-
  Use this data source to query detailed information of images
---
# volcengine_images
Use this data source to query detailed information of images
## Example Usage
```hcl
data "volcengine_images" "foo" {
  os_type          = "Linux"
  visibility       = "public"
  instance_type_id = "ecs.g1.large"
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Image IDs.
* `image_name` - (Optional) The name of Image.
* `instance_type_id` - (Optional) The specification of  Instance.
* `is_support_cloud_init` - (Optional) Whether the Image support cloud-init.
* `is_tls` - (Optional) Whether the Image maintained for a long time.
* `name_regex` - (Optional) A Name Regex of Image.
* `os_type` - (Optional) The operating system type of Image.
* `output_file` - (Optional) File name where to save data source results.
* `platform` - (Optional) The platform of Image.
* `status` - (Optional) A list of Image status, the value can be `available` or `creating` or `error`.
* `tags` - (Optional) Tags.
* `visibility` - (Optional) The visibility of Image.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `images` - The collection of Image query.
    * `architecture` - The architecture of Image.
    * `boot_mode` - The boot mode of Image.
    * `created_at` - The create time of Image.
    * `description` - The description of Image.
    * `image_id` - The ID of Image.
    * `image_name` - The name of Image.
    * `is_support_cloud_init` - Whether the Image support cloud-init.
    * `os_name` - The name of Image operating system.
    * `os_type` - The operating system type of Image.
    * `platform_version` - The platform version of Image.
    * `platform` - The platform of Image.
    * `share_status` - The share mode of Image.
    * `size` - The size(GiB) of Image.
    * `status` - The status of Image.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `updated_at` - The update time of Image.
    * `visibility` - The visibility of Image.
* `total_count` - The total count of Image query.


