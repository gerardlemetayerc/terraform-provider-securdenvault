---
layout: "securdenvault"
page_title: "Provider: Securden Vault"
description: |-
  The Securden Vault provider is used to interact with the resources supported by Securden Password Vault.
---

# Securden Vault Provider

The Securden Vault provider is used to manage the lifecycle of resources in a Securden Password Vault. The provider allows Terraform to interact with Securden API and manage its resources programmatically.

## Example Usage

```hcl
provider "securdenvault" {
  securden_host  = "yourserver.example.com"
  securden_token = "your-securden-token"
  tls_unsecure   = false
}

# Use the provider to create a new password entry
resource "securdenvault_password" "example" {
  # ...
}
```

## Authentication

The Securden Vault provider requires a host URL and an API token to interact with Securden. These credentials can be supplied through the configuration file or environment variables.

## Argument Reference

The following arguments are supported:

- `securden_host` - (Required) The host URL of the Securden Password Vault.
- `securden_token` - (Required) The API token for access to Securden.
- `tls_unsecure` - (Optional) Whether to perform TLS cert verification. Defaults to `false`.
- `use_proxy` - (Optional) Enable usage of proxy server. Defaults to `true`.
- `proxy_server` - (Optional) Custom proxy server URL.

## Building from Source

If you wish to work on the provider, clone the repository, and build from source:

```shell
git clone https://github.com/your-github-account/terraform-provider-securdenvault.git
cd terraform-provider-securdenvault
go build
```

## Contributing

Contributions to the Securden Vault provider are welcome. Please see [CONTRIBUTING.md](https://github.com/gerardlemetayerc/terraform-provider-securdenvault/CONTRIBUTING.md) for more details.

## Issues

If you find an issue with the provider, please report it in the [issue tracker](https://github.com/gerardlemetayerc/terraform-provider-securdenvault/issues).