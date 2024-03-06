package provider

import (
	"testing"
	"text/template"

	"github.com/etleap/terraform-provider-etleap/internal/provider/testutils"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

const DBT_SCHEDULE_TEMPLATE_STRING = `
resource "etleap_dbt_schedule" "my_dbtschedule" {
	connection_id             = "{{.RedshiftConnectionId}}"
	cron                      = "0 * * * *"
	name                      = "{{.Name}}"
	selector                  = "test"
	skip_build_if_no_new_data = true
	target_schema             = "caius"
  }
`

var dbtScheduleTemplate, _ = template.New("DbtSchedule").Parse(DBT_SCHEDULE_TEMPLATE_STRING)

type DbtScheduleConfig struct {
	Name string
	testutils.TestConstants
}

func TestAccDbtSchedule(t *testing.T) {
	dbtScheduleName := acctest.RandomWithPrefix("DbtSchedule")

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Resource creation
				Config: GetProviderDefinition() + getDbtScheduleConfig(&DbtScheduleConfig{
					dbtScheduleName,
					*testutils.Constants,
				}),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should create the resource
						plancheck.ExpectResourceAction("etleap_dbt_schedule.my_dbtschedule", plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_dbt_schedule.my_dbtschedule", "name", dbtScheduleName),
				),
			},
			{ // Resource Modification
				Config: GetProviderDefinition() + getDbtScheduleConfig(&DbtScheduleConfig{
					dbtScheduleName + " 2",
					*testutils.Constants,
				}),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should modify the resource
						plancheck.ExpectResourceAction("etleap_dbt_schedule.my_dbtschedule", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_dbt_schedule.my_dbtschedule", "name", dbtScheduleName+" 2"),
				),
			},
		},
	})
}

func getDbtScheduleConfig(config *DbtScheduleConfig) string {
	return testutils.RunTemplate(dbtScheduleTemplate, config)
}
