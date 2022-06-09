data "vestack_scaling_activities" "default"{
  scaling_group_id = "scg-ybmf7aaabe72q1xm64rv"
  ids = ["sga-yboqhv8pub72q1ovhbi3","sga-yboqf9lh7v72q1thez6r"]
}

output "data" {
  value = data.vestack_scaling_activities.default
}