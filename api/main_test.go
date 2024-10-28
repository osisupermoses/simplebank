package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // tells gin to use test mode, helps make logging tests cleaner

	os.Exit(m.Run())
}