package example

import (
	"errors"
	"fmt"
	"go.uber.org/fx"
	"net/http"
	"os"
)

func test() {
	// A module that provides a HTTP server depends on
	// the $PORT environment variable. If the variable
	// is unset, the module returns an fx.Error option.
	newHTTPServer := func() fx.Option {
		port := os.Getenv("PORT")
		if port == "" {
			return fx.Error(errors.New("$PORT is not set"))
		}
		return fx.Provide(&http.Server{
			Addr: fmt.Sprintf(":%s", port),
		})
	}

	app := fx.New(
		newHTTPServer(),
		fx.Invoke(func(s *http.Server) error { return s.ListenAndServe() }),
	)

	fmt.Println(app.Err())
}
