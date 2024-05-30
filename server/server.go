package server

import (
	"context"
	"log"
	"net"
	"strings"

	pb "myProject1/protoc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBeefCounterServer
}

func (s *server) CountBeef(ctx context.Context, req *pb.BeefRequest) (*pb.BeefResponse, error) {
	meatSet := make(map[string]struct{})
	beefCount := make(map[string]int32)

	meats := []string{
		"bacon", "jerky", "fatback", "boudin", "sausage", "ribs", "picanha",
		"andouille", "jowl", "beef", "biltong", "venison", "shank", "frankfurter",
		"pork", "salami", "chuck", "meatball", "swine", "rump", "doner", "alcatra",
		"tongue", "ball tip", "ham hock", "chislic", "turducken", "cow", "porchetta",
		"tenderloin", "kielbasa", "tail", "tri-tip", "ribeye", "sirloin", "biltong",
		"flank", "landjaeger", "meatloaf", "prosciutto", "short loin", "strip steak",
		"capicola", "short ribs", "pancetta", "shoulder", "corned beef", "buffalo",
		"picanha", "filet mignon", "bresaola", "hamburger", "turkey", "ground round",
		"chicken", "brisket", "drumstick", "boudin", "shank", "spare ribs", "leberkas",
	}

	for _, meat := range meats {
		meatSet[strings.ToLower(meat)] = struct{}{}
	}

	meatList := strings.FieldsFunc(req.Data, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.'
	})

	for _, meat := range meatList {
		meat = strings.ToLower(meat)
		if _, exists := meatSet[meat]; exists {
			beefCount[meat]++
		}
	}
	return &pb.BeefResponse{Beef: beefCount}, nil
}

func Servers() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	go func() {
		pb.RegisterBeefCounterServer(s, &server{})

		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}
