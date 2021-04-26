package dbops

import "testing"

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
		{name: "1", args: args{loginName: "avenssi"}, wantErr: false},
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


func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", TestAddUserCredential)
	t.Run("Get", TestGetUserCredential)
	t.Run("Del", TestDeleteUser)
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
}