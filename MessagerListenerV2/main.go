package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetClient() *azservicebus.Client {
	namespace, ok := os.LookupEnv("SBCONNSTR")
	if !ok {
		panic("SBCONNSTR environment variable not found")
	}
	client, err := azservicebus.NewClientFromConnectionString(namespace, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func insertMktLead(name string, email string) {
	mailMktConnStr, ok := os.LookupEnv("MKTDBCONNSTR")
	if !ok {
		panic("MKTDBCONNSTR environment variable not found")
	}
	Mktconn, err := sql.Open("mssql", mailMktConnStr)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = Mktconn.PingContext(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	tsql := fmt.Sprintf("SELECT [Name] FROM [EmailMarketing].[MailList] where [email] = '" + email + "';")

	// Execute query
	rows, err := Mktconn.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Erro ao pesquisar lead:", err)
	}

	defer rows.Close()

	// Iterate through the result set.
	if rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println("Erro ao carregar o lead", err)
		}
		fmt.Println("Lead:", name)
	} else {
		//ctx := context.Background()
		fmt.Println("Lead not found! ", email)
		tsql := "INSERT INTO [EmailMarketing].[MailList] ([Name], [email]) VALUES ('" + name + "', '" + email + "'); select isNull(SCOPE_IDENTITY(), -1);"

		stmt, err := Mktconn.Prepare(tsql)
		if err != nil {
			fmt.Println("Erro ao conectar para inserir lead:", err)
			return
		}
		defer stmt.Close()

		row := stmt.QueryRowContext(
			ctx)
		var newID int64
		err = row.Scan(&newID)
		if err != nil {
			fmt.Println("Erro ao inserir lead:", err)
			return
		}
		fmt.Printf("novo lead inserido: nome: %s id:%d", name, newID)
	}
}

func GetMessage(count int, client *azservicebus.Client) {
	queueName, ok := os.LookupEnv("QUEUENAME")
	if !ok {
		panic("QUEUENAME environment variable not found")
	}
	receiver, err := client.NewReceiverForQueue(queueName, nil)
	if err != nil {
		panic(err)
	}
	defer receiver.Close(context.TODO())

	for {
		messages, err := receiver.ReceiveMessages(context.TODO(), count, nil)
		if err != nil {
			fmt.Println("Erro ao receber mensagem:", err)
			continue
		}

		for _, message := range messages {
			body := message.Body
			fmt.Printf("%s\n", string(body))
			customer := strings.Split(string(body), ";")
			insertMktLead(customer[0], customer[1])
			err = receiver.CompleteMessage(context.TODO(), message, nil)
			if err != nil {
				panic(err)
			}
		}

	}
}

func main() {
	client := GetClient()

	fmt.Println("ms nop to mail marketing - geting messages")
	GetMessage(1, client)
}
