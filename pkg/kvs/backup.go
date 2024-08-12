package kvs

/*
func (d *dynamoKVS) SaveConditional(ctx context.Context, item interface{}) error {
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
}*/

/*
func (d *dynamoKVS) SaveConditional(ctx context.Context, item interface{}) error {
	// Marshall the item into a DynamoDB attribute map
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	// Create the condition expression to ensure the external_id does not already exist
	condition := "attribute_not_exists(external_id)"

	// Prepare the input for the PutItem operation
	input := &dynamodb.PutItemInput{
		TableName:           aws.String(d.config.TableName),
		Item:                av,
		ConditionExpression: aws.String(condition),
	}

	// Attempt to put the item into DynamoDB with the condition
	_, err = d.client.PutItem(ctx, input)
	if err != nil {
		// Check if the error is a conditional check failure
		var cfe *types.ConditionalCheckFailedException
		if errors.As(err, &cfe) {
			return appErrors.NewAlreadyExistError("external_id already exists")
		}
		return err
	}

	return nil
}*/
