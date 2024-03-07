// https://github.com/webview/webview_go/blob/master/examples/bind/main.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	webview "github.com/webview/webview_go"
)

func main() {

    w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(800, 800, webview.HintNone)
	

    // Fonction Go pour concaténer "Hello " avec le nom
    sayHello := func(data string) string {
        name := struct {
            Name string `json:"name"`
        }{}
        err := json.Unmarshal([]byte(data), &name)
        if err != nil {
            log.Println("Erreur lors du décodage JSON :", err)
            return ""
        }

        greeting := "Hello " + name.Name
        response, _ := json.Marshal(struct {
            Greeting string `json:"greeting"`
        }{greeting})

        return string(response)
    }

    // Lier la fonction Go à JavaScript
    w.Bind("sayHello", sayHello)

    // Charger le contenu HTML à partir du fichier
    html, err := ioutil.ReadFile("index.html")
    if err != nil {
        log.Fatal(err)
    }
    w.SetHtml(string(html))

    w.Run()
}