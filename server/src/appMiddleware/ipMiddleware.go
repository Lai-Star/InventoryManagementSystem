package app_middleware

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type contextKey string

const contextUserKey contextKey = "user_ip"

// Receive user's ip address from the context
func IpFromContext(ctx context.Context) string {
	return ctx.Value(contextUserKey).(string)
}

func AddIPToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a context
		var ctx = context.Background()

		// get the ip (as accurately as possible)
		ip, err := GetIP(r)
		if err != nil {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
			if len(ip) == 0 {
				ip = "unknown"
			}
			ctx = context.WithValue(r.Context(), contextUserKey, ip)
		} else {
			ctx = context.WithValue(r.Context(), contextUserKey, ip)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetIP(r *http.Request) (string, error) {
	// split to the specific ip and port (but don't need the port number)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown", err
	}

	// Checking if the ip address that we got can be parsed (e.g., 192.168.0.104:8080)
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return "", fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	}

	forward := r.Header.Get("X-Forwarded-For")
	if len(forward) > 0 {
		ip = forward
	}

	// in case there is no ip
	if len(ip) == 0 {
		ip = "forward"
	}

	return ip, nil
}
