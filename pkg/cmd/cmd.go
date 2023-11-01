package cmd

import (
	"log/slog"
	"os"
	"time"
	"truestudent_network_login/pkg/config"
	"truestudent_network_login/pkg/truestudent"
	"truestudent_network_login/pkg/version"
)

func Execute() error {

	// parse args
	runningConfig := config.Parse()

	// set log level
	var logger *slog.Logger
	if runningConfig["verbose"].(bool) {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	slog.SetDefault(logger)

	// print version
	slog.Info("truestudent network login",
		slog.String("version", version.GetVersion()),
		slog.String("commit", version.GetCommit()),
		slog.String("buildDate", version.GetBuildDate()),
	)

	// args verbose output
	slog.Debug("running with config",
		slog.Bool("verbose", runningConfig["verbose"].(bool)),
		slog.Bool("run_once", runningConfig["runOnce"].(bool)),
		slog.Bool("force_login", runningConfig["forceLogin"].(bool)),
		slog.Duration("check_duration", runningConfig["checkDuration"].(time.Duration)),
		slog.Int("retry_count", runningConfig["retryCount"].(int)),
	)

	for {
		needAutoLogin := false

		// check connectivity
		slog.Info("checking connectivity")

		connectivity := false
		for i := 0; i < runningConfig["retryCount"].(int); i++ {
			if truestudent.CheckConnectivity() {
				connectivity = true
				break
			}
		}

		if connectivity {
			slog.Info("connection normal")
		} else {
			slog.Info("connection lost")
			needAutoLogin = true
		}

		// force autologin
		if runningConfig["forceLogin"].(bool) {
			slog.Info("force autologin")
			needAutoLogin = true
		}

		// auto login
		if needAutoLogin {
			slog.Info("perform autologin")

			loginSucc := false
			for i := 0; i < runningConfig["retryCount"].(int); i++ {
				if truestudent.AutoLogin() {
					loginSucc = true
					break
				}
			}
			if loginSucc {
				slog.Info("autologin perform success")
			} else {
				slog.Info("autologin perform failed")
			}
		}

		if runningConfig["runOnce"].(bool) {
			slog.Info("run once is set to true, exiting")
			break
		}

		// sleep
		slog.Info("wait for next check", slog.Duration("duration", runningConfig["checkDuration"].(time.Duration)))
		time.Sleep(runningConfig["checkDuration"].(time.Duration))
	}

	return nil
}
