package commands

import (
	render "github.com/SvenDowideit/gendoc/render"

	"github.com/codegangsta/cli"
)

var Render = cli.Command{
	Name:  "render",
	Usage: "render files from the ./docs dir",
	Action: func(context *cli.Context) error {
		render.RenderDocsDir()
		return nil
	},
}
