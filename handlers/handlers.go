package handlers

import (
	"net/http"

	dashboard "github.com/fdaygon/drift/frontend"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	dashboardTemplate := dashboard.Dashboard()
	dashboardTemplate.Render(r.Context(), w)

}
