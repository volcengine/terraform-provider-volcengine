---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_key_archive"
sidebar_current: "docs-volcengine-resource-kms_key_archive"
description: |-
  Provides a resource to manage kms key archive
---
# volcengine_kms_key_archive
Provides a resource to manage kms key archive
## Example Usage
```hcl
resource "volcengine_kms_keyring" "foo" {
  keyring_name = "tf-test"
  description  = "tf-test"
  project_name = "default"
}

resource "volcengine_kms_key" "foo" {
  keyring_name = volcengine_kms_keyring.foo.keyring_name
  key_name     = "mrk-tf-key-mod"
  description  = "tf test key-mod"
  tags {
    key   = "tfkey3"
    value = "tfvalue3"
  }
}

resource "volcengine_kms_key_archive" "foo" {
  key_id = volcengine_kms_key.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `key_id` - (Optional, ForceNew) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of the key.
* `keyring_name` - (Optional, ForceNew) The name of the keyring.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `key_state` - The state of the key.


## Import
KmsKeyArchive can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_archive.default resource_id
or
$ terraform import volcengine_kms_key_archive.default key_name:keyring_name
```

