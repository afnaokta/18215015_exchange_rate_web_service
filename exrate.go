package main

import(
_ "github.com/go-sql-driver/mysql"
"github.com/gin-gonic/gin"
"database/sql"
"net/http"
"log")

const (
    DB_HOST = "tcp(167.205.67.251:3306)"
    DB_NAME = "exchange_rate"
    DB_USER = "root"
    DB_PASS = ""
)

func main(){
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()

    type Rate struct {
    	idCountry int
    	Country string
    	Currency string
    	Exchange float64
    	Inverse float64
    }

    router := gin.Default()

    router.GET("/Rate/:Country", func(c *gin.Context){
    	var(
    		Rate Rate
    		result gin.H
    	)

    	Country := c.Param("Country")
    	row := db.QueryRow("select Country, Exchange, Inverse from exchange where Country like '%", Country, "%';")
    	err = row.Scan(&Rate.Country, &Rate.Exchange, &Rate.Inverse)
		if err != nil {
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": Rate,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result) 
    })
    router.Run(":1076")
}