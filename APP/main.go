package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	sq *sqs.SQS
)

func pollMessages(channel chan<- *sqs.Message) {

	for {
		output, err := sq.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String("https://sqs.eu-central-1.amazonas.com/12345678/suaQueue"),
			MaxNumberOfMessages: aws.Int64(2),
			WaitTimeSeconds:     aws.Int64(15),
		})
		if err != nil {
			fmt.Println("failed to fetch sqs message %v", err)
		}

		for _, message := range output.Messages {

			channel <- message
		}
	}

}

func deleteMessage(msg *sqs.Message) {

}

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{Config: aws.Config{Region: aws.String("eu-central-1")}, Profile: "xyz"}))

	sq = sqs.New(sess)

	channelOfMessagens := make(chan *sqs.Message, 2)

	go pollMessages(channelOfMessagens)

	for message := range channelOfMessagens {
		fmt.Println("RECEIVING MESSAGE >>> ")
		fmt.Println(message.Body)

		deleteMessage(message)

	}

}
