
# 示例1：使用资源方式获取和管理 TLS 账号
# 资源创建会自动激活 TLS 账号（如果未激活）
resource "volcengine_tls_account" "example" {
}

# 输出资源结果
output "account_resource_arch_version" {
  value = volcengine_tls_account.example.arch_version
  description = "日志服务版本：1.0（老架构）或 2.0（新架构）"
}

output "account_resource_status" {
  value = volcengine_tls_account.example.status
  description = "日志服务状态：Activated（已开通）或 NonActivated（未开通）"
}
