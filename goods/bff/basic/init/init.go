package init

import (
	"flag"
	"gospacex/goods/bff/basic/config"
	payment "gospacex/proto/payment"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	GrpcInit()
}
func GrpcInit() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//config.ProductClient = __.NewStreamGreeterClient(conn)
	config.PaymentClient = payment.NewStreamGreeterClient(conn)

}
