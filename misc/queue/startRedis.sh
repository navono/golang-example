#! /bin/bash
# startRedis.sh

docker run -p 6379:6379 --name queue-redis -d redis redis-server --appendonly yes
