package vault

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	Host  string       = "SECURDEN_HOST"
	Token string       = "SECURDEN_TOKEN"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"securden_host": {
				Type        : schema.TypeString,
				Required    : true,
				DefaultFunc : schema.EnvDefaultFunc(
					Host,
					"",
				),
				Description  : "Securden server to interact with.",
			},
			"securden_token": {
				Type        :     schema.TypeString,
				Required    : true,
				DefaultFunc : schema.EnvDefaultFunc(
					Token,
					"",
				),
				Description  : "Securden API Token.",
			},
			"tls_unsecure": {
				Type        : schema.TypeBool,
				Optional    : true,
				Default     : false,
				Description : "Whether to perform TLS cert verification on the server's certificate. (default : false)",
			},
			"use_proxy": {
				Type        : schema.TypeBool,
				Optional    : true,
				Default     : true,
				Description : "Should use a proxy. (default : false)",
			},
			"proxy_server": {
				Type        : schema.TypeString,
				Optional    : true,
				Default     : "",
				Description : "Custom proxy server to use (if not, HTTP_PROXY environment variable will be used)",
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
				Description: "Timeout in seconds for API requests.",
			},
		},
		ConfigureFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"securdenvault_password":      datasourceVaultPassword(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"securdenvault_account":       resourceAccount(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	host         := d.Get("securden_host").(string)
	token        := d.Get("securden_token").(string)
	tlsInsecure  := d.Get("tls_unsecure").(bool)
	useProxy     := d.Get("use_proxy").(bool)
	proxyServer  := d.Get("proxy_server").(string)
	timeout      := d.Get("timeout").(int)

	if host == "" {
		return nil, fmt.Errorf("no Securden host was provided")
	}

	if token == "" {
		return nil, fmt.Errorf("no token was provided")
	}

	timeout = d.Get("timeout").(int)

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: tlsInsecure})
	client.SetBaseURL(fmt.Sprintf("https://%s/api", host))

	if(useProxy){
		if(proxyServer != ""){
			client.SetProxy(proxyServer)
		}
	}else{
		client.RemoveProxy()
	}
	client.SetHeader("authtoken", token)
	client.SetTimeout(time.Duration(timeout) * time.Second)

	return client, nil
}
