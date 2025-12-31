# 创建一个 VMP 集成任务
resource "volcengine_vmp_integration_task" "foo1" {
  name        = "tf_test_integration_task3"
  type        = "CloudMonitor"
  environment = "Managed"
  workspace_id = "38046480-06a4-4d0e-bbd0-fed4eb358385"
  params      = jsonencode({
    JobName              = "test1"
    ScrapeInterval       = "1m"
    Regions = {
      cn-beijing = {
        VCM_ALB = {}
        VCM_ECS = {}
      }
    }
    MetaLabelWithNewMetric = false
  })
}

resource "volcengine_vmp_integration_task" "foo2" {
  name        = "tf_test_integration_task4"
  type        = "CloudMonitor"
  environment = "Managed"
  workspace_id = "38046480-06a4-4d0e-bbd0-fed4eb358385"
  params      = jsonencode({
    JobName              = "test1"
    ScrapeInterval       = "1m"
    Regions = {
      cn-beijing = {
        VCM_ALB = {}
        VCM_ECS = {}
      }
    }
    MetaLabelWithNewMetric = false
  })
}


# 查询 VMP 集成任务列表
data "volcengine_vmp_integration_tasks" "foo" {
  ids = [volcengine_vmp_integration_task.foo1.id,volcengine_vmp_integration_task.foo2.id]
}

output "integration_task_id" {
  value = join(",", [volcengine_vmp_integration_task.foo1.id, volcengine_vmp_integration_task.foo2.id])
}

output "integration_task_name" {
  value = join(",", [volcengine_vmp_integration_task.foo1.name, volcengine_vmp_integration_task.foo2.name])
}