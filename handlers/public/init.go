package handlers_public

import (
	"log"

	"github.com/liip/sheriff"
)

var optionPublic = &sheriff.Options{
	Groups: []string{"public"},
}

func init() {
	// Init database.
	log.Println("Public handler : Initializing \n")

}
