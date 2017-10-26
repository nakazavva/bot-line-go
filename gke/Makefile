PROJECT_ID = civil-honor-183803
CLUSTER = bot
SECRET = line
IMAGE = gcr.io/$(PROJECT_ID)/go-bot-line
PORT ?= 9000
CHANNEL_SECRET=$(shell cat .channel_secret)
CHANNEL_TOKEN=$(shell cat .channel_token)

.PHONY: all build run push deploy clean destroy

all: build

build:
	docker build -t $(IMAGE) ../src

run:
	docker run --rm -p $(PORT):9000 \
		--env CHANNEL_SECRET=$(CHANNEL_SECRET) \
		--env CHANNEL_TOKEN=$(CHANNEL_TOKEN) \
		$(IMAGE)

push:
	gcloud docker -- push $(IMAGE)

deploy:
	- kubectl create secret generic $(SECRET) --from-file=.channel_secret --from-file=.channel_token
	kubectl apply -f deployment.yml
	kubectl apply -f service.yml

clean:
	kubectl delete -f service.yml
	kubectl delete -f deployment.yml
	kubectl delete secret $(SECRET)

destroy:
	gcloud beta container images delete $(IMAGE)
	docker rmi $(IMAGE)
