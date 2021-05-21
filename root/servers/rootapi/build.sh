echo "create go executable for linux"
GOOS=linux go build

echo "build docker container"
docker build -t ngoa2/rootapi .

echo "delete go executable"
go clean