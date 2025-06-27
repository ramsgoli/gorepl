# gorepl

## usage
```go
func main() {
	command := make(chan string)
	cont := make(chan bool)
	go repl.StartRepl(command, cont, "> ", ".exit")

	for t := range command {
        fmt.Printf("got command: %s\n", t)
        cont <- true
	}
	fmt.Println("done!")
}

```
* `.exit` to exit the REPL

