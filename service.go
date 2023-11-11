package main

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetCatFact(c *fiber.Ctx) error
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (service *CatFactService) GetCatFact(c *fiber.Ctx) error {
	resp, err := http.Get(service.url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	fact := &CatFact{}

	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return err
	}

	return c.SendString(string(fact.Fact))

}
