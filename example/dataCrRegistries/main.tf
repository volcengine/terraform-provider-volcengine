data "volcengine_cr_registries" "foo"{
  # names=["tf-1"]
  # types=["Enterprise"]
  statuses {
      phase="Running"
      conditions=["Ok"]
    }
}