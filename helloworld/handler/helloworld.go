package handler

import (
	"context"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/context/metadata"

	"github.com/micro/micro/v3/service/logger"
	helloworld "github.com/micro/services/helloworld/proto"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	logger.Info("Received Helloworld.Call request")
	acc,ok:= auth.AccountFromContext(ctx)
	if ok {
		logger.Infof("Account is %+v", acc)
	}
	callTok, ok := metadata.Get(ctx, "CallerToken")
	if ok {
		acc,err:= auth.Inspect(callTok)
		if err!=nil {
			logger.Errorf("Error inspecting token %s", err)
		} else {
			logger.Infof("Caller account is %+v", *acc)
		}

	} else {
		logger.Errorf("Can't find caller token")
	}
	rsp.Msg = "Hello " + req.Name
	return nil
}
