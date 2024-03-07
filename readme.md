# go hello world

requirements

```bash
# Ubuntu 22.04
sudo apt-get install libxxf86vm-dev
```

```bash
go mod init github.com/xl714/go-webview-hello-world

go get github.com/webview/webview_go

```

```go
package main

import webview "github.com/webview/webview_go"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetHtml("Thanks for using webview!")
	w.Run()
}
```

Execute the following command to run the code:

```bash
go run main.go
```

Formatting: For better readability, use the go fmt command to automatically format your code:

```bash
go fmt main.go
```

Building a Binary: To create a standalone executable file, use the go build command:

```bash
go build -o hello main.go
```

This will create an executable file named hello in your project directory. You can run it directly without the go run command.

Using Modules: For larger projects, consider using Go modules for dependency management.
