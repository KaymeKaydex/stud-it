package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetAgeInfo(gCtx *gin.Context) {
	val := gCtx.Query("age")
	if val == "" {
		log.Println("empty age val")

		gCtx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(val)
	if err != nil {
		log.Println("error convert value")
		fmt.Println(err)

		gCtx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	gCtx.String(http.StatusOK, c.srv.GetYoungOrOld(age))
}
