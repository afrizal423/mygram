package middlewares

import (
	"fmt"
	"net/http"

	"github.com/afrizal423/mygram/app/models"
	"github.com/afrizal423/mygram/configs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func PhotoAuthorizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := configs.GormPostgresConn()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["user_id"].(float64))
		photo := models.Photo{}

		result := db.Where("user_id = ?", userID).Order("id desc").Take(&photo)
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintln("There is no data photo"),
			})
			return
		}

		// for protect unregistered user to login
		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func SingleDataPhotoAuthorizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := configs.GormPostgresConn()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["user_id"].(float64))
		photo := models.Photo{}
		photoId := c.Param("photoId")

		result := db.Where("user_id = ? AND id = ?", userID, photoId).Order("id desc").Take(&photo)
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": fmt.Sprintln("There is no data photo"),
			})
			return
		}

		// for protect unregistered user to login
		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/afrizal423/mygram/app/models"
// 	"github.com/afrizal423/mygram/configs"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func CreateProductAuthorizations() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		roles := uint(userData["roles"].(float64))
// 		fmt.Println(roles)
// 		// jika bukan admin atau user maka tidak bisa create product
// 		if roles != 1 && roles != 2 {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this data",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func DetailDataProductAuthorizations() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := configs.GormPostgresConn()
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		roles := uint(userData["roles"].(float64))
// 		userID := uint(userData["user_id"].(float64))
// 		productId := c.Param("productId")
// 		product := models.Product{}

// 		err := db.Select("user_id").Where("id = ?", productId).Order("id desc").First(&product).Error
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Data Not Found",
// 				"message": fmt.Sprintln("There is no data"),
// 			})
// 			return
// 		}

// 		// jika user
// 		if roles == 2 && product.UserID != userID {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this data",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func DeleteDataProductAuthorizations() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := configs.GormPostgresConn()
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		roles := uint(userData["roles"].(float64))
// 		productId := c.Param("productId")
// 		product := models.Product{}

// 		err := db.Select("user_id").Where("id = ?", productId).Order("id desc").First(&product).Error
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Data Not Found",
// 				"message": fmt.Sprintln("There is no data"),
// 			})
// 			return
// 		}

// 		// harus admin yang bisa delete produk
// 		if roles != 1 {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this data",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func UpdateProductAuthorizations() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := configs.GormPostgresConn()
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		roles := uint(userData["roles"].(float64))
// 		productId := c.Param("productId")
// 		product := models.Product{}

// 		err := db.Select("user_id").Where("id = ?", productId).Order("id desc").First(&product).Error
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Data Not Found",
// 				"message": fmt.Sprintln("There is no data"),
// 			})
// 			return
// 		}

// 		// harus admin yang bisa update produk
// 		if roles != 1 {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this data",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }
