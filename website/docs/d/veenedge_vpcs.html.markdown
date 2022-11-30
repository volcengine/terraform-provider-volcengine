---
subcategory: "VEENEDGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_veenedge_vpcs"
sidebar_current: "docs-volcengine-datasource-veenedge_vpcs"
description: |-
  Use this data source to query detailed information of veenedge vpcs
---
# volcengine_veenedge_vpcs
Use this data source to query detailed information of veenedge vpcs
## Example Usage
```hcl
data "volcengine_veenedge_vpcs" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of vpc IDs.
* `name_regex` - (Optional) A Name Regex of Vpc.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of Vpc query.
* `vpc_instances` - The collection of Vpc query.
    * `account_identity` - The id of account.
    * `cluster_vpc_id` - The cluster vpc id.
    * `cluster` - The cluster info.
        * `alias` - The alias of cluster.
        * `city` - The city of cluster.
        * `cluster_name` - The name of cluster.
        * `country` - The country of cluster.
        * `isp` - The isp of cluster.
        * `level` - The level of cluster.
        * `province` - The province of cluster.
        * `region` - The region of cluster.
    * `create_time` - The create time of VPC.
    * `desc` - The description of VPC.
    * `id` - The ID of VPC.
    * `is_default` - Is default vpc.
    * `resource_statistic` - The resource statistic info.
        * `veen_instance_count` - The count of instance.
        * `veew_lb_instance_count` - The count of load balancers.
        * `veew_sg_instance_count` - The count of security groups.
    * `status` - The status of VPC.
    * `sub_nets` - The subnets info.
        * `account_identity` - The account id.
        * `cidr_ip` - The ip of cidr.
        * `cidr_mask` - The mask of cidr.
        * `create_time` - The creation time.
        * `subnet_identity` - The id of subnet.
        * `update_time` - The update time.
        * `user_identity` - The id of user.
    * `update_time` - The update time of VPC.
    * `user_identity` - The id of user.
    * `vpc_identity` - The ID of VPC.
    * `vpc_name` - The name of VPC.
    * `vpc_ns` - The namespace of VPC.


