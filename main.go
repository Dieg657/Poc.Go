package main

import (
	"context"

	"github.com/Dieg657/Poc.Go/cmd/web"
)

//	@Title			Golang WebAPI POC
//	@Version		1.0
//	@Description	WebAPI in Golang for educational purposes

//	@contact.name	Diego dos Santos Soares
//	@contact.url	http://www.diegosoares.dev.br
//	@contact.email	dieg657@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @Host		http://localhost:8080
// @BasePath	/api/v1
func main() {
	webApi := web.WebApi{}
	context := context.Background()
	webApi.Start(context)
}
