package area

import (
	"go-mssql/internal/domain"
	"go-mssql/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListHandler(areaService services.AreaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		areasFound, err := areaService.ListAreas()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, &domain.Response{Message: "error : " + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, areasFound)
	}
}
