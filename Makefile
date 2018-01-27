.Phony: build clean deploy package start-api test zip

CFN_INPUT_TEMPLATE := ./cloudformation/template.yml
CFN_PKG_TEMPLATE := ./artifacts/packaged-template.yml

artifacts/main: main.go
	GOOS=linux go build -o artifacts/main

build: artifacts/main

clean:
	rm -r artifacts

deploy: package
	aws cloudformation deploy \
		--template-file $(CFN_PKG_TEMPLATE) \
		--stack-name $(STACK_NAME) \
		--capabilities CAPABILITY_IAM

package: zip
	aws cloudformation package \
		--template-file $(CFN_INPUT_TEMPLATE) \
		--s3-bucket $(S3_BUCKET) \
		--output-template-file $(CFN_PKG_TEMPLATE)

start-api: zip
	sam local start-api --template $(CFN_INPUT_TEMPLATE)

test:
	go test

zip: build
	zip -j artifacts/build.zip artifacts/main
