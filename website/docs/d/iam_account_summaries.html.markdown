---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_account_summaries"
sidebar_current: "docs-volcengine-datasource-iam_account_summaries"
description: |-
  Use this data source to query detailed information of iam account summaries
---
# volcengine_iam_account_summaries
Use this data source to query detailed information of iam account summaries
## Example Usage
```hcl
data "volcengine_iam_account_summaries" "default" {}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `account_summaries` - The collection of account summaries.
    * `access_keys_per_account_quota` - The quota of access keys per account.
    * `access_keys_per_user_quota` - The quota of access keys per user.
    * `attached_policies_per_group_quota` - The quota of attached policies per group.
    * `attached_policies_per_role_quota` - The quota of attached policies per role.
    * `attached_policies_per_user_quota` - The quota of attached policies per user.
    * `attached_system_policies_per_group_quota` - The quota of attached system policies per group.
    * `attached_system_policies_per_role_quota` - The quota of attached system policies per role.
    * `attached_system_policies_per_user_quota` - The quota of attached system policies per user.
    * `groups_per_user_quota` - The quota of groups per user.
    * `groups_quota` - The quota of groups.
    * `groups_usage` - The usage of groups.
    * `policies_quota` - The quota of policies.
    * `policies_usage` - The usage of policies.
    * `policy_size` - The size of policy.
    * `roles_quota` - The quota of roles.
    * `roles_usage` - The usage of roles.
    * `users_quota` - The quota of users.
    * `users_usage` - The usage of users.
* `total_count` - The total count of query.


