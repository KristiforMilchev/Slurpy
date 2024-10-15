package implementations

import "Slurpy/interfaces"

type Locator struct {
	DeploymentService interfaces.DeploymentService
	EncoderService    interfaces.EncoderService
	RpcService        interfaces.RpcService
	Storage           interfaces.Storage
	WalletService     interfaces.WalletService
	NetworkService    interfaces.NetworkService
}
