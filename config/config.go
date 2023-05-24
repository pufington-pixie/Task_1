package config
 
import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)
 
func ConnectDB() (*sql.DB,error) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "ganesh"
    dbName := "project"
    
    
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName )
    if err != nil {
        panic(err.Error())
    }
    return db,nil
    
}