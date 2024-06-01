package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

type QueueService struct {
	svc      *sqs.SQS
	queueURL string
}

func NewSQSClient(queueURL string) *QueueService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	}))

	svc := sqs.New(sess)
	return &QueueService{svc: svc, queueURL: queueURL}
}

func (c *QueueService) SendMessage(ctx context.Context, messageBody string) error {
	_, err := c.svc.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(c.queueURL),
		MessageBody: aws.String(messageBody),
	})
	return err
}

func (c *QueueService) ReceiveMessages(ctx context.Context) ([]*sqs.Message, error) {
	result, err := c.svc.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(c.queueURL),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	})
	if err != nil {
		return nil, err
	}
	return result.Messages, nil
}

func (c *QueueService) DeleteMessage(ctx context.Context, receiptHandle *string) error {
	_, err := c.svc.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(c.queueURL),
		ReceiptHandle: receiptHandle,
	})
	return err
}
