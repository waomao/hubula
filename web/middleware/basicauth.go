package middleware

import (
	_ "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
)


// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "pass",
	},
})
