package pkg

import (
	"errors"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

var ErrEmptyBody = errors.New("empty body")

func ParseJSON(c fiber.Ctx, output interface{}) error {
	body := c.Body()

	if len(body) == 0 {
		return ErrEmptyBody
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return err
	}

	return nil
}
