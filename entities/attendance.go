package entities

type GetAllAttendancesByUserIdRequest struct {
	UserId int64  `json:"user_id"`
	Filter string `json:"filter"`
}

type CheckInRequest struct {
	UserId     int64 `json:"user_id"`
	LocationId int64 `json:"location_id"`
}

type CheckOutRequest struct {
	UserId     int64 `json:"user_id"`
	LocationId int64 `json:"location_id"`
}
