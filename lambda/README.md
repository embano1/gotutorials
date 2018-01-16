# Inside AWS Lambda with Go

Simple function to respond Go runtime environment stats, like `NumCPU`. Uses AWS API Gateway.

Follow [this tutorial](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/) from AWS to get started with setting up your function.  
**Note:** the name of the Lambda handler must be `hello` (matching the name of the Go binary in the .zip)


## Build function
```bash
$ git clone https://github.com/embano1/gotutorials/
$ cd lambda
$ sh build.sh
```

Then upload `lambda.zip`.

# Query AWS API Gateway endpoint
```bash
# Uses a 2s timeout
$ curl -I -m 2 https://<ENDPOINT>
```