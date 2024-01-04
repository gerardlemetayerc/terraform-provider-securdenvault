package vault

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type datasourceD42BusinessAppApi struct {
	Password   string    `json:"password"`
}

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
	client.SetDebug(true)
	var resp datasourceD42BusinessAppApi
	_, err := client.R().
		SetResult(&datasourceD42BusinessAppApi).
		Get(uri)

	if err != nil {
		log.Printf("[WARN] No password found for account %s", d.Get("name").(string))
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] datasourceVaultPasswordRead - Result: %#v", resp.Result())
	if len(resp.Password) == 1 {
		d.SetId(accountId)
		d.Set("password", resp.Password)
	}
	return nil
}
