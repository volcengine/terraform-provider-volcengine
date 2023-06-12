---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_zone"
sidebar_current: "docs-volcengine-resource-privatelink_vpc_endpoint_zone"
description: |-
  Provides a resource to manage privatelink vpc endpoint zone
---
# volcengine_privatelink_vpc_endpoint_zone
Provides a resource to manage privatelink vpc endpoint zone
## Example Usage
```hcl
resource "volcengine_privatelink_vpc_endpoint_zone" "foo" {
  endpoint_id        = "ep-2byz5nlkimc5c2dx0ef9g****"
  subnet_id          = "subnet-2bz47q19zhx4w2dx0eevn****"
  private_ip_address = "172.16.0.251"
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_id` - (Required, ForceNew) The endpoint id of vpc endpoint zone.
* `subnet_id` - (Required, ForceNew) The subnet id of vpc endpoint zone.
* `private_ip_address` - (Optional, ForceNew) The private ip address of vpc endpoint zone.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `network_interface_id` - The network interface id of vpc endpoint.
* `zone_domain` - The domain of vpc endpoint zone.
* `zone_id` - The Id of vpc endpoint zone.
* `zone_status` - The status of vpc endpoint zone.


## Import
VpcEndpointZone can be imported using the endpointId:subnetId, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_zone.default ep-3rel75r081l345zsk2i59****:subnet-2bz47q19zhx4w2dx0eevn****
```

