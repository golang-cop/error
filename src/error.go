package Error

type Interface interface {
	Message() string
	IsNull() bool
}

type data struct {
	message string
}

func New(message string) Interface {
	return &data{
		message: message,
	}
}

func (d data) Message() string {
	return d.message
}

func (d data) IsNull() bool {
	return false
}
