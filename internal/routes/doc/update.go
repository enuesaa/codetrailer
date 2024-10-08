package doc

import (
	"strings"

	"github.com/enuesaa/codetrailer/internal/router/ctx"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	cc := ctx.Use(c)

	var reqbody UpdateRequestBody
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	readme := strings.Join(reqbody.Slides, "\n\n---\n\n")
	if err := cc.Repos.Fs.Create(cc.Repos.Config.DocPath, strings.NewReader(readme)); err != nil {
		return err
	}

	res := Creation{
		Ok: true,
	}
	return cc.JSON(200, res)
}
