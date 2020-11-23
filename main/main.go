//  Golang service API for HotelScanner
//
//  Swagger spec.
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//  - multipart/form-data
//
//  Produces:
//	- application/json
//  swagger:meta
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/pkg/crawler"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/server"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
)

func initRelativePath() string {
	_, fileName, _, _ := runtime.Caller(0)
	return filepath.ToSlash(filepath.Dir(filepath.Dir(fileName))) + "/"
}

func main() {
	var serverVar bool
	var crawlerVar bool
	flag.BoolVar(&crawlerVar, "fill", false, "crawl a sites with hotels to fill bd")
	flag.BoolVar(&serverVar, "server", false, "start server")
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	configs.Init()
	if err := configs.ExportConfig(); err != nil {
		log.Fatalln(err)
	}
	db := server.InitDB()
	s3 := server.InitS3Session()

	configs.PrefixPath = initRelativePath()
	logOutput, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer logOutput.Close()

	log := logger.NewLogger(logOutput)

	if crawlerVar {
		crawler.StartCrawler(db, s3, log)
	}
	if serverVar {
		server.StartServer(db, log)
	}
}
