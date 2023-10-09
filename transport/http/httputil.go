package httputil

import (
	"context"
	"net/http"
)

// QueryHandler is a handler for usecases with a return value.
func QueryHandler[Req any, Res any](
	requestReader func(w http.ResponseWriter, r *http.Request) (Req, error),
	handler func(ctx context.Context, req Req) (Res, error),
	reponseWriter func(w http.ResponseWriter, r *http.Request, res Res),
	errorHandler func(w http.ResponseWriter, r *http.Request, err error),
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req, err := requestReader(w, r)
		if err != nil {
			errorHandler(w, r, err)
			return
		}

		res, err := handler(r.Context(), req)
		if err != nil {
			errorHandler(w, r, err)
			return
		}

		reponseWriter(w, r, res)
	}
}

// CommandHandler is a handler for usecases without a return value.
func CommandHandler[Req any](
	requestReader func(w http.ResponseWriter, r *http.Request) (Req, error),
	handler func(ctx context.Context, req Req) error,
	reponseWriter func(w http.ResponseWriter, r *http.Request),
	errorHandler func(w http.ResponseWriter, r *http.Request, err error),
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req, err := requestReader(w, r)
		if err != nil {
			errorHandler(w, r, err)
			return
		}

		if err := handler(r.Context(), req); err != nil {
			errorHandler(w, r, err)
			return
		}

		reponseWriter(w, r)
	}
}

func Handler[Res any](
	handler func(ctx context.Context) (Res, error),
	reponseWriter func(w http.ResponseWriter, r *http.Request, res Res),
	errorHandler func(w http.ResponseWriter, r *http.Request, err error),
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		res, err := handler(r.Context())
		if err != nil {
			errorHandler(w, r, err)
			return
		}

		reponseWriter(w, r, res)
	}
}
