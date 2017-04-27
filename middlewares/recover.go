package middlewares

import (
	"net/http"
	"runtime"

	"github.com/ThanFX/G3/handlers"
	"github.com/ThanFX/G3/log"
)

const stackSize = 1024 * 4

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			l, _ := log.GetLogger("text", "debug")
			if err := recover(); err != nil {
				stack := make([]byte, stackSize)
				stack = stack[:runtime.Stack(stack, false)]
				l.Errorf("middleware: PANIC: %s\n%s", err, stack)
				handlers.SendJsonResponse(w, r, http.StatusInternalServerError, "", 0, "Error!")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
