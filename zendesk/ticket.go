package zendesk

import "time"

type ticket struct {
	ID                  int64       `json:"id,omitempty"`
	URL                 string      `json:"url,omitempty"`
	ExternalID          string      `json:"external_id,omitempty"`
	Type                string      `json:"type,omitempty"`
	Subject             string      `json:"subject,omitempty"`
	RawSubject          string      `json:"raw_subject,omitempty"`
	Description         string      `json:"description,omitempty"`
	Priority            string      `json:"priority,omitempty`
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
	Via                 interface{} `json:"via,omitempty"`
	CustomFields        []string    `json:"custom_fileds,omitempty"`
	SatisfactionRating  interface{} `json:"satisfaction_rating,omitempty"`
	SharingAgreementIDs []int64     `json:"sharing_agreement_ids",omitempty"`
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

func getTickets() []ticket {
	//TODO
	return
}
