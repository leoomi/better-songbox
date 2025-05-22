package pages

import (
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type Home struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *Home) Render() app.UI {
	return app.Main().Class().Body(app.H1().Text("Hello World!"), app.Button().Text("Click me").OnClick(h.CallApiTest))
}

func (h *Home) CallApiTest(ctx app.Context, e app.Event) {
	resp, err := http.Get("/api/search")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	fmt.Println("response:", string(body))
}
