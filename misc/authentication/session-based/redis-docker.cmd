docker run -d -p 6379:6379 --name redis1 ^
  -v E:\\data\\redis:/data ^
  redis:latest redis-server --appendonly yes

REM docker exec -it redis1 redis-cli
