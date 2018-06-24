package main

import (
	"Zendesk-Exporter/src/config"
	"Zendesk-Exporter/src/zendesk"

	"github.com/prometheus/client_golang/prometheus"
)

type collector struct {
	zenClient *zendesk.Client
	filter    *config.Filter
}

func (c collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc("dummy", "dummy", nil, nil)
}

func (c collector) Collect(ch chan<- prometheus.Metric) {

	atf, err := c.zenClient.SetAllTicketField()
	if err != nil {
		log.Errorln(err)
		return
	}

	rt, err := c.zenClient.GetTicketStats(atf, c.filter.CustomFields)
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

	if c.filter.Priority {
		p := rt.GetPriority()
		for priority, value := range p {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("zendesk_ticket_priority", "Tickets Priority Statistics", []string{"priority"}, nil),
				prometheus.GaugeValue,
				value,
				priority)
		}
	}

	if c.filter.Status {
		s := rt.GetStatus()
		for status, value := range s {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("zendesk_ticket_status", "Tickets Status Statistics", []string{"status"}, nil),
				prometheus.GaugeValue,
				value,
				status)
		}
	}

	if c.filter.Channel {
		v := rt.GetVia()
		for key, value := range v {
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("zendesk_ticket_channel", "Tickets Channel Statistics", []string{"channel"}, nil),
				prometheus.GaugeValue,
				value,
				key)
		}
	}

}
