package middlewares

import (
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/Braly-Ltd/voice-changer-api-public/properties"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/exception"
	"github.com/golibs-starter/golib/web/context"
	"github.com/golibs-starter/golib/web/response"
	"strings"
)

func Authenticate(port ports.AuthenticationPort, props *properties.MiddlewaresProperties) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !props.AuthenticationEnabled {
			c.Next()
			return
		}
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, response.Error(exception.New(401, "missing access token")))
			return
		}

		agent := c.GetHeader("X-Agent")
		if agent != "" && strings.ToLower(agent[:3]) == "ios" {
			agent = "ios"
		} else {
			agent = "android"
		}
		tokenData, err := port.Authenticate(c, agent, token)
		if err != nil {
			c.AbortWithStatusJSON(401, response.Error(exception.New(403, "invalid access token")))
			return
		}

		c.Request.Header.Set("X-User-ID", tokenData.UserID)
		context.GetOrCreateRequestAttributes(c.Request).SecurityAttributes.UserId = tokenData.UserID

		c.Next()
	}
}
