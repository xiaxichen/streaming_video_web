package dbops

import (
	"streaming_video_web/api/def"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestAddUserCredential(t *testing.T) {
	type args struct {
		loginName string
		pwd       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				loginName: "avenssi",
				pwd:       "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddUserCredential(tt.args.loginName, tt.args.pwd); (err != nil) != tt.wantErr {
				t.Errorf("AddUserCredential() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetUserCredential(t *testing.T) {
	type args struct {
		loginName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1", args: args{loginName: "avenssi"}, want: "123", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserCredential(tt.args.loginName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserCredential() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDeleteUser(t *testing.T) {
	type args struct {
		loginName string
		pwd       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				loginName: "avenssi",
				pwd:       "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.loginName, tt.args.pwd); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var displayCtime string

func TestAddNewVideo(t *testing.T) {
	type args struct {
		aid  int
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *def.VideoInfo
		wantErr bool
	}{
		{
			name: "add new video",
			args: args{
				aid:  1,
				name: "my-video",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddNewVideo(tt.args.aid, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddNewVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			displayCtime = got.DisplayCtime
		})
	}
}
func TestGetVideoInfo(t *testing.T) {
	type args struct {
		vid string
	}
	tests := []struct {
		name    string
		args    args
		want    *def.VideoInfo
		wantErr bool
	}{
		{name: "get video", args: args{vid: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetVideoInfo(tt.args.vid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideoInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
func TestDeleteVideoInfo(t *testing.T) {
	type args struct {
		vid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "del video",
			args:    args{"1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteVideoInfo(tt.args.vid); (err != nil) != tt.wantErr {
				t.Errorf("DeleteVideoInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserWorkFlow(t *testing.T) {
	t.Log("------------user------------")
	t.Run("Add", TestAddUserCredential)
	t.Run("Get", TestGetUserCredential)
	t.Run("Del", TestDeleteUser)
	t.Log("------------videos------------")
	t.Run("Add", TestAddUserCredential)
	t.Run("AddVideo", TestAddNewVideo)
	t.Run("GetVideo", TestGetVideoInfo)
	t.Run("DelVideo", TestDeleteVideoInfo)
	t.Log("------------comment------------")
	t.Run("AddVideo", TestAddNewVideo)
	t.Run("AddComment", TestAddNewComments)
	t.Run("GetCommentList", TestListComments)
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
}

func TestAddNewComments(t *testing.T) {
	type args struct {
		vid     string
		aid     int
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "add comment",
			args:    args{aid: 1, vid: "1", content: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddNewComments(tt.args.vid, tt.args.aid, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("AddNewComments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListComments(t *testing.T) {
	type args struct {
		vid  string
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []*def.Comment
		wantErr bool
	}{
		{
			name: "test get comment list",
			args: args{
				vid:  "1",
				from: 0,
				to:   int(time.Now().Unix()),
			},
			wantErr: false,
			want:    []*def.Comment{{Id: "1", VideoId: "1", Author: "avenssi", Content: "test"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ListComments(tt.args.vid, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
