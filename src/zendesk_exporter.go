package main

import (
	"Zendesk-Exporter/src/config"
	"Zendesk-Exporter/src/zendesk"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	promlog "github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	sc = config.SafeConfig{
		C: &config.Config{},
	}

	log promlog.Logger

	configFile    = kingpin.Flag("config.file", "Compteur configuration file.").Default("zendesk.yml").String()
	listenAddress = kingpin.Flag("web.listen-address", "The address to listen on for HTTP requests.").Default(":9146").String()
	logLevel      = kingpin.Flag("log.level", "Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal]").Default("info").String()
)

func init() {
	prometheus.MustRegister(version.NewCollector("zendesk_exporter"))
}

func setZendeskClient(c *config.Config) (*zendesk.Client, error) {
	if c.Token != "" {
		return zendesk.NewClientByToken(c.URL, c.Login, c.Token)
	}
	return zendesk.NewClientByPassword(c.URL, c.Login, c.Password)
}

func zendeskHandler(w http.ResponseWriter, r *http.Request, z *zendesk.Client) {
	registry := prometheus.NewRegistry()
	collector := collector{zenClient: z}
	registry.MustRegister(collector)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	log = promlog.Base()
	if err := log.SetLevel(*logLevel); err != nil {
		log.Fatal("Error: ", err)
	}

	log.Infoln("Starting Zendesk-Exporter")

	if err := sc.ReloadConfig(*configFile); err != nil {
		log.Fatal("Error loading config: ", err)
		os.Exit(1)
	}
	log.Infoln("Loaded config file")
	sc.Lock()
	conf := sc.C
	sc.Unlock()

	hup := make(chan os.Signal)
	reloadCh := make(chan chan error)
	signal.Notify(hup, syscall.SIGHUP)

	go func() {
		for {
			select {
			case <-hup:
				if err := sc.ReloadConfig(*configFile); err != nil {
					log.Errorln("Error reloading config:", err)
					continue
				}
				log.Infoln("Reloaded config file")
			case rc := <-reloadCh:
				if err := sc.ReloadConfig(*configFile); err != nil {
					log.Errorln("Error reloading config:", err)
					rc <- err
				} else {
					log.Infoln("Reloaded config file")
					rc <- nil
				}
			}
		}
	}()

	zen, err := setZendeskClient(conf)
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/-/reload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "This endpoint requires a POST request.\n")
			return
		}

		rc := make(chan error)
		reloadCh <- rc
		if err := <-rc; err != nil {
			http.Error(w, fmt.Sprintf("Failed to reload config: %s", err), http.StatusInternalServerError)
			return
		}
		tmp, err := setZendeskClient(conf)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to reload config: %s", err), http.StatusInternalServerError)
			return
		}
		zen = tmp
	})

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/zendesk", func(w http.ResponseWriter, r *http.Request) {
		zendeskHandler(w, r, zen)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<html>
				<head>
					<title>Zendesk-Exporter</title>
				</head>
				<body>
					<h1>Zendesk-Exporter</h1>
					<p><a href="/zendesk">Zendesk Statistics</a></p>
				</body>
			</html>`))
	})

	log.Infoln("Listening on:", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal("Error: Can't starting HTTP server: ", err)
		os.Exit(1)
	}
	m, err := zen.GetTicketStats()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m)
	}
}
