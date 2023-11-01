package truestudent

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func login(ip, nasid, username, password string) error {

	postUrl := "ip=" + ip + "&nasid=" + nasid + "&username=" + username + "&password=" + password + "&r20cookie="

	resp, err := http.Post("https://clblgn.sceur.ch/goform/HtmlAjaxLoginRequest", "application/x-www-form-urlencoded", strings.NewReader(postUrl))
	if err != nil {
		slog.Debug("failed to request url post", slog.String("reason", err.Error()))
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
	if resp.StatusCode != http.StatusOK {
		slog.Debug("unexpected http status", slog.Int("status", resp.StatusCode))
		return errors.New("unexpected http status " + string(rune(resp.StatusCode)))
	}

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Debug("failed in read body", slog.String("reason", err.Error()))
		return err
	}

	if strings.Contains(string(body), "success") {
		slog.Debug("login success")
		return nil
	}

	return errors.New("other error")
}
