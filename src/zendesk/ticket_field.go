package zendesk

import (
	"encoding/json"
	"errors"
	"time"
)

type ticketFieldSystemFieldOption struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Position int64  `json:"position"`
	RawName  string `json:"raw_name"`
	URL      string `json:"url"`
	Value    string `json:"value"`
}

type ticketFieldCustomFieldOption struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name"`
	Position int64  `json:"position,omitempty"`
	RawName  string `json:"raw_name,omitempty"`
	URL      string `json:"url,omitempty"`
	Value    string `json:"value"`
}

//TicketField Contains result of request ticket_field api
type TicketField struct {
	ID                  int64                          `json:"id,omitempty"`
	URL                 string                         `json:"url,omitempty"`
	Type                string                         `json:"type"`
	Title               string                         `json:"title"`
	RawTitle            string                         `json:"raw_title,omitempty"`
	Description         string                         `json:"description,omitempty"`
	RawDescription      string                         `json:"raw_description,omitempty"`
	Position            int64                          `json:"position,omitempty"`
	Active              bool                           `json:"active,omitempty"`
	Required            bool                           `json:"required,omitempty"`
	CollapsedForAgents  bool                           `json:"collapsed_for_agents,omitempty"`
	RegexpForValidation string                         `json:"regexp_for_validation,omitempty"`
	TitleInPortal       string                         `json:"title_in_portal,omitempty"`
	RawTitleInPortal    string                         `json:"raw_title_in_portal,omitempty"`
	VisibleInPortal     bool                           `json:"visible_in_portal,omitempty"`
	EditableInPortal    bool                           `json:"editable_in_portal,omitempty"`
	RequiredInPortal    bool                           `json:"required_in_portal,omitempty"`
	Tag                 string                         `json:"tag,omitempty"`
	CreatedAt           *time.Time                     `json:"created_at,omitempty"`
	UpdatedAt           *time.Time                     `json:"updated_at,omitempty"`
	SystemFieldOptions  []ticketFieldSystemFieldOption `json:"system_field_options,omitempty"`
	CustomFieldOptions  []ticketFieldCustomFieldOption `json:"custom_field_options,omitempty"`
	SubTypeID           int64                          `json:"sub_type_id,omitempty"`
	Removable           bool                           `json:"removable,omitempty"`
	AgentDescription    string                         `json:"agent_description,omitempty"`
}

type resultTicketField struct {
	TicketFields []TicketField `json:"ticket_fields"`
}

func getTicketFieldByID(id int64, atf []TicketField) (TicketField, error) {
	for _, tf := range atf {
		if tf.ID == id {
			return tf, nil
		}
	}
	return TicketField{}, errors.New("Unknown ticketField id")
}

//SetAllTicketField Set all ticketField of zendesk
func (c *Client) SetAllTicketField() ([]TicketField, error) {
	var tmp resultTicketField

	body, err := c.Get("/ticket_fields.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &tmp)
	if err != nil {
		return nil, err
	}

	return tmp.TicketFields, nil
}
