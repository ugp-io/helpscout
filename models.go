package helpscout

var accessCode string
var conversationsURL string = "https://api.helpscout.net/v2/conversations"
var customersURL string = "https://api.helpscout.net/v2/customers"
var tagsURL string = "https://api.helpscout.net/v2/tags"

type HelpScoutConversationRequest struct {
	URL       *string
	Emails    *[]string
	Mailboxes *[]string
	Tags      *[]string
	Folder    *string
	Status    *string
}

type HelpScoutTagUpdate struct {
	ConversationID string
	Tags           []string
}

type HelpScoutConversationsResponse struct {
	Embedded struct {
		Conversations []struct {
			Embedded struct {
				Threads *[]interface{} `json:"threads,omitempty"`
			} `json:"_embedded,omitempty"`
			Links struct {
				ClosedBy struct {
					Href *string `json:"href,omitempty"`
				} `json:"closedBy,omitempty"`
				CreatedByCustomer struct {
					Href *string `json:"href,omitempty"`
				} `json:"createdByCustomer,omitempty"`
				Mailbox struct {
					Href *string `json:"href,omitempty"`
				} `json:"mailbox,omitempty"`
				PrimaryCustomer struct {
					Href *string `json:"href,omitempty"`
				} `json:"primaryCustomer,omitempty"`
				Self struct {
					Href *string `json:"href,omitempty"`
				} `json:"self,omitempty"`
				Threads struct {
					Href *string `json:"href,omitempty"`
				} `json:"threads,omitempty"`
				Web struct {
					Href *string `json:"href,omitempty"`
				} `json:"web,omitempty"`
			} `json:"_links,omitempty"`
			Bcc          *[]interface{} `json:"bcc,omitempty"`
			Cc           *[]interface{} `json:"cc,omitempty"`
			ClosedBy     *int           `json:"closedBy,omitempty"`
			ClosedByUser struct {
				Email *string `json:"email,omitempty"`
				First *string `json:"first,omitempty"`
				ID    *int    `json:"id,omitempty"`
				Last  *string `json:"last,omitempty"`
				Type  *string `json:"type,omitempty"`
			} `json:"closedByUser,omitempty"`
			CreatedAt *string `json:"createdAt,omitempty"`
			CreatedBy struct {
				Email    *string `json:"email,omitempty"`
				First    *string `json:"first,omitempty"`
				ID       *int    `json:"id,omitempty"`
				Last     *string `json:"last,omitempty"`
				PhotoURL *string `json:"photoUrl,omitempty"`
				Type     *string `json:"type,omitempty"`
			} `json:"createdBy,omitempty"`
			CustomFields         *[]interface{} `json:"customFields,omitempty"`
			CustomerWaitingSince struct {
				Friendly *string `json:"friendly,omitempty"`
				Time     *string `json:"time,omitempty"`
			} `json:"customerWaitingSince,omitempty"`
			FolderID        *int    `json:"folderId,omitempty"`
			ID              *int    `json:"id,omitempty"`
			MailboxID       *int    `json:"mailboxId,omitempty"`
			Number          *int    `json:"number,omitempty"`
			Preview         *string `json:"preview,omitempty"`
			PrimaryCustomer struct {
				Email    *string `json:"email,omitempty"`
				First    *string `json:"first,omitempty"`
				ID       *int    `json:"id,omitempty"`
				Last     *string `json:"last,omitempty"`
				PhotoURL *string `json:"photoUrl,omitempty"`
				Type     *string `json:"type,omitempty"`
			} `json:"primaryCustomer,omitempty"`
			Source struct {
				Type *string `json:"type,omitempty"`
				Via  *string `json:"via,omitempty"`
			} `json:"source,omitempty"`
			State         *string                   `json:"state,omitempty"`
			Status        *string                   `json:"status,omitempty"`
			Subject       *string                   `json:"subject,omitempty"`
			Tags          *[]map[string]interface{} `json:"tags,omitempty"`
			Threads       *int                      `json:"threads,omitempty"`
			Type          *string                   `json:"type,omitempty"`
			UserUpdatedAt *string                   `json:"userUpdatedAt,omitempty"`
		} `json:"conversations,omitempty"`
	} `json:"_embedded,omitempty"`
	Links struct {
		Next struct {
			Href *string `json:"href,omitempty"`
		} `json:"next,omitempty"`
		First struct {
			Href *string `json:"href,omitempty"`
		} `json:"first,omitempty"`
		Last struct {
			Href *string `json:"href,omitempty"`
		} `json:"last,omitempty"`
		Page struct {
			Href      *string `json:"href,omitempty"`
			Templated *bool   `json:"templated,omitempty"`
		} `json:"page,omitempty"`
		Self struct {
			Href *string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Page struct {
		Number        *int `json:"number,omitempty"`
		Size          *int `json:"size,omitempty"`
		TotalElements *int `json:"totalElements,omitempty"`
		TotalPages    *int `json:"totalPages,omitempty"`
	} `json:"page,omitempty"`
}

type HelpScoutThreadsResponse struct {
	Embedded struct {
		Threads []struct {
			Embedded struct {
				Attachments []interface{} `json:"attachments,omitempty"`
			} `json:"_embedded,omitempty"`
			Links struct {
				AssignedTo struct {
					Href *string `json:"href,omitempty"`
				} `json:"assignedTo,omitempty"`
				CreatedByCustomer struct {
					Href *string `json:"href,omitempty"`
				} `json:"createdByCustomer,omitempty"`
				Customer struct {
					Href *string `json:"href,omitempty"`
				} `json:"customer,omitempty"`
			} `json:"_links,omitempty"`
			Action struct {
				AssociatedEntities struct{} `json:"associatedEntities,omitempty"`
				Type               *string  `json:"type,omitempty"`
			} `json:"action,omitempty"`
			AssignedTo struct {
				Email *string `json:"email,omitempty"`
				First *string `json:"first,omitempty"`
				ID    *int    `json:"id,omitempty"`
				Last  *string `json:"last,omitempty"`
			} `json:"assignedTo,omitempty"`
			Bcc       *[]interface{} `json:"bcc,omitempty"`
			Body      *string        `json:"body,omitempty"`
			Cc        *[]interface{} `json:"cc,omitempty"`
			CreatedAt *string        `json:"createdAt,omitempty"`
			CreatedBy struct {
				Email    *string `json:"email,omitempty"`
				First    *string `json:"first,omitempty"`
				ID       *int    `json:"id,omitempty"`
				Last     *string `json:"last,omitempty"`
				PhotoURL *string `json:"photoUrl,omitempty"`
				Type     *string `json:"type,omitempty"`
			} `json:"createdBy,omitempty"`
			Customer struct {
				Email    *string `json:"email,omitempty"`
				First    *string `json:"first,omitempty"`
				ID       *int    `json:"id,omitempty"`
				Last     *string `json:"last,omitempty"`
				PhotoURL *string `json:"photoUrl,omitempty"`
			} `json:"customer,omitempty"`
			ID           *int `json:"id,omitempty"`
			SavedReplyID *int `json:"savedReplyId,omitempty"`
			Source       struct {
				Type *string `json:"type,omitempty"`
				Via  *string `json:"via,omitempty"`
			} `json:"source,omitempty"`
			State  *string   `json:"state,omitempty"`
			Status *string   `json:"status,omitempty"`
			To     []*string `json:"to,omitempty"`
			Type   *string   `json:"type,omitempty"`
		} `json:"threads,omitempty"`
	} `json:"_embedded,omitempty"`
	Links struct {
		First struct {
			Href *string `json:"href,omitempty"`
		} `json:"first,omitempty"`
		Last struct {
			Href *string `json:"href,omitempty"`
		} `json:"last,omitempty"`
		Page struct {
			Href      *string `json:"href,omitempty"`
			Templated bool    `json:"templated,omitempty"`
		} `json:"page,omitempty"`
		Self struct {
			Href *string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Page struct {
		Number        *int `json:"number,omitempty"`
		Size          *int `json:"size,omitempty"`
		TotalElements *int `json:"totalElements,omitempty"`
		TotalPages    *int `json:"totalPages,omitempty"`
	} `json:"page,omitempty"`
}

type HelpScoutTagsResponse struct {
	Embedded struct {
		Tags []struct {
			Color       *string `json:"color,omitempty"`
			CreatedAt   *string `json:"createdAt,omitempty"`
			ID          *int    `json:"id,omitempty"`
			Name        *string `json:"name,omitempty"`
			Slug        *string `json:"slug,omitempty"`
			TicketCount *int    `json:"ticketCount,omitempty"`
			UpdatedAt   *string `json:"updatedAt,omitempty"`
		} `json:"tags,omitempty"`
	} `json:"_embedded,omitempty"`
	Links struct {
		First struct {
			Href *string `json:"href,omitempty"`
		} `json:"first,omitempty"`
		Last struct {
			Href *string `json:"href,omitempty"`
		} `json:"last,omitempty"`
		Page struct {
			Href      *string `json:"href,omitempty"`
			Templated *bool   `json:"templated,omitempty"`
		} `json:"page,omitempty"`
		Self struct {
			Href *string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Page struct {
		Number        *int `json:"number,omitempty"`
		Size          *int `json:"size,omitempty"`
		TotalElements *int `json:"totalElements,omitempty"`
		TotalPages    *int `json:"totalPages,omitempty"`
	} `json:"page,omitempty"`
}
