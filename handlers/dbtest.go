package handlers
import (
	"github.com/gin-gonic/gin"
	"server/constants"
)
/*
simple test query to db, initdb at core already 
pings the db so this is optional but if you want to test this, then:
paste into pgsql:
create table test(
	id SERIAL PRIMARY KEY,
	body varchar(50)
)

*/


type test struct {
	Id    uint   `gorm:"primaryKey"`
	Body  string
}

// Map struct to existing table "test"
func (test) TableName() string {
	return "test"
}


func TestDb(c *gin.Context) {
	var t test
	// Query first record
	if err := constants.Db.DB.First(&t).Error; err != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}

	// Return Id and Body as JSON
	c.JSON(200, gin.H{
		"id":   t.Id,
		"body": t.Body,
	})
}
