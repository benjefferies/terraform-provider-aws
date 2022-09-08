// Code generated by "internal/generate/listpages/main.go -ListOps=GetApis,GetDomainNames"; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
)

func getAPIsPages(conn *apigatewayv2.ApiGatewayV2, input *apigatewayv2.GetApisInput, fn func(*apigatewayv2.GetApisOutput, bool) bool) error {
	return getAPIsPagesWithContext(context.Background(), conn, input, fn)
}

func getAPIsPagesWithContext(ctx context.Context, conn *apigatewayv2.ApiGatewayV2, input *apigatewayv2.GetApisInput, fn func(*apigatewayv2.GetApisOutput, bool) bool) error {
	for {
		output, err := conn.GetApisWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getDomainNamesPages(conn *apigatewayv2.ApiGatewayV2, input *apigatewayv2.GetDomainNamesInput, fn func(*apigatewayv2.GetDomainNamesOutput, bool) bool) error {
	return getDomainNamesPagesWithContext(context.Background(), conn, input, fn)
}

func getDomainNamesPagesWithContext(ctx context.Context, conn *apigatewayv2.ApiGatewayV2, input *apigatewayv2.GetDomainNamesInput, fn func(*apigatewayv2.GetDomainNamesOutput, bool) bool) error {
	for {
		output, err := conn.GetDomainNamesWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getApiMappingsPages(conn *apigatewayv2.ApiGatewayV2, input *apigatewayv2.GetApiMappingsInput, fn func(*apigatewayv2.GetApiMappingsOutput, bool) bool) error {
	return getApiMappingsPagesWithContext(context.Background(), conn, input, fn)
}

func getApiMappingsPagesWithContext(ctx context.Context, conn *apigatewayv2.ApiGatewayV2, input *apigatewayv2.GetApiMappingsInput, fn func(*apigatewayv2.GetApiMappingsOutput, bool) bool) error {
	for {
		output, err := conn.GetApiMappingsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func getStagesPages(conn *apigatewayv2.ApiGatewayV2, input *apigatewayv2.GetStagesInput, fn func(*apigatewayv2.GetStagesOutput, bool) bool) error {
	return getStagesPagesWithContext(context.Background(), conn, input, fn)
}

func getStagesPagesWithContext(ctx context.Context, conn *apigatewayv2.ApiGatewayV2, input *apigatewayv2.GetStagesInput, fn func(*apigatewayv2.GetStagesOutput, bool) bool) error {
	for {
		output, err := conn.GetStagesWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}