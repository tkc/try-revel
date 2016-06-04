package controllers

import "github.com/revel/revel"
// import "net/http"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) TryPage() revel.Result {
  return c.Render()
}

func (c App) JsonReturn() revel.Result {
	return c.Render()
}
