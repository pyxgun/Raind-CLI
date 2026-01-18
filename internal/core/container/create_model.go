package container

type ServiceCreateModel struct {
	Image   string
	Command []string
	Network string
	Volume  []string
	Publish []string
	Tty     bool
}

type CreateRequestModel struct {
	Image   string   `json:"image,omitempty"`
	Command []string `json:"command,omitempty"`
	Network string   `json:"network,omitempty"`
	Volume  []string `json:"mount,omitempty"`
	Publish []string `json:"port,omitempty"`
	Tty     bool     `json:"tty,omitempty"`
}

type CreateResponseDataModel struct {
	Id string `json:"id"`
}

type CreateResponseModel struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    CreateResponseDataModel `json:"data,omitempty"`
}
