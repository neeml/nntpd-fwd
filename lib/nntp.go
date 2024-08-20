package lib

import (
	nntp "github.com/dustin/go-nntp"
	"github.com/dustin/go-nntp/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"sync"
)

type DBBackend struct {
	DB         *gorm.DB
	NewsGroups map[string]*nntp.Group
	Lock       sync.Mutex
}

type InConnection struct {
	Conn *nntp.Conn
	DB   *DBBackend
	Lock sync.Mutex
}
