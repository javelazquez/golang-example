package kvs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"golang-example/internal/appErrors"
)

type KVS interface {
	GetItem(ctx context.Context, keys map[string]any, m interface{}) error
	SaveConditional(ctx context.Context, item interface{}) error
	Save(ctx context.Context, item interface{}) error
	GetByExternalID(ctx context.Context, externalID string) (map[string]types.AttributeValue, error)
}

type dynamoKVS struct {
	client *dynamodb.Client
	config Config
}

func NewDynamoKVS(config Config) *dynamoKVS {
	//client := dynamodb.NewFromConfig(config.Credential.GetConfig())
	client := dynamodb.New(dynamodb.Options{
		Credentials:  config.Credential.GetConfig().Credentials,
		Region:       config.Credential.GetConfig().Region,
		BaseEndpoint: aws.String(config.AWSEndpoint),
	})
	return &dynamoKVS{
		client: client,
		config: config,
	}
}

func (d *dynamoKVS) GetItem(ctx context.Context, keys map[string]any, m interface{}) error {
	av, err := attributevalue.MarshalMap(keys)
	if err != nil {
		return err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(d.config.TableName),
		Key:       av,
	}

	result, err := d.client.GetItem(ctx, input)
	if err != nil {
		return err
	}

	if result.Item == nil {
		return nil
	}

	err = attributevalue.UnmarshalMap(result.Item, m)
	if err != nil {
		return err
	}

	return nil
}

func (d *dynamoKVS) Save(ctx context.Context, item interface{}) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(d.config.TableName),
		Item:      av,
	}

	_, err = d.client.PutItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (d *dynamoKVS) SaveConditional(ctx context.Context, item interface{}) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	// Extract the external_id from the item
	var externalID string
	if v, ok := av["external_id"]; ok {
		if s, ok := v.(*types.AttributeValueMemberS); ok {
			externalID = s.Value
		} else {
			return fmt.Errorf("unexpected type for external_id")
		}
	} else {
		return fmt.Errorf("external_id not found in item")
	}

	// Check if the external_id already exists
	existingItem, err := d.GetByExternalID(ctx, externalID)
	if err != nil {
		var notFoundErr *appErrors.NotFoundError
		if !errors.As(err, &notFoundErr) {
			return err
		}
	}

	if existingItem != nil {
		return appErrors.NewAlreadyExistError(fmt.Sprintf("item already exists: %s", externalID))
	}

	// Proceed to put the item into DynamoDB
	putInput := &dynamodb.PutItemInput{
		TableName: aws.String(d.config.TableName),
		Item:      av,
	}

	_, err = d.client.PutItem(ctx, putInput)
	if err != nil {
		return err
	}

	return nil
}

func (d *dynamoKVS) GetByExternalID(ctx context.Context, externalID string) (map[string]types.AttributeValue, error) {
	// Prepare the query input
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(d.config.TableName),
		IndexName:              aws.String("ExternalIdIndex"), // Specify the index name
		KeyConditionExpression: aws.String("external_id = :external_id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":external_id": &types.AttributeValueMemberS{Value: externalID},
		},
	}

	// Execute the query
	result, err := d.client.Query(ctx, queryInput)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, appErrors.NewNotFoundError(fmt.Sprintf("item with external_id %s not found", externalID))
	}

	// Assuming only one item per external_id
	return result.Items[0], nil
}
