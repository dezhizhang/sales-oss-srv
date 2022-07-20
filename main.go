package main

import (
	"fmt"
	"os"
	"sales-oss/router"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func main() {

	// 初始化路由
	r := router.Router()

	r.Run()

}
