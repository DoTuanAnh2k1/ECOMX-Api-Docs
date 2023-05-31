package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Data map[string]interface{} `yaml:",inline"`
}

func main() {
	// Đọc file YAML
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Giải mã cấu trúc YAML vào biến config
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Khởi tạo router Gin
	router := gin.Default()

	// Định nghĩa route
	router.GET("/", func(c *gin.Context) {
		// Truy cập các giá trị từ file YAML
		for key, value := range config.Data {
			c.String(http.StatusOK, "%s: %v\n", key, value)
		}
	})

	// Khởi động máy chủ HTTP
	log.Fatal(router.Run(":1234"))
}
