package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

// Status writer is crucial to get the status code of the request.
func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

// Logger logs the incoming HTTP request and its duration.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("err: %v", err)
				log.Printf("trace: %v", debug.Stack())
			}
		}()

		sw := &statusWriter{ResponseWriter: w}

		start := time.Now()
		next.ServeHTTP(sw, r)
		log.Printf("status=%v method=%v path=%v duration=%v", sw.status, r.Method, r.URL.EscapedPath(), time.Since(start))
	})
}