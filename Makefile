build-image:
	docker build --file=Dockerfile --tag=test-hello:latest .

start-container:
	docker run --env-file .env -p 8080:8080 test-hello:latest