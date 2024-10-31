package handlerUser

import "goWallet/internal/ports"

type Handler struct {
	User ports.UserSrv
	Account ports.AccountSrv
}
