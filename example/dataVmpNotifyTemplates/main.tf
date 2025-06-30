resource "volcengine_vmp_notify_template" "foo" {
  name        = "acc-test-vmp-notify-template"
  description = "acc-test-vmp"
  channel     = "WeComBotWebhook"
  active {
    title   = "acc-test-active-template-title"
    content = "acc-test-active-template-content"
  }
  resolved {
    title   = "acc-test-resolved-template-title"
    content = "acc-test-resolved-template-content"
  }
}

data "volcengine_vmp_notify_templates" "default" {
  ids = [volcengine_vmp_notify_template.foo.id]
}