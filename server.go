package main
import(
	"app/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error Loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}