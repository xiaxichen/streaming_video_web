package dbops

import "errors"

func AddUserCredential(loginName, pwd string) error {
	prepare, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) Values (?, ?)")
	if err != nil {
		return err
	}
	_, err = prepare.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	prepare, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		return "", err
	}
	var pwd string
	err = prepare.QueryRow(loginName).Scan(pwd)
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func DeleteUser(loginName, pwd string) error {
	prepare, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? and pwd=?")
	if err != nil {
		return err
	}
	exec, err := prepare.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		return err
	}
	if affected==1{
		return nil
	}else {
		return errors.New("没有影响的行数！")
	}
}
