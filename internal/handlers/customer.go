package customerhdl

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	serviceport "github.com/nextchanupol/go-fiber-server/internal/core/services/ports"
)

type customerHandler struct {
	custService serviceport.CustomerService
}

func NewCustomerHandler(custService serviceport.CustomerService) customerHandler {
	return customerHandler{custService: custService}
}

func (h customerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.custService.GetCustomers()
	if err != nil {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			// Override status code if fiber.Error type
			code = e.Code
		}
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		return c.Status(code).SendString(err.Error())
	}
	return c.JSON(&fiber.Map{
		"customers": customers,
	})
}

func (h customerHandler) GetCustomerByID(c *fiber.Ctx) error {

	code := fiber.StatusInternalServerError

	customerID := c.Params("customerID")

	id, _ := strconv.ParseInt(customerID, 10, 64)
	customer, err := h.custService.GetCustomerByID(id)

	if err != nil {

		if e, ok := err.(*fiber.Error); ok {
			// Override status code if fiber.Error type
			code = e.Code
		}
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		return c.Status(code).SendString(err.Error())
	}
	return c.JSON(&fiber.Map{
		"customer": customer,
	})

	// customerID, err := strconv.Atoi(mux.Vars(r)["customerID"])
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Println(w, err)
	// 	return
	// }
	// customer, err := h.custService.GetCustomerByID(int64(customerID))
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Println(w, err)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(customer)
}
