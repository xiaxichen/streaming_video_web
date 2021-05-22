package dbops

import (
	"database/sql"
	"errors"
	"streaming_video_web/api/def"
	"streaming_video_web/api/utils"
	"time"
)

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
	err = prepare.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
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
	if affected == 1 {
		return nil
	} else {
		return errors.New("没有影响的行数！")
	}
}

func AddNewVideo(aid int, name string) (*def.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info 
		(id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &def.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*def.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &def.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*def.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*def.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &def.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}
