---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_access_key"
sidebar_current: "docs-volcengine-resource-iam_access_key"
description: |-
  Provides a resource to manage iam access key
---
# volcengine_iam_access_key
Provides a resource to manage iam access key
## Example Usage
```hcl
resource "volcengine_iam_user" "foo" {
  user_name    = "acc-test-user"
  description  = "acc-test"
  display_name = "name"
}

resource "volcengine_iam_access_key" "foo" {
  user_name   = volcengine_iam_user.foo.user_name
  secret_file = "./sk"
  status      = "active"
  #  pgp_key = "keybase:some_person_that_exists"
}
```
## Argument Reference
The following arguments are supported:
* `pgp_key` - (Optional, ForceNew) Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`.
* `secret_file` - (Optional, ForceNew) The file to save the access id and secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.
* `status` - (Optional) The status of the access key, Optional choice contains `active` or `inactive`.
* `user_name` - (Optional, ForceNew) The user name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_date` - The create date of the access key.
* `encrypted_secret` - The encrypted secret of the access key by pgp key, base64 encoded.
* `key_fingerprint` - The key fingerprint of the encrypted secret.
* `secret` - The secret of the access key.


## Import
Iam access key don't support import

