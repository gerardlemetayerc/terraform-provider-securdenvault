---
layout: "securdenvault"
page_title: "Securden Vault Datasource: Resource"
subcategory: "Account Management"
description: |-
  This resource allows you to manage accounts in your Securden Vault instance using Terraform.
---

# securdenvault_account (Resource)

The `securdenvault_account` resource permits the management of accounts in a Securden Vault instance. It can be used to create, update, read, and delete accounts.

## Example Usage 

```hcl
resource "securdenvault_account" "example" {
  account_title       = "Example Account"
  account_name        = "example account"
  account_type        = "user"
  personal_account    = false
  folder_id           = "folder123"
  password_policy_name = "default_policy"
  generate_password   = true
}
```

The following arguments are supported:

* ```account_title``` (String - Required) - The title of the account.
* ```account_name``` (String - Required) - The name of the account.
* ```account_type``` (String - Required) - The type of the account.
* ```personal_account``` (Bool - Optional, Default: false) - Whether the account is a personal account.
* ```folder_id (String - Required) - The ID of the folder in which the account resides.
* ```password_policy_name``` (String - Optional) - The name of the policy to use to generate the account password.
* ```password``` (String - Optional, Sensitive) - The password of the account. Either provide a password or set generate_password to true.
* ```generate_password``` (Bool - Optional, Default: false) - Whether to generate a new password for the account. If true, a password is generated automatically if not provided.
* ```password_change_reason``` (String - Optional, Default: "Updated by Terraform") - Reason for the password change.

### In addition to all arguments above, the following attributes are exported:

* ```id``` - The ID of the account in Securden Vault.

```
--- 
Note: The ressource store the password in the Terraform state file. Consider secure the state file in an encrypted place.
```