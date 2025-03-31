package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	def "github.com/MarbleTree/MaruBoard/internal/api"
	"github.com/MarbleTree/MaruBoard/internal/bbs"
	"github.com/gin-gonic/gin"
)

func defineEntryPoints() {
	router := gin.Default()
	//MaruBoard APIs
	//BBS boards
	router.POST("/maru/bbs", bbs.CreateMaruBoard)

	//BBS posts
	// router.GET("/maru/bbs/:bid/post/:id", getObjectById)
	// router.POST("/maru/bbs/:bid/post/", dms.PostUploadDocumnt)
	// //multiple file operation
	// router.GET("/api/dms/objects", getObjectList)
	// router.POST("/api/dms/objects", dms.PostUploadDocumnts)

	// router.POST("/api/dms/folder/current", dms.SetCurrentFolder)
	// router.GET("/api/dms/folder/current", dms.GetCurrentFolder)

	// //ADMIN APIs
	// router.POST("/api/admin/tenant", admin.PostCreateTenant)
	// router.GET("/api/admin/tenant", admin.GetTenants)
	// router.GET("/api/admin/tenant/:id", admin.GetTenantById)
	// router.GET("/api/admin/tenant/:id", admin.GetTenantById)

	router.Run("localhost:9132")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func initializeDatabase() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", def.Host, def.Port, def.User, def.Password, def.Dbname)

	// Connect to database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	//see if table is already exist
	dbstmt := `SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename  = 'maru_boards');`
	rows, e := db.Query(dbstmt)
	CheckError(e)

	// Foreach movie
	var exist string
	rows.Next()
	err = rows.Scan(&exist)
	// check errors
	CheckError(err)
	//table is not exist, let's initialize it
	if exist == "false" {
		dbstmt := `CREATE TABLE maru_boards(bid uuid NOT NULL, name varchar(255) NOT NULL, owner varchar(255) NOT NULL, options json NOT NULL, PRIMARY KEY (bid));`
		_, e := db.Exec(dbstmt)
		CheckError(e)
		dbstmt = `CREATE TABLE maru_contents(cid varchar(64) NOT NULL, uid varchar(64) NOT NULL, bid uuid NOT NULL, uid_name varchar(64) NOT NULL, title varchar(255) NOT NULL, view integer NOT NULL, content text NOT NULL, PRIMARY KEY (cid));`
		_, e = db.Exec(dbstmt)
		CheckError(e)
	}

	// insertStmt := `CREATE TABLE dms_root(id varchar(64) NOT NULL, description varchar(255) NOT NULL, email varchar(255) NOT NULL, PRIMARY KEY (id));`
	// _, e = db.Exec(insertStmt)
	// CheckError(e)

	// rows, err := db.Query(`SELECT * FROM "dms_tenants"`)
	// CheckError(err)

	// defer rows.Close()
	// for rows.Next() {
	// 	var name string
	// 	var roll int

	// 	err = rows.Scan(&name, &roll)
	// 	CheckError(err)

	// 	fmt.Println(name, roll)
	// }
	// CheckError(err)

	// insertStmt := `insert into "dms_tenants"("id", "description", "root") values('HSI', 'John', 'E:/test')`
	// _, e = db.Exec(insertStmt)
	// CheckError(e)

	// check db
	err = db.Ping()
	CheckError(err)

}

func main() {
	initializeDatabase()

	defineEntryPoints()
}

func getObjectById(c *gin.Context) {
	// id := c.Param("id")

}
