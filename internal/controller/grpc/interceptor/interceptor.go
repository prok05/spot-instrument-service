package interceptor

import "github.com/prok05/spot-instrument-service/pkg/logger"

type Interceptor struct {
	L logger.Interface
}

func New(l logger.Interface) *Interceptor {
	return &Interceptor{L: l}
}
