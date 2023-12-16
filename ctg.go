package main

import (
	"flag"
	"fmt"

	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	cfn_inc "github.com/aws/aws-cdk-go/awscdk/v2/cloudformationinclude"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type config struct {
	template string
	Region   string
	Account  string
	Debug    bool
}

type CtgStackProps struct {
	cdk.StackProps
}

func CtgStack(scope constructs.Construct,
	id string, cfg config, props *CtgStackProps) cdk.Stack {

	var sprops cdk.StackProps

	if props != nil {
		sprops = props.StackProps
	}

	stack := cdk.NewStack(scope, &id, &sprops)

	cfnTemplate := cfn_inc.NewCfnInclude(
		stack, jsii.String("Template"), &cfn_inc.CfnIncludeProps{
			TemplateFile: jsii.String(cfg.template),
		})

	fmt.Println(cfnTemplate.Stack())

	return stack
}

type ValidateS3IsPrefixAspect struct {
	Prefix string
}

func (vpa *ValidateS3IsPrefixAspect) Visit(node constructs.IConstruct) {
	fmt.Println("Visiting resource: " + *node.Node().Id())

	if bucket, ok := node.(awss3.CfnBucket); ok { //&& strings.HasPrefix(*bucket.BucketName(), vpa.Prefix) {
		fmt.Println("BucketName " + *bucket.BucketName())
		cdk.Annotations_Of(node).AddInfo(jsii.String("Annotations: Each S3 Bucket name needs to start with " + vpa.Prefix))
	}
}

func NewValidateS3IsPrefixAspect(prefix string) *ValidateS3IsPrefixAspect {
	return &ValidateS3IsPrefixAspect{Prefix: prefix}
}

func main() {
	var cfg config

	flag.StringVar(&cfg.template, "template", "templates/s3.yaml", "The name of the cloudformation template to include")
	flag.StringVar(&cfg.Region, "region", "us-east-1", "The AWS region to deploy to")
	flag.StringVar(&cfg.Account, "account", "123456789012", "The AWS account to deploy to")

	defer jsii.Close()
	flag.Parse()

	app := cdk.NewApp(nil)
	CtgStack(app, "CdkGoStack", cfg, &CtgStackProps{
		cdk.StackProps{
			Env: &cdk.Environment{
				Account: jsii.String(cfg.Account),
				Region:  jsii.String(cfg.Region),
			},
		},
	})

	cdk.Aspects_Of(app).Add(NewValidateS3IsPrefixAspect("myPrefix"))

	app.Synth(nil)
}
