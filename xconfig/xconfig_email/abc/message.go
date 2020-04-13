package abc

// IMessage mail
type IMessage interface {
	Send(tos []string, templateTitle, templateContent string, templateParams []string) error
}
