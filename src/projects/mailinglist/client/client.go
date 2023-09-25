package main

import (
	"context"
	"log"
	pb "mailinglist/proto"
	"time"

	"github.com/alexflint/go-arg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func logResponse(res *pb.EmailResponse, err error) {
	if err != nil {
		log.Fatalf(" error: %v", err)
	} else {
		log.Printf(" response:: %v\n", res)
	}
}

func createEmail(client pb.MailingListServiceClient, email string) *pb.EmailEntry {
	log.Println("create email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateEmail(ctx, &pb.CreateEmailRequest{EmailAddr: email})
	logResponse(res, err)

	return res.EmailEntry
}

func getEmail(client pb.MailingListServiceClient, email string) *pb.EmailEntry {
	log.Println("create email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetEmail(ctx, &pb.GetEmailRequest{EmailAddr: email})
	logResponse(res, err)
	return res.EmailEntry
}

func getEmailBatch(client pb.MailingListServiceClient, count, page int32) {
	log.Println("get email batch")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetEmailBatch(ctx, &pb.GetEmailBatchRequest{Count: count, Page: page})
	if err != nil {
		log.Fatalf(" error: %v", err)
	}

	log.Println(" response:")
	for i := 0; i < len(res.EmailEntries); i++ {
		log.Printf("  %v\n", res.EmailEntries[i])
	}
}

func updateEmail(client pb.MailingListServiceClient, entry pb.EmailEntry) *pb.EmailEntry {
	log.Println("update email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UpdateEmail(ctx, &pb.UpdateEmailRequest{EmailEntry: &entry})
	logResponse(res, err)
	return res.EmailEntry
}

func deleteEmail(client pb.MailingListServiceClient, email string) *pb.EmailEntry {
	log.Println("delete email")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DeleteEmail(ctx, &pb.DeleteEmailRequest{EmailAddr: email})
	logResponse(res, err)
	return res.EmailEntry
}

var args struct {
	GrpcAddress string `arg:"env:GRPC_ADDRESS"`
}

func main() {

	arg.MustParse(&args)
	if args.GrpcAddress == "" {
		args.GrpcAddress = ":8081"
	}

	conn, err := grpc.Dial(args.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewMailingListServiceClient(conn)

	// newEmail := createEmail(client, "12999@99.99")
	// confirmedAt := int64(10000)
	// newEmail.ConfirmedAt = &confirmedAt
	// updateEmail(client, *newEmail)
	// deleteEmail(client, *newEmail.Email)

	getEmailBatch(client, 5, 1)
}
