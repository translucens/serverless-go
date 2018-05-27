# serverless-go
Apps for serverless architecture

## Prerequisites

```bash
pip install...

```

1. Create SNS topic for SES; it should be located in same region as SES, i.e. us-east-1, us-west-2, or eu-west-1
1. Setup SES Email Receiving
    * Create rule and specify SNS topic; Encoding: `BASE64`
1. Create SNS topic for proxying; it should be located in same region as other Lambda functions
1. Create S3 Bucket for deploying; it should be located in same region as Lambda functions
1. Create Slack incoming webhook URL via <https://my.slack.com/services/new/incoming-webhook/>

```bash
S3BUCKET_BIN="YOUR_S3_BACKET_NAME"

SNS_TOPIC_PUB_MAIL="SNS_TOPIC_NAME"
SNS_TOPIC_ARN_PUB_MAIL="arn:aws:sns:REGION:ACCOUNT_NUMBER:$SNS_TOPIC_PUB_MAIL"

SLACK_WEBHOOK="https://hooks.slack.com/YOUR_WEBHOOK_URL"
SNS_TOPIC_ARN_SUB_SLACK=$SNS_TOPIC_ARN_PUB_MAIL

GOOS=linux
GOARCH=amd64

cd sns2slack
go build -o bin/sns2slack
sam package \
    --template-file sns2slack-sam.yaml \
    --output-template-file packaged.yaml \
    --s3-bucket $S3BUCKET_BIN
sam deploy \
    --template-file packaged.yaml \
    --stack-name lambda-sns2slack \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides SlackWebHookURL=$SLACK_WEBHOOK \
    PostSlackTemplate=hogehoge \
    EventSourceSNSARN=$SNS_TOPIC_ARN_SUB_SLACK

cd ../ses2sns
go build -o bin/ses2sns
sam package \
    --template-file ses2sns-sam.yaml \
    --output-template-file packaged.yaml \
    --s3-bucket $S3BUCKET_BIN
sam deploy \
    --template-file packaged.yaml \
    --stack-name lambda-ses2sns \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides SNSTopicARNPubMail=$SNS_TOPIC_ARN_PUB_MAIL \
    SNSTopicPubMail=$SNS_TOPIC_PUB_MAIL
```

## Required IAM Permissions for deploying

Console or CI server should have the permissions below:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "cloudformation:CreateChangeSet",
                "cloudformation:ExecuteChangeSet",
                "cloudformation:GetTemplateSummary"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "iam:DetachRolePolicy",
                "iam:CreateRole",
                "iam:DeleteRole",
                "iam:PutRolePolicy",
                "iam:AttachRolePolicy",
                "iam:DeleteRolePolicy"
            ],
            "Resource": "arn:aws:iam::*:role/lambda-*"
        }
    ]
}
```
