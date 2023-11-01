package truestudent

import (
	"errors"
	"html"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func getCookie(url string) (error, string) {
	cookie := ""

	resp, err := http.Get(url)
	if err != nil {
		slog.Debug("get cookie failed in http get", slog.String("reason", err.Error()))
		return err, cookie
	}
	for i, headers := range resp.Header {
		// find cookie
		if i == "Set-Cookie" {
			for _, j := range headers {
				// find cookie start with symfony=
				if strings.Contains(j, "symfony=") {
					cookie = html.UnescapeString(strings.Split(j, ";")[0])
					slog.Debug("get cookie success", slog.String("cookie", cookie))
					return nil, cookie
				}
			}
		}
	}
	// cookie not found
	return errors.New("failed to find cookie in http response header"), cookie
}

func gookieLogin(cookieStr string, getUrl string) error {

	// replace url
	getUrl = strings.Replace(getUrl, "/hscp.php?", "/hscp.php/index/login?", 1)

	// create client and request
	client := http.Client{}
	req, _ := http.NewRequest("GET", getUrl, nil)

	// set cookie
	cookieName := strings.Split(cookieStr, "=")[0]
	cookieVal := strings.Split(cookieStr, "=")[1]
	cookie := &http.Cookie{Name: cookieName, Value: cookieVal}
	req.AddCookie(cookie)

	// start request
	slog.Debug("perform http get", slog.String("url", getUrl), slog.String("cookie", cookieStr))
	resp, err := client.Do(req)
	if err != nil {
		slog.Debug("failed to request url get", slog.String("reason", err.Error()))
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Debug("failed to close http body", slog.String("reason", err.Error()))
			return
		}
	}(resp.Body)

	// check status code
	if resp.StatusCode == http.StatusOK {
		slog.Debug("cookie login success", slog.String("cookie", cookieStr))
		return nil
	}

	return errors.New("other error")
}
