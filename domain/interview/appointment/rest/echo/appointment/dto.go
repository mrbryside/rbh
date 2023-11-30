package appointment

type CreatePayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdatePayload struct {
	Name        *string `json:"name" validate:"required"`
	Description *string `json:"description" validate:"required"`
	Status      *string `json:"status" validate:"required"`
	Enabled     *bool   `json:"enabled" validate:"required"`
}

type Paginate struct {
	Results []GetAllResp `json:"results"`
	Next    bool         `json:"next"`
}

type GetAllResp struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Creator     Creator `json:"creator"`
	CreatedAt   string  `json:"created_at"`
}

type Response struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Enabled     bool      `json:"enabled"`
	Creator     Creator   `json:"creator"`
	Comments    []Comment `json:"comments"`
	Histories   []History `json:"histories"`
	CreatedAt   string    `json:"created_at"`
}

type UpdateResponse struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Enabled     bool    `json:"enabled"`
	Creator     Creator `json:"creator"`
	CreatedAt   string  `json:"created_at"`
}

type Comment struct {
	Id        uint    `json:"id"`
	Message   string  `json:"message"`
	Creator   Creator `json:"creator"`
	CreatedAt string  `json:"created_at"`
}

type History struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Creator struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
