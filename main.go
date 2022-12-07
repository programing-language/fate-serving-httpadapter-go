package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/feature", FeatureHandler)
	if err := r.Run(); err != nil {
		panic(err)
	}
}

// 分析请求参数
func AnalysisRequest(c *gin.Context) {
	// 查询请求URL后面拼接的参数
	token := c.Query("token")
	// 查询请求URL后面的参数，如果没有填写默认值
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	// 输出 URL 后面拼接的参数
	fmt.Printf("URL参数 >>>>>> token: %s; page: %s; size: %s", token, page, size)
	fmt.Println()

	// 从表单中查询参数
	name := c.PostForm("name")
	//从取得URL中参数，此处URL中没有name字段
	name = c.Param("name")
	// 从表单中查询参数,，如果没有就获取默认值
	message := c.DefaultPostForm("message", "default")
	fmt.Printf("表单参数 >>>>>> name: %s; message: %s;", name, message)
	fmt.Println()

	// 获取header指定字段
	appKey := c.Request.Header.Get("app_key")
	accessToken := c.Request.Header.Get("access_token")
	expireTime := c.Request.Header.Get("expire_time")
	fmt.Printf("header >>>>>> app_key: %s; access_token: %s; expire_time: %s;", appKey, accessToken, expireTime)
	fmt.Println()
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	fmt.Println()

	// 获取 Body 值
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("Body >>>>>> bodyBytes:%s", bodyBytes)
}

// curl -v localhost:8080/feature
func FeatureHandler(c *gin.Context) {
	AnalysisRequest(c)

	x_data := map[string]interface{}{}
	rand.Seed(time.Now().UnixNano())
	floats_list := randFloats(-5.0, 5.0, 10)
	for i := range floats_list {
		value := fmt.Sprintf("%.6f", floats_list[i])
		key := fmt.Sprintf("x%d", i)
		x_data[key] = value
	}

	data := map[string]interface{}{
		"data": x_data,
		"code": 200,
	}
	c.JSON(200, data)
}

// 根据区间产生随机数
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}
