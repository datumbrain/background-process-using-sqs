# Running Lambda Asynchorously Using Amazon SQS
Here we are going to use `Amazon SQS` to run labmbda in background

## Prerequisites
- You’ll need an AWS account for this. If you don’t yet have one, [sign up for a free account here](https://aws.amazon.com/free/?all-free-tier.sort-by=item.additionalFields.SortRank&all-free-tier.sort-order=asc&awsf.Free%20Tier%20Types=*all&awsf.Free%20Tier%20Categories=*all).
- If you don’t have Go installed yet, you can either [download an installer](https://golang.org/dl/) from the official website or use your favorite package manager to install it.
- For building and deploying your functions, you’ll be using the `Serverless Framework`. Assuming you have a recent version of [Node.js](https://nodejs.org/en/) installed, you can install the `Serverless CLI` with the following npm command
	
	```console
	  $ npm install -g serverless
	```
	Once you have the Serverless CLI installed, you must configure it to use the AWS access keys of your account
	```console
	  $ serverless config credentials --provider aws --key <access key ID> --secret <secret access key>
	```
- Create SQS in your amazon account and `replace your sqs arn` with the last line of `serverless.yml` where it mentioned.
## Build and Deploy
```console
  $ scripts/deploy.sh
```
