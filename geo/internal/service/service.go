package service

import (
	"encoding/json"
	"fmt"
	"geo/internal/models"
	"io"
	"log"
	"net/http"
	"strings"
)

type GeoService struct {
}

func NewGeoService() *GeoService {
	return &GeoService{}
}

func (g *GeoService) Search(input string) ([]byte, error) {
	var data = strings.NewReader(fmt.Sprintf("[ \"%s\" ]", input))

	req, err := http.NewRequest("POST", "https://cleaner.dadata.ru/api/v1/clean/address", data)
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+"cd6b57c7195fe08a623d3ab0bd8b398256163779")
	req.Header.Set("X-Secret", "51aae3bef5d26db8afb51fcc0c36c84c9aa79209")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	defer resp.Body.Close()
	log.Println("dadata status code -", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}

	var adrs models.AddressSearch
	err = json.Unmarshal(body, &adrs)
	if err != nil {
		log.Println("err unmarshal Q")
	}

	var pr models.ResponseAddress

	for _, s := range adrs {

		prs := []models.AddressSearchEl{{
			Result: s.Result,
			GeoLat: s.GeoLat,
			GeoLon: s.GeoLon,
		},
		}

		pr.Addresses = prs
	}

	addresses, err := json.Marshal(pr)
	if err != nil {
		log.Println("err marshal :", err)
	}

	return addresses, nil
}

func (g *GeoService) Geocode(lat, lng string) ([]byte, error) {
	var data = strings.NewReader(fmt.Sprintf("{ \"lat\":%s, \"lon\":%s }", lat, lng))
	req, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token cd6b57c7195fe08a623d3ab0bd8b398256163779")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("dadata status code", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err read response body", err)
	}
	var adrs models.GeocodeResponse
	err = json.Unmarshal(body, &adrs)
	if err != nil {
		log.Println("err unmarshal ", err)
	}

	var pr models.ResponseGeocode

	for _, v := range adrs.Suggestions {
		as := []models.AddressSearchEl{
			{
				Result: v.Value,
				GeoLat: v.Data.GeoLat,
				GeoLon: v.Data.GeoLon,
			},
		}
		ps := models.ResponseGeocode{
			Value:             v.Value,
			UnrestrictedValue: v.UnrestrictedValue,
			Data:              as,
		}

		pr = ps
	}

	address, err := json.Marshal(pr)
	if err != nil {
		log.Println("err marshal :", err)
	}

	return address, nil
}
