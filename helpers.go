package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andersjanmyr/battlesnake/api"
	"github.com/andersjanmyr/battlesnake/pkg/core"
	"github.com/andersjanmyr/battlesnake/pkg/empty"
	"github.com/andersjanmyr/battlesnake/pkg/horry"
	"github.com/andersjanmyr/battlesnake/pkg/randy"
)

func initSnake(kind string) api.BattleSnake {
	switch kind {
	case "empty":
		return empty.New()
	case "horry":
		return horry.New()
	case "randy":
		return randy.New()
	default:
		return randy.New()
	}
}

var snakes = map[string]api.BattleSnake{}

func getBattleSnake(kind, id string) api.BattleSnake {
	key := fmt.Sprintf("%s-%s", kind, id)
	if snake := snakes[key]; snake != nil {
		return snake
	}
	snake := initSnake(kind)
	snakes[key] = snake
	return snake
}

func GetSnakeKinds() string {
	return "randy, horry"
}

func LocalhostToIP(next http.Handler) http.Handler {
	ip := IP()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, port := getHostPort(r.Host)
		url := fmt.Sprintf("http://%s%s%s", ip, port, r.URL.Path)
		if host == "127.0.0.1" || host == "localhost" {
			http.Redirect(w, r, url, http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getHostPort(s string) (string, string) {
	ss := strings.Split(s, ":")
	if len(ss) < 2 {
		return ss[0], ""
	}
	return ss[0], ":" + ss[1]
}

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

var files = map[string]*os.File{}

func getFile(req *api.SnakeRequest) *os.File {
	key := fmt.Sprintf("./tmp/%s-%s.csv", req.You.Name, req.You.ID)
	if file := files[key]; file != nil {
		return file
	}
	file, err := os.Create(key)
	if err != nil {
		fmt.Println(err)
	}
	files[key] = file
	return file
}

func record(req *api.SnakeRequest, moveResponse *api.MoveResponse) {
	move := "end"
	if moveResponse != nil {
		move = string(moveResponse.Move)
	}

	file := getFile(req)
	fmt.Fprintf(file, "%s,%s,%d,%t\n", boardToString(req.Board, req.You), move, req.Turn, isAlive(req))
}

func closeRecord(req *api.SnakeRequest) {
	file := getFile(req)
	file.Close()
}

func isAlive(req *api.SnakeRequest) bool {
	for _, s := range req.Board.Snakes {
		if s.ID == req.You.ID {
			return true
		}
	}
	return false
}

func boardToString(b api.Board, you api.Snake) string {
	arr := []string{}
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			v := core.ValueAt(b, you, api.Coord{X: x, Y: y})
			arr = append(arr, strconv.Itoa(int(v)))
		}
	}
	return strings.Join(arr, ",")
}
