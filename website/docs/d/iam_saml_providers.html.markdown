---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_saml_providers"
sidebar_current: "docs-volcengine-datasource-iam_saml_providers"
description: |-
  Use this data source to query detailed information of iam saml providers
---
# volcengine_iam_saml_providers
Use this data source to query detailed information of iam saml providers
## Example Usage
```hcl
data "volcengine_iam_saml_providers" "foo" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `providers` - The collection of query.
    * `create_date` - Identity provider creation time, such as 20150123T123318Z.
    * `description` - The description of the SAML provider.
    * `encoded_saml_metadata_document` - Metadata document, encoded in Base64.
    * `saml_provider_name` - The name of the SAML provider.
    * `sso_type` - SSO types, 1. Role-based SSO, 2. User-based SSO.
    * `status` - User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
    * `trn` - The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.
    * `update_date` - Identity provider update time, such as: 20150123T123318Z.
* `total_count` - The total count of query.


