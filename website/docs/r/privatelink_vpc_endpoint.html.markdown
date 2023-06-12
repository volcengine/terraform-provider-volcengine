---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint"
sidebar_current: "docs-volcengine-resource-privatelink_vpc_endpoint"
description: |-
  Provides a resource to manage privatelink vpc endpoint
---
# volcengine_privatelink_vpc_endpoint
Provides a resource to manage privatelink vpc endpoint
## Example Usage
```hcl
resource "volcengine_privatelink_vpc_endpoint" "endpoint" {
  security_group_ids = ["sg-2d5z8cr53k45c58ozfdum****"]
  service_id         = "epsvc-2byz5nzgiansw2dx0eehh****"
  endpoint_name      = "tf-test-ep"
  description        = "tf-test"
}

resource "volcengine_privatelink_vpc_endpoint_zone" "zone" {
  endpoint_id        = volcengine_privatelink_vpc_endpoint.endpoint.id
  subnet_id          = "subnet-2bz47q19zhx4w2dx0eevn****"
  private_ip_address = "172.16.0.252"
}
```
## Argument Reference
The following arguments are supported:
* `security_group_ids` - (Required, ForceNew) the security group ids of vpc endpoint.
* `service_id` - (Required, ForceNew) The id of vpc endpoint service.
* `description` - (Optional) The description of vpc endpoint.
* `endpoint_name` - (Optional) The name of vpc endpoint.
* `service_name` - (Optional, ForceNew) The name of vpc endpoint service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `business_status` - Whether the vpc endpoint is locked.
* `connection_status` - The connection  status of vpc endpoint.
* `creation_time` - The create time of vpc endpoint.
* `deleted_time` - The delete time of vpc endpoint.
* `endpoint_domain` - The domain of vpc endpoint.
* `endpoint_type` - The type of vpc endpoint.
* `status` - The status of vpc endpoint.
* `update_time` - The update time of vpc endpoint.
* `vpc_id` - The vpc id of vpc endpoint.


## Import
VpcEndpoint can be imported using the id, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint.default ep-3rel74u229dz45zsk2i6l****
```

