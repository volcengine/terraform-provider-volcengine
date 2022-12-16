data "volcengine_vke_kubeconfigs" "default"{
  cluster_ids = ["cce7hb97qtofmj1oi4udg"]
  types = ["Private", "Public"]
}