resource "volcengine_cdn_shared_config" "foo" {
    config_name = "tftest"
    config_type = "allow_ip_access_rule"
    allow_ip_access_rule {
        rules = ["1.1.1.1", "2.2.2.0/24", "3.3.3.3"]
    }
    deny_ip_access_rule {
        rules = ["1.1.1.1", "2.2.2.0/24"]
    }
    common_match_list {
        common_type {
            rules = ["1.1.1.1", "2.2.2.0/24"]
        }
    }
    allow_referer_access_rule {
        common_type {
            rules = ["1.1.1.1", "2.2.2.0/24"]
        }
    }
    deny_referer_access_rule {
        common_type {
            rules = ["1.1.1.1", "2.2.2.0/24"]
        }
    }
}