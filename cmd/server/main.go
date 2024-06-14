package main

import (
	"GO-ECOMMERCE-BACKEND-API/internal/routers"
	"fmt"
)

func main() {
	fmt.Println("starting")
	r := routers.NewRouter()
	r.Run(":3000") // listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
}
