# grpcsimple

deploy:
docker-compose up

test:
curl -XGET 'http://localhost:8001/fibonacci?from=1&to=10'
