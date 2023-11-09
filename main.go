package main

import (
	"github.com/yaza-putu/online-test-dot/src/core"
)

func main() {
	// load env
	core.Env()

	// init database
	core.Database()

	// redis
	core.Redis()
	// start server
	core.HttpServe()
}
