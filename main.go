package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	backendchallenge1 "myProject1/backend-challenge1"
	backendchallenge2 "myProject1/backend-challenge2"
	pb "myProject1/protoc"
	"os"
	"time"

	"myProject1/server"

	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc"
)

func client() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBeefCounterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	apiURL := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	client := resty.New()
	resp, err := client.R().Get(apiURL)
	if err != nil {
		return
	}
	text := resp.String()
	result, err := c.CountBeef(ctx, &pb.BeefRequest{Data: text})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Print(result)

}

func main() {
	//----------------------------1---------------------------------
	resultTriangle, err := backendchallenge1.PrintTriangle()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ผลลัพธ์:", resultTriangle)
	//----------------------------1---------------------------------

	//----------------------------2---------------------------------
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("input : ")
	scanner.Scan()
	encoded := scanner.Text()

	decoded := backendchallenge2.Decode(encoded)
	fmt.Println("Decoded sequence:", decoded)
	fmt.Println("Sum of decoded sequence:", backendchallenge2.Sum(decoded))
	//----------------------------2---------------------------------

	//----------------------------3---------------------------------
	server.Servers()
	client()
	//----------------------------3---------------------------------

}
