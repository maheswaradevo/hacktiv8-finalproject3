package service

type PingService struct {
}

func (p PingService) Ping() string {
	return "pong"
}

func NewPingService() PingService {
	return PingService{}
}
