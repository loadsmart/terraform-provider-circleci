package circleci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	var p *schema.Provider
	p = &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIRCLE_TOKEN", nil),
				Description: "CircleCI API token",
			},
			"organization": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIRCLE_PROJECT_USERNAME", nil),
				Description: "CircleCI username/organization",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"circleci_project": resourceCircleciProject(),
		},

		DataSourcesMap: map[string]*schema.Resource{},
	}
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			Token:        d.Get("token").(string),
			Organization: d.Get("organization").(string),
		}

		return config.Client()
	}
}
