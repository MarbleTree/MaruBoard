


func defineEntryPoints() {
	router := gin.Default()
	//MaruBoard APIs
	//BBS posts
	router.GET("/maru/bbs/:bid/post/:id", getObjectById)
	router.POST("/maru/bbs/:bid/post/", dms.PostUploadDocumnt)
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

var db *gorm.DB
var err error

func initializeDatabase() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", def.Host, def.Port, def.User, def.Password, def.Dbname)

	// Connect to database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	createtable := `CREATE TABLE dms_tenants(id varchar(64) NOT NULL, description varchar(255) NOT NULL, email varchar(255) NOT NULL, PRIMARY KEY (id));`
	_, e := db.Exec(createtable)
	fmt.Printf("Error = [%s]", e)
	//CheckError(e)

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
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
