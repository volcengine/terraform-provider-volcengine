data "volcengine_scaling_instances" "default" {
  scaling_group_id = "scg-ybtawtznszgh9yv8agcp"
  ids = ["i-ybzl4u5uogl8j07hgcbg", "i-ybyncrcpzpgh9zmlct0w", "i-ybyncrcpzogh9z4ax9tv"]
  scaling_configuration_id = "scc-ybtawzucw95pkgon0wst"
  status = "InService"
}