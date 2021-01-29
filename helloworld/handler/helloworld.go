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
	md, ok:= metadata.FromContext(ctx)
	if !ok {
		logger.Errorf("No metadata found")
	} else {
		for k, v:=range md {
			logger.Infof("Key %s val %s", k, v)
		}
	}
	callTok, ok := metadata.Get(ctx, "Caller-Token")
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
