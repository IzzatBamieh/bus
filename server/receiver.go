package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/pkg/errors"
)

type Receiver struct {
	logger    *Logger
	sqsClient *sqs.SQS
	queueURL  string
	available []*sqs.Message
}

type Message struct {
	Body []byte
	Ack  func() error
}

func (receiver *Receiver) Next() (*Message, error) {
	for len(receiver.available) == 0 {
		result, err := receiver.sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(receiver.queueURL),
			MaxNumberOfMessages: aws.Int64(10),
			VisibilityTimeout:   aws.Int64(10),
			WaitTimeSeconds:     aws.Int64(10),
		})
		if err != nil {
			return nil, err
		}

		receiver.available = append(receiver.available, result.Messages...)
	}

	next := receiver.available[0]
	receiver.available = receiver.available[1:]

	return newMessage(*next.Body, func() error {
		_, err := receiver.sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      aws.String(receiver.queueURL),
			ReceiptHandle: next.ReceiptHandle,
		})
		return errors.Wrap(err, "problem ACKing message")
	}), nil
}

func newReceiver(logger *Logger, sqsClient *sqs.SQS, queueURL string) *Receiver {
	return &Receiver{
		logger:    logger,
		sqsClient: sqsClient,
		queueURL:  queueURL,
		available: []*sqs.Message{},
	}
}

func newMessage(body string, ack func() error) *Message {
	return &Message{
		Body: []byte(body),
		Ack:  ack,
	}
}
