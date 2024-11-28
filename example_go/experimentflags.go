// AUTOMATICALLY GENERATED BY OPENFEATURE CODEGEN, DO NOT EDIT.
package experimentflags

import (
	"context"
	"github.com/open-feature/go-sdk/openfeature"
)

type BooleanProvider func(ctx context.Context, evalCtx openfeature.EvaluationContext) (bool, error)
type BooleanProviderDetails func(ctx context.Context, evalCtx openfeature.EvaluationContext) (openfeature.BooleanEvaluationDetails, error)
type FloatProvider func(ctx context.Context, evalCtx openfeature.EvaluationContext) (float64, error)
type FloatProviderDetails func(ctx context.Context, evalCtx openfeature.EvaluationContext) (openfeature.FloatEvaluationDetails, error)
type IntProvider func(ctx context.Context, evalCtx openfeature.EvaluationContext) (int64, error)
type IntProviderDetails func(ctx context.Context, evalCtx openfeature.EvaluationContext) (openfeature.IntEvaluationDetails, error)
type StringProvider func(ctx context.Context, evalCtx openfeature.EvaluationContext) (string, error)
type StringProviderDetails func(ctx context.Context, evalCtx openfeature.EvaluationContext) (openfeature.StringEvaluationDetails, error)

var client openfeature.IClient = nil
// Discount percentage applied to purchases.
var DiscountPercentage = struct {
    // Value returns the value of the flag DiscountPercentage,
    // as well as the evaluation error, if present.
    Value FloatProvider

    // ValueWithDetails returns the value of the flag DiscountPercentage,
    // the evaluation error, if any, and the evaluation details.
    ValueWithDetails FloatProviderDetails
}{
    Value: func(ctx context.Context, evalCtx openfeature.EvaluationContext) (float64, error) {
        return client.FloatValue(ctx, "discountPercentage", 0.15, evalCtx)
    },
    ValueWithDetails: func(ctx context.Context, evalCtx openfeature.EvaluationContext) (openfeature.FloatEvaluationDetails, error){
        return client.FloatValueDetails(ctx, "discountPercentage", 0.15, evalCtx)
    },
}
// Controls whether Feature A is enabled.
var EnableFeatureA = struct {
    // Value returns the value of the flag EnableFeatureA,
    // as well as the evaluation error, if present.
    Value BooleanProvider

    // ValueWithDetails returns the value of the flag EnableFeatureA,
    // the evaluation error, if any, and the evaluation details.
    ValueWithDetails BooleanProviderDetails
}{
    Value: func(ctx context.Context, evalCtx openfeature.EvaluationContext) (bool, error) {
        return client.BooleanValue(ctx, "enableFeatureA", false, evalCtx)
    },
    ValueWithDetails: func(ctx context.Context, evalCtx openfeature.EvaluationContext) (openfeature.BooleanEvaluationDetails, error){
        return client.BooleanValueDetails(ctx, "enableFeatureA", false, evalCtx)
    },
}
// The message to use for greeting users.
var GreetingMessage = struct {
    // Value returns the value of the flag GreetingMessage,
    // as well as the evaluation error, if present.
    Value StringProvider

    // ValueWithDetails returns the value of the flag GreetingMessage,
    // the evaluation error, if any, and the evaluation details.
    ValueWithDetails StringProviderDetails
}{
    Value: func(ctx context.Context, evalCtx openfeature.EvaluationContext) (string, error) {
        return client.StringValue(ctx, "greetingMessage", "Hello there!", evalCtx)
    },
    ValueWithDetails: func(ctx context.Context, evalCtx openfeature.EvaluationContext) (openfeature.StringEvaluationDetails, error){
        return client.StringValueDetails(ctx, "greetingMessage", "Hello there!", evalCtx)
    },
}
// Maximum allowed length for usernames.
var UsernameMaxLength = struct {
    // Value returns the value of the flag UsernameMaxLength,
    // as well as the evaluation error, if present.
    Value IntProvider

    // ValueWithDetails returns the value of the flag UsernameMaxLength,
    // the evaluation error, if any, and the evaluation details.
    ValueWithDetails IntProviderDetails
}{
    Value: func(ctx context.Context, evalCtx openfeature.EvaluationContext) (int64, error) {
        return client.IntValue(ctx, "usernameMaxLength", 50, evalCtx)
    },
    ValueWithDetails: func(ctx context.Context, evalCtx openfeature.EvaluationContext) (openfeature.IntEvaluationDetails, error){
        return client.IntValueDetails(ctx, "usernameMaxLength", 50, evalCtx)
    },
}

func init() {
	client = openfeature.GetApiInstance().GetClient()
}
