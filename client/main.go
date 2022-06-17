package main

import (
	"context"
	"fmt"
	"go_gRPC/api"
	"io"

	"google.golang.org/grpc"
)

func main() {

	addr := "localhost:8080"
	//Normalde https üstünden kullanılır.
	//grpc.WithInsecure() ile bunu şu anlık göz artı ettik.
	//Şifresiz bağlantı oluşturduk.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//Bağlantı kullanarak havadurumu servis clienti verir.
	client := api.NewWeatherServiceClient(conn)
	//
	ctx := context.Background()

	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Cities:")
	for _, city := range resp.Items {
		fmt.Printf("\n%s: %s", city.GetCityCode(), city.CityName)
	}

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{
		CityCode: "tr_ank",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("Weather in Ankara:")
	fmt.Println()
	for {
		//Hava durumu ekrana yazdırıyoruz.
		//Streamdan recive edilir.
		//Cevap varsa yolluyor yoksa hata alıyor.
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("\t temperature: %.2f\n", msg.GetTemperature())
	}
	fmt.Println("Server stopped sending")

}
