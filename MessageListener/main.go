package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	servicebus "github.com/Azure/azure-service-bus-go"
)

type Worker struct {
	queueClient          *servicebus.Queue
	mailMarketingService MailMarketingService
}

func NewWorker(mailMarketingService MailMarketingService, connectionString string, queueName string) (*Worker, error) {

	namespace, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connectionString))
	if err != nil {
		return nil, err
	}
	queueClient, err := namespace.NewQueue(queueName)

	//queueClient, err := servicebus.NewQueueClientWithNamespace(connectionString, queueName)
	if err != nil {
		return nil, err
	}

	return &Worker{
		queueClient:          queueClient,
		mailMarketingService: mailMarketingService,
	}, nil
}

func (w *Worker) Run() error {
	handlerOptions := servicebus.HandlerFuncOptions{
		MaxConcurrentCalls:   1,
		MaxAutoRenewDuration: 10 * time.Minute,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w.queueClient.Receive(ctx, w.ProcessMessage, handlerOptions)

	return nil
}

func (w *Worker) Stop() error {
	return w.queueClient.Close(context.Background())
}

func (w *Worker) ProcessMessage(ctx context.Context, message *servicebus.Message) error {
	messageStr := string(message.Data)
	fmt.Println(messageStr)

	var messageData string
	if err := json.Unmarshal(message.Data, &messageData); err != nil {
		log.Printf("Error unmarshaling message data: %v", err)
		return err
	}

	ret, err := w.mailMarketingService.AddMailMarketing(messageData)
	if err != nil {
		log.Printf("Error adding mail marketing: %v", err)
		return err
	}

	if ret {
		return message.Complete(ctx)
	} else {
		return message.DeadLetter(ctx, "invalid data")
	}
}

type MailMarketingService interface {
	AddMailMarketing(data string) (bool, error)
}

type MailMarketingServiceImpl struct {
}

func (s *MailMarketingServiceImpl) AddMailMarketing(data string) (bool, error) {
	log.Printf("AddMailMarketing data: %v", data)
	return true, nil
}

func main() {
	connectionString := os.Getenv("QueueConnStr")
	queueName := os.Getenv("QueueNameStr")

	mailMarketingService := &MailMarketingServiceImpl{}

	worker, err := NewWorker(mailMarketingService, connectionString, queueName)
	if err != nil {
		log.Fatalf("Error creating worker: %v", err)
	}

	if err := worker.Run(); err != nil {
		log.Fatalf("Error running worker: %v", err)
	}
}
