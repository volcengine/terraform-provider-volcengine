---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_saml_provider"
sidebar_current: "docs-volcengine-resource-iam_saml_provider"
description: |-
  Provides a resource to manage iam saml provider
---
# volcengine_iam_saml_provider
Provides a resource to manage iam saml provider
## Example Usage
```hcl
resource "volcengine_iam_saml_provider" "foo" {
  encoded_saml_metadata_document = "your document"
  saml_provider_name             = "terraform"
  sso_type                       = 2
  status                         = 1
}
```
## Argument Reference
The following arguments are supported:
* `encoded_saml_metadata_document` - (Required) Metadata document, encoded in Base64.
* `saml_provider_name` - (Required, ForceNew) The name of the SAML provider.
* `sso_type` - (Required) SSO types, 1. Role-based SSO, 2. User-based SSO.
* `description` - (Optional) The description of the SAML provider.
* `status` - (Optional) User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_date` - Identity provider creation time, such as 20150123T123318Z.
* `trn` - The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.
* `update_date` - Identity provider update time, such as: 20150123T123318Z.


## Import
IamSamlProvider can be imported using the id, e.g.
```
$ terraform import volcengine_iam_saml_provider.default SAMLProviderName
```

