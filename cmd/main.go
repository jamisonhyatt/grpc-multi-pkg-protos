package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jamisonhyatt/grpc-multi-pkg-protos/pkg/external/weather"
	"github.com/jamisonhyatt/grpc-multi-pkg-protos/pkg/weatherman"
	desktop "github.com/jamisonhyatt/grpc-multi-pkg-protos/pkg/weatherman/desktop_svc"
	mobile "github.com/jamisonhyatt/grpc-multi-pkg-protos/pkg/weatherman/mobile_svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("starting grpc server")
	host := fmt.Sprintf("localhost:%d", 8000)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	mobile.RegisterMobileServer(grpcServer, &mobileService{})
	desktop.RegisterDesktopServer(grpcServer, &desktopService{})
	weatherman.RegisterWeathermanServer(grpcServer, &weathermanService{})

	reflection.Register(grpcServer)
	fmt.Printf("listening on %s\n", host)
	grpcServer.Serve(lis)

	fmt.Println("exiting...")
}

type desktopService struct{}

func (d *desktopService) GetWeather(ctx context.Context, req *desktop.GetWeatherRequest) (*desktop.GetWeatherResponse, error) {
	return &desktop.GetWeatherResponse{
		Forecast: &weather.Forecast{
			Date: "2018-01-01",
			Rain: true,
		},
		ExtendedForecast: &weather.ExtendedForecast{
			Forecasts: []*weather.Forecast{
				{Date: "2018-01-02",
					Rain: true},
				{Date: "2018-01-03",
					Rain: false},
			},
		},
	}, nil
}

type mobileService struct{}

func (m *mobileService) GetWeather(ctx context.Context, req *mobile.GetWeatherRequest) (*mobile.GetWeatherResponse, error) {
	return &mobile.GetWeatherResponse{
		Forecast: &weather.Forecast{
			Date: "2018-01-01",
			Rain: true,
		},
	}, nil
}

type weathermanService struct{}

func (w *weathermanService) Healthcheck(ctx context.Context, req *weatherman.HealthCheckRequest) (*weatherman.HealthCheckResponse, error) {
	return &weatherman.HealthCheckResponse{
		Healthy: true,
	}, nil
}
