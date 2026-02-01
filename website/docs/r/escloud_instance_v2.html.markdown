---
subcategory: "ESCLOUD"
layout: "volcengine"
page_title: "Volcengine: volcengine_escloud_instance_v2"
sidebar_current: "docs-volcengine-resource-escloud_instance_v2"
description: |-
  Provides a resource to manage escloud instance v2
---
# volcengine_escloud_instance_v2
Provides a resource to manage escloud instance v2
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
## Example Usage
```hcl
# query available zones in current region
data "volcengine_zones" "foo" {
}

# create vpc
resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
  dns_servers  = ["8.8.8.8", "114.114.114.114"]
  project_name = "default"
}

# create subnet
resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

# create escloud instance
resource "volcengine_escloud_instance_v2" "foo" {
  instance_name       = "acc-test-escloud-instance"
  version             = "V7_10"
  zone_ids            = [data.volcengine_zones.foo.zones[0].id, data.volcengine_zones.foo.zones[1].id, data.volcengine_zones.foo.zones[2].id]
  subnet_id           = volcengine_subnet.foo.id
  enable_https        = false
  admin_password      = "Password@@123"
  charge_type         = "PostPaid"
  auto_renew          = false
  period              = 1
  configuration_code  = "es.standard"
  enable_pure_master  = true
  deletion_protection = false
  project_name        = "default"

  node_specs_assigns {
    type               = "Master"
    number             = 3
    resource_spec_name = "es.x2.medium"
    storage_spec_name  = "es.volume.essd.pl0"
    storage_size       = 20
  }
  node_specs_assigns {
    type               = "Hot"
    number             = 6
    resource_spec_name = "es.x2.medium"
    storage_spec_name  = "es.volume.essd.flexpl-standard"
    storage_size       = 500
    extra_performance {
      throughput = 65
    }
  }
  node_specs_assigns {
    type               = "Kibana"
    number             = 1
    resource_spec_name = "kibana.x2.small"
    storage_spec_name  = ""
    storage_size       = 0
  }

  network_specs {
    type      = "Elasticsearch"
    bandwidth = 1
    is_open   = true
    spec_name = "es.eip.bgp_fixed_bandwidth"
  }

  network_specs {
    type      = "Kibana"
    bandwidth = 1
    is_open   = true
    spec_name = "es.eip.bgp_fixed_bandwidth"
  }

  tags {
    key   = "k1"
    value = "v1"
  }
}

# create escloud ip white list
resource "volcengine_escloud_ip_white_list" "foo" {
  instance_id = volcengine_escloud_instance_v2.foo.id
  type        = "public"
  component   = "es"
  ip_list     = ["172.16.0.10", "172.16.0.11", "172.16.0.12"]
}
```
## Argument Reference
The following arguments are supported:
* `admin_password` - (Required) The password of administrator account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `charge_type` - (Required) The charge type of ESCloud instance, valid values: `PostPaid`, `PrePaid`.
* `configuration_code` - (Required, ForceNew) Configuration code used for billing.
* `enable_https` - (Required, ForceNew) Whether Https access is enabled.
* `instance_name` - (Required) The name of ESCloud instance.
* `node_specs_assigns` - (Required) The number and configuration of various ESCloud instance node. Kibana NodeSpecsAssign should not be modified.
* `subnet_id` - (Required, ForceNew) The id of subnet, the subnet must belong to the AZ selected.
* `version` - (Required, ForceNew) The version of instance. When creating ESCloud instance, the valid value is `V6_7` or `V7_10`. When creating OpenSearch instance, the valid value is `OPEN_SEARCH_2_9`.
* `zone_ids` - (Required, ForceNew) The zone id of the ESCloud instance. Support specifying multiple availability zones.
 The first zone id is the primary availability zone, while the rest are backup availability zones.
* `auto_renew` - (Optional) Whether to automatically renew in prepaid scenarios. Default is false.
* `deletion_protection` - (Optional) Whether enable deletion protection for ESCloud instance. Default is false.
* `enable_pure_master` - (Optional, ForceNew) Whether the Master node is independent.
* `maintenance_day` - (Optional) The maintainable day for the instance. Valid values: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`. Works only on modified scenes.
* `maintenance_time` - (Optional) The maintainable time period for the instance. Works only on modified scenes.
* `network_specs` - (Optional, ForceNew) The public network config of the ESCloud instance.
* `period` - (Optional) Purchase duration in prepaid scenarios. Unit: Monthly.
* `project_name` - (Optional) The project name to which the ESCloud instance belongs.
* `tags` - (Optional) Tags.

The `extra_performance` object supports the following:

* `throughput` - (Required) When your data node chooses to use FlexPL storage type and the storage specification configuration is 500GiB or above, it supports purchasing bandwidth packages to increase disk bandwidth.
The unit is MiB, and the adjustment step size is 10MiB.

The `network_specs` object supports the following:

* `bandwidth` - (Required, ForceNew) The bandwidth of the eip. Unit: Mbps.
* `is_open` - (Required, ForceNew) Whether the eip is opened.
* `spec_name` - (Required, ForceNew) The spec name of public network.
* `type` - (Required, ForceNew) The type of public network, valid values: `Elasticsearch`, `Kibana`.

The `node_specs_assigns` object supports the following:

* `number` - (Required) The number of node.
* `resource_spec_name` - (Required) The name of compute resource spec.
* `storage_size` - (Required) The size of storage. Unit: GiB. the adjustment step size is 10GiB. Default is 100 GiB. Kibana NodeSpecsAssign should specify this field to 0.
* `storage_spec_name` - (Required) The name of storage spec. Kibana NodeSpecsAssign should specify this field to ``.
* `type` - (Required) The type of node, valid values: `Master`, `Hot`, `Cold`, `Warm`, `Kibana`, `Coordinator`.
* `extra_performance` - (Optional) The extra performance of FlexPL storage spec.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `cerebro_private_domain` - The cerebro private domain of instance.
* `cerebro_public_domain` - The cerebro public domain of instance.
* `es_eip_id` - The eip id associated with the instance.
* `es_eip` - The eip address of instance.
* `es_private_domain` - The es private domain of instance.
* `es_private_endpoint` - The es private endpoint of instance.
* `es_private_ip_whitelist` - The whitelist of es private ip.
* `es_public_domain` - The es public domain of instance.
* `es_public_endpoint` - The es public endpoint of instance.
* `es_public_ip_whitelist` - The whitelist of es public ip.
* `kibana_eip_id` - The eip id associated with kibana.
* `kibana_eip` - The eip address of kibana.
* `kibana_private_domain` - The kibana private domain of instance.
* `kibana_private_ip_whitelist` - The whitelist of kibana private ip.
* `kibana_public_domain` - The kibana public domain of instance.
* `kibana_public_ip_whitelist` - The whitelist of kibana public ip.
* `main_zone_id` - The main zone id of instance.
* `status` - The status of instance.


## Import
EscloudInstanceV2 can be imported using the id, e.g.
```
$ terraform import volcengine_escloud_instance_v2.default resource_id
```

