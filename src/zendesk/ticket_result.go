package zendesk

import (
	"errors"
)

var (
	listPriority = []string{"urgent", "high", "normal", "low", "undefined"}
	listStatus   = []string{"new", "open", "pending", "hold", "solved", "closed"}
	listVia      = []string{"web", "mobile", "rule", "system", "twitter", "email", "chat"}
)

func getEmptyGlobal() []Global {
	t := make([]Global, 0)

	for _, p := range listPriority {
		for _, s := range listStatus {
			for _, v := range listVia {
				t = append(t, Global{
					Labels: map[string]string{
						"priority": p,
						"status":   s,
						"via":      v,
					},
					Count: 0,
				})
			}
		}
	}

	return t
}

func getEmptyPriority() map[string]float64 {
	m := make(map[string]float64)
	for _, p := range listPriority {
		m[p] = 0
	}
	return m
}

func getEmptyStatus() map[string]float64 {
	m := make(map[string]float64)
	for _, s := range listStatus {
		m[s] = 0
	}
	return m
}

func getEmptyVia() map[string]float64 {
	m := make(map[string]float64)
	for _, v := range listVia {
		m[v] = 0
	}
	return m
}

//Global Stock result with all label in parameters
type Global struct {
	Labels map[string]string
	Count  float64
}

//GetGlobal Return Global of []Global with right labels
func GetGlobal(gs *[]Global, labels map[string]string) (*Global, int, error) {
	ok := false
	for k, g := range *gs {
		for key, value := range g.Labels {
			if value != labels[key] {
				ok = false
				break
			} else {
				ok = true
			}
		}
		if ok {
			return &g, k, nil
		}
	}

	return nil, -1, errors.New("Global not found")
}

//ResultTicket Stock result
type ResultTicket struct {
	count    float64
	global   []Global
	priority map[string]float64
	status   map[string]float64
	via      map[string]float64
}

//NewResultTicket Create new ResultTicket
func NewResultTicket() *ResultTicket {
	return &ResultTicket{
		count:    0,
		global:   getEmptyGlobal(),
		priority: getEmptyPriority(),
		status:   getEmptyStatus(),
		via:      getEmptyVia(),
	}
}

//SetCount Set number of tickets
func (rt *ResultTicket) SetCount(c float64) {
	rt.count = c
}

//GetCount Get number of tickets
func (rt *ResultTicket) GetCount() float64 {
	return rt.count
}

//SetGlobals Set number of ticket by all filters
func (rt *ResultTicket) SetGlobals(t *[]Global) {
	rt.global = *t
}

//GetGlobals Get number of ticket by all filters
func (rt *ResultTicket) GetGlobals() *[]Global {
	return &rt.global
}

//SetPriority Set number of ticket by priority
func (rt *ResultTicket) SetPriority(m map[string]float64) {
	rt.priority = m
}

//GetPriority Get number of ticket by status
func (rt *ResultTicket) GetPriority() map[string]float64 {
	return rt.priority
}

//SetStatus Set number of ticket by status
func (rt *ResultTicket) SetStatus(m map[string]float64) {
	rt.status = m
}

//GetStatus Get number of ticket by status
func (rt *ResultTicket) GetStatus() map[string]float64 {
	return rt.status
}

//SetVia Set number of tickets by source type
func (rt *ResultTicket) SetVia(v map[string]float64) {
	rt.via = v
}

//GetVia Get number of tickets by source type
func (rt *ResultTicket) GetVia() map[string]float64 {
	return rt.via
}
