package main

import (
	"log"
	"net/http"
	"os"
	// "embed"
	"fmt"
	// "math"
	"math/big"
	// "log"
	"regexp"
	"strconv"
	"strings"
	// "text/template"
	"time"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func baseURL(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	query := r.URL.Query()
	filters, present := query["id"]
	if !present || len(filters) == 0 {
		fmt.Println("filters not present")
	} else {
		val, err := strconv.Atoi(filters[0])
		if err == nil {
			// numbers := make([]int, 0)
			var sb strings.Builder

			for i := 0; i < val; i++ {
				sb.WriteString("a")
			}
			sb.WriteString("c")
			regexp.MatchString("^((a)*)+b$", sb.String())

		}
	}

	elapsed := time.Since(start)
	concatenated := fmt.Sprintf("total time  taken = %v ns\n", elapsed.Nanoseconds())

	fmt.Fprintf(w, concatenated)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		start := time.Now()
		val := c.DefaultQuery("id", "0")
		id, err := strconv.Atoi(val)

		if err == nil {
			// p := 178542003245811211274167228297361192303886321036074276889145691522634525820185614278499562592134188995169731066418203258297035264969457638591284906658912408319763156912951486020761069099132619194489006875108217247513715271974383296142805846405783845170862140174184507256128825312324419293575432423822703857091;

			var p, _ = new(big.Int).SetString("178542003245811211274167228297361192303886321036074276889145691522634525820185614278499562592134188995169731066418203258297035264969457638591284906658912408319763156912951486020761069099132619194489006875108217247513715271974383296142805846405783845170862140174184507256128825312324419293575432423822703857091", 0)
			// fmt.Println(p)
			// fmt.Println(id)
			r, _ := new(big.Int).SetString("187",0);
			// fmt.Println(r)
			for i := 0; i < id; i++ {
				p.Mul(p,p)
				p.Mod(p, r)
			}
		}
		elapsed := time.Since(start)
		concatenated := fmt.Sprintf("total time  taken = %v ns\n", elapsed.Nanoseconds())

		c.String(http.StatusOK, concatenated)
	})

	router.Run(":" + port)
}
