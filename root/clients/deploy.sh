echo "build web client"
bash build.sh

echo "push container image to dockerhub"
docker push ngoa2/root_client

echo "ssh into root client server"
ssh ec2-user@root.quest << EOF

echo "stop existing container"
docker rm -f root_client

echo "pull from dockerhub"
docker pull ngoa2/root_client

export TLSCERT=/etc/letsencrypt/live/ngoa.tech/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/ngoa.tech/privkey.pem

echo $TLSKEY
echo $TLSCERT

echo "run new instance of container"
docker run -d \
--name root_client \
-p 443:443 \
-p 80:80 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
ngoa2/root_client

EOF