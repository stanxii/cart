package main

import (
	"log"
	"net/http"

	"github.com/marcusolsson/coverage-cravings/sql"
)

func main() {
	sqlOrders, err := sql.NewOrderRepository(
		sql.SetUser("pqtest"),
		sql.SetDBName("orders"),
		sql.SetSSLMode("verify-full"),
	)
	if err != nil {
		log.Fatal(err)
	}

	srv := server{
		orders: sqlOrders,
	}

	mux := http.NewServeMux()
	srv.register(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
