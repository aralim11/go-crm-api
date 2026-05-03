package router

import (
	"net/http"

	"github.com/aralim11/go-crm-api/internal/auth"
	"github.com/aralim11/go-crm-api/internal/report"
	"github.com/aralim11/go-crm-api/internal/user"

	"github.com/jmoiron/sqlx"
)

func RegisterModules(mux *http.ServeMux, db *sqlx.DB) {
	user.RegisterModule(mux, db)
	report.RegisterModule(mux, db)
	auth.RegisterModule(mux, db)
}
