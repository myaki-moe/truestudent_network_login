package truestudent

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

func getAccountInfo(cookieStr string, postUrl string) (error, accountInfo) {
	var accountinfo accountInfo

	// replace url
	postUrl = strings.Replace(postUrl, "/hscp.php?", "/hscp.php/ajax/doFta?", 1)

	// create client and request
	client := http.Client{}
	req, _ := http.NewRequest("POST", postUrl, nil)

	// set cookie
	cookieName := strings.Split(cookieStr, "=")[0]
	cookieVal := strings.Split(cookieStr, "=")[1]
	cookie := &http.Cookie{Name: cookieName, Value: cookieVal}
	req.AddCookie(cookie)

	// start request
	slog.Debug("perform http post", slog.String("url", postUrl), slog.String("cookie", cookieStr))
	resp, err := client.Do(req)
	if err != nil {
		slog.Debug("failed to request url post", slog.String("reason", err.Error()))
		return err, accountinfo
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Debug("failed to close http body", slog.String("reason", err.Error()))
			return
		}
	}(resp.Body)

	// check status code
	if resp.StatusCode != http.StatusOK {
		slog.Debug("unexpected http status", slog.Int("status", resp.StatusCode))
		return errors.New("unexpected http status " + string(rune(resp.StatusCode))), accountinfo
	}

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Debug("failed in read body", slog.String("reason", err.Error()))
		return err, accountinfo
	}

	// parse json
	err = json.Unmarshal(body, &accountinfo)
	if err != nil {
		slog.Debug("failed in parse json", slog.String("reason", err.Error()))
		return err, accountinfo
	}

	// unescape login
	unescaped, err := url.QueryUnescape(accountinfo.Login)
	if err != nil {
		slog.Debug("failed in unescape", slog.String("reason", err.Error()))
		return err, accountinfo
	}
	accountinfo.Login = unescaped

	// unescape password
	unescaped, err = url.QueryUnescape(accountinfo.Password)
	if err != nil {
		slog.Debug("failed in unescape", slog.String("reason", err.Error()))
		return err, accountinfo
	}
	accountinfo.Password = unescaped

	return nil, accountinfo
}
