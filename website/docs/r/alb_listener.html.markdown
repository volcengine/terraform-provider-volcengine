---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_listener"
sidebar_current: "docs-volcengine-resource-alb_listener"
description: |-
  Provides a resource to manage alb listener
---
# volcengine_alb_listener
Provides a resource to manage alb listener
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_alb" "foo" {
  address_ip_version = "IPv4"
  type               = "private"
  load_balancer_name = "acc-test-alb-private"
  description        = "acc-test"
  subnet_ids         = [volcengine_subnet.foo.id]
  project_name       = "default"
  delete_protection  = "off"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_alb_server_group" "foo" {
  vpc_id            = volcengine_vpc.foo.id
  server_group_name = "acc-test-server-group"
  description       = "acc-test"
  server_group_type = "instance"
  scheduler         = "wlc"
  project_name      = "default"
  health_check {
    enabled  = "on"
    interval = 3
    timeout  = 3
    method   = "GET"
  }
  sticky_session_config {
    sticky_session_enabled = "on"
    sticky_session_type    = "insert"
    cookie_timeout         = "1100"
  }
}

resource "volcengine_alb_certificate" "foo" {
  description = "tf-test"
  public_key  = "public key"
  private_key = "private key"
}

resource "volcengine_alb_listener" "foo" {
  load_balancer_id   = volcengine_alb.foo.id
  listener_name      = "acc-test-listener"
  protocol           = "HTTPS"
  port               = 6666
  enabled            = "off"
  certificate_source = "alb"
  #  cert_center_certificate_id = "cert-***"
  certificate_id                               = volcengine_alb_certificate.foo.id
  server_group_id                              = volcengine_alb_server_group.foo.id
  description                                  = "acc test listener"
  access_log_record_customized_headers_enabled = "off"
  ca_certificate_source                        = "alb"
  ca_certificate_id                            = "cert-xoekc6lpu9s054ov5eo*****"
  domain_extensions {
    domain             = "example.com"
    certificate_source = "alb"
    certificate_id     = "cert-1pf4a8k8tokcg845wf******"
  }
  tags {
    key   = "key1"
    value = "value2"
  }
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The Id of the load balancer.
* `port` - (Required, ForceNew) The port receiving request of the Listener, the value range in 1~65535.
* `protocol` - (Required, ForceNew) The protocol of the Listener. Optional choice contains `HTTP`, `HTTPS`.
* `server_group_id` - (Required) The server group id associated with the listener.
* `access_log_record_customized_headers_enabled` - (Optional) Whether to enable custom headers in access logs. Default is `off`.
* `acl_ids` - (Optional) The id list of the Acl. When the AclStatus parameter is configured as on, AclType and AclIds.N are required.
* `acl_status` - (Optional) The enable status of Acl. Optional choice contains `on`, `off`. Default is `off`.
* `acl_type` - (Optional) The type of the Acl. Optional choice contains `white`, `black`. When the AclStatus parameter is configured as on, AclType and AclIds.N are required.
* `ca_certificate_id` - (Optional) The CA certificate id associated with the listener. When the value of ca_certificate_source is alb, the ca_certificate_id parameter must be specified.
* `ca_certificate_source` - (Optional) The source of the CA certificate associated with the listener. This parameter is only valid for HTTPS listeners and is used for two-way authentication. Valid values: `alb`, `pca_root`, `pca_sub`.
* `cert_center_certificate_id` - (Optional) The certificate id associated with the listener. Source is `cert_center`.
* `certificate_id` - (Optional) The certificate id associated with the listener. Source is `alb`.
* `certificate_source` - (Optional) The source of the certificate. Valid values: `alb`, `cert_center`. Default is `alb`.
* `customized_cfg_id` - (Optional) Personalized configuration ID, with a value of " " when not bound.
* `description` - (Optional) The description of the Listener.
* `domain_extensions` - (Optional) The domain extensions of the Listener. Only HTTPS listener is effective.
* `enable_http2` - (Optional) The HTTP2 feature switch,valid value is on or off. Default is `off`.
* `enable_quic` - (Optional) The QUIC feature switch,valid value is on or off. Default is `off`.
* `enabled` - (Optional) The enable status of the Listener. Optional choice contains `on`, `off`. Default is `on`.
* `listener_name` - (Optional) The name of the Listener.
* `pca_leaf_certificate_id` - (Optional) The CA certificate id associated with the listener. When the value of ca_certificate_source is pca_leaf, pca_leaf_certificate_id parameter must be specified.
* `pca_root_ca_certificate_id` - (Optional) The CA certificate id associated with the listener. When the value of ca_certificate_source is pca_root, pca_root_ca_certificate_id parameter must be specified.
* `pca_sub_ca_certificate_id` - (Optional) The CA certificate id associated with the listener. When the value of ca_certificate_source is pca_sub, pca_sub_ca_certificate_id parameter must be specified.
* `tags` - (Optional) Tags.

The `domain_extensions` object supports the following:

* `cert_center_certificate_id` - (Optional) The server certificate ID used by the domain name. Valid when the certificate_source is `cert_center`.
* `certificate_id` - (Optional) The server certificate ID used by the domain name. Valid when the certificate_source is `alb`.
* `certificate_source` - (Optional) The source of the certificate. Valid values: `alb`, `cert_center`.
* `domain_extension_id` - (Optional) The extended domain ID, required only for deletion and modification.
* `domain` - (Optional) The domain name.
* `pca_leaf_certificate_id` - (Optional) The server certificate ID used by the domain name. Valid when the certificate source is `pca_leaf`.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `listener_id` - The ID of the Listener.


## Import
AlbListener can be imported using the id, e.g.
```
$ terraform import volcengine_alb_listener.default lsn-273yv0mhs5xj47fap8sehiiso
```

