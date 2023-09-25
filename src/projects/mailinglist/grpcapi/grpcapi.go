package grpcapi

import (
	"context"
	"database/sql"
	"log"
	"net"
	"time"

	"mailinglist/mdb"
	pb "mailinglist/proto"

	"google.golang.org/grpc"
)

type MailServer struct {
	pb.UnimplementedMailingListServiceServer
	db *sql.DB
}

func emailResponse(db *sql.DB, email string) (*pb.EmailResponse, error) {
	entry, err := mdb.GetEmail(db, email)
	if err != nil {
		return nil, err
	}

	if entry == nil {
		return &pb.EmailResponse{}, nil
	}

	res := mdbEntryToPbEntry(entry)
	return &pb.EmailResponse{EmailEntry: &res}, nil
}

func (s MailServer) CreateEmail(ctx context.Context, req *pb.CreateEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gpc CreateEmail: %v\n", req)

	err := mdb.CreateEmail(s.db, req.EmailAddr)
	if err != nil {
		return &pb.EmailResponse{}, err
	}

	return emailResponse(s.db, req.EmailAddr)
}

func (s MailServer) GetEmail(ctx context.Context, req *pb.GetEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("GetEmail: %v\n", req)
	return emailResponse(s.db, *&req.EmailAddr)
}

func (s MailServer) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gpc UpdateEmail: %v\n", req)

	entry := pbEntryToMdbEntry(req.EmailEntry)
	err := mdb.UpdateEmail(s.db, entry)
	if err != nil {
		return &pb.EmailResponse{}, err
	}

	return emailResponse(s.db, entry.Email)
}

func (s MailServer) DeleteEmail(ctx context.Context, req *pb.DeleteEmailRequest) (*pb.EmailResponse, error) {
	log.Printf("gpc DeleteEmail: %v\n", req)

	err := mdb.DeleteEmail(s.db, req.EmailAddr)
	if err != nil {
		return &pb.EmailResponse{}, err
	}

	return emailResponse(s.db, req.EmailAddr)
}

func (s MailServer) GetEmailBatch(ctx context.Context, req *pb.GetEmailBatchRequest) (*pb.GetEmailBatchResponse, error) {
	log.Printf("GetEmailBatch: %v\n", req)

	params := mdb.GetEmailBatchQueryParams{
		Page:  int(req.Page),
		Count: int(req.Count),
	}

	entries, err := mdb.GetEmailBatch(s.db, params)
	if err != nil {
		return &pb.GetEmailBatchResponse{}, err
	}

	pbEntries := make([]*pb.EmailEntry, len(entries))

	for i := 0; i < len(entries); i++ {

		entry := mdbEntryToPbEntry(&entries[i])

		log.Printf("  %v\n", &entry.Email)
		pbEntries = append(pbEntries, &entry)
	}

	return &pb.GetEmailBatchResponse{EmailEntries: pbEntries}, nil
}

func pbEntryToMdbEntry(pbEntry *pb.EmailEntry) mdb.EmailEntry {

	t := time.Unix(*pbEntry.ConfirmedAt, 0)

	return mdb.EmailEntry{
		Id:          *pbEntry.Id,
		Email:       *pbEntry.Email,
		ConfirmedAt: &t,
		OptOut:      *pbEntry.OptOut,
	}
}

func mdbEntryToPbEntry(mdbEntry *mdb.EmailEntry) pb.EmailEntry {

	var confirmedAt int64
	if mdbEntry.ConfirmedAt != nil {
		confirmedAt = mdbEntry.ConfirmedAt.Unix()
	}

	return pb.EmailEntry{
		Id:          &mdbEntry.Id,
		Email:       &mdbEntry.Email,
		ConfirmedAt: &confirmedAt,
		OptOut:      &mdbEntry.OptOut,
	}
}

func Serve(db *sql.DB, bind string) {
	log.Printf("Starting grpc server on %s\n", bind)

	listener, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	gprcServer := grpc.NewServer()
	mailServer := MailServer{db: db}

	pb.RegisterMailingListServiceServer(gprcServer, mailServer)
	log.Printf("Serving grpc on %s\n", bind)

	if err := gprcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
