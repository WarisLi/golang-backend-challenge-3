package tests

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/WarisLi/golang-backend-challenge-3/adapters"
	"github.com/WarisLi/golang-backend-challenge-3/core"
	"github.com/WarisLi/golang-backend-challenge-3/tests/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBeefsSuccess(t *testing.T) {
	app := fiber.New()

	mockBeefRepo := new(mocks.MockBeefRepository)
	beefService := core.NewBeefService(mockBeefRepo)
	beefHandler := adapters.NewHttpBeefHandler(beefService)
	app.Get("/beef/summary", beefHandler.GetBeefs)

	mockBeefRepo.On("GetData").Return([]byte("Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."), nil)
	t.Run("Success case", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/beef/summary", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, 10000)
		defer resp.Body.Close()
		byteValue, _ := io.ReadAll(resp.Body)

		experctResponse := []byte("{\"beef\":{\"bresaola\":1,\"enim\":1,\"fatback\":1,\"jowl\":1,\"meatloaf\":1,\"pastrami\":1,\"pork\":1,\"t-bone\":4}}")
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		assert.Equal(t, string(experctResponse), string(byteValue))
	})
	mockBeefRepo.AssertExpectations(t)
}

func TestGetBeefsFailed(t *testing.T) {
	app := fiber.New()

	mockBeefRepo := new(mocks.MockBeefRepository)
	beefService := core.NewBeefService(mockBeefRepo)
	beefHandler := adapters.NewHttpBeefHandler(beefService)
	app.Get("/beef/summary", beefHandler.GetBeefs)

	mockBeefRepo.On("GetData").Return([]byte{}, errors.New("error"))
	t.Run("Failed case", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/beef/summary", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, 10000)

		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	mockBeefRepo.AssertExpectations(t)
}
