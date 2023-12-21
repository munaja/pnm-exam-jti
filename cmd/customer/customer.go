package main

import (
	"github.com/karincake/apem"

	"github.com/munaja/pnm-exam-jti/internal/handler/customer"
)

func main() {
	apem.Run("pnm-exam-jti/customer", customer.SetRoutes())
}
