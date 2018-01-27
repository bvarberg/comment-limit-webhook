# comment-limit-webhook

> A personally-deployable webhook for warning when comments on a GitHub issue
> exceed a given limit.

## Concept

A webhook for GitHub *could* be simple to deploy and maintain by taking
advantage of AWS services like Lambda, API Gateway, and CloudFormation.

While this project implements a specific webhook, the goals of this experiment
are more general. Any kind of functionality could be used in place of what is
included here. The focus of this project is the infrastructure that enables and
supports that functionality.

## Local development

One option for development and debugging is to use the [`sam` CLI][sam-local]
tool. The package, installed via NPM, can be used to run an API Gateway instance
configured by the template included with this project. For convenience, a `make
start-api` target is included; however, it expects that the `sam` tool is
already installed on your device.

```sh
$ make start-api
```

## Deployment

In order to deploy this webhook, the executable must be generated, zipped, and
prepared for deployment by the `aws` CLI. These steps are combined in the `make
deploy` target.

Input must be provided to specify which S3 bucket should hold the packaged code
bundle, along with the name of the CloudFormation stack to create or update.

```sh
$ S3_BUCKET=bvarberg-webhooks STACK_NAME=CommentLimitWebhook make deploy
```

> **Note:** The stack will attempt to create API Gateway, Lambda, and IAM
> resources. To learn more about what will be created, see the [SAM][aws-sam]
> documentation.

[aws-sam]: https://github.com/awslabs/serverless-application-model
[sam-local]: https://github.com/awslabs/aws-sam-local
