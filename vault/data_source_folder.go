package vault

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type datasourceVaulFolderReadApi struct {
	FolderName   string    `json:"folder_name"`
	Id           string    `json:"id"`
	Parent       string    `json:"parent"`
}

func datasourceVaultFolder() *schema.Resource {
	return &schema.Resource{
		Read:        datasourceVaultFolderRead,
		Description: "Read password.",
		Schema: map[string]*schema.Schema{
			"folder_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Folder name in Securden.",
			},
			"parent": {
				Type:        schema.TypeString,
				Computed:    true;
				Description: "Account ID in Securden.",				
			},
		},
	}
}

func datasourceVaultFolderRead(d *schema.ResourceData, m interface{}) error {
	client     := m.(*resty.Client)
	folderName :=  d.Get("folder_name").(string)
	uri        := fmt.Sprintf("/get_folders?search_text=%s", accountId)
	log.Printf("[DEBUG] datasourceVaultFolderRead - Query: %s", uri)
	client.SetDebug(true)
	var resp datasourceVaulFolderReadApi
	_, err := client.R().
		SetResult(&resp).
		Get(uri)

	if err != nil {
		log.Printf("[WARN] No folder found for name %s", FolderName)
		d.SetId("")
		return nil
	}

	log.Printf("[DEBUG] datasourceVaultFolderRead - Result: %#v", resp)
	d.SetId(resp.Id)
	d.Set("folder_name", resp.FolderName)
	d.Set("parent", resp.FolderName)
	return nil
}
