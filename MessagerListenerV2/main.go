package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func GetClient() *azservicebus.Client {
	//namespace, ok := os.LookupEnv("AZURE_SERVICEBUS_HOSTNAME") //ex: myservicebus.servicebus.windows.net
	//if !ok {
	//	panic("AZURE_SERVICEBUS_HOSTNAME environment variable not found")
	//}
	namespace := "viverorganizer.servicebus.windows.net"

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	client, err := azservicebus.NewClient(namespace, cred, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func SendMessage(message string, client *azservicebus.Client) {
	sender, err := client.NewSender("nop.to.mailmarketing.qa", nil)
	if err != nil {
		panic(err)
	}
	defer sender.Close(context.TODO())

	sbMessage := &azservicebus.Message{
		Body: []byte(message),
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		panic(err)
	}
}

func GetMessage(count int, client *azservicebus.Client) {
	receiver, err := client.NewReceiverForQueue("myqueue", nil)
	if err != nil {
		panic(err)
	}
	defer receiver.Close(context.TODO())

	messages, err := receiver.ReceiveMessages(context.TODO(), count, nil)
	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		body := message.Body
		fmt.Printf("%s\n", string(body))

		err = receiver.CompleteMessage(context.TODO(), message, nil)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	client := GetClient()

	fmt.Println("send a single message...")
	SendMessage("firstMessage", client)

	fmt.Println("\nget all three messages:")
	GetMessage(3, client)
}
