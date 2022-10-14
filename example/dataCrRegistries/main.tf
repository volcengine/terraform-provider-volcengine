data "volcengine_cr_registries" "foo"{
  # names=["liaoliuqing-prune-test"]
  # types=["Enterprise"]
  statuses {
      phase="Running"
      condition="Ok"
    }
}