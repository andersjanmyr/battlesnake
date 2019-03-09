package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func ips() []string {
	ifaces, _ := net.Interfaces()
	ips := []string{}
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ips = append(ips, ip.String())
		}
	}
	return ips
}

func IP() string {
	for _, ip := range ips() {
		if strings.Contains(ip, ".") && ip != "127.0.0.1" {
			return ip
		}
	}
	return "missing"
}

func respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
	res.Write([]byte("\n"))
}

func dump(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
}

const LogFormat = `"%s %s %s %d %d" %f`

func LoggingHandler(next http.Handler) http.Handler {
	var logger = log.New(os.Stderr, "", log.LstdFlags)

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		loggingWriter := &LoggingResponseWriter{res, http.StatusOK}

		startTime := time.Now()
		next.ServeHTTP(loggingWriter, req)
		elapsedTime := time.Now().Sub(startTime)

		logger.Printf(
			LogFormat,
			req.Method, req.RequestURI, req.Proto,
			loggingWriter.statusCode, 0, elapsedTime.Seconds(),
		)

	})
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (res *LoggingResponseWriter) WriteHeader(code int) {
	res.statusCode = code
	res.ResponseWriter.WriteHeader(code)
}
