package dbops

import "database/sql"

func openConn() *sql.DB {
	return nil
}
func AddUserCredential(loginName, pwd string) error {
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	return "", nil
}
