resource "volcengine_vke_default_node_pool" "default" {
    cluster_id = "ccbpngpnqtofrqpms2kcg"
#    instances {
#        instance_id = "i-ybvwvaswar8rx7rxcwrd"
#    }
    instances {
        instance_id = "i-ybvwvaswas8rx7n0fow5"
        keep_instance_name = true
    }
    node_config {
        security {
            login {
                password = "amw4WTdVcTRJVVFsUXpVTw=="
            }
        }
    }

}