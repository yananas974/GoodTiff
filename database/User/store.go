package database

import ("database/sql"
"example.com/go/demoHTTP/models"
)


func NewUserStore(db *sql.DB) *UserStore {
    return &UserStore{
        DB: db,
    }
}


type UStore struct {

	models.UserStoreInterface
    
}