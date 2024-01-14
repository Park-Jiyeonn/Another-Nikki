package consumer

type Consumer interface {
	Start()
	Close()
}
