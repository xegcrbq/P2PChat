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

type TalkMeWebHook struct {
	Data TalkMeWebHookData `json:"data"`
}
type TalkMeWebHookData struct {
	Message TalkMeMessage `json:"message"`
	Client  TalkMeClient  `json:"client"`
}
type TalkMeClient struct {
	ClientId string `json:"clientId"`
}
