package main


import(
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang_api_crud_sample/controller"
	"github.com/golang_api_crud_sample/repository"
	"github.com/golang_api_crud_sample/service"
)

func setupEnvironment() map[string]string {
	params, err := godotenv.Read(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	return params
}

func setupLogOutput(app_name string) {
	f, _ := os.Create(app_name + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	envs := setupEnvironment()
	setupLogOutput(envs["APP_NAME"])
	var db = repository.NewConnection(envs["DB_HOST"], envs["DB_USER"], envs["DB_PASSWORD"], envs["DB_NAME"], envs["DB_PORT"])

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Welcome!",
		})
	})

	var userRepository repository.UserRepository = repository.NewUserRepository(db)
	var userService service.UserService = service.New(userRepository)
	var userController controller.UserController = controller.New(userService)

	apiRoutes := server.Group("/v1")
	{
		apiRoutes.GET("/users", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, userController.FindAll())
		})

		apiRoutes.POST("/users", func(ctx *gin.Context) {
			user, err := userController.Save(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, user)
			}
		})

		apiRoutes.PUT("/users/:id", func(ctx *gin.Context) {
			user, err := userController.Update(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, user)
			}
		})

		apiRoutes.DELETE("/users/:id", func(ctx *gin.Context) {
			err := userController.Delete(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
			}
		})

	}

	server.Run(":" + envs["APP_PORT"])
}
