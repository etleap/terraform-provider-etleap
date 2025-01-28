package provider

import (
	"testing"
	"text/template"

	"github.com/etleap/terraform-provider-etleap/internal/provider/testutils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

type PipelineConfig struct {
	Name                   string
	AutomaticSchemaChanges bool
	testutils.TestConstants
}

const PIPELINE_CONFIG_TEMPLATE_STRING = `
resource "etleap_pipeline" "test_mysql_pipeline" {
	name = "{{.Name}}"
  
	source = {
	  mysql = {
		type                = "MYSQL"
		connection_id       = "{{.MysqlConnectionId}}"
		database            = "etleap_test"
		table               = "date_pk_test"
		primary_key_columns = ["date_pk"]
	  }
	}
  
	destination = {
	  redshift = {
		type          = "REDSHIFT"
		connection_id = "{{.RedshiftConnectionId}}"
		schema        = "caius"
		table         = "test_tf_mysql"

		automatic_schema_changes = {{.AutomaticSchemaChanges}}
  
		distribution_style = {
		  one = "AUTO"
		}
	  }
	}
  }
`

var pipelineConfigTemplate, _ = template.New("PipelineConfig").Parse(PIPELINE_CONFIG_TEMPLATE_STRING)

func TestAccMysqlToRedshiftPipeline(t *testing.T) {
	var pipelineName = acctest.RandomWithPrefix("Mysql Test")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Resource creation
				Config: GetProviderDefinition() + getPipelineConfig(&PipelineConfig{
					pipelineName,
					false,
					*testutils.Constants,
				}),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should create the resource
						plancheck.ExpectResourceAction("etleap_pipeline.test_mysql_pipeline", plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_pipeline.test_mysql_pipeline", "name", pipelineName),
					resource.TestCheckResourceAttr("etleap_pipeline.test_mysql_pipeline", "source.mysql.connection_id", testutils.Constants.MysqlConnectionId),
					resource.TestCheckResourceAttr("etleap_pipeline.test_mysql_pipeline", "destination.redshift.connection_id", testutils.Constants.RedshiftConnectionId),
					resource.TestCheckResourceAttr("etleap_pipeline.test_mysql_pipeline", "destination.redshift.automatic_schema_changes", "false"),
				),
			},
			{ // Modify name and schema changes
				Config: GetProviderDefinition() + getPipelineConfig(&PipelineConfig{
					pipelineName + " 2",
					true,
					*testutils.Constants,
				}),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should update the resource, and not destroy and recreate it
						plancheck.ExpectResourceAction("etleap_pipeline.test_mysql_pipeline", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_pipeline.test_mysql_pipeline", "name", pipelineName+" 2"),
					resource.TestCheckResourceAttr("etleap_pipeline.test_mysql_pipeline", "destination.redshift.automatic_schema_changes", "true"),
				),
			},
		},
	})
}

func getPipelineConfig(config *PipelineConfig) string {
	return testutils.RunTemplate(pipelineConfigTemplate, config)
}
