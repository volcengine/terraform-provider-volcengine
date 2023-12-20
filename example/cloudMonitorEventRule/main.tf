resource "volcengine_cloud_monitor_event_rule" "foo" {
    status = "enable"
    contact_methods = ["Phone", "TLS", "MQ"]
    event_source = "ecs"
    level = "notice"
    rule_name = "tftest1"
    effective_time {
        start_time = "01:00"
        end_time = "22:00"
    }
    event_type = ["ecs:Disk:DiskError.Redeploy.Canceled"]
    contact_group_ids = ["1737941730782699520", "1737940985502777344"]
    filter_pattern {
        type = ["ecs:Disk:DiskError.Redeploy.Canceled"]
        source = "ecs"
    }
    message_queue {
        instance_id = "kafka-cnoe4rfrsqfb1d64"
        vpc_id = "vpc-2d68hz41j7qio58ozfd6jxgtb"
        type = "kafka"
        region = "*****"
        topic = "tftest"
    }
    tls_target {
        project_name = "tf-test"
        region_name_cn = "*****"
        region_name_en = "*****"
        project_id = "17ba378d-de43-495e-8906-03ae6567b376"
        topic_id = "7ce12237-6670-44a7-9d79-2e36961586e6"
    }
}
