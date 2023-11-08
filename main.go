package main

import (
	"github.com/yaza-putu/online-test-dot/src/core"
)

func main() {
	// load env
	core.Env()

	// init database
	core.Database()

	// start server
	core.HttpServe()
}
