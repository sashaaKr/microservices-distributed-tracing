# microservices-distributed-tracing

## Getting started

```bash
mkdir -p $GOPATH/src/github.com/sashaaKr
cd $GOPATH/src/github.com/sashaaKr
git clone git@github.com:sashaaKr/microservices-distributed-tracing.git

cd go/
dep ensure

docker run -d --name mysql56 -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=mysqlpwd mysql:5.6

# populate db
docker exec -i mysql56 mysql -uroot -pmysqlpwd < ./database.sql

docker run -d --name jaeger \
    -p 6831:6831/udp \
    -p 16686:16686 \
    -p 14268:14268 \
    jaegertracing/all-in-one:1.6

curl http://localhost:8080/sayHello/Nefario
```