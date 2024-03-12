package middlewares

import (
	"github.com/golibs-starter/golib/web/constant"
	"github.com/golibs-starter/golib/web/context"
	"net/http"
	"strings"
)

func AddCustomHeaders() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqAttr := context.GetOrCreateRequestAttributes(r)

			if sessionID := r.Header.Get("X-Session-ID"); sessionID != "" {
				r.Header.Set(constant.HeaderDeviceSessionId, sessionID)
				reqAttr.DeviceSessionId = sessionID
			}

			if agent := r.Header.Get("X-Agent"); agent != "" {
				r.Header.Set(constant.HeaderUserAgent, agent)
				reqAttr.UserAgent = agent
			}

			if appVer := r.Header.Get("X-App-Version"); appVer != "" {
				r.Header.Set(constant.HeaderServiceClientName, appVer)
				reqAttr.CallerId = appVer
			}

			remoteAddress := r.Header.Get(constant.HeaderClientIpAddress)
			if remoteAddress == "" {
				remoteAddress = r.RemoteAddr
			}
			addressOnly := strings.Split(remoteAddress, ":")[0]
			r.Header.Set(constant.HeaderClientIpAddress, addressOnly)
			reqAttr.ClientIpAddress = addressOnly

			next.ServeHTTP(w, r)
		})
	}
}
