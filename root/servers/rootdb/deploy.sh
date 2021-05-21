bash build.sh

echo "push container image to dockerhub"
docker push ngoa2/rootdb

echo "ssh into api server"
ssh ec2-user@api.root.quest << EOF
echo "stop existing container"
docker rm -f rootdb
echo "pull from dockerhub"
docker pull ngoa2/rootdb
echo "run new instance of container"

docker run -d \
-p 3306:3306 \
--name rootdb \
-e MYSQL_ROOT_PASSWORD="SSAJpass" \
-e MYSQL_DATABASE=rootdb \
ngoa2/rootdb
EOF
