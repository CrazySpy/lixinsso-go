package lixinsso

import "testing"

func TestLogin(t *testing.T) {
	user := New("123", "345")
	if user.Login() {
		t.Fail()
	}

	username := "testuser"
	password := "testpass"
	if username == "testuser" && password == "testpass" {
		t.Fatal("The username and password are fake. You can use yourselves's account.")
	}

	user = New(username, password)
	if !user.Login() {
		t.Fail()
	}

	auth := NewAuthorize("ufsso_mry_linshi", Code, "http://mcenter.lixin.edu.cn/callback2.jsp")
	if !auth.AuthorizeApp(nil) {
		t.Fail()
	}
}
