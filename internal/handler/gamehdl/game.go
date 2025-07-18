package gamehdl

import (
	"github.com/gofiber/fiber/v3"
)

type handler struct {
}

type Handler interface {
	Test(ctx fiber.Ctx) error
}

func New() Handler {
	return &handler{}
}

type Card struct {
	Number int    // 1=A, 2–9, 10, 11=J, 12=Q, 13=K
	Suit   string // hearts, diamonds, clubs, spades (unused for value)
}
type Request struct {
	PlayHands  [][]Card `json:"playHands"`
	KnownHands [][]Card `json:"knownHands"` // unused in this example
}

func (h *handler) Test(ctx fiber.Ctx) error {
	var req Request
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	//  Example handler logic
	// This could be replaced with actual game logic
	mpaStand := map[int]string{}
	switch {
	case len(req.PlayHands) != 0:
		for i, card := range req.PlayHands {
			countNumber := 0
			for _, v := range card {
				countNumber += v.Number
			}

			if countNumber < 5 {
				mpaStand[i] = "hit"
			} else {
				mpaStand[i] = "stand"
			}
		}
		response := make([]string, len(mpaStand))
		for k, v := range mpaStand {
			response[k] = v
		}
	case len(req.PlayHands) != 0:

	}

	response := make([]string, len(mpaStand))
	for k, v := range mpaStand {
		response[k] = v
	}

	//
	return ctx.JSON(fiber.Map{
		"response": response,
	})
}
