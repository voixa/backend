package handlers

type Request struct {
	ID        string `json:"id"`        // Unique identifier for a request
	AudioPath string `json:"audioPath"` // Path to uploaded audio file
}

func (rq *Request) SetID(id string) {
	rq.ID = id
}

func (rq *Request) SetAudioPath(audioPath string) {
	rq.AudioPath = audioPath
}
