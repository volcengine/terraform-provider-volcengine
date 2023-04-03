resource "volcengine_ecs_instance_state" "foo" {
  instance_id = "i-ycc01lmwecgh9z3sqqfl"
  action = "ForceStop"
  stopped_mode = "KeepCharging"
  //stopped_mode = "StopCharging"
}