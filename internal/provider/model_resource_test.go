package provider

import (
	"testing"
	"text/template"

	"github.com/etleap/terraform-provider-etleap/internal/provider/testutils"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

const MODEL_TEMPLATE_STRING = `
resource "etleap_model" "my_model" {
	deletion_of_export_products = true
	name                        = "<<.Name>>"
	query_and_triggers = {
	  query    = "SELECT * from {{<<.PipelineId>>}};"
	  triggers = ["<<.PipelineId>>"]
	}
	update_schedule = {
	  daily = {
		hour_of_day = 4
		mode        = "DAILY"
	  }
	}
	warehouse = {
	  redshift = {
		connection_id = "<<.RedshiftConnectionId>>"
		distribution_style = {
		  one = "ALL"
		}
		materialized_view           = true
		schema                      = "public"
		table                       = "tf_test_table"
		type                        = "REDSHIFT"
		wait_for_update_preparation = false
	  }
	}
  }
`

// Using << and >> as template delimiters to not conflict with the model query syntax
var modelTemplate, _ = template.New("ModelTemplate").Delims(`<<`, `>>`).Parse(MODEL_TEMPLATE_STRING)

type ModelConfig struct {
	Name string
	testutils.TestConstants
}

func TestAccModel(t *testing.T) {
	modelName := acctest.RandomWithPrefix("Model")

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Resource creation
				Config: GetProviderDefinition() + getModelConfig(&ModelConfig{
					modelName,
					*testutils.Constants,
				}),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should create the resource
						plancheck.ExpectResourceAction("etleap_model.my_model", plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_model.my_model", "name", modelName),
				),
			},
			{ // Resource Modification
				Config: GetProviderDefinition() + getModelConfig(&ModelConfig{
					modelName + " 2",
					*testutils.Constants,
				}),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should modify the resource
						plancheck.ExpectResourceAction("etleap_model.my_model", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_model.my_model", "name", modelName+" 2"),
				),
			},
		},
	})
}

func getModelConfig(config *ModelConfig) string {
	return testutils.RunTemplate(modelTemplate, config)
}
