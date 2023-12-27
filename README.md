# golang

This is my go lang tutorial

First you need to install go (https://go.dev/)
I am using VS Code (https://code.visualstudio.com/docs/)
After installing VSCode install go extension (https://marketplace.visualstudio.com/items?itemName=golang.go)

The package fmt documentation for printing
https://pkg.go.dev/fmt

Sample Demos are placed on directories:
Code credits: 
- Basic-Lang session (https://www.tutorialspoint.com/go/index.htm)
- Go-Routines session (https://golangbot.com/channels/) (https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca)

Gargabe Collector (https://blog.twitch.tv/en/2016/07/05/gos-march-to-low-latency-gc-a6fa96f06eb7/)

In order to compile and execute samples go to sample directory and execute (go run main.go)

Fisrt Sample Hello world 

go-lang-tutorialv2\Basic-Lang\1-Hello-Word> go run main.go

If some package is needed into import that was not got yet you need to run the following command

go-lang-tutorialv2\Basic-Lang\8-Strings> go get strings

to import azure sql db lib run

go-lang-tutorialv2\MessagerListenerV2> go get -u "github.com/denisenkom/go-mssqldb"