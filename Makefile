build-server:
	@docker buildx build -t chat-server -f ./server/dockerfile ./server

build-client:

clean-images:
	@docker rmi chat-server

clean-containers:
	@docker stop chat-server

run-server:
	@docker run -d --name chat-server -p 6969:6969 --rm chat-server

run: