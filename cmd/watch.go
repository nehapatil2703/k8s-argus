package cmd

import (
	"fmt"
	"net/http"
	"os"

	argus "github.com/logicmonitor/k8s-argus/pkg"
	"github.com/logicmonitor/k8s-argus/pkg/config"
	"github.com/logicmonitor/k8s-argus/pkg/connection"
	"github.com/logicmonitor/k8s-argus/pkg/constants"
	"github.com/logicmonitor/k8s-argus/pkg/cronjob"
	"github.com/logicmonitor/k8s-argus/pkg/filters"
	"github.com/logicmonitor/k8s-argus/pkg/healthz"
	lmlog "github.com/logicmonitor/k8s-argus/pkg/log"
	"github.com/logicmonitor/k8s-argus/pkg/permission"
	ratelimiter "github.com/logicmonitor/k8s-argus/pkg/rl"
	util "github.com/logicmonitor/k8s-argus/pkg/utilities"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{ // nolint: exhaustivestruct
	Use:   "watch",
	Short: "Watch Kubernetes events",
	Long:  `Monitors a cluster autonomously by watching events and translating them to LogicMonitor objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: All objects created by the Watcher should add a property that
		//  indicates the Watcher version. This can be useful in migrations from one
		//  version to the next.

		// Add hook to log pod id and goroutine id in log context
		hook := &lmlog.DefaultFieldHook{}
		logrus.AddHook(hook)

		if kubeConfigFile != "" {
			err := os.Setenv(constants.IsLocal, "true")
			if err != nil {
				logrus.Errorf("Failed to set IsLocal environment")
			}
		}

		cronjob.Init()
		fmt.Printf("kubeconfig file path: %s\n", kubeConfigFile) // nolint: forbidigo
		if err := config.Init(kubeConfigFile); err != nil {
			fmt.Println("Failed to initialise Kubernetes client: %w", err) // nolint: forbidigo
			os.Exit(constants.ConfigInitK8sClientExitCode)
		}
		// Application configuration
		if err := config.Load(); err != nil {
			fmt.Println("failed to load application config from configmaps") // nolint: forbidigo
			os.Exit(constants.ConfigInitExitCode)
		}
		conf, err := config.GetConfig()
		if err != nil {
			fmt.Println("Failed to get config: %w", err) // nolint: forbidigo
			os.Exit(constants.GetConfigExitCode)
		}

		// Once minimal configuration gets loaded, start config watcher to watch on events
		config.Run()

		filters.Init()
		ratelimiter.Init()

		if logLevel := os.Getenv("LOG_LEVEL"); util.IsLocal() && logLevel != "" {
			level, err := logrus.ParseLevel(logLevel)
			if err != nil {
				fmt.Println("Incorrect log level, setting to info") // nolint: forbidigo
				logrus.SetLevel(logrus.InfoLevel)
			} else {
				logrus.SetLevel(level)
			}
		} else {
			logrus.SetLevel(conf.LogLevel)
			// Monitor config for log level change
			lmlog.MonitorConfig()
		}

		// Instantiate the base struct.
		base, err := argus.NewBase(conf)
		if err != nil {
			logrus.Fatal(err.Error())
			return
		}

		// Init the permission component
		permission.Init(base.K8sClient)

		if util.IsLocal() {
			logrus.SetFormatter(&logrus.TextFormatter{ // nolint: exhaustivestruct
				ForceColors: false,
			})
		} else {
			// Set up a gRPC connection and CSC Client.
			connection.Initialize(conf)
			connection.CreateConnectionHandler()
		}

		// Instantiate the application and add watchers.
		argus, err := argus.NewArgus(base)
		if err != nil {
			logrus.Fatal(err.Error())
			return
		}

		// Invoke the watcher.
		argus.Watch()

		// To update K8s & Helm properties in cluster device group periodically with the server
		cronjob.UpdateTelemetryCron(base)

		http.Handle("/metrics", promhttp.Handler())
		go func() {
			addr := ":" + conf.GetOpenmetricsPort()
			logrus.Fatal(http.ListenAndServe(addr, nil))
		}()

		// Health check.
		http.HandleFunc("/healthz", healthz.HandleFunc)

		logrus.Fatal(http.ListenAndServe(":8080", nil))
	},
}

// nolint: gochecknoinits
func init() {
	RootCmd.AddCommand(watchCmd)
}
