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
	"github.com/andersjanmyr/battlesnake/pkg/hungry"
	"github.com/andersjanmyr/battlesnake/pkg/randy"
)

var moveMap map[api.Move]int

func init() {
	moveMap = map[api.Move]int{}
	moveMap[api.Up] = 1
	moveMap[api.Down] = 2
	moveMap[api.Left] = 3
	moveMap[api.Right] = 4
}

func initSnake(kind string) api.BattleSnake {
	switch kind {
	case "empty":
		return empty.New()
	case "horry":
		return horry.New()
	case "randy":
		return randy.New()
	case "hungry":
		return hungry.New()
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
	return "randy, horry, hungry"
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

type Data struct {
	values *os.File
	labels *os.File
}

var datas = map[string]Data{}

func getFile(req *api.SnakeRequest) Data {
	key := fmt.Sprintf("./tmp/%s-%s.csv", req.You.Name, req.You.ID)
	labelsKey := fmt.Sprintf("./tmp/%s-%s-labels.csv", req.You.Name, req.You.ID)
	if data, ok := datas[key]; ok {
		return data
	}
	values, err := os.Create(key)
	labels, err := os.Create(labelsKey)
	data := Data{values: values, labels: labels}
	if err != nil {
		fmt.Println(err)
	}
	datas[key] = data
	return data
}

func record(req *api.SnakeRequest, moveResponse *api.MoveResponse) {

	data := getFile(req)
	fmt.Fprintln(data.values, boardToString(req.Board, req.You))
	if moveResponse == nil {
		fmt.Fprintln(data.labels, 0)
	} else {
		fmt.Fprintln(data.labels, moveToNumber(moveResponse.Move))
	}
}

func moveToNumber(move api.Move) int {
	return moveMap[move]
}

func closeRecord(req *api.SnakeRequest) {
	data := getFile(req)
	data.values.Close()
	data.labels.Close()
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
