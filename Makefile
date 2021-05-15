run:
	go run cmd/app/main.go start -e env.yaml
all:
	make build & make push
build:
	CGO_ENABLED=0 GOOS=linux go build -o app cmd/app/main.go
	docker build . -t 注册用户名/镜像名:版本
	rm -rf app
push:
    docker push 注册用户名/镜像名:版本
start:
	docker pull 注册用户名/镜像名:版本
	docker run -d --name tia -p 8000:8000 -p 8001:8001 注册用户名/镜像名:版本 app start
stop:
	docker rm -f tia