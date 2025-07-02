resource "volcengine_traffic_mirror_filter" "foo" {
  traffic_mirror_filter_name = "acc-test-traffic-mirror-filter"
  description                = "acc-test"
  project_name               = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}