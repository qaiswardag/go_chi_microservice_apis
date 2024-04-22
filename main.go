package main

import (
	"context"
	"fmt"

	application "github.com/qaiswardag/go_chi_microservice_apis/app"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start the app: ", err)
	}
}
