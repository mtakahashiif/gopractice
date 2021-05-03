#!/bin/bash -ux

BASE_DIR=$(cd $(dirname $0); pwd)

docker stop test-web-server
docker rm test-web-server
docker run \
    --name test-web-server \
    -v ${BASE_DIR}/nginx/docroot/:/usr/share/nginx/html:ro \
    -v ${BASE_DIR}/nginx/conf.d/:/etc/nginx/conf.d \
    -v ${BASE_DIR}/nginx/cert/:/etc/ssl/localcerts/nignx \
    -d \
    -p 80:80 \
    -p 443:443 \
    nginx
