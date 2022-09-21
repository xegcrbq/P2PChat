package models

type Dialogue struct {
	User1    string
	User2    string
	Messages []*Message
}

func (d *Dialogue) String() string {
	result := ""
	for _, message := range d.Messages {
		result += message.String() + "\n"
	}
	return result
}
