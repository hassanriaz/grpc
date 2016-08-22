package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "github.com/hassanriaz/grpc/client1/add"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/gorilla/mux"
)

const (
	server = "localhost:50000"
)

func main() {
	port := flag.String("port", ":8000", "HTTP Port")
	flag.Parse()
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddServiceClient(conn)

	r := mux.NewRouter()
	r.HandleFunc("/add/{first}/{second}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		first, _ := strconv.ParseInt(vars["first"], 10, 32)
		second, _ := strconv.ParseInt(vars["second"], 10, 32)
		reply, err := c.Add(context.Background(), &pb.AddRequest{First: int32(first), Second: int32(second)})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		log.Printf("Sum: %v", reply.Sum)
		fmt.Fprintf(w, "Sum: %v", reply.Sum)
	})

	log.Fatal(http.ListenAndServe(*port, r))

}
