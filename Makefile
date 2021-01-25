
BUCKET="cloudformation-package-bucket-imishinist"

.PHONY: build
build: main.go
	go build -o build/test/test .

.PHONY: package
package: template.yaml
	aws cloudformation package \
		--template-file template.yaml \
		--s3-bucket $(BUCKET) \
		--s3-prefix "sqs-batch-window-sample" \
		--output-template-file .template.yaml

.PHONY: deploy
deploy: .template.yaml
	aws cloudformation deploy \
		--stack-name sqs-batch-window-sample \
		--template-file .template.yaml \
		--capabilities CAPABILITY_IAM \
		--no-fail-on-empty-changeset
