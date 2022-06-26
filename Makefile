GO=${GOROOT}/bin/go

run-mongodb:
	docker run -it --rm --name mongodb-server -e MONGODB_DATABASE=test -p 27017:27017 mongo