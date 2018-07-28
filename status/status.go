package status

import (
	"fmt"
	"net/http"
)

// GetStatus godoc
// @Summary Show status
// @Description Show current system state
// @Tags status
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /api/v1/status [get]
func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
