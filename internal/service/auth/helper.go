package auth

import (
	dg "github.com/karincake/apem/databasegorm"
	"gorm.io/gorm"
)

// just return the error code
func getAndCheck(input, condition any) (eCode string) {
	result := dg.I.Where(condition).First(input)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "record-not-found"
		} else {
			return "fetch-fail"
		}
	} else if result.RowsAffected == 0 {
		return "auth-login-incorrect"
	}

	return ""
}
