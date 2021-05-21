echo "run build.sh"
bash build.sh

echo "push container image"
docker push ngoa2/rootapi

echo "ssh into api-server"
ssh ec2-user@api.root.quest << EOF

echo "stop existing container"
docker rm -f rootapi

echo "pull from dockerhub"
docker pull ngoa2/rootapi

echo "run new instance of container"
docker run -d \
--name rootapi \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=/etc/letsencrypt/live/api.root.quest/fullchain.pem \
-e TLSKEY=/etc/letsencrypt/live/api.root.quest/privkey.pem \
ngoa2/rootapi
EOF