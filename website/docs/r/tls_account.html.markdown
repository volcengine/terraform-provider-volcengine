---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_account"
sidebar_current: "docs-volcengine-resource-tls_account"
description: |-
  Provides a resource to manage tls account
---
# volcengine_tls_account
Provides a resource to manage tls account
## Example Usage
```hcl
# 示例1：使用资源方式获取和管理 TLS 账号
# 资源创建会自动激活 TLS 账号（如果未激活）
resource "volcengine_tls_account" "example" {
}

# 输出资源结果
output "account_resource_arch_version" {
  value       = volcengine_tls_account.example.arch_version
  description = "日志服务版本：1.0（老架构）或 2.0（新架构）"
}

output "account_resource_status" {
  value       = volcengine_tls_account.example.status
  description = "日志服务状态：Activated（已开通）或 NonActivated（未开通）"
}
```
## Argument Reference
The following arguments are supported:


## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `arch_version` - The version of the log service architecture. Valid values: 2.0 (new architecture), 1.0 (old architecture).
* `status` - The status of the log service. Valid values: Activated (already activated), NonActivated (not activated).


## Import
The TlsAccount is not support import.

