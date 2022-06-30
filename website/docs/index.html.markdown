---
layout: "volcengine"
page_title: "Provider: volcengine"
sidebar_current: "docs-volcengine-index"
description: |-
The volcengine provider is used to interact with many resources supported by Volcengine. The provider needs to be configured with the proper credentials before it can be used.
---

# Volcengine Provider

The Volcengine provider is used to interact with many resources supported by [Volcengine](https://www.volcengine.com/).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

-> **Note:** This guide requires an available Volcengine account or sub-account with project to create resources.

## Example Usage
```hcl
# Configure the Volcengine Provider
provider "volcengine" {
  access_key = "your ak"
  secret_key = "your sk"
  session_token = "sts token"
  region = "cn-beijing"
}

# Query Vpc
data "volcengine_vpcs" "default"{
  ids = ["vpc-mizl7m1kqccg5smt1bdpijuj"]
}

#Create vpc
resource "volcengine_vpc" "foo" {
  vpc_name = "tf-test-1"
  cidr_block = "172.16.0.0/16"
  dns_servers = ["8.8.8.8","114.114.114.114"]
}

```

## Authentication

The Volcengine provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding an `public_key` and `private_key` in-line in the
volcengine provider block:

Usage:

```hcl
provider "volcengine" {
   access_key = "your ak"
   secret_key = "your sk"
   region = "cn-beijing"
}
```

### Environment variables

You can provide your credentials via `VOLCENGINE_ACCESS_KEY` and `VOLCENGINE_SECRET_KEY`
environment variables, representing your volcengine public key and private key respectively.
`VOLCENGINE_REGION` is also used, if applicable:

```hcl
provider "volcengine" {
  
}
```

Usage:

```hcl
$ export VOLCENGINE_ACCESS_KEY="your_public_key"
$ export VOLCENGINE_SECRET_KEY="your_private_key"
$ export VOLCENGINE_REGION="cn-beijing"
$ terraform plan
```

