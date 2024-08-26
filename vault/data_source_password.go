package vault

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceVaultPassword() *schema.Resource {
	return &schema.Resource{
		Read:        datasourceVaultPasswordRead,
		Description: "Read password.",
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account ID in Securden.",
			},
			"password": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Computed:    true,
				Description: "Password of the account.",
			},
		},
	}
}

func datasourceVaultPasswordRead(d *schema.ResourceData, m interface{}) error {
	client    := m.(*resty.Client)
	accountId :=  d.Get("account_id").(string)
	uri       := fmt.Sprintf("/get_password?account_id=%s", accountId)
	log.Printf("[DEBUG] datasourceVaultPasswordRead - Query: %s", uri)
	var resp datasourceVaultPasswordReadApi
	_, err := client.R().
		SetResult(&resp).
		Get(uri)

	if err != nil {
		log.Printf("[WARN] No password found for account %s", accountId)
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] datasourceVaultPasswordRead - Result: %#v", resp)
	d.SetId(accountId)
	d.Set("password", resp.Password)
	return nil
}
