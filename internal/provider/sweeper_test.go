package provider

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/etleap/terraform-provider-etleap/internal/sdk"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/operations"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Names of the pipelines that the acceptance tests create, including the " 2" suffix that the update step
// adds. A test run that stops before Terraform destroys its resources leaves these pipelines behind. A
// leftover pipeline still owns its destination table, so the next run cannot create a pipeline that writes to
// the same table.
var sweptPipelineName = regexp.MustCompile(`^tf-acc-test-pipeline-\d+( 2)?$`)

// Pipelines that are newer than this can belong to a test run that is still in progress.
const sweepMinimumAge = time.Hour

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	resource.AddTestSweepers("etleap_pipeline", &resource.Sweeper{
		Name: "etleap_pipeline",
		F:    sweepPipelines,
	})
}

func sweepPipelines(_ string) error {
	ctx := context.Background()
	client := sdk.New(sdk.WithSecurity(shared.Security{
		Username: credentials.AccessKey,
		Password: credentials.SecretKey,
	}))

	res, err := client.Pipeline.GetPipelines(ctx, operations.GetPipelinesRequest{})
	if err != nil {
		return fmt.Errorf("could not list pipelines: %w", err)
	}
	if res.PipelinesOutput == nil {
		return fmt.Errorf("could not list pipelines: got response code %d", res.StatusCode)
	}

	for _, pipeline := range res.PipelinesOutput.Pipelines {
		if !sweptPipelineName.MatchString(pipeline.Name) || time.Since(pipeline.CreateDate) < sweepMinimumAge {
			continue
		}
		deleteRes, err := client.Pipeline.Delete(ctx, operations.DeletePipelineRequest{ID: pipeline.ID})
		if err != nil {
			return fmt.Errorf("could not delete pipeline %s (%s): %w", pipeline.ID, pipeline.Name, err)
		}
		if deleteRes.StatusCode != 204 {
			return fmt.Errorf("could not delete pipeline %s (%s): got response code %d", pipeline.ID, pipeline.Name, deleteRes.StatusCode)
		}
		log.Printf("[INFO] Deleted leftover test pipeline %s (%s), created %s", pipeline.ID, pipeline.Name, pipeline.CreateDate)
	}
	return nil
}
