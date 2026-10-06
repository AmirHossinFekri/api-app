package middleware

import (
	"log"
	"net/http"
	"time"
)


type statusRecorder struct {
	http.ResponseWriter
	status int
	size	 int
}


func (r *statusRecorder) WriteHeader(code int){
	r.status=code
	r.ResponseWriter.WriteHeader(code) 
}

func (r *statusRecorder) Write(b []byte) (int ,error){
	n , err := r.ResponseWriter.Write(b)
	r.size +=n

	return n,err
}


const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	gray    = "\033[90m"
)


func statusColor(code int) string {
	switch {
	case code >= 500:
		return red
	case code >= 400:
		return yellow
	case code >= 300:
		return cyan
	default:
		return green
	}
}

func methodColor(m string) string {
	switch m {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return green
	case http.MethodPut, http.MethodPatch:
		return yellow
	case http.MethodDelete:
		return red
	default:
		return magenta
	}
}


func Logging(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rec,r)

		log.Printf("%s%-7s%s %s %s%d%s %s%v%s",
					methodColor(r.Method), r.Method, reset,
					r.URL.Path,
					statusColor(rec.status), rec.status, reset,
					gray, time.Since(start).Round(time.Microsecond), reset,
		)
	})
}