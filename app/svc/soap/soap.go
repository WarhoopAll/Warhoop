package soap

import (
	"github.com/gofiber/fiber/v2"
	"warhoop/app/log"
)

type SoapService struct {
	logger *log.Logger
	client *fiber.Client
}

func New(logger *log.Logger) *SoapService {
	return &SoapService{
		logger: logger,
	}
}
