package truestudent

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

func getUserInfo() (error, userInfo) {
	var userinfo userInfo

	// get userinfo json
	resp, err := http.Get("https://clblgn.sceur.ch/goform/userinfos.json")
	if err != nil {
		slog.Debug("get userinfo failed in http get", slog.String("reason", err.Error()))
		return err, userinfo
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Debug("get userinfo failed in close body", slog.String("reason", err.Error()))
			return
		}
	}(resp.Body)

	// check http status
	if resp.StatusCode != http.StatusOK {
		slog.Debug("get userinfo failed", slog.Int("status", resp.StatusCode))
		return errors.New("http status expect 200, got " + string(rune(resp.StatusCode))), userinfo
	}

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Debug("get userinfo failed in read body", slog.String("reason", err.Error()))
		return err, userinfo
	}

	// parse json
	err = json.Unmarshal(body, &userinfo)
	if err != nil {
		slog.Debug("get userinfo failed in parse json", slog.String("reason", err.Error()))
		return err, userinfo
	}

	// unescape login url
	unescaped, err := url.QueryUnescape(userinfo.LoginURL)
	if err != nil {
		slog.Debug("get userinfo failed in unescape login url", slog.String("reason", err.Error()))
		return err, userinfo
	}
	userinfo.LoginURL = unescaped

	slog.Debug("get userinfo success")
	return nil, userinfo
}
