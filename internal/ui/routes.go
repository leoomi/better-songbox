package ui

import (
	"github.com/leoomi/better-songbox/internal/ui/pages"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func SetupRoutes() {
	app.Route("/", func() app.Composer { return &pages.Home{} })
}
