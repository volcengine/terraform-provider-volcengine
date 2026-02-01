---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_upstream_sources"
sidebar_current: "docs-volcengine-datasource-apig_upstream_sources"
description: |-
  Use this data source to query detailed information of apig upstream sources
---
# volcengine_apig_upstream_sources
Use this data source to query detailed information of apig upstream sources
## Example Usage
```hcl
data "volcengine_apig_upstream_sources" "foo" {
  gateway_id = "gd13d8c6eq1emkiunq6p0"
}
```
## Argument Reference
The following arguments are supported:
* `enable_ingress` - (Optional) The enable ingress of apig upstream source.
* `gateway_id` - (Optional) The id of api gateway.
* `name` - (Optional) The name of nacos source.
* `output_file` - (Optional) File name where to save data source results.
* `source_type` - (Optional) The source type of apig upstream source. Valid values: `K8S`, `Nacos`.
* `status` - (Optional) The status of apig upstream source. Valid values: `Syncing`, `SyncedSucceed`, `SyncedFailed`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `upstream_sources` - The collection of query.
    * `comments` - The comments of apig upstream source.
    * `create_time` - The create time of apig upstream source.
    * `gateway_id` - The id of api gateway.
    * `id` - The id of apig upstream source.
    * `ingress_settings` - The ingress settings of apig upstream source.
        * `enable_all_ingress_classes` - Whether to enable all ingress classes.
        * `enable_all_namespaces` - Whether to enable all namespaces.
        * `enable_ingress_without_ingress_class` - Whether to enable ingress without ingress class.
        * `enable_ingress` - Whether to enable ingress.
        * `ingress_classes` - The ingress classes of ingress settings.
        * `update_status` - The update status of ingress settings.
        * `watch_namespaces` - The watch namespaces of ingress settings.
    * `source_spec` - The source spec of apig upstream source.
        * `k8s_source` - The k8s source of apig upstream source.
            * `cluster_id` - The cluster id of k8s source.
            * `cluster_type` - The cluster type of k8s source.
        * `nacos_source` - The nacos source of apig upstream source.
            * `address` - The address of nacos source.
            * `auth_config` - The auth config of nacos source.
                * `basic` - The basic auth config of nacos source.
                    * `password` - The password of basic auth config.
                    * `username` - The username of basic auth config.
            * `context_path` - The context path of nacos source.
            * `grpc_port` - The grpc port of nacos source.
            * `http_port` - The http port of nacos source.
            * `nacos_id` - The nacos id of nacos source.
            * `nacos_name` - The nacos name of nacos source.
    * `source_type` - The source type of apig upstream source.
    * `status_message` - The status message of apig upstream source.
    * `status` - The status of apig upstream source.
    * `update_time` - The update time of apig upstream source.


