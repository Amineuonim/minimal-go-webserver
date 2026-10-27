package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"server/constants"

	"github.com/gin-gonic/gin"
)

type test struct {
	ID   int64  `db:"id"`
	Body string `db:"body"`
}

func TestDb(c *gin.Context) {
	var t test
	err := constants.Db.DB.GetContext(
		c.Request.Context(),
		&t,
		`SELECT id, body FROM test LIMIT 1`,
	)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   t.ID,
		"body": t.Body,
	})
}
