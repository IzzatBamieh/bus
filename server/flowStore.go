package main

import (
	"bufio"
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/s3"
	set "github.com/deckarep/golang-set"
	jsoniter "github.com/json-iterator/go"
)

type FlowStore struct {
	s3Client *s3.S3
}

type FlowState struct {
	Services []string
}

func NewFlowStore(provider client.ConfigProvider) *FlowStore {
	set.NewSet()
	return &FlowStore{
		s3Client: s3.New(provider),
	}
}

func NewFlowState() *FlowState {
	return &FlowState{
		Services: []string{},
	}
}

func (store *FlowStore) GetFlowState() (*FlowState, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String("keithhassen-flow"),
		Key:    aws.String("state.json"),
	}
	result, err := store.s3Client.GetObject(input)
	if err != nil {
		if awsError, ok := err.(awserr.Error); ok {
			switch awsError.Code() {
			case s3.ErrCodeNoSuchBucket, "BucketRegionError":
				_, err := store.s3Client.CreateBucket(&s3.CreateBucketInput{
					Bucket: aws.String("keithhassen-flow"),
				})
				if err != nil {
					return nil, err
				}
			case s3.ErrCodeNoSuchKey:
			default:
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	state := NewFlowState()
	jsoniter.NewDecoder(result.Body).Decode(state)

	return state, nil
}

func (store *FlowStore) SetFlowState(state *FlowState) error {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	err := jsoniter.NewEncoder(writer).Encode(state)
	if err != nil {
		return err
	}
	input := &s3.PutObjectInput{
		Bucket: aws.String("flow"),
		Key:    aws.String("state.json"),
		Body:   bytes.NewReader(b.Bytes()),
	}
	_, err = store.s3Client.PutObject(input)
	return err
}
