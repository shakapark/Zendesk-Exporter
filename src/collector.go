package main

import (
	"Zendesk-Exporter/src/config"
	"Zendesk-Exporter/src/zendesk"
	"reflect"

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

	rt, err := c.zenClient.GetTicketStats(atf, *c.filter)
	if err != nil {
		log.Errorln(err)
		return
	}
	log.Debugln("ResultTicket: ", rt)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("zendesk_ticket_count", "Tickets Number", nil, nil),
		prometheus.GaugeValue,
		rt.GetCount())

	if c.filter.Global {
		gs := rt.GetGlobals()
		for _, g := range *gs {
			var cf []reflect.Value
			if c.filter.CustomFields.Enable {
				cf = reflect.ValueOf(c.filter.CustomFields.Fields).MapKeys()
			} else {
				cf = []reflect.Value{}
			}
			switch len(cf) {
			case 0:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel"}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"])
			case 1:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel", cf[0].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"], g.Labels[cf[0].String()])
			case 2:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()])
			case 3:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String(), cf[2].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()], g.Labels[cf[2].String()])
			case 4:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String(), cf[2].String(), cf[3].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()], g.Labels[cf[2].String()], g.Labels[cf[3].String()],
					g.Labels[cf[4].String()])
			case 5:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String(), cf[2].String(), cf[3].String(),
							cf[4].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()], g.Labels[cf[2].String()], g.Labels[cf[3].String()],
					g.Labels[cf[4].String()])
			case 6:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String(), cf[2].String(), cf[3].String(),
							cf[4].String(), cf[5].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()], g.Labels[cf[2].String()], g.Labels[cf[3].String()],
					g.Labels[cf[4].String()], g.Labels[cf[5].String()])
			case 7:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String(), cf[2].String(), cf[3].String(),
							cf[4].String(), cf[5].String(), cf[6].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()], g.Labels[cf[2].String()], g.Labels[cf[3].String()],
					g.Labels[cf[4].String()], g.Labels[cf[5].String()], g.Labels[cf[6].String()])
			case 8:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String(), cf[2].String(), cf[3].String(),
							cf[4].String(), cf[5].String(), cf[6].String(), cf[7].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()], g.Labels[cf[2].String()], g.Labels[cf[3].String()],
					g.Labels[cf[4].String()], g.Labels[cf[5].String()], g.Labels[cf[6].String()], g.Labels[cf[7].String()])
			case 9:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String(), cf[2].String(), cf[3].String(),
							cf[4].String(), cf[5].String(), cf[6].String(), cf[7].String(),
							cf[8].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()], g.Labels[cf[2].String()], g.Labels[cf[3].String()],
					g.Labels[cf[4].String()], g.Labels[cf[5].String()], g.Labels[cf[6].String()], g.Labels[cf[7].String()],
					g.Labels[cf[8].String()])
			case 10:
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel",
							cf[0].String(), cf[1].String(), cf[2].String(), cf[3].String(),
							cf[4].String(), cf[5].String(), cf[6].String(), cf[7].String(),
							cf[8].String(), cf[9].String()}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"],
					g.Labels[cf[0].String()], g.Labels[cf[1].String()], g.Labels[cf[2].String()], g.Labels[cf[3].String()],
					g.Labels[cf[4].String()], g.Labels[cf[5].String()], g.Labels[cf[6].String()], g.Labels[cf[7].String()],
					g.Labels[cf[8].String()], g.Labels[cf[9].String()])

			default:
				log.Errorln("Error: Too much Custom Fields")
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc("zendesk_ticket", "Tickets Statistics",
						[]string{"priority", "status", "channel"}, nil),
					prometheus.GaugeValue,
					g.Count,
					g.Labels["priority"], g.Labels["status"], g.Labels["via"])

			}
		}
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
