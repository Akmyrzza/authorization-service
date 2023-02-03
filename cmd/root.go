package cmd

import (
	"github.com/Akmyrzza/authorization-service/internal/adapters/server/https"
	"github.com/Akmyrzza/authorization-service/internal/domain/core"
	"github.com/Akmyrzza/authorization-service/internal/domain/usecases"
)

func Execute() {
	var err error

	app := struct {
		core       *core.St
		ucs        *usecases.St
		restApiSrv *https.St
		//lg         *dopLoggerZap.St
		//cache      dopCache.Cache
		//db         *dopDbPg.St
		//repo       *pg.St
		//jwts       dopJwt.Jwt
		//sms        dopSms.Sms
	}{}

	confLoad()
}
