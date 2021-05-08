package structs

type RoomInfo struct {
	RoomNo   int            `json:"roomNo"`
	Users    [2]string      `json:"users"`
	Contents []RoomContents `json:"contents"`
}

type RoomContents struct {
	User    string `json:"user"`
	Message string `json:"message"`
	Date    int    `json:"date"`
}
