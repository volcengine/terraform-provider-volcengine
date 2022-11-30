---
subcategory: "VEENEDGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_veenedge_cloud_server"
sidebar_current: "docs-volcengine-resource-veenedge_cloud_server"
description: |-
  Provides a resource to manage veenedge cloud server
---
# volcengine_veenedge_cloud_server
Provides a resource to manage veenedge cloud server
## Example Usage
```hcl
resource "volcengine_veenedge_cloud_server" "foo" {
  image_id          = "image*****viqm"
  cloudserver_name  = "tf-test"
  spec_name         = "veEN****rge"
  server_area_level = "region"
  secret_type       = "KeyPair"
  secret_data       = "sshkey-47*****wgc"
  network_config {
    bandwidth_peak = 5
  }
  schedule_strategy {
    schedule_strategy = "dispersion"
    price_strategy    = "high_priority"
    network_strategy  = "region"
  }
  billing_config {
    computing_billing_method = "MonthlyPeak"
    bandwidth_billing_method = "MonthlyP95"
  }
  storage_config {
    system_disk {
      storage_type = "CloudBlockSSD"
      capacity     = 40
    }
    data_disk_list {
      storage_type = "CloudBlockSSD"
      capacity     = 20
    }
  }
  default_area_name = "C******na"
  default_isp       = "CMCC"
}
```
## Argument Reference
The following arguments are supported:
* `cloudserver_name` - (Required, ForceNew) The name of cloud server.
* `default_area_name` - (Required) The name of default area.
* `default_isp` - (Required) The default isp info.
* `image_id` - (Required, ForceNew) The image id of cloud server.
* `network_config` - (Required, ForceNew) The config of the network.
* `schedule_strategy` - (Required, ForceNew) The schedule strategy.
* `secret_type` - (Required, ForceNew) The type of secret. The value can be `KeyPair` or `Password`.
* `server_area_level` - (Required, ForceNew) The server area level. The value can be `region` or `city`.
* `spec_name` - (Required, ForceNew) The spec name of cloud server.
* `storage_config` - (Required, ForceNew) The config of the storage.
* `billing_config` - (Optional, ForceNew) The config of the billing.
* `custom_data` - (Optional, ForceNew) The custom data.
* `default_cluster_name` - (Optional) The name of default cluster.
* `secret_data` - (Optional, ForceNew) The data of secret. The value can be Password or KeyPair ID.

The `billing_config` object supports the following:

* `bandwidth_billing_method` - (Required, ForceNew) The method of bandwidth billing. The value can be `MonthlyP95` or `DailyPeak`.
* `computing_billing_method` - (Required, ForceNew) The method of computing billing. The value can be `MonthlyPeak` or `DailyPeak`.

The `custom_data` object supports the following:

* `data` - (Required, ForceNew) The custom data info.

The `data_disk_list` object supports the following:

* `capacity` - (Required, ForceNew) The capacity of storage.
* `storage_type` - (Required, ForceNew) The type of storage. The value can be `CloudBlockHDD` or `CloudBlockSSD`.

The `network_config` object supports the following:

* `bandwidth_peak` - (Required, ForceNew) The peak of bandwidth.
* `custom_external_interface_name` - (Optional, ForceNew) The name of custom external interface.
* `custom_internal_interface_name` - (Optional, ForceNew) The name of custom internal interface.
* `enable_ipv6` - (Optional, ForceNew) Whether enable ipv6.
* `internal_bandwidth_peak` - (Optional, ForceNew) The internal peak of bandwidth.

The `schedule_strategy` object supports the following:

* `network_strategy` - (Required, ForceNew) The network strategy.
* `price_strategy` - (Required, ForceNew) The price strategy. The value can be `high_priority` or `low_priority`.
* `schedule_strategy` - (Required, ForceNew) The type of schedule strategy. The value can be `dispersion` or `concentration`.

The `storage_config` object supports the following:

* `system_disk` - (Required, ForceNew) The disk info of system.
* `data_disk_list` - (Optional, ForceNew) The disk list info of data.

The `system_disk` object supports the following:

* `capacity` - (Required, ForceNew) The capacity of storage.
* `storage_type` - (Required, ForceNew) The type of storage. The value can be `CloudBlockHDD` or `CloudBlockSSD`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `default_instance_id` - The default instance id generate by cloud server.


## Import
CloudServer can be imported using the id, e.g.
```
$ terraform import volcengine_veenedge_cloud_server.default cloudserver-n769ewmjjqyqh5dv
```

After the veenedge cloud server is created, a default edge instance will be created, we recommend managing this default instance as follows
```
resource "volcengine_veenedge_instance" "foo1" {
  instance_id = volcengine_veenedge_cloud_server.foo.default_instance_id
}
```

