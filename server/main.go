package main

import (
	"context"
	"fmt"
	"go_gRPC/api"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//HTTP 2.0 kullanıyor ve Tcp üstünden iletişim sağlanıyor.
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	//Grpc sunucusu
	srv := grpc.NewServer()
	//Hava servisinin sunucusunu register et.
	api.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("Starting server...")
	panic(srv.Serve(lis))

}

type myWeatherService struct {
	//Protodaki Weather Service'yi ekliyor.
	api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context, req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "tr_Ankara",
				CityName: "Ankara"},
			&api.CityEntry{CityCode: "tr_Istanbul",
				CityName: "İstanbul"},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(req *api.WeatherRequest, resp api.WeatherService_QueryWeatherServer) error {
	for {
		err := resp.Send(&api.WeatherResponse{Temperature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
		time.Sleep(time.Second)

	}
	return nil
}
