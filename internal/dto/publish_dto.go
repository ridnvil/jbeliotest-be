package dto

type PublishDto struct {
	Key        string `json:"key"`
	MasterData bool   `json:"master_data"`
	ClientID   string `json:"client_id"`
}
