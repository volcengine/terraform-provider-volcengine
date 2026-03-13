resource "volcengine_nlb_backend_servers" "foo" {
  server_group_id = "rsp-3rezryqxrtatc5zsk2ieoy3t6"
  backend_servers {
    type        = "ecs"
    instance_id = "i-yegwcgr1fkr9cxxxxxxx"
    ip          = "your-ip"
    port        = 80
    weight      = 30
    description = "nlb server test by tf"
    zone_id     = "cn-guilin-a"
  }
      backend_servers {
        type        = "ecs"
        instance_id = "i-yegwcgr1fkr9cxxxxxxx"
        ip          = "your-ip"
        port        = 80
        weight      = 10
        description = "nlb server test 2 by tf"
        zone_id     = "cn-guilin-a"
      }
}
