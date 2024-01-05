# Welcome to cloud-development-libraries-tests

# cdk-aws

```
$ npm install aws-cdk-lib
$ go mod tidy
$ cdk synth
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