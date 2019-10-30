aws dynamodb create-table \
  --table-name cache_dev \
  --attribute-definitions \
      AttributeName=key,AttributeType=S \
  --key-schema AttributeName=key,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
  --endpoint-url http://localhost:8000
