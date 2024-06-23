package main 

import (
	"context"
    "fmt"
	"os"
	"os/signal"

	"4hoeschele/go_chi_project.git/application"
)

func main() {
	app := application.New()

	// should only be used by the main function, initializations and tests
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// with defer this function will be called when the main function ends
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}

}

