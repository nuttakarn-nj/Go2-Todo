.PHONY: build maria

build:
	go build -o app --ldflags "-X main.buildCommit=`git rev-parse --short HEAD` -X main.buildTime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`"

maria:
	docker run -p 127.0.0.1:3306:3306  --name some-mariadb \
	-e MARIADB_ROOT_PASSWORD=my-secret-pw -e MARIADB_DATABASE=myapp -d mariadb:latest

image:
	docker build -t todo:test -f Dockerfile .

container:
	docker run -p:8080:8080 --env-file ./local.env --link some-mariadb:db \
	--name myapp todo:test