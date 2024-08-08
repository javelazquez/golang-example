package kvs

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type KVS interface {
	GetItem(ctx context.Context, keys map[string]any, m interface{}) error
	Save(ctx context.Context, item interface{}) error
}

type dynamoKVS struct {
	client *dynamodb.Client
	config Config
}

func NewDynamoKVS(config Config) *dynamoKVS {
	client := dynamodb.NewFromConfig(config.Credential.GetConfig())
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
		return fmt.Errorf("item not found")
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
