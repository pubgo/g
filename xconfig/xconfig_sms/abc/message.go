package abc

type IMessage interface {
	Send(tos []string, templateTitle, templateContent string, templateParams []string) error
}
