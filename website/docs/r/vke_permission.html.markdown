---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_permission"
sidebar_current: "docs-volcengine-resource-vke_permission"
description: |-
  Provides a resource to manage vke permission
---
# volcengine_vke_permission
Provides a resource to manage vke permission
## Example Usage
```hcl
# query available zones in current region
data "volcengine_zones" "foo" {
}

# create vpc
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

# create subnet
resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

# create security group
resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

# create vke cluster
resource "volcengine_vke_cluster" "foo" {
  name                      = "acc-test-1"
  description               = "created by terraform"
  project_name              = "default"
  delete_protection_enabled = false
  cluster_config {
    subnet_ids                       = [volcengine_subnet.foo.id]
    api_server_public_access_enabled = true
    api_server_public_access_config {
      public_access_network_config {
        billing_type = "PostPaidByBandwidth"
        bandwidth    = 1
      }
    }
    resource_public_access_default_enabled = true
  }
  pods_config {
    pod_network_mode = "VpcCniShared"
    vpc_cni_config {
      subnet_ids = [volcengine_subnet.foo.id]
    }
  }
  services_config {
    service_cidrsv4 = ["172.30.0.0/18"]
  }
  tags {
    key   = "tf-k1"
    value = "tf-v1"
  }
}

resource "volcengine_vke_permission" "foo" {
  role_name    = "vke:visitor"
  grantee_id   = 385500000
  grantee_type = "User"
  role_domain  = "cluster"
  cluster_id   = volcengine_vke_cluster.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `grantee_id` - (Required, ForceNew) The ID of the grantee.
* `grantee_type` - (Required, ForceNew) The type of the grantee. Valid values: `User`.
* `role_domain` - (Required, ForceNew) The types of permissions granted to IAM users or roles. Valid values: `namespace`, `cluster`, `all_clusters`. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `role_name` - (Required, ForceNew) The name of RBAC role. The following RBAC permissions can be granted: custom role name, system preset role names.
* `cluster_id` - (Optional, ForceNew) The cluster ID that needs to be authorized to IAM users or roles.
* `is_custom_role` - (Optional, ForceNew) Whether the RBAC role is a custom role. Default is false.
* `namespace` - (Optional, ForceNew) The namespace that needs to be authorized to IAM users or roles.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `authorized_at` - The authorized time of the RBAC Permission.
* `authorizer_id` - The ID of the Authorizer.
* `authorizer_name` - The name of the Authorizer.
* `authorizer_type` - The type of the Authorizer.
* `granted_at` - The granted time of the RBAC Permission.
* `kube_role_binding_name` - The name of the Kube Role Binding.
* `message` - The message of the RBAC Permission.
* `revoked_at` - The revoked time of the RBAC Permission.
* `status` - The status of the RBAC Permission.


## Import
VkePermission can be imported using the id, e.g.
```
$ terraform import volcengine_vke_permission.default resource_id
```

