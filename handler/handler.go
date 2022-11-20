package handler

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	DEFAULT_VERSION = "demo"
	log             = logrus.New()
)

func init() {
	log.SetLevel(logrus.GetLevel())
	log.Out = &lumberjack.Logger{
		Filename:   "./logs/handler.log",
		MaxSize:    128,
		MaxBackups: 2,
	}
}

func Register(ctx context.Context, r *gin.Engine) error {

	/*接收客户端 request，并将 request 中带的 header 写入 response header*/
	r.GET("/header", HeaderHandler())

	/*读取当前系统的环境变量中的 VERSION 配置，并写入 response header*/
	r.GET("/headerVersion", HeaderVersionHandler())

	/*Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出*/
	r.GET("/requestInfo", RequestInfoHandler())

	/*当访问 localhost/healthz 时，应返回 200*/
	r.GET("/healthz", HealthzHandler())
	return nil
}

//header handler
func HeaderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rheader := c.Request.Header
		for k, v := range rheader {
			c.Writer.Header().Set(k, strings.Join(v, ", "))
			log.WithFields(logrus.Fields{
				"header_k": k,
				"header_v": strings.Join(v, ", "),
			}).Infoln("HeaderHandler")
		}

		log.WithFields(logrus.Fields{
			"message": "headers ok!",
		}).Infoln("HeaderHandler")

		c.JSON(http.StatusOK, gin.H{
			"message": "headers ok!",
		})
	}
}

//header version handler
func HeaderVersionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rheader := c.Request.Header
		for k, v := range rheader {
			c.Writer.Header().Set(k, strings.Join(v, ", "))
		}

		version := os.Getenv("VERSION")
		if version == "" {
			version = DEFAULT_VERSION
		}
		c.Writer.Header().Set("VERSION", version)

		log.WithFields(logrus.Fields{
			"version": version,
			"message": "headers with version ok!",
		}).Infoln("HeaderVersionHandler")

		c.JSON(http.StatusOK, gin.H{
			"message": "headers with version ok!",
		})
	}
}

//request info handler
func RequestInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			log.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("split host port err")

			logrus.WithError(err).Errorf("split host port err")
		}
		fmt.Fprintf(os.Stdout, "request url %s, ip %s, status: %d", c.Request.URL, ip, http.StatusOK)

		log.WithFields(logrus.Fields{
			"request url": c.Request.URL,
			"ip":          ip,
			"status":      http.StatusOK,
			"message":     "requestInfo ok!",
		}).Infoln("RequestInfoHandler")

		c.JSON(http.StatusOK, gin.H{
			"message": "requestInfo ok!",
		})
	}
}

//healthz handler
func HealthzHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.WithFields(logrus.Fields{
			"message": "200",
		}).Infoln("HealthzHandler")

		c.JSON(http.StatusOK, gin.H{
			"message": "200",
		})
	}
}
