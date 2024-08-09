---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_resolver_rule"
sidebar_current: "docs-volcengine-resource-private_zone_resolver_rule"
description: |-
  Provides a resource to manage private zone resolver rule
---
# volcengine_private_zone_resolver_rule
Provides a resource to manage private zone resolver rule
## Example Usage
```hcl
resource "volcengine_private_zone_resolver_rule" "foo" {
  endpoint_id = 346
  name        = "tf0"
  type        = "OUTBOUND"
  vpcs {
    region = "cn-beijing"
    vpc_id = "vpc-13f9uuuqfdjb43n6nu5p1****"
  }
  forward_ips {
    ip   = "10.199.38.19"
    port = 33
  }
  zone_name = ["www.baidu.com"]
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) The name of the rule.
* `type` - (Required, ForceNew) Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
* `vpcs` - (Required) The parameter name <region> is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
* `endpoint_id` - (Optional, ForceNew) Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
* `forward_ips` - (Optional) IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
* `line` - (Optional) The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
* `zone_name` - (Optional, ForceNew) Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.

The `forward_ips` object supports the following:

* `ip` - (Required) IP address of the external DNS server. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
* `port` - (Optional) The port of the external DNS server. Default is 53. This parameter is only valid and optional when the Type parameter is OUTBOUND.

The `vpcs` object supports the following:

* `vpc_id` - (Required) The id of the bind vpc.
* `region` - (Optional) The region of the bind vpc. The default value is the region of the default provider config.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
PrivateZoneResolverRule can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_resolver_rule.default resource_id
```

