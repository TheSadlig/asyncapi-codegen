// Universal parts generation
//go:generate go run ../../../cmd/asyncapi-codegen -g client -p generated -i ../asyncapi.yaml -o ./generated/client.gen.go
//go:generate go run ../../../cmd/asyncapi-codegen -g types -p generated -i ../asyncapi.yaml -o ./generated/types.gen.go

package main

import (
	"context"
	"time"

	"github.com/lerenn/asyncapi-codegen/examples/ping/client/generated"
	"github.com/lerenn/asyncapi-codegen/pkg/broker/controllers"
	"github.com/lerenn/asyncapi-codegen/pkg/log"
	"github.com/lerenn/asyncapi-codegen/pkg/middleware"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		panic(err)
	}

	// Create a new client controller
	ctrl, err := generated.NewClientController(controllers.NewNATS(nc))
	if err != nil {
		panic(err)
	}
	defer ctrl.Close(context.Background())

	// Attach a logger (optional)
	logger := log.NewECS()
	ctrl.SetLogger(logger)
	ctrl.AddMiddlewares(middleware.Logging(logger))

	// Make a new ping message
	req := generated.NewPingMessage()
	req.Payload = "ping"

	// Create the publication function to send the message
	// Note: it will indefinitely wait to publish as context has no timeout
	publicationFunc := func(ctx context.Context) error {
		return ctrl.PublishPing(ctx, req)
	}

	// The following function will subscribe to the 'pong' channel, execute the publication
	// function and wait for a response. The response will be detected through its
	// correlation ID.
	//
	// This function is available only if the 'correlationId' field has been filled
	// for any channel in the AsyncAPI specification. You will then be able to use it
	// with the form WaitForXXX where XXX is the channel name.
	//
	// Note: it will indefinitely wait for messages as context has no timeout
	_, err = ctrl.WaitForPong(context.Background(), req, publicationFunc)
	if err != nil {
		panic(err)
	}

	// Wait for the message to be received
	time.Sleep(time.Second)
}
