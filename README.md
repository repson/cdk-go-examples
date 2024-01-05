# Welcome to cloud-development-libraries-tests

# cdk-aws

## python

```
$ source .venv/bin/activate
$ python -m pip install -r requirements.txt

$ cdk synth
$ cdk deploy
```

Using localstack to deploy the resources:

```
$ npm install -g aws-cdk-local
$ source .venv/bin/activate
$ python -m pip install -r requirements.txt
$ pip install localstack

$ echo 'export AWS_ENDPOINT_URL="http://localhost:4566"' >> .bashrc
$ source $HOME/.bashrc
$ localstack start

$ cdklocal synth
$ cdklocal bootstrap
$ cdklocal deploy

$ awslocal s3 ls
```

## golang

```
$ npm install aws-cdk-lib
$ go mod tidy
$ go run ctg.go -template templates/s3.yaml
```

# cdk-tf

```
$ npm install --global cdktf-cli@latest
$ cdktf init --template=go --providers=kreuzwerker/docker --local
$ go mod tidy

$ cdktf deploy
$ docker ps
$ cdktf destroy
```

# pulumi

```
$ curl -fsSL https://get.pulumi.com | sh
$ pulumi new go
$ go mod tidy
$ pulumi config set gophersAPIPort 8080
$ pulumi config set gophersAPIWatcherPort 8000
$ pulumi up
$ docker container ls
$ curl localhost:8080/gophers
$ pulumi destroy
```