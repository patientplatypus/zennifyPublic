package util

type Mail struct{
	Header string
	Subject string
}

type ReturnMessages struct{
	SentMail []Mail
	ReceivedMail []Mail
}