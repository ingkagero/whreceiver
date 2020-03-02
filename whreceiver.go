package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/"
)

func main() {
	hook, _ := github.New(github.Options.Secret("SomeSecret"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("%+v", r)
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				//fmt.Printf("%+v", err)
				//fmt.Printf("%+v", payload)
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.CreatePayload:
			createRef := payload.(github.CreatePayload)
			// Do whatever you want from here...
			fmt.Println("create payload")
			fmt.Printf("%+v", createRef)

		case github.DeletePayload:
			deleteRef := payload.(github.DeletePayload)
			// Do whatever you want from here...
			fmt.Println("delete payload")
			fmt.Printf("%+v", deleteRef)

		case github.PushPayload:
			push := payload.(github.PushPayload)
			// Do whatever you want from here...
			fmt.Println("push payload")
			fmt.Printf("%+v", push)
		}
	})
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
