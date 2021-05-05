package structs

type RoomInfo struct {
	RoomNo   int       `json:"roomNo"`
	Users    [2]string `json:"users"`
	Contents []struct {
		User    string `json:"user"`
		Message string `json:"message"`
		Date    int    `json:"date"`
	} `json:"contents"`
}
