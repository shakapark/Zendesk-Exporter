package zendesk

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type via struct {
	Channel string      `json:"channel"`
	Source  interface{} `json:"source"`
}

type ticket struct {
	ID                  int64       `json:"id,omitempty"`
	URL                 string      `json:"url,omitempty"`
	ExternalID          string      `json:"external_id,omitempty"`
	Type                string      `json:"type,omitempty"`
	Subject             string      `json:"subject,omitempty"`
	RawSubject          string      `json:"raw_subject,omitempty"`
	Description         string      `json:"description,omitempty"`
	Priority            string      `json:"priority,omitempty"`
	Status              string      `json:"status,omitempty"`
	Recipient           string      `json:"recipient,omitempty"`
	RequesterID         int64       `json:"requester_id"`
	SubmitterID         int64       `json:"submitter_id,omitempty"`
	AssigneeID          int64       `json:"assignee_id,omitempty"`
	OrganizationID      int64       `json:"organization_id,omitempty"`
	GroupID             int64       `json:"group_id,omitempty"`
	CollaboratorIDs     []int64     `json:"collaborator_ids,omitempty"`
	Collaborators       []string    `json:"collaborators,omitempty"`
	FollowerIDs         []int64     `json:"follower_ids,omitempty"`
	ForumTopicID        int64       `json:"forum_topic_id,omitempty"`
	ProblemID           int64       `json:"problem_id,omitempty"`
	HasIncidents        bool        `json:"has_incidents,omitempty"`
	DueAt               *time.Time  `json:"due_at,omitempty"`
	Tags                []string    `json:"tags,omitempty"`
	Via                 via         `json:"via,omitempty"`
	CustomFields        []string    `json:"custom_fileds,omitempty"`
	SatisfactionRating  interface{} `json:"satisfaction_rating,omitempty"`
	SharingAgreementIDs []int64     `json:"sharing_agreement_ids,omitempty"`
	FollowupIDs         []int64     `json:"followup_ids,omitempty"`
	ViaFollowupSourceID int64       `json:"via_followup_source_id,omitempty"`
	MacroIds            []int64     `json:"macro_ids,omitempty"`
	TicketFormID        int64       `json:"ticket_form_id,omitempty"`
	BrandID             int64       `json:"brand_id,omitempty"`
	AllowChannelback    bool        `json:"allow_channelback,omitempty"`
	AllowAttachements   bool        `json:"allow_attachments,omitempty"`
	IsPublic            bool        `json:"is_public,omitempty"`
	CreatedAt           *time.Time  `json:"created_at,omitempty"`
	UpdatedAt           *time.Time  `json:"updated_at,omitempty"`
}

//Tickets Object get by api
type Tickets struct {
	List     []ticket `json:"tickets"`
	Next     string   `json:"next_page"`
	Previous string   `json:"previous_page"`
	Count    int64    `json:"count"`
}

func (c *Client) getTickets() ([]ticket, error) {
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
func (c *Client) GetTicketStats() (*ResultTicket, error) {
	rt := NewResultTicket()

	list, err := c.getTickets()
	if err != nil {
		return nil, err
	}

	rt.SetCount(float64(len(list)))

	status := rt.GetStatus()
	fmt.Println("Debug: ", status)
	for _, t := range list {
		if t.Priority == "" {
			status[t.Status]["undefined"]++
			fmt.Println("Debug: ", status)
		} else {
			status[t.Status][t.Priority]++
			fmt.Println("Debug: ", status)
		}
	}
	rt.SetStatus(status)

	via := rt.GetVia()
	for _, t := range list {
		if _, ok := via[t.Via.Channel]; ok {
			via[t.Via.Channel]++
		} else {
			fmt.Println("Error:", t.Via.Channel, "channel is not know")
		}
	}
	rt.SetVia(via)

	return rt, nil
}
