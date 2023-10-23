resource "volcengine_transit_router_vpc_attachment" "foo" {
  transit_router_id = "tr-2d6fr7f39unsw58ozfe1ow21x"
  vpc_id            = "vpc-2bysvq1xx543k2dx0eeulpeiv"
  attach_points {
    subnet_id = "subnet-3refsrxdswsn45zsk2hmdg4zx"
    zone_id   = "cn-beijing-a"
  }
  attach_points {
    subnet_id = "subnet-2d68bh74345q858ozfekrm8fj"
    zone_id   = "cn-beijing-a"
  }
  transit_router_attachment_name = "tfname1"
  description = "desc"
}