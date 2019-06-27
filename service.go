package dexter

import (
	"context"
	"log"
	"net"
	"github.com/jinzhu/gorm"
	"github.com/davecgh/go-spew/spew"
	grpc "google.golang.org/grpc"
	pb "github.com/whiteblock/dexter/api/alerts"
)

type dexterAlertsServer struct {
	db *gorm.DB
}

func (s *dexterAlertsServer) CreateAlert(ctx context.Context, alert *pb.Alert) (*pb.Alert, error) {
	newAlert := &pb.Alert{}
	return newAlert, nil
}

func (s *dexterAlertsServer) ListAlerts(ctx context.Context, opts *pb.ListAlertsRequest) (*pb.ListAlertsResponse, error) {
	response := &pb.ListAlertsResponse{}
	return response, nil
}

func (s *dexterAlertsServer) GetAlert(ctx context.Context, opts *pb.GetAlertRequest) (*pb.Alert, error) {
	alert := &pb.Alert{}
	return alert, nil
}

func (s *dexterAlertsServer) UpdateAlert(ctx context.Context, alert *pb.Alert) (*pb.Alert, error) {
	updatedAlert := &pb.Alert{}
	return updatedAlert, nil
}

func (s *dexterAlertsServer) DeleteAlert(ctx context.Context, opts *pb.DeleteAlertRequest) (*pb.DeleteAlertResponse, error) {
	response := &pb.DeleteAlertResponse{}
	return response, nil
}

func (s *dexterAlertsServer) ListIndicators(ctx context.Context, opts *pb.ListIndicatorsRequest) (*pb.ListIndicatorsResponse, error) {
	response := &pb.ListIndicatorsResponse{}
	log.Printf("ListIndicators")
	var indicators []IndicatorSpec
	s.db.Find(&indicators)
	spew.Dump(&indicators)
	return response, nil
}

// StartServer starts the gRPC service for alert management
func StartServer(listen string, db *gorm.DB) {
	var opts []grpc.ServerOption
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Listening on %s\n", listen)
	}
	server := &dexterAlertsServer{
		db: db,
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAlertsServer(grpcServer, server)
	grpcServer.Serve(listener)
}