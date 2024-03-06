package provider

import (
	"encoding/json"
	"os"
	"testing"
	"text/template"

	"github.com/etleap/terraform-provider-etleap/internal/provider/testutils"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

const MYSQL_CONNECTION_TEMPLATE_STRING = `
resource "etleap_connection_mysql" "test_mysql"{
    type = "MYSQL"
    name = "{{.Name}}"
    address = "{{.Address}}"
    database = "etleap_test"
    port = 3306
    username = "{{.Username}}"
    password = "{{.Password}}"
}
`

var mysqlConnectionTemplate, _ = template.New("MySQLConnection").Parse(MYSQL_CONNECTION_TEMPLATE_STRING)

type MysqlConnectionConfig struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var mysqlConfig = MysqlConnectionConfig{}

func init() {
	conf, ok := os.LookupEnv(`ETLEAP_MYSQL_CONFIG`)
	if !ok {
		panic("Set ETLEAP_MYSQL_CONFIG env variable")
	}
	json.Unmarshal([]byte(conf), &mysqlConfig)
}

func TestAccMysqlConnection(t *testing.T) {
	connectionName := acctest.RandomWithPrefix("MySQL Connection")
	mysqlConfig.Name = connectionName
	mysqlConfigUpdate := mysqlConfig
	mysqlConfigUpdate.Name = connectionName + " 2"

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{ // Resource creation
				Config: GetProviderDefinition() + getMySqlConnectionConfig(&mysqlConfig),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should create the resource
						plancheck.ExpectResourceAction("etleap_connection_mysql.test_mysql", plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_connection_mysql.test_mysql", "name", connectionName),
				),
			},
			{ // Resource Modification
				PreConfig: func() {
					mysqlConfig.Name = connectionName + " 2"
				},
				Config: GetProviderDefinition() + getMySqlConnectionConfig(&mysqlConfigUpdate),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// should modify the resource
						plancheck.ExpectResourceAction("etleap_connection_mysql.test_mysql", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("etleap_connection_mysql.test_mysql", "name", connectionName+" 2"),
				),
			},
		},
	})
}

func getMySqlConnectionConfig(config *MysqlConnectionConfig) string {
	return testutils.RunTemplate(mysqlConnectionTemplate, config)
}
