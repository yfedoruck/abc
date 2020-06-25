package env

import (
	"log"
	"path/filepath"
	"runtime"
)

func BasePath() string {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic("Caller error")
	}
	env := filepath.Dir(b)
	pkg := filepath.Dir(env)
	app := filepath.Dir(pkg)
	return app
}