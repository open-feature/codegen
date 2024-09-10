package experimentflags

import (
	"codegen/providers"
	"context"
	"github.com/open-feature/go-sdk/openfeature"
)

var client *openfeature.Client = nil
// This is a flag.
var MyOpenFeatureFlag = struct {
    Value providers.BooleanProvider    
}{
    Value: func(ctx context.Context) (bool, error) {
        return client.BooleanValue(ctx, "myOpenFeatureFlag", false, openfeature.EvaluationContext{})
    },
}

func init() {
	client = openfeature.NewClient("experimentflags")
}
