package api

import (
	"net/http"
	"time"

	"github.com/delivery-much/dm-go/logger"
	"github.com/go-chi/chi/middleware"
)

type req struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	ID     string `json:"id"`
	IP     string `json:"ip"`
}

type res struct {
	StatusCode int `json:"status_code"`
	Length     int `json:"length"`
}

// HTTPLogInfo is a HTTP request log
type HTTPLogInfo struct {
	responseTime int64
	req
	res
}

type responseObserver struct {
	http.ResponseWriter
	code        int
	length      int
	wroteHeader bool
}

// Logger is logger middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			o = &responseObserver{ResponseWriter: w}

			start = time.Now()
			reqID = middleware.GetReqID(r.Context())
			reqIP = r.RemoteAddr
		)

		defer func() {
			httpInfo := HTTPLogInfo{
				responseTime: time.Since(start).Milliseconds(),
				req: req{
					Method: r.Method,
					URL:    r.URL.Path,
					ID:     reqID,
					IP:     reqIP,
				},
				res: res{
					StatusCode: o.code,
					Length:     o.length,
				},
			}

			logger.Debugw("request completed",
				"responseTime", httpInfo.responseTime,
				"req", httpInfo.req,
				"res", httpInfo.res,
			)
		}()

		next.ServeHTTP(o, r)
	})
}

func (o *responseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.length += n
	return
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.code = code
}
