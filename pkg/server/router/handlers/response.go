package handlers

type Response struct {
	ID      string `json:"id"`      // ID of the request
	Message string `json:"message"` // The response generated
	Status  int    `json:"status"`  // Status code for HTTP response
}

func (rp *Response) SetStatus(status int) {
	rp.Status = status
}

func (rp *Response) SetMessage(response string) {
	rp.Message = response
}
