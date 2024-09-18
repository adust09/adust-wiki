package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=pass dbname=go_todo port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Todo{})

	r := gin.Default()

	r.POST("/todos", createTodo)
	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodoByID)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Run()
}

type Todo struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Deadline string `json:"deadline"`
	Weather string `json:"weather"`
	Project string `json:"project"`
}

// モック用のAPI関数
func fetchCalendarAPI() (string, error) {
	time.Sleep(2 * time.Second) // 遅延をシミュレート
	return "2024-09-30", nil
}

func fetchWeatherAPI() (string, error) {
	time.Sleep(1 * time.Second) // 遅延をシミュレート
	return "Sunny", nil
}

func fetchProjectAPI() (string, error) {
	time.Sleep(3 * time.Second) // 遅延をシミュレート
	return "Project X", nil
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func createTodoByAPI(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	todo := &ToDO {
		Title: "New Task",
		Status: "pending"
	}

	wg.Add(3)

	// 	•	無名関数の()には引数を渡すことができ、引数を受け取るように定義して、その場で値を渡して実行できます。
	// 	•	**ifブロックの最後の()**は、関数呼び出しを示している場合が多く、無名関数をその場で定義してすぐに実行していることを意味します。
	go func() {
		defer wg.Done()
		deadline, err :- fetchCalendarAPI()
		if err != nil {
			log.Println("Error fetching API:", err)
			return
		}
		todo.Deadline = deadline
	}()

	go func(){
		defer wg.Done()
		weather, err := fetchWeatherAPI()
		if err != nil {
			log.Println("Error fetching weather API:", err)
			return
		}
		todo.Weather = weather
	}()


	go func() {
		defer wg.Done()
		project, err := fetchProjectAPI()
		if err != nill{
			log.Println("Error fetching project API:", err) 
				return
			todo.Project = project
		}()
	}
	// 全てのゴルーチンの完了を待機

	wg.Wait()
	// ToDoをレスポンスとして返す

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func getTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
	var todo Todo
	if err := db.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
	var todo Todo
	if err := db.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	var todo Todo
	if err := db.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
