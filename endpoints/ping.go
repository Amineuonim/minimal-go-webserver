package endpoints
import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(200, "")
}

