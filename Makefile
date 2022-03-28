-include .env

VERSION_FILE := config/version
VERSION := $(shell cat ${VERSION_FILE})

# build project -out ./jwter
build:
	- rm ./jwter-*
	go build -o jwter-${VERSION}

# compile and run jwter
run:
	go run .

docker-build:
	 docker build --tag jwter-docker .

docker-run:
	docker run -it --net=boomer_default -p $(PROXY_PORT):$(PROXY_PORT) --name jwter-proxy jwter-docker

docker-up: docker-build docker-run

docker-stop:
	docker stop jwter-proxy

docker-down: docker-stop
	docker rm jwter-proxy

# remove built
down:
	- rm ./jwter-*

# generate signing and public certificates in cert/ dir
certs:
	- mkdir certs 2>/dev/null
	openssl genrsa -out certs/jwter-signing-key.pem 4096
	openssl req -x509 -new -key certs/jwter-signing-key.pem -days 365 -out certs/jwter-pub-cert.pem -subj "/"

# run interactive script to set up .env
conf:
	sh config/config_setup.sh

# increment version from version file
#update-version:
#	sed -i '1s/.*/${new_ver}/' version

release: build
	git add ./jwter-*
	git commit jwter -m 'new jwter build - no need to thank me ;)'
	git push
