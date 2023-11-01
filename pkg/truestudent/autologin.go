package truestudent

import (
	"log/slog"
)

func AutoLogin() bool {

	// get login url
	err, userinfo := getUserInfo()
	if err != nil {
		slog.Debug("failed to get login url", slog.String("reason", err.Error()))
		return false
	}

	// get cookie
	err, cookie := getCookie(userinfo.LoginURL)
	if err != nil {
		slog.Debug("failed to get cookie", slog.String("reason", err.Error()))
		return false
	}

	// cookie login
	err = gookieLogin(cookie, userinfo.LoginURL)
	if err != nil {
		slog.Debug("failed to get login cookie", slog.String("reason", err.Error()))
		return false
	}

	// get account info
	err, accountinfo := getAccountInfo(cookie, userinfo.LoginURL)
	if err != nil {
		slog.Debug("failed to get account info", slog.String("reason", err.Error()))
		return false
	}

	// auto login
	err = login(userinfo.CustIP, userinfo.Nasid, accountinfo.Login, accountinfo.Password)
	if err != nil {
		slog.Debug("failed to login", slog.String("reason", err.Error()))
		return false
	}

	return true
}
