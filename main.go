package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/james-maloney/templates"
)

func main() {
	e := gin.New()

	e.GET("/", Home)
	e.GET("/hello", Hello)
	e.GET("/hello/:name", Hello)

	if err := e.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}

func Home(ctx *gin.Context) {
	// templates is a package I created to more easily use
	// Go's template package
	templates.MustRenderOne(ctx.Writer, "home", nil)
}

// HelloResponse is used as the JSON response for the /hello GET route
type HelloResponse struct {
	Name     string `json:"name"`
	Gretting string `json:"greeting"`
	Usage    string `json:"usage,omitempty"`
	URL      string `json:"url"`
}

func Hello(ctx *gin.Context) {
	// check if we have any form values

	// ctx.Request represents the HTTP request
	// FormValue in this context (GET) looks for a url param
	// named 'name'
	name := ctx.Request.FormValue("name")
	if len(name) == 0 {
		// Since we are using Hello in two routes check for 'name' as a Param. ctx.Param checks the path of the url
		name = ctx.Param("name")
	}

	usage := ""
	if len(name) == 0 {
		name = "World"
		usage = "Make your request again but include your name, like '/hello?name=james'"
	} else {
		// Capitalize the name
		name = strings.Title(name)
	}

	res := HelloResponse{
		Name:     name,
		Gretting: fmt.Sprintf("Hello, %s!", name),
		Usage:    usage,
		URL:      ctx.Request.URL.String(),
	}

	// Marshal response to JSON and set appropriate content type headers
	ctx.JSON(200, res)
}

// init get called before main. All go packages can have an init or multiple init functions.
func init() {
	templates.AddView("home", `
<!DOCTYPE>
<html>
	<head>
		<title>GIN Demo</title>
	</head>
	<body>
		<h1>Gin API Demo</h1>
		<h3>Hello, World!</h3>
		<div><a href="/hello">/hello</a></div>
		<h3>Greeting Route using url param</h3>
		<div>
			<form action="/hello" method="GET">
				/hello?name=<input placeholder="Name" name="name" />
				<button type="submit">Submit</button>
			</form>
		</div>
		<h3>Greeting Route using a path param</h3>
		<form id="param-form" action="/hello" method="GET" onSubmit="return updateAction()">
			/hello/<input placeholder="Name" name="name" />
			<button type="submit">Submit</button>
		</form>

		<script>
			document.getElementById("param-form").onsubmit = function() {
				location.href="/hello/" + this.querySelector("input").value;
				return false;
			};
		</script>
	</body>
</html>
	`)
}
