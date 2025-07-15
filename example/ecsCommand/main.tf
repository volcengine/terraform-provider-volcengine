resource "volcengine_ecs_command" "foo" {
  name            = "tf-test"
  description     = "tf"
  working_dir     = "/home"
  username        = "root"
  timeout         = 100
  command_content = base64encode("#!/bin/bash\n\n\necho \"{{ test_str }} {{ test_num }} operation success!\"")
  project_name    = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  enable_parameter = true
  parameter_definitions {
    name       = "test_str"
    type       = "String"
    required   = true
    min_length = 1
    max_length = 100
  }
  parameter_definitions {
    name      = "test_num"
    type      = "Digit"
    required  = false
    min_value = "-10"
    max_value = "100"
  }
}
