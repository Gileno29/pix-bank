package grpc

import (
	"os"

	"github.com/Gileno29/pix-bank/application/grpc"
	"github.com/Gileno29/pix-bank/infraestruture/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 55051)
}
