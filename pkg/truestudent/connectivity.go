package truestudent

import "log/slog"

func CheckConnectivity() bool {
	err, userinfo := getUserInfo()
	if err != nil {
		slog.Debug("check connectivity failed", slog.String("reason", err.Error()))
		return false
	}

	return userinfo.IsLoggedIn == "yes"
}
