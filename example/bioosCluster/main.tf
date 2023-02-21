resource "volcengine_bioos_cluster" "foo" {
  name = "test-cluster" //必填
  description = "test-description" //选填
#  vke_config { //选填，和shared_config二者中必填一个
#    cluster_id = "ccerdh8fqtofh16uf6q60" //也可替换成volcengine_vke_cluster.example.id
#    storage_class = "ebs-ssd"
#  }
  shared_config {
    enable = true
  }
}