package zendesk

import (
	"Zendesk-Exporter/src/config"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type via struct {
	Channel string      `json:"channel"`
	Source  interface{} `json:"source"`
}

type customField struct {
	ID    int64 `json:"id,omitempty"`
	Name  string
	Value string `json:"value,omitempty"`
}

type ticket struct {
	ID                  int64         `json:"id,omitempty"`
	URL                 string        `json:"url,omitempty"`
	ExternalID          string        `json:"external_id,omitempty"`
	Type                string        `json:"type,omitempty"`
	Subject             string        `json:"subject,omitempty"`
	RawSubject          string        `json:"raw_subject,omitempty"`
	Description         string        `json:"description,omitempty"`
	Priority            string        `json:"priority,omitempty"`
	Status              string        `json:"status,omitempty"`
	Recipient           string        `json:"recipient,omitempty"`
	RequesterID         int64         `json:"requester_id"`
	SubmitterID         int64         `json:"submitter_id,omitempty"`
	AssigneeID          int64         `json:"assignee_id,omitempty"`
	OrganizationID      int64         `json:"organization_id,omitempty"`
	GroupID             int64         `json:"group_id,omitempty"`
	CollaboratorIDs     []int64       `json:"collaborator_ids,omitempty"`
	Collaborators       []string      `json:"collaborators,omitempty"`
	FollowerIDs         []int64       `json:"follower_ids,omitempty"`
	ForumTopicID        int64         `json:"forum_topic_id,omitempty"`
	ProblemID           int64         `json:"problem_id,omitempty"`
	HasIncidents        bool          `json:"has_incidents,omitempty"`
	DueAt               *time.Time    `json:"due_at,omitempty"`
	Tags                []string      `json:"tags,omitempty"`
	Via                 via           `json:"via,omitempty"`
	CustomFields        []customField `json:"custom_fields,omitempty"`
	SatisfactionRating  interface{}   `json:"satisfaction_rating,omitempty"`
	SharingAgreementIDs []int64       `json:"sharing_agreement_ids,omitempty"`
	FollowupIDs         []int64       `json:"followup_ids,omitempty"`
	ViaFollowupSourceID int64         `json:"via_followup_source_id,omitempty"`
	MacroIds            []int64       `json:"macro_ids,omitempty"`
	TicketFormID        int64         `json:"ticket_form_id,omitempty"`
	BrandID             int64         `json:"brand_id,omitempty"`
	AllowChannelback    bool          `json:"allow_channelback,omitempty"`
	AllowAttachements   bool          `json:"allow_attachments,omitempty"`
	IsPublic            bool          `json:"is_public,omitempty"`
	CreatedAt           *time.Time    `json:"created_at,omitempty"`
	UpdatedAt           *time.Time    `json:"updated_at,omitempty"`
}

//Tickets Object get by api
type Tickets struct {
	List     []ticket `json:"tickets"`
	Next     string   `json:"next_page"`
	Previous string   `json:"previous_page"`
	Count    int64    `json:"count"`
}

func (c *Client) getTickets(allTicketField []TicketField) ([]ticket, error) {
	var tickets Tickets
	var list []ticket

	i := 1
	for {
		body, err := c.Get("/tickets.json?page=" + strconv.Itoa(i))
		if err != nil {
			return []ticket{}, err
		}

		err = json.Unmarshal(body, &tickets)
		if err != nil {
			return []ticket{}, err
		}

		for _, t := range tickets.List {
			cfs := t.CustomFields
			var cfss []customField
			for _, cf := range cfs {
				tf, err := getTicketFieldByID(cf.ID, allTicketField)
				if err != nil {
					return []ticket{}, err
				}
				cf.Name = tf.Title
				if cf.Value == "" {
					cf.Value = "undefined"
				}
				cfss = append(cfss, cf)
			}
			t.CustomFields = cfss
			list = append(list, t)
		}

		if tickets.Next == "" {
			break
		}

		tickets = Tickets{
			List:     []ticket{},
			Next:     "",
			Previous: "",
			Count:    0,
		}
		i++
	}
	return list, nil
}

//GetTicketStats Return statistics of all tickets in a map
func (c *Client) GetTicketStats(allTicketField []TicketField, cfs config.CustomFields) (*ResultTicket, error) {
	getEmptyGlobalWithCustomField(cfs.Fields)
	return nil, nil

	rt := NewResultTicket()

	list, err := c.getTickets(allTicketField)
	if err != nil {
		return nil, err
	}

	rt.SetCount(float64(len(list)))

	global := rt.GetGlobals()
	priority := rt.GetPriority()
	status := rt.GetStatus()
	via := rt.GetVia()

	for _, t := range list {
		if t.Priority == "" {
			t.Priority = "undefined"
		}
		//Set Global
		m := map[string]string{
			"priority": t.Priority,
			"status":   t.Status,
			"via":      t.Via.Channel,
		}
		if cfs.Enable {
			for cf := range cfs.Fields {
				for _, tmp := range t.CustomFields {
					if tmp.Name == cf {
						m[cf] = tmp.Value
					}
				}
			}
		}
		g, k, err := GetGlobal(global, m)
		if err != nil {
			return nil, err
		}
		g.Count++
		(*global)[k] = *g

		//Set Priority
		if _, ok := priority[t.Priority]; ok {
			priority[t.Priority]++
		} else if t.Priority == "" {
			priority["undefined"]++
		} else {
			return nil, errors.New("Error: " + t.Priority + " priority is not know")
		}

		//Set Status
		if _, ok := status[t.Status]; ok {
			status[t.Status]++
		} else {
			return nil, errors.New("Error: " + t.Status + " status is not know")
		}

		//Set Via
		if _, ok := via[t.Via.Channel]; ok {
			via[t.Via.Channel]++
		} else {
			return nil, errors.New("Error: " + t.Via.Channel + " channel is not know")
		}
	}

	return rt, nil
}
