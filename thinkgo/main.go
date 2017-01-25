package main

import (
	"github.com/henrylee2cn/thinkgo"
	"time"
)

type Index struct {
	Id        int      `param:"<in:path> <required> <desc:ID> <range: 0:10>"`
	Title     string   `param:"<in:query> <nonzero>"`
	Paragraph []string `param:"<in:query> <name:p> <len: 1:10> <regexp: ^[\\w]*$>"`
	Cookie    string   `param:"<in:cookie> <name:thinkgoID>"`
	// Picture         *multipart.FileHeader `param:"<in:formData> <name:pic> <maxmb:30>"`
}

func (i *Index) Serve(ctx *thinkgo.Context) error {
	if ctx.CookieParam("thinkgoID") == "" {
		ctx.SetCookie("thinkgoID", time.Now().String())
	}
	return ctx.JSON(200, i)
}

func main() {
	app := thinkgo.New("myapp", "0.1")

	// Register the route in a chain style
	app.GET("/index/:id", new(Index))

	// Register the route in a tree style
	// app.Route(
	//   app.NewGET("/index/:id", new(Index)),
	// )

	// Start the service
	app.Run()
}

