package provider

import (
	"testing"
	"text/template"

	"github.com/etleap/terraform-provider-etleap/internal/provider/testutils"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

const TEAM_TEMPLATE_STRING = `
resource "etleap_team" "my_team" {
	name        = "{{.Name}}"
	description = "A team managed via TF"
}
`

var teamTemplate, _ = template.New("ModelTemplate").Parse(TEAM_TEMPLATE_STRING)

type TeamConfig struct {
	Name string
}

func TestAccTeamResource(t *testing.T) {
	teamName := acctest.RandomWithPrefix("Team")

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Resource creation
				Config: GetProviderDefinition() + getTeamConfig(&TeamConfig{
					Name: teamName,
				}),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should create the resource
						plancheck.ExpectResourceAction("etleap_team.my_team", plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_team.my_team", "name", teamName),
				),
			},
			{ // Resource Modification
				Config: GetProviderDefinition() + getTeamConfig(&TeamConfig{
					Name: teamName + " 2",
				}),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should replace the resource; the api does not have a `PATCH` option
						plancheck.ExpectResourceAction("etleap_team.my_team", plancheck.ResourceActionReplace),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_team.my_team", "name", teamName+" 2"),
				),
			},
		},
	})
}

func getTeamConfig(config *TeamConfig) string {
	return testutils.RunTemplate(teamTemplate, config)
}
