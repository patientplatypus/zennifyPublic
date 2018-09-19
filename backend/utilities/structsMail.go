package util

type Mail struct{
	Header string
	Subject string
	ID string
}

type ReturnMessages struct{
	SentMail []Mail
	ReceivedMail []Mail
}

type ReturnMsg struct {
	Subject string
	Content string
	ID string
}