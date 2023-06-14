// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// GetTag fetches an individual dynamodb service tag for a resource.
// Returns whether the key value and any errors. A NotFoundError is used to signal that no value was found.
// This function will optimise the handling over listTags, if possible.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func GetTag(ctx context.Context, conn dynamodbiface.DynamoDBAPI, identifier, key string) (*string, error) {
	listTags, err := listTags(ctx, conn, identifier)

	if err != nil {
		return nil, err
	}

	if !listTags.KeyExists(key) {
		return nil, tfresource.NewEmptyResultError(nil)
	}

	return listTags.KeyValue(key), nil
}

// listTags lists dynamodb service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn dynamodbiface.DynamoDBAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &dynamodb.ListTagsOfResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsOfResourceWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, "ResourceNotFoundException") {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists dynamodb service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).DynamoDBConn(ctx), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns dynamodb service tags.
func Tags(tags tftags.KeyValueTags) []*dynamodb.Tag {
	result := make([]*dynamodb.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &dynamodb.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from dynamodb service tags.
func KeyValueTags(ctx context.Context, tags []*dynamodb.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns dynamodb service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*dynamodb.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets dynamodb service tags in Context.
func SetTagsOut(ctx context.Context, tags []*dynamodb.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates dynamodb service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn dynamodbiface.DynamoDBAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.DynamoDB)
	if len(removedTags) > 0 {
		input := &dynamodb.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     aws.StringSlice(removedTags.Keys()),
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.DynamoDB)
	if len(updatedTags) > 0 {
		input := &dynamodb.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates dynamodb service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).DynamoDBConn(ctx), identifier, oldTags, newTags)
}
