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
data "volcengine_images" "default" {
  ids = ["image-cm9ssb4eqmhdas306zlp", "image-ybkzct2rtj4ay5rmlfc3"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Image IDs.
* `instance_type_id` - (Optional) The specification of  Instance.
* `is_support_cloud_init` - (Optional) Whether the Image support cloud-init.
* `name_regex` - (Optional) A Name Regex of Image.
* `os_type` - (Optional) The operating system type of Image.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) A list of Image status.
* `visibility` - (Optional) The visibility of Image.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `images` - The collection of Image query.
  * `architecture` - The architecture of Image.
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
  * `updated_at` - The update time of Image.
  * `visibility` - The visibility of Image.
* `total_count` - The total count of Image query.


