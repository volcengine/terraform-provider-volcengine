resource "volcengine_ecs_command" "foo" {
  name = "tf-test"
  description = "tf"
  working_dir = "/home"
  username = "root"
  timeout = 100
  command_content = "IyEvYmluL2Jhc2gKCgplY2hvICJvcGVyYXRpb24gc3VjY2VzcyEi"
}