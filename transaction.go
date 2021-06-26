package thepeer

import "time"

type Transaction struct {
	Message     string `json:"message"`
	Event       string `json:"event"`
	Transaction struct {
		ID     string `json:"id"`
		Remark string `json:"remark"`
		Amount int64  `json:"amount"`
		Type   string `json:"type"`
		Status string `json:"status"`
		User   struct {
			Name           string `json:"name"`
			Identifier     string `json:"identifier"`
			IdentifierType string `json:"identifier_type"`
			Email          string `json:"email"`
			Reference      string `json:"reference"`
		} `json:"user"`
		Mode      string `json:"mode"`
		Reference string `json:"reference"`
		Peer      struct {
			Business struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
				Logo  string `json:"logo"`
			} `json:"business"`
			User struct {
				Name           string `json:"name"`
				Identifier     string `json:"identifier"`
				IdentifierType string `json:"identifier_type"`
				Email          string `json:"email"`
				Reference      string `json:"reference"`
			} `json:"user"`
		} `json:"peer"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"transaction"`
}

type IndexUserOptions struct {
	FullName   string `json:"full_name"`
	Identifier string `json:"identifier"`
	Email      string `json:"email"`
}

type UpdateUserOptions struct {
	Identifier string `json:"identifier"`
}

type DeIndexUserOptions struct {
	UserReference string `json:"user_reference"`
}

type indexedUserResponse struct {
	IndexedUser IndexedUser `json:"indexed_user"`
}

type IndexedUser struct {
	Name           string `json:"name"`
	Identifier     string `json:"identifier"`
	IdentifierType string `json:"identifier_type"`
	Email          string `json:"email"`
	Reference      string `json:"reference"`
}

type Receipt struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
	User   struct {
		Name           string `json:"name"`
		Identifier     string `json:"identifier"`
		IdentifierType string `json:"identifier_type"`
		Email          string `json:"email"`
		Reference      string `json:"reference"`
	} `json:"user"`
	Peer struct {
		User struct {
			Name           string `json:"name"`
			Identifier     string `json:"identifier"`
			IdentifierType string `json:"identifier_type"`
			Email          string `json:"email"`
			Reference      string `json:"reference"`
		} `json:"user"`
		Business struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			Email          string `json:"email"`
			Logo           string `json:"logo"`
			IdentifierType string `json:"identifier_type"`
		} `json:"business"`
	} `json:"peer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type receiptResponse struct {
	Receipt Receipt `json:"receipt"`
}
