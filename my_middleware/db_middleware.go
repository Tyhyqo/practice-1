package my_middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DbContextKey string

const dbContextKey DbContextKey = "db"

func DbMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := context.WithValue(req.Context(), dbContextKey, db)
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}

func GetDB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(dbContextKey).(*gorm.DB)
	if !ok {
		return nil
	}
	return db
}
