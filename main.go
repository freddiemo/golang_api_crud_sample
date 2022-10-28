package main


import(
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
	"net/http"

	"github.com/golang_api_crud_sample/service"
	"github.com/golang_api_crud_sample/controller"
)

func setupEnvironment() map[string]string {
	params, err := godotenv.Read(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	return params
}


func main() {

	envs := setupEnvironment()

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Welcome!",
		})
	})

	var userService service.UserService = service.New()
	var userController controller.UserController = controller.New(userService)

	server.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.Save(ctx))
	})

	server.Run(":" + envs["APP_PORT"])
}
