package main

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/pkg/errors"
)

type Bus struct {
	logger    *Logger
	snsClient *sns.SNS
	sqsClient *sqs.SQS
	topics    sync.Map
	queues    sync.Map
}

func NewBus(logger *Logger, provider client.ConfigProvider) (*Bus, error) {
	return &Bus{
		logger:    logger,
		snsClient: sns.New(provider),
		sqsClient: sqs.New(provider),
		topics:    sync.Map{},
		queues:    sync.Map{},
	}, nil
}

func (bus *Bus) Send(topicName string, body []byte) error {
	topicArn, err := bus.ensureTopic(topicName)
	if err != nil {
		return err
	}
	_, err = bus.snsClient.Publish(&sns.PublishInput{
		TopicArn: aws.String(topicArn),
		Message:  aws.String(string(body)),
	})
	return errors.Wrap(err, "problem publishing message")
}

func (bus *Bus) Receive(topicName string, consumerGroupName string) (*Receiver, error) {
	queueName := fmt.Sprintf("%s_%s", topicName, consumerGroupName)
	topicArn, err := bus.ensureTopic(topicName)
	if err != nil {
		return nil, err
	}
	queueURL, err := bus.ensureQueue(topicArn, queueName)
	if err != nil {
		return nil, err
	}
	return newReceiver(bus.logger, bus.sqsClient, queueURL), nil
}

func (bus *Bus) ensureTopic(topicName string) (string, error) {
	if topic, ok := bus.topics.Load(topicName); ok {
		return *topic.(*sns.CreateTopicOutput).TopicArn, nil
	}
	input := &sns.CreateTopicInput{
		Name: aws.String(topicName),
	}
	result, err := bus.snsClient.CreateTopic(input)
	if err != nil {
		return "", errors.Wrap(err, "problem creating topic")
	}

	bus.topics.Store(topicName, result)
	return *result.TopicArn, nil
}

func (bus *Bus) ensureQueue(topicArn string, queueName string) (string, error) {
	if queue, ok := bus.queues.Load(queueName); ok {
		return *queue.(*sqs.CreateQueueOutput).QueueUrl, nil
	}

	input := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}
	result, err := bus.sqsClient.CreateQueue(input)
	if err != nil {
		return "", errors.Wrap(err, "problem creating queue")
	}
	queueUrl := result.QueueUrl
	attributes, err := bus.sqsClient.GetQueueAttributes(&sqs.GetQueueAttributesInput{
		QueueUrl: queueUrl,
		AttributeNames: []*string{
			aws.String("QueueArn"),
		},
	})
	if err != nil {
		// TODO: remove queue
		return "", errors.Wrap(err, "problem getting queue attributes")
	}
	queueArn, ok := attributes.Attributes["QueueArn"]
	if !ok {
		return "", errors.New("problem getting queue attributes")
	}

	subscription, err := bus.snsClient.Subscribe(&sns.SubscribeInput{
		TopicArn: aws.String(topicArn),
		Protocol: aws.String("sqs"),
		Endpoint: queueArn,
		Attributes: map[string]*string{
			"RawMessageDelivery": aws.String("true"),
		},
	})
	if err != nil {
		// TODO: remove queue
		return "", errors.Wrap(err, "problem subscribing queue to topic")
	}
	_, err = bus.snsClient.SetSubscriptionAttributes(&sns.SetSubscriptionAttributesInput{
		SubscriptionArn: subscription.SubscriptionArn,
		AttributeName:   aws.String("RawMessageDelivery"),
		AttributeValue:  aws.String("true"),
	})
	if err != nil {
		return "", errors.Wrap(err, "problem configuring topic subscription")
	}

	policy := fmt.Sprintf(`{
		"Version": "2008-10-17",
		"Statement": [{
			"Effect": "Allow",
			"Principal": "*",
			"Action": ["sqs:SendMessage"],
			"Resource":["%s"],
			"Condition": {
				"ArnLike":{"aws:SourceArn":["%s"]}
			}
		}]
	}
	`, *queueArn, topicArn)
	_, err = bus.sqsClient.SetQueueAttributes(&sqs.SetQueueAttributesInput{
		QueueUrl: queueUrl,
		Attributes: map[string]*string{
			"Policy": aws.String(policy),
		},
	})
	if err != nil {
		// TODO: remove queue
		return "", errors.Wrap(err, "problem setting queue attributes")
	}

	bus.queues.Store(queueName, result)
	return *result.QueueUrl, nil
}
