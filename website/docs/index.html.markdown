---
layout: "vestack"
page_title: "Provider: vestack"
sidebar_current: "docs-vestack-index"
description: |-
The vestack provider is used to interact with many resources supported by Vestack. The provider needs to be configured with the proper credentials before it can be used.
---

# Vestack Provider

The Vestack provider is used to interact with many resources supported by [Vestack](https://www.volcengine.com/).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

-> **Note:** This guide requires an available Vestack account or sub-account with project to create resources.

## Example Usage
```hcl
# Configure the Vestack Provider
provider "vestack" {
  access_key = "your ak"
  secret_key = "your sk"
  session_token = "sts token"
  region = "cn-beijing"
}

# Query Vpc
data "vestack_vpcs" "default"{
  ids = ["vpc-mizl7m1kqccg5smt1bdpijuj"]
}

#Create vpc
resource "vestack_vpc" "foo" {
  vpc_name = "tf-test-1"
  cidr_block = "172.16.0.0/16"
  dns_servers = ["8.8.8.8","114.114.114.114"]
}

```

## Authentication

The Vestack provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding an `public_key` and `private_key` in-line in the
vestack provider block:

Usage:

```hcl
provider "vestack" {
   access_key = "your ak"
   secret_key = "your sk"
   region = "cn-beijing"
}
```

### Environment variables

You can provide your credentials via `VESTACK_ACCESS_KEY` and `VESTACK_SECRET_KEY`
environment variables, representing your vestack public key and private key respectively.
`VESTACK_REGION` is also used, if applicable:

```hcl
provider "vestack" {
  
}
```

Usage:

```hcl
$ export VESTACK_ACCESS_KEY="your_public_key"
$ export VESTACK_SECRET_KEY="your_private_key"
$ export VESTACK_REGION="cn-beijing"
$ terraform plan
```

