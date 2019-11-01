# credit https://github.com/dwmkerr/docker-dynamodb/blob/master/samples/test-data/setup.sh

DATADIR="$(pwd)/data"
rm -rf $DATADIR; mkdir $DATADIR
ID=$(docker run -d -p 8000:8000 -v $DATADIR:/data/ amazon/dynamodb-local -dbPath /data/)

aws dynamodb create-table \
  --table-name cache_dev \
  --attribute-definitions \
      AttributeName=key,AttributeType=S \
  --key-schema AttributeName=key,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
  --endpoint-url http://localhost:8000

docker stop $ID && docker rm $ID || true
