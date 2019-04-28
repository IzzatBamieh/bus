package main

import (
	"io/ioutil"
	"sync"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaNodeForgeService struct {
	client    *lambda.Lambda
	flowStore *FlowStore
	inventory *sync.Map
	roleARN   string
}

type CreateServiceCommand interface {
	GetName() string
}

func initializeRole(client client.ConfigProvider) (string, error) {
	var role *iam.Role
	iamClient := iam.New(client)
	getRoleResult, err := iamClient.GetRole(&iam.GetRoleInput{
		RoleName: aws.String("bus"),
	})
	if err != nil {
		if awsError, ok := err.(awserr.Error); ok {
			switch awsError.Code() {
			case iam.ErrCodeEntityAlreadyExistsException:
				break
			case iam.ErrCodeNoSuchEntityException:
				createRoleResult, err := iamClient.CreateRole(&iam.CreateRoleInput{
					RoleName:                 aws.String("bus"),
					Path:                     aws.String("/"),
					AssumeRolePolicyDocument: aws.String("AWSLambdaExecute"),
				})
				if err != nil {
					return "", err
				}
				role = createRoleResult.Role
			default:
				return "", err
			}
		} else {
			return "", err
		}
	} else {
		role = getRoleResult.Role
	}

	return *role.Arn, nil
}

func NewLambdaNodeForgeService(client client.ConfigProvider, flowStore *FlowStore) (*LambdaNodeForgeService, error) {
	roleARN, err := initializeRole(client)
	if err != nil {
		return nil, err
	}

	service := &LambdaNodeForgeService{
		client:    lambda.New(client),
		flowStore: flowStore,
		roleARN:   roleARN,
	}
	inventory, err := service.findAll()
	if err != nil {
		return nil, err
	}
	service.inventory = &sync.Map{}
	for _, name := range inventory {
		service.inventory.Store(name, true)
	}
	return service, nil
}

func (forge *LambdaNodeForgeService) findAll() ([]string, error) {
	names := []string{}
	result, err := forge.client.ListFunctions(&lambda.ListFunctionsInput{})
	if err != nil {
		return nil, errors.Wrap(err, "problem listing Lambda functions")
	}
	for _, function := range result.Functions {
		names = append(names, *function.FunctionName)
	}

	return names, nil
}

func (forge *LambdaNodeForgeService) delete(name string) error {
	if _, ok := forge.inventory.Load(name); !ok {
		return nil
	}
	_, err := forge.client.DeleteFunction(&lambda.DeleteFunctionInput{
		FunctionName: aws.String(name),
	})
	if err != nil {
		return errors.Wrap(err, "problem trying to overwrite function")
	}
	forge.inventory.Delete(name)

	return nil
}

func (forge *LambdaNodeForgeService) create(serviceName string) error {
	err := forge.delete(serviceName)
	if err != nil {
		return err
	}
	zip, err := ioutil.ReadFile("./server.zip")
	if err != nil {
		return err
	}
	input := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{
			ZipFile: zip,
		},
		Description:  aws.String("Auto generated Lambda function"),
		FunctionName: aws.String(serviceName),
		Handler:      aws.String("server"),
		MemorySize:   aws.Int64(128),
		Publish:      aws.Bool(true),
		Role:         aws.String(forge.roleARN),
		Runtime:      aws.String("go1.x"),
		Timeout:      aws.Int64(15),
		VpcConfig:    &lambda.VpcConfig{},

		Environment: &lambda.Environment{
			Variables: map[string]*string{
				"LAMBDA": aws.String("TRUE"),
			},
		},
	}
	_, err = forge.client.CreateFunction(input)
	return err
}

func (forge *LambdaNodeForgeService) CreateService(command CreateServiceCommand) error {
	return forge.create(command.GetName())
}

func (forge *LambdaNodeForgeService) ListServices() []string {
	services := []string{}
	forge.inventory.Range(func(key interface{}, value interface{}) bool {
		services = append(services, key.(string))
		return true
	})
	return services
}

func (forge *LambdaNodeForgeService) DeleteService(name string) error {
	return forge.delete(name)
}
