package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ReturnErrorWithData returns an error response with additional data.
func ReturnErrorWithData(c *gin.Context, data responseData, result interface{}) {
	data.Timestamp = time.Now().Unix()
	data.Data = result
	c.JSON(http.StatusOK, data)
	// Abort the request
	c.Abort()
}

// ReturnData returns a normal response.
func ReturnData(c *gin.Context, result interface{}) {
	data := Success
	data.Timestamp = time.Now().Unix()
	data.Data = result
	c.JSON(http.StatusOK, data)
	// Abort the request
	c.Abort()
}

// ReturnDataWithCount returns a normal response with a count.
func ReturnDataWithCount(c *gin.Context, count int, result interface{}) {
	data := Success
	data.Timestamp = time.Now().Unix()
	data.Data = result
	data.Count = &count
	c.JSON(http.StatusOK, data)
	// Abort the request
	c.Abort()
}

// ReturnError returns an error response with a description.
func ReturnError(c *gin.Context, data responseData, description string) {
	data.Timestamp = time.Now().Unix()
	data.Message = func() string {
		if description == "" {
			return data.Message
		}
		return description
	}()
	c.JSON(http.StatusOK, data)
	// Abort the request
	c.Abort()
}

// ReturnSuccess returns a success response.
func ReturnSuccess(c *gin.Context) {
	data := Success
	data.Timestamp = time.Now().Unix()
	c.JSON(http.StatusOK, data)
	// Abort the request
	c.Abort()
}
