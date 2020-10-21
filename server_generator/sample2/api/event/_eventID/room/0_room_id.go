package room

// GetRoomRequest ...
type GetRoomRequest struct {
	RoomID  string `param:"roomID"`
	EventID string `param:"eventID"`
}

// GetRoomResponse ...
type GetRoomResponse struct {
}
