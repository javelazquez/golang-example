env-build:
	docker run -d -p 4566:4566 --name localstack \
    		-e SERVICES=sqs,sns,s3,dynamodb -e AWS_ACCESS_KEY_ID=user-test -e AWS_SECRET_ACCESS_KEY=pass-test -e AWS_DEFAULT_REGION=eu-west-1 \
    		localstack/localstack

env-setup:
	docker exec localstack aws dynamodb --endpoint-url=http://localhost:4566 create-table \
        	--table-name payout_table \
        	--attribute-definitions \
        	    AttributeName=id,AttributeType=S \
        	--key-schema \
        	    AttributeName=id,KeyType=HASH \
    		--provisioned-throughput \
        	    ReadCapacityUnits=10,WriteCapacityUnits=5

env-update-table:
	docker exec localstack aws dynamodb --endpoint-url=http://localhost:4566 update-table \
        --table-name payout_table \
        --attribute-definitions \
            AttributeName=id,AttributeType=S \
            AttributeName=external_id,AttributeType=S \
        --global-secondary-index-updates \
        "[{\"Create\":{\"IndexName\":\"ExternalIdIndex\",\"KeySchema\":[{\"AttributeName\":\"external_id\",\"KeyType\":\"HASH\"}],\"Projection\":{\"ProjectionType\":\"ALL\"},\"ProvisionedThroughput\":{\"ReadCapacityUnits\":10,\"WriteCapacityUnits\":5}}}]"



env-up:
	-docker start localstack

env-down:
	-docker stop localstack

env-clear:
	-docker rm localstack

get-table:
	docker exec localstack aws dynamodb --endpoint-url=http://localhost:4566 list-tables


get-all-dynamo:
	docker exec localstack aws dynamodb scan \
        --table-name payout_table \
        --endpoint-url=http://localhost:4566

get-by-dynamo-id:
	docker exec localstack aws dynamodb get-item \
    --table-name payout_table \
    --key '{"id": {"S": "6a9efb88-6c5c-484d-a35a-7deeed5b0526"}}' \
    --endpoint-url=http://localhost:4566
