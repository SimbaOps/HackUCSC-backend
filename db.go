package sccommute

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
)

func init() {
	go initClientPool(2)
}

type LocPing struct {
	BusId string  `json:"id"`
	Lat   float32 `json:"lat"`
	Long  float32 `json:"lon"`
	Type  string  `json:"type"`
}

var postChannel chan LocPing = make(chan LocPing)

//Creates a pool of clients
func initClientPool(num int) {
	for i := 0; i < num; i++ {
		go initPost(postChannel)
	}
}

//Initializes GoRoutine that listens to postChannel
//and posts all inputs to the DB
func initPost(channel chan LocPing) {
	/*db, err := sql.Open("mysql", "root:@/confessions")
	if err != nil {
		log.Fatal("Unable to open mysql connection")
	}
	defer db.Close()
	stmt, err := db.Prepare("")
	if err != nil {
		log.Fatal("Unable to create prepared statement", err)
	}
	defer stmt.Close()*/
	for loc := range channel {
		/*_, err := stmt.Exec()
		if err != nil {
			fmt.Println(conf, err)
		}*/
		postIM(loc)
		fmt.Println(loc)
	}
}

//Takes location object and update it in the DB
func Post(loc LocPing) {
	postChannel <- loc
}

//Gets location of the object with this objId
func Get(id string) LocPing {
	return getIM(id)
}
