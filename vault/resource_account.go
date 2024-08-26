package vault

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountCreate,
		Read:   resourceAccountRead,
		Update: resourceAccountUpdate,
		Delete: resourceAccountDelete,

		Schema: map[string]*schema.Schema{
			"account_title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The title of the account.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the account.",
			},
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the account.",
			},
			"personal_account": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the account is a personal account.",
			},
			"folder_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the folder in which the account resides.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password of the account. Either provide a password or set generate_password to true.",
			},
			"generate_password": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to generate a new password for the account. If true, a password is generated automatically if not provided.",
			},
			"password_change_reason": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Updated by Terraform",
				Description: "Reason for the password change.",
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*resty.Client)
	account := map[string]interface{}{
		"account_title":    d.Get("account_title").(string),
		"account_name":     d.Get("account_name").(string),
		"account_type":     d.Get("account_type").(string),
		"personal_account": d.Get("personal_account").(bool),
		"folder_id":        d.Get("folder_id").(string),
	}

	if generate, ok := d.GetOk("generate_password"); ok && generate.(bool) {
		password, err := generatePassword(client)
		if err != nil {
			return fmt.Errorf("error generating password: %s", err)
		}
		account["password"] = password
	} else {
		account["password"] = d.Get("password").(string)
	}

	uri := "/api/accounts/create"
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(account).
		Post(uri)

	if err != nil {
		return fmt.Errorf("error creating account: %s", err)
	}

	d.SetId(resp.String())
	return resourceAccountRead(d, m)
}

func resourceAccountRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*resty.Client)
	id := d.Id() // L'identifiant unique du compte

	// Construire l'URL pour l'endpoint de récupération des détails du compte
	uri := fmt.Sprintf("/api/accounts/%s", id) // Assurez-vous que cet endpoint est correct

	log.Printf("[DEBUG] Reading account details for ID: %s", id)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(uri)

	if err != nil {
		return fmt.Errorf("failed to retrieve account details: %s", err)
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("API responded with non-200 status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	var result struct {
		Account struct {
			AccountTitle    string `json:"account_title"`
			AccountName     string `json:"account_name"`
			AccountType     string `json:"account_type"`
			PersonalAccount bool   `json:"personal_account"`
			FolderID        string `json:"folder_id"`
			Password        string `json:"password,omitempty"` // Si le mot de passe est retourné, ce qui est peu probable pour des raisons de sécurité
		} `json:"account"`
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return fmt.Errorf("failed to parse account details: %s", err)
	}

	// Mettre à jour l'état Terraform avec les données reçues
	d.Set("account_title", result.Account.AccountTitle)
	d.Set("account_name", result.Account.AccountName)
	d.Set("account_type", result.Account.AccountType)
	d.Set("personal_account", result.Account.PersonalAccount)
	d.Set("folder_id", result.Account.FolderID)
	// Conditionnellement définir le mot de passe si celui-ci est effectivement retourné
	if result.Account.Password != "" {
		d.Set("password", result.Account.Password)
	}

	return nil
}

func resourceAccountUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*resty.Client)
	id := d.Id() // L'identifiant unique du compte

	updateData := map[string]interface{}{
		"account_title":    d.Get("account_title").(string),
		"account_name":     d.Get("account_name").(string),
		"account_type":     d.Get("account_type").(string),
		"personal_account": d.Get("personal_account").(bool),
		"folder_id":        d.Get("folder_id").(string),
	}

	if d.HasChange("password") {
		passwordData := map[string]interface{}{
			"password":               d.Get("password").(string),
			"password_change_reason": d.Get("password_change_reason").(string),
		}
		// Appel API pour mise à jour du mot de passe
		passwordUri := fmt.Sprintf("/api/accounts/%s/password", id)
		log.Printf("[DEBUG] Updating password for account ID: %s", id)
		_, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(passwordData).
			Post(passwordUri)

		if err != nil {
			return fmt.Errorf("failed to update password for account ID %s: %s", id, err)
		}
	}

	// Appel API pour mise à jour des informations générales du compte
	uri := fmt.Sprintf("/api/accounts/%s", id)
	log.Printf("[DEBUG] Updating account ID: %s", id)
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(updateData).
		Put(uri)

	if err != nil {
		return fmt.Errorf("failed to update account ID %s: %s", id, err)
	}

	// Relecture de l'état de la ressource pour synchroniser les changements
	return resourceAccountRead(d, m)
}

func resourceAccountDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*resty.Client)
	id := d.Id()
	uri := fmt.Sprintf("/api/accounts/%s/delete", id)

	resp, err := client.R().Delete(uri)
	if err != nil {
		return fmt.Errorf("failed to delete account: %s", err)
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("API responded with non-200 status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	d.SetId("")
	return nil
}

func generatePassword(client *resty.Client) (string, error) {
	uri := "/api/generate_password"
	resp, err := client.R().Get(uri)
	if err != nil {
		return "", fmt.Errorf("error generating password: %s", err)
	}

	// Parse the JSON response to extract the generated password
	var result struct {
		Password string `json:"password"`
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", fmt.Errorf("failed to parse the password generation response: %s", err)
	}

	return result.Password, nil
}
