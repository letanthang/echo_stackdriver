package migrate

import (
	"encoding/json"
	"fmt"
	"github.com/letanthang/echo_stackdriver/types"
	"gopkg.in/mgo.v2/bson"
)

func InsertIntoPSQL(obj types.ProfileFull) {
	db := GetDB()
	sqlStr := "INSERT INTO contact(id, account_id, type, fullname, phone, address, location, district, province, region_id, email) VALUES "
	for _, c := range obj.Contact {
		contact := types.Contact{}
		db.Raw("select * from contact where id=?", c.ID).Scan(&contact)
		if contact.ID == 0 {
			c.CustomerID = obj.CustomerID
			sqlStr += fmt.Sprintf("(%d, %d, %d, '%s', '%s', '%s',  '{%v,%v}', '%s', '%s', %d, '%s'),",
				c.ID, c.CustomerID, c.Type, c.Fullname, c.Phone, c.Address, c.Location[0], c.Location[1], c.District, c.Province, c.RegionID, c.Email)
		}
	}
	sqlStr = sqlStr[0:len(sqlStr)-1]
	db.Exec(sqlStr)
	fmt.Printf("Migrated:  %s", obj.Phone)
}

