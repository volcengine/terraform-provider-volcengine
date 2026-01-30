data "volcengine_tls_alarm_content_templates" "foo" {
  # Filter by name (fuzzy matching)
  # alarm_content_template_name = "test-alarm"
  
  # Filter by specific IDs
  # ids = ["alarm-content-template-123456", "alarm-content-template-789012"]
}
