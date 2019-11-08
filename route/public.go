package route

import (
	"github.com/labstack/echo"
	handlers_public "github.com/letanthang/echo_stackdriver/handlers/public"
)

func Public(e *echo.Echo) {

	g := e.Group("/echo_stackdriver")

	g.GET("/health", handlers_public.HealthCheck)
	g.GET("/push_metric", handlers_public.PushMetric)
	g.GET("/push_metric_custom", handlers_public.PushCustomMetric)
}
