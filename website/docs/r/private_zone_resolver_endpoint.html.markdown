---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_resolver_endpoint"
sidebar_current: "docs-volcengine-resource-private_zone_resolver_endpoint"
description: |-
  Provides a resource to manage private zone resolver endpoint
---
# volcengine_private_zone_resolver_endpoint
Provides a resource to manage private zone resolver endpoint
## Example Usage
```hcl
resource "volcengine_private_zone_resolver_endpoint" "foo" {
  name         = "tf-test"
  vpc_id       = "vpc-13f9uuuqfdjb43n6nu5p1****"
  vpc_region   = "cn-beijing"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  ip_configs {
    az_id     = "cn-beijing-a"
    subnet_id = "subnet-mj2o4co2m2v45smt1bx1****"
    ip        = "172.16.0.2"
  }
  ip_configs {
    az_id     = "cn-beijing-a"
    subnet_id = "subnet-mj2o4co2m2v45smt1bx1****"
    ip        = "172.16.0.3"
  }
  ip_configs {
    az_id     = "cn-beijing-a"
    subnet_id = "subnet-mj2o4co2m2v45smt1bx1****"
    ip        = "172.16.0.4"
  }
  ip_configs {
    az_id     = "cn-beijing-a"
    subnet_id = "subnet-mj2o4co2m2v45smt1bx1****"
    ip        = "172.16.0.5"
  }
}
```
## Argument Reference
The following arguments are supported:
* `ip_configs` - (Required) Availability zones, subnets, and IP configurations of terminal nodes.
* `name` - (Required) The name of the private zone resolver endpoint.
* `vpc_id` - (Required, ForceNew) The VPC ID of the endpoint.
* `vpc_region` - (Required, ForceNew) The VPC region of the endpoint.
* `direction` - (Optional, ForceNew) DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
* `project_name` - (Optional) The project name of the private zone resolver endpoint.
* `security_group_id` - (Optional, ForceNew) The security group ID of the endpoint.
* `tags` - (Optional) Tags.
* `vpc_trns` - (Optional) The vpc trns of the private zone resolver endpoint. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.

The `ip_configs` object supports the following:

* `az_id` - (Required) Id of the availability zone.
* `ip` - (Required) Source IP address of traffic. You can add up to 6 IP addresses at most. To ensure high availability, you must add at least two IP addresses.
* `subnet_id` - (Required) Id of the subnet.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
PrivateZoneResolverEndpoint can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_resolver_endpoint.default resource_id
```

