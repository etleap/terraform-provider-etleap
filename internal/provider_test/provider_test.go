package provider

import (
	"os"
	"text/template"
	"../provider"

	"github.com/etleap/terraform-provider-etleap/internal/provider/testutils"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const PROVIDER_CONFIG_TEMPLATE_STRING = `
provider "etleap" { 
	username    = "{{.AccessKey}}" 
	password    = "{{.SecretKey}}" 
}
`

var providerConfigTemplate, _ = template.New("providerConfig").Parse(PROVIDER_CONFIG_TEMPLATE_STRING)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"etleap": providerserver.NewProtocol6WithError(New("test")()),
}

type Credentials struct {
	AccessKey string
	SecretKey string
}

var credentials = &Credentials{}

func init() {
	var (
		hasAccessKey bool
		hasSecretKey bool
	)

	credentials.AccessKey, hasAccessKey = os.LookupEnv(`ETLEAP_ACCESS_KEY`)
	credentials.SecretKey, hasSecretKey = os.LookupEnv(`ETLEAP_SECRET_KEY`)

	if !hasAccessKey || !hasSecretKey {
		panic(`Credentials not set; Please set the environment variables "ETLEAP_ACCESS_KEY" and "ETLEAP_SECRET_KEY"`)
	}
}

func GetProviderDefinition() string {
	return testutils.RunTemplate(providerConfigTemplate, credentials)
}
