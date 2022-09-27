package models

type TalkMeMessageGetListAnswer struct {
	Success bool                         `json:"success"`
	Result  []TalkMeMessageGetListResult `json:"result"`
}
type TalkMeMessage struct {

	//Operator string `json:"operator"`
	Id          int32  `json:"id"`
	WhoSend     string `json:"whoSend"`
	Text        string `json:"text"`
	DateTime    string `json:"dateTimeUTC"`
	MessageType string `json:"messageType"`
}
type TalkMeMessageGetListResult struct {
	Messages []TalkMeMessage `json:"messages"`
	ClientId string          `json:"clientId"`
}
