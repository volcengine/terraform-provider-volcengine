---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_image_import"
sidebar_current: "docs-volcengine-resource-image_import"
description: |-
  Provides a resource to manage image import
---
# volcengine_image_import
Provides a resource to manage image import
## Example Usage
```hcl
resource "volcengine_image_import" "foo" {
  platform     = "CentOS"
  url          = "https://*****_system.qcow2"
  image_name   = "acc-test-image"
  description  = "acc-test"
  boot_mode    = "UEFI"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `image_name` - (Required) The name of the custom image.
* `platform` - (Required, ForceNew) The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
* `url` - (Required, ForceNew) The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `architecture` - (Optional, ForceNew) The architecture of the custom image. Valid values: `amd64`, `arm64`.
* `boot_mode` - (Optional) The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
* `description` - (Optional) The description of the custom image.
* `license_type` - (Optional, ForceNew) The license type of the custom image. Valid values: `VolcanoEngine`.
* `os_type` - (Optional, ForceNew) The os type of the custom image. Valid values: `linux`, `Windows`.
* `platform_version` - (Optional, ForceNew) The platform version of the custom image.
* `project_name` - (Optional) The project name of the custom image.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - The create time of Image.
* `is_support_cloud_init` - Whether the Image support cloud-init.
* `os_name` - The name of Image operating system.
* `share_status` - The share mode of Image.
* `size` - The size(GiB) of Image.
* `status` - The status of Image.
* `updated_at` - The update time of Image.
* `visibility` - The visibility of Image.


## Import
ImageImport can be imported using the id, e.g.
```
$ terraform import volcengine_image_import.default resource_id
```

