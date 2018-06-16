package main

import (
	"Zendesk-Exporter/src/zendesk"

	"github.com/prometheus/client_golang/prometheus"
)

type collector struct {
	zenClient *zendesk.Client
}

func (c collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc("dummy", "dummy", nil, nil)
}

func (c collector) Collect(ch chan<- prometheus.Metric) {

	m, err := c.zenClient.GetTicketStats()
	if err != nil {
		log.Errorln(err)
		return
	}

	for key, value := range m {
		if key == "count" {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("zendesk_ticket_count", "Tickets Number", []string{"status"}, nil),
				prometheus.GaugeValue,
				value,
				key)
		} else {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("zendesk_ticket", "Tickets Statistics", []string{"status"}, nil),
				prometheus.GaugeValue,
				value,
				key)
		}
	}
}
