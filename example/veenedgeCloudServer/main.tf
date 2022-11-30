resource "volcengine_veenedge_cloud_server" "foo" {
  image_id = "image*****viqm"
  cloudserver_name = "tf-test"
  spec_name = "veEN****rge"
  server_area_level = "region"
  secret_type = "KeyPair"
  secret_data = "sshkey-47*****wgc"
  network_config {
      bandwidth_peak = 5
  }
  schedule_strategy {
      schedule_strategy = "dispersion"
      price_strategy = "high_priority"
      network_strategy = "region"
  }
  billing_config {
      computing_billing_method = "MonthlyPeak"
      bandwidth_billing_method = "MonthlyP95"
  }
  storage_config {
      system_disk {
         storage_type = "CloudBlockSSD"
         capacity = 40
      }
      data_disk_list {
         storage_type = "CloudBlockSSD"
         capacity = 20
      }
  }
  default_area_name = "C******na"
  default_isp = "CMCC"
}