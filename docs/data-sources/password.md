---
layout: "securdenvault"
page_title: "Securden Vault Datasource: password"
description: |-
  Use this data source to get information about a specific password stored in Securden Password Vault.
---

# Securden Vault Datasource: password

The `password` data source allows you to retrieve information about a specific password stored in the Securden Password Vault. This can be used to fetch passwords for use in other Terraform resources.

## Example Usage

```hcl
data "securdenvault_password" "example" {
  account_id = "example-account-id"
}

output "password" {
  value = data.securdenvault_password.example.password
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The ID of the account for which the password is to be retrieved.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `password` - The password of the specified account.

## Timeouts

The `password` data source supports the following `timeouts` configuration options:

- `read` - (Optional) The time to wait for the API to return the password. Default is 30 seconds.

```
--- 
Note: The data source store the password in the Terraform state file. Consider secure the state file in an encrypted place.
```

## Import

Password data cannot be imported; this data source is for retrieval of existing passwords only.
