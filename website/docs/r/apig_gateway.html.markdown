---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_gateway"
sidebar_current: "docs-volcengine-resource-apig_gateway"
description: |-
  Provides a resource to manage apig gateway
---
# volcengine_apig_gateway
Provides a resource to manage apig gateway
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo1" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_subnet" "foo2" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.1.0/24"
  zone_id     = data.volcengine_zones.foo.zones[1].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_apig_gateway" "foo" {
  name         = "acc-test-apig"
  type         = "standard"
  comments     = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  network_spec {
    vpc_id     = volcengine_vpc.foo.id
    subnet_ids = [volcengine_subnet.foo1.id, volcengine_subnet.foo2.id]
  }
  resource_spec {
    replicas                    = 2
    instance_spec_code          = "1c2g"
    clb_spec_code               = "small_1"
    public_network_billing_type = "bandwidth"
    public_network_bandwidth    = 1
    network_type {
      enable_public_network  = true
      enable_private_network = true
    }
  }
  log_spec {
    enable     = true
    project_id = "d3cb87c0-faeb-4074-b1ee-9bd747865a76"
    topic_id   = "d339482e-d86d-4bd8-a9bb-f270417f00a1"
  }
  monitor_spec {
    enable       = true
    workspace_id = "4ed1caf3-279d-4c5f-8301-87ea38e92ffc"
  }
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The name of the api gateway.
* `network_spec` - (Required, ForceNew) The network spec of the api gateway.
* `backend_spec` - (Optional, ForceNew) The backend spec of the api gateway.
* `comments` - (Optional) The comments of the api gateway.
* `log_spec` - (Optional) The log spec of the api gateway.
* `monitor_spec` - (Optional, ForceNew) The monitor spec of the api gateway.
* `project_name` - (Optional, ForceNew) The project name of the api gateway.
* `resource_spec` - (Optional) The resource spec of the api gateway.
* `tags` - (Optional) Tags.
* `type` - (Optional, ForceNew) The type of the api gateway. Valid values: `standard`, `serverless`.

The `backend_spec` object supports the following:

* `is_vke_with_flannel_cni_supported` - (Required, ForceNew) Whether the api gateway support vke flannel cni.
* `vke_pod_cidr` - (Required, ForceNew) The vke pod cidr of the api gateway.

The `log_spec` object supports the following:

* `enable` - (Required) Whether the api gateway enable tls log.
* `project_id` - (Optional) The project id of the tls. This field is required when `enable` is true.
* `topic_id` - (Optional) The topic id of the tls.

The `monitor_spec` object supports the following:

* `enable` - (Required, ForceNew) Whether the api gateway enable monitor.
* `workspace_id` - (Optional, ForceNew) The workspace id of the monitor. This field is required when `enable` is true.

The `network_spec` object supports the following:

* `subnet_ids` - (Required, ForceNew) The subnet ids of the network spec.
* `vpc_id` - (Required, ForceNew) The vpc id of the network spec.

The `network_type` object supports the following:

* `enable_private_network` - (Required, ForceNew) Whether the api gateway enable private network.
* `enable_public_network` - (Required, ForceNew) Whether the api gateway enable public network.

The `resource_spec` object supports the following:

* `instance_spec_code` - (Required) The instance spec code of the resource spec. Valid values: `1c2g`, `2c4g`, `4c8g`, `8c16g`.
* `replicas` - (Required) The replicas of the resource spec.
* `clb_spec_code` - (Optional, ForceNew) The clb spec code of the resource spec. Valid values: `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
* `network_type` - (Optional) The network type of the resource spec. The default values for both `enable_public_network` and `enable_private_network` are true.
* `public_network_bandwidth` - (Optional, ForceNew) The public network bandwidth of the resource spec.
* `public_network_billing_type` - (Optional, ForceNew) The public network billing type of the resource spec. Valid values: `traffic`, `bandwidth`.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the api gateway.
* `message` - The error message of the api gateway.
* `status` - The status of the api gateway.
* `version` - The version of the api gateway.


## Import
ApigGateway can be imported using the id, e.g.
```
$ terraform import volcengine_apig_gateway.default resource_id
```

