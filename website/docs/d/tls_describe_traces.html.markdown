---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_describe_traces"
sidebar_current: "docs-volcengine-datasource-tls_describe_traces"
description: |-
  Use this data source to query detailed information of tls describe traces
---
# volcengine_tls_describe_traces
Use this data source to query detailed information of tls describe traces
## Example Usage
```hcl
data "volcengine_tls_describe_traces" "example" {
  trace_instance_id = "ac368174-2353-4e5d-859d-84c8bd255590"
  trace_id          = "b4de47438f9d052ca71106bb9bace34e"
}
```
## Argument Reference
The following arguments are supported:
* `trace_id` - (Required) Trace ID.
* `trace_instance_id` - (Required) Trace instance ID.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of tls trace query.
* `traces` - The collection of tls trace query.
    * `spans` - The collection of spans.
        * `attributes` - Span attributes.
            * `key` - Attribute key.
            * `value` - Attribute value.
        * `end_time` - Span end time.
        * `events` - Span events.
            * `attributes` - Event attributes.
                * `key` - Attribute key.
                * `value` - Attribute value.
            * `name` - Event name.
            * `timestamp` - Event timestamp.
        * `instrumentation_library` - Instrumentation library information.
            * `name` - Library name.
            * `version` - Library version.
        * `kind` - Span type.
        * `links` - Span links.
            * `attributes` - Link attributes.
                * `key` - Attribute key.
                * `value` - Attribute value.
            * `span_id` - Span ID.
            * `trace_id` - Trace ID.
            * `trace_state` - Trace state.
        * `name` - Span name.
        * `parent_span_id` - Parent Span ID.
        * `resource` - Resource information.
            * `attributes` - Resource attributes.
                * `key` - Attribute key.
                * `value` - Attribute value.
        * `span_id` - Span ID.
        * `start_time` - Span start time.
        * `status` - Span status.
            * `code` - Status code.
            * `message` - Error message.
        * `trace_id` - Trace ID.
        * `trace_state` - Trace state.
    * `trace_id` - Trace ID.


