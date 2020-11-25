package config

import(
	"github.com/go-pg/pg"
	"fmt"
	"log"
   "github.com/joho/godotenv"
   "os"
   "strings"
)

func DbConn() (con *pg.DB) {
   godotenvErr := godotenv.Load()

   if godotenvErr != nil {
      log.Fatal("Error loading .env file")
   }

   errorList := []string{}
   errorFormat := "ERROR - %s is not found in environment variable"

   HostDatabase := os.Getenv("HOST_DATABASE")
   if HostDatabase == "" {
      errorList = append(errorList, fmt.Sprintf(errorFormat, "HOST_DATABASE"))
   }

   PortDatabase := os.Getenv("PORT_DATABASE")
   if PortDatabase == "" {
      errorList = append(errorList, fmt.Sprintf(errorFormat, "PORT_DATABASE"))
   }

   UserDatabase := os.Getenv("USER_DATABASE")
   if UserDatabase == "" {
      errorList = append(errorList, fmt.Sprintf(errorFormat, "USER_DATABASE"))
   }

   PasswordDatabase := os.Getenv("PASSWORD_DATABASE")
   if PasswordDatabase == "" {
      errorList = append(errorList, fmt.Sprintf(errorFormat, "PASSWORD_DATABASE"))
   }

   NameDatabase := os.Getenv("NAME_DATABASE")
   if NameDatabase == "" {
      errorList = append(errorList, fmt.Sprintf(errorFormat, "NAME_DATABASE"))
   }

   if(len(errorList) != 0){
      errorMessage := strings.Join(errorList, "\n")
      log.Println(errorMessage)
      os.Exit(1)
   }

   address := fmt.Sprintf("%s:%s", HostDatabase, PortDatabase)
   options := &pg.Options{
      User:     UserDatabase,
      Password: PasswordDatabase,
      Addr:     address,
      Database: NameDatabase,
      PoolSize: 50,
   }
   con = pg.Connect(options)
   if con == nil {
      log.Fatal("cannot connect to postgres")
   }

   log.Println("DB Connected...")

   return
}