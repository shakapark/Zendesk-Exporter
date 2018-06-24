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

	gs := rt.GetGlobals()
	for _, g := range *gs {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
				[]string{"priority", "status", "channel"}, nil),
			prometheus.GaugeValue,
			g.Count,
			g.Labels["priority"], g.Labels["status"], g.Labels["via"])
	}

	p := rt.GetPriority()
	for priority, value := range p {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc("zendesk_ticket_priority", "Tickets Priority Statistics", []string{"priority"}, nil),
			prometheus.GaugeValue,
			value,
			priority)
	}

	s := rt.GetStatus()
	for status, value := range s {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc("zendesk_ticket_status", "Tickets Status Statistics", []string{"status"}, nil),
			prometheus.GaugeValue,
			value,
			status)
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
