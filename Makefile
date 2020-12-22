NAME=connect-service
TAG=diogomonte/homeautomation-connect-service
VER=v0.0.1

build-prod-connect:
	docker build -t $(TAG) -t $(TAG):$(VER) -f services/connect-service/prod.dockerfile .

push-prod-connect: build-prod-connect
	docker push $(TAG):$(VER)

local:
	docker run -d -p 8080:8080 --name=$(NAME) $(TAG)

stop:
	docker stop $(NAME)
	docker rm $(NAME)