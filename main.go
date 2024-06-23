package main 

import (
	"context"
    "fmt"

	"4hoeschele/go_chi_project.git/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}

