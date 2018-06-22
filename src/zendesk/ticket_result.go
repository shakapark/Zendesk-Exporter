package zendesk

//ResultTicket Stock result
type ResultTicket struct {
	count  float64
	status map[string]float64
	via    map[string]float64
}

//NewResultTicket Create new ResultTicket
func NewResultTicket() *ResultTicket {
	emptyStatus := map[string]float64{
		"new":     0,
		"open":    0,
		"pending": 0,
		"hold":    0,
		"solved":  0,
		"closed":  0,
	}
	emptyVia := map[string]float64{
		"web":     0,
		"mobile":  0,
		"rule":    0,
		"system":  0,
		"twitter": 0,
	}
	return &ResultTicket{
		count:  0,
		status: emptyStatus,
		via:    emptyVia,
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
