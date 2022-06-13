build-image:
	docker build --file=Dockerfile --tag=gurygury/test-hello:latest .

start-container:
	docker run --env-file .env -p 8080:8080 gurygury/test-hello:latest