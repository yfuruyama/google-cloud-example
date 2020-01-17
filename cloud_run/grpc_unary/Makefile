IMAGE=gcr.io/${PROJECT_ID}/cloudrun_grpc_unary

proto:
	protoc proto/hello.proto --go_out=plugins=grpc:.

build:
	docker build -t ${IMAGE} .

push: build
	docker push ${IMAGE}

deploy: push
	gcloud run deploy --image=${IMAGE} --platform=managed --region=us-central1 --project=${PROJECT_ID} --allow-unauthenticated cloudrun-grpc-unary
