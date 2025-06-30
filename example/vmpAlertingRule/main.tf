resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "Pass123456"
}

resource "volcengine_vmp_contact" "foo" {
  name  = "acc-test-contact"
  email = "acctest1@tftest.com"

  webhook {
    address = "https://www.acctest1.com"
  }

  lark_bot_webhook {
    address = "https://www.acctest1.com"
  }

  ding_talk_bot_webhook {
    address    = "https://www.dingacctest1.com"
    at_mobiles = ["18046891812"]
  }
  phone_number {
    country_code = "+86"
    number       = "18310101010"
  }
}

resource "volcengine_vmp_contact" "foo1" {
  name  = "acc-test-contact2"
  email = "acctest2@tftest.com"

  webhook {
    address = "https://www.acctest2.com"
  }

  lark_bot_webhook {
    address = "https://www.acctest2.com"
  }

  ding_talk_bot_webhook {
    address    = "https://www.dingacctest2.com"
    at_mobiles = ["18046891813"]
  }
  phone_number {
    country_code = "+86"
    number       = "18310101011"
  }
}

resource "volcengine_vmp_contact_group" "foo" {
  name        = "acc-test"
  contact_ids = [volcengine_vmp_contact.foo.id]
}

resource "volcengine_vmp_contact_group" "foo1" {
  name        = "acc-test-1"
  contact_ids = [volcengine_vmp_contact.foo1.id]
}

resource "volcengine_vmp_notify_policy" "foo" {
  name        = "acc-test-1"
  description = "acc-test-1"
  levels {
    level             = "P1"
    contact_group_ids = [volcengine_vmp_contact_group.foo.id]
    channels          = ["Email", "Webhook"]
  }
  levels {
    level             = "P0"
    contact_group_ids = [volcengine_vmp_contact_group.foo1.id]
    channels          = ["LarkBotWebhook"]
  }
}

resource "volcengine_vmp_notify_group_policy" "foo" {
  name        = "acc-test-1"
  description = "acc-test-1"
  levels {
    level           = "P2"
    group_by        = ["__rule__"]
    group_wait      = "35"
    group_interval  = "60"
    repeat_interval = "70"
  }
  levels {
    level           = "P0"
    group_by        = ["__rule__"]
    group_wait      = "30"
    group_interval  = "60"
    repeat_interval = "70"
  }
  levels {
    level           = "P1"
    group_by        = ["__rule__"]
    group_wait      = "40"
    group_interval  = "75"
    repeat_interval = "75"
  }
}

resource "volcengine_vmp_alerting_rule" "foo" {
  name                   = "acc-test-1"
  description            = "acc-test-1"
  notify_policy_id       = volcengine_vmp_notify_policy.foo.id
  notify_group_policy_id = volcengine_vmp_notify_group_policy.foo.id
  query {
    workspace_id = volcengine_vmp_workspace.foo.id
    prom_ql      = "sum(up)"
  }
  levels {
    level      = "P0"
    for        = "0s"
    comparator = ">="
    threshold  = 2.0
  }
  levels {
    level      = "P1"
    for        = "0s"
    comparator = ">="
    threshold  = 1.0
  }
  levels {
    level      = "P2"
    for        = "0s"
    comparator = ">="
    threshold  = 0.5
  }
  annotations {
    name  = "annotation"
    value = "acc-test"
  }
  labels {
    name  = "label"
    value = "acc-test"
  }
}
