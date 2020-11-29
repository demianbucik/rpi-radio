package utils

import (
	"bytes"
	"context"
	"net/http"
	"runtime/debug"

	"github.com/apex/log"
	"github.com/go-chi/chi/middleware"
)

func RequestCtx(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RequestCtxKey, &RequestContext{})
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func CORS(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				if requestCtx, _ := r.Context().Value(RequestCtxKey).(*RequestContext); requestCtx != nil {
					requestCtx.StackTrace = getStackTrace(0)
					requestCtx.Panic = rvr
				}
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			status := ww.Status()
			entry := log.WithFields(log.Fields{
				"url":    r.URL,
				"method": r.Method,
				"status": status,
			})
			if requestCtx, _ := r.Context().Value(RequestCtxKey).(*RequestContext); requestCtx != nil {
				if requestCtx.Error != nil {
					entry = entry.WithField("error", requestCtx.Error)
				}
				if requestCtx.StackTrace != "" {
					entry = entry.WithField("stackTrace", requestCtx.StackTrace)
				}
				if requestCtx.Panic != nil {
					entry.WithField("panic", requestCtx.Panic).Error("HTTP request panicked")
					return
				}
			}
			if status < 500 {
				entry.Debug("HTTP request")
			} else {
				entry.Warn("HTTP request")
			}
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}

func getStackTrace(skipCalls int) string {
	parts := (2 * skipCalls) + 6
	stackTrace := debug.Stack()
	split := bytes.SplitAfterN(stackTrace, []byte("\n"), parts)
	if len(split) < parts {
		return string(stackTrace)
	}
	return string(split[parts-1])
}
