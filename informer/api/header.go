package api

const(
	queue = "Informer"
)

type Header struct {
	Queue string
}

func GetServiceName() string {
	return queue
}
