package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"util/agent"
	"util/encode"
	"util/logger"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var (
	optPort       int
	optDebugPort  int
	optGinMode    string
	optLogToFile  bool
	optPathPrefix string
)

const (
	dstLogFile = "/tmp/www.log"

	HTTP_UNAUTHORIZE_ACCESS  = 401
	HTTP_FORBIDDEN_ACCESS    = 403
	HTTP_FOUND               = 302
	HTTP_STATUS_SUCCESS      = 200
	HTTP_SERVICE_UNAVAILABLE = 503
	HTTP_NOTFOUND            = 404
)

func init() {
	flag.IntVar(&optPort, "port", 3333, "Http running on the port")
	flag.IntVar(&optDebugPort, "debugPort", 3330, "Port to listen to for debugging, set to 0 to disable")
	flag.StringVar(&optGinMode, "ginMode", "release", "Gin webframework running on release mode")
	flag.BoolVar(&optLogToFile, "logToFile", true, "Log write to file")
}

func setupLogging() io.WriteCloser {
	w := logger.MustWriteTo(dstLogFile)
	log.SetOutput(w)
	log.SetPrefix(fmt.Sprintf("%d ", os.Getpid()))
	return w
}

func Flags() int {
	return log.Ldate | log.LUTC | log.Lmicroseconds | log.Lshortfile
}

func main() {
	flag.Parse()
	log.SetFlags(Flags())
	var logfile io.WriteCloser

	logfile = os.Stderr
	if optLogToFile {
		logfile = setupLogging()
	}
	defer agent.Listen().Close()
	if optDebugPort > 0 {
		go func() {
			log.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", optDebugPort), nil))
		}()
	}
	fmt.Println("hello world")

	gin.SetMode(optGinMode)
	route := gin.Default()
	//// This needs to be able to called without requiring access
	route.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: true,
	}))
	route.GET(":id", func(ctx *gin.Context) {
		rand := ctx.Param("id")
		ctx.String(HTTP_STATUS_SUCCESS, rand)
	})
	route.GET("/encode/:id", func(ctx *gin.Context) {
		str := ctx.Param("id")
		ret := encode.EncodeString(str)
		ctx.String(HTTP_STATUS_SUCCESS, ret)
	})
	route.GET("/decode/:id", func(ctx *gin.Context) {
		str := ctx.Param("id")
		ret := encode.DecodeString(str)
		ctx.String(HTTP_STATUS_SUCCESS, ret)
	})
	route.GET("/knockknock", func(ctx *gin.Context) {
		ctx.String(HTTP_STATUS_SUCCESS, "Service is running...\n\r")
	})

	endless.DefaultMaxHeaderBytes = 1 << 20 //1 MB
	//It will avoid hanging connections that the client has no intention of closing
	if err := endless.ListenAndServe(fmt.Sprintf(":%d", optPort), route); err != nil {
		log.Printf("Http not running: %v\n", err)
		if optLogToFile {
			logfile.Close()
		}
		os.Exit(1)
	} else {
		log.Println("shutting down")
	}
	if optLogToFile {
		logfile.Close()
	}
	os.Exit(0)
}
