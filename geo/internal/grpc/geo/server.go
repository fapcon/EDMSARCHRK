package geo

import (
	"context"
	"fmt"
	"geo/internal/grpc/grpcclients"
	"geo/internal/service"
	geopr "github.com/fapcon/MSHUGOprotos/protos/geo/gen"
)

type Geo interface {
	Search(input string) ([]byte, error)
	Geocode(lat, lng string) ([]byte, error)
}

type ServerGeo struct {
	geopr.UnimplementedGeoServiceServer
	gcl *grpcclients.ClientAuth
	geo service.GeoService
}

func (s *ServerGeo) Search(context context.Context, req *geopr.SearchRequest) (*geopr.SearchResponse, error) {

	address, err := s.geo.Search(req.Input)
	if err != nil {
		return nil, fmt.Errorf("err service:%v", err)
	}

	return &geopr.SearchResponse{Data: address}, nil
}

func (s *ServerGeo) Geocode(context context.Context, req *geopr.GeocodeRequest) (*geopr.GeocodeResponse, error) {
	address, err := s.geo.Geocode(req.Lat, req.Lon)
	if err != nil {
		return nil, fmt.Errorf("err service:%v", err)
	}
	return &geopr.GeocodeResponse{Data: address}, nil
}
