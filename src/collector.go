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

	rt, err := c.zenClient.GetTicketStats()
	if err != nil {
		log.Errorln(err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("zendesk_ticket_count", "Tickets Number", nil, nil),
		prometheus.GaugeValue,
		rt.GetCount())

	m := rt.GetStatus()
	for key, value := range m {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc("zendesk_ticket", "Tickets Statistics", []string{"status"}, nil),
			prometheus.GaugeValue,
			value,
			key)
	}

	v := rt.GetVia()
	for key, value := range v {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc("zendesk_ticket_channel", "Tickets Channel Statistics", []string{"channel"}, nil),
			prometheus.GaugeValue,
			value,
			key)
	}

}
