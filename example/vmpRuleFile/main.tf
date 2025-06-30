resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "**********"
}

resource "volcengine_vmp_rule_file" "foo" {
  name         = "acc-test-1"
  workspace_id = volcengine_vmp_workspace.foo.id
  description  = "acc-test-1"
  content      = <<EOF
groups:
    - interval: 10s
      name: recording_rules
      rules:
        - expr: sum(irate(container_cpu_usage_seconds_total{image!=""}[5m])) by (pod) *100
          labels:
            team: operations
          record: pod:cpu:useage
EOF
}