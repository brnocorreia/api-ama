package api

import "github.com/brnocorreia/api-ama/internal/store/pgstore"

const (
	MessageKindMessageCreated          = "message_created"
	MessageKindMessageRactionIncreased = "message_reaction_increased"
	MessageKindMessageRactionDecreased = "message_reaction_decreased"
	MessageKindMessageAnswered         = "message_answered"
)

type MessageMessageReactionIncreased struct {
	ID    string `json:"id"`
	Count int64  `json:"count"`
}

type MessageMessageReactionDecreased struct {
	ID    string `json:"id"`
	Count int64  `json:"count"`
}

type MessageMessageAnswered struct {
	ID string `json:"id"`
}

type MessageMessageCreated struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type Message struct {
	Kind   string `json:"kind"`
	Value  any    `json:"value"`
	RoomID string `json:"-"`
}

type ResponseID struct {
	ID string `json:"id"`
}

type ResponseRooms struct {
	Rooms []pgstore.Room `json:"rooms"`
}

type ResponseMessages struct {
	Messages []pgstore.Message `json:"messages"`
}

type ResponseCount struct {
	Count int64 `json:"count"`
}
