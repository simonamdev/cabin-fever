rm -rf cabinserver-prod
docker build --no-cache -t cabinfever:prod .
container_id=$(docker create cabinfever:prod)
docker cp $container_id:/app/cabinserver ./cabinserver-prod
docker rm $container_id