package room

import "time"

// GetRoomRequest ...
type GetRoomRequest struct {
	RoomID  string `param:"roomID"`
	EventID string `param:"eventID"`
}

// GetRoomResponse ...
type GetRoomResponse struct {
	RequestTime time.Time
}
