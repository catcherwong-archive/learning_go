package main

import(
	"github.com/gin-gonic/gin"
	"learning_go/gin-demo/models"
	"strconv"
)

func index(c *gin.Context){
	c.String(200,"welcome to gin demo")
}

func getAllPersons(c *gin.Context){
	persons := models.GetAllPersons()
	c.JSON(200,persons)
}

func getPersonById(c *gin.Context){
	personId , _ := strconv.Atoi(c.Param("id")) 
	person := models.GetPersonById(personId)
	c.JSON(200,person)
}

func main(){
	r := gin.Default()
	
	r.GET("/",index)
	r.GET("/api/persons",getAllPersons)
	r.GET("/api/persons/:id",getPersonById)

	r.Run(":8000")
}

