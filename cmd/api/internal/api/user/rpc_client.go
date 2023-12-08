package user

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap"
	"tiktok-api/internal/api"
	"tiktok-api/internal/global"
	"tiktok-api/internal/model"
	"tiktok/kitex_gen/user/userservice"
)

// InitRpcUserClient initializes the RPC user client
func InitRpcUserClient() {
	// Create a new etcd resolver
	r, err := etcd.NewEtcdResolver(global.Settings.Etcd.Addr)
	if err != nil {
		panic(err)
	}

	// Create a new OpenTelemetry provider
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(global.Settings.Jaeger.ServerName[model.TiktokUser]),
		provider.WithExportEndpoint(global.Settings.Jaeger.RPCExportEndpoint),
		provider.WithInsecure(),
	)

	// Create a new RPC user client
	c, err := userservice.NewClient(
		model.RpcUser,
		client.WithSuite(tracing.NewClientSuite()),
		//client.WithHostPorts(global.Settings.Rpc.ServerAddrs[model.RpcUser]),
		//client.WithMiddleware(rpcmiddleware.CommonMiddleware),
		//client.WithInstanceMW(rpcmiddleware.ClientMiddleware),
		client.WithResolver(r),
	)
	if err != nil {
		zap.L().Error("apiServer InitRpc User Client err:", zap.Error(err))
		panic(err)
	}
	// Set the user client
	api.UserClient = c
}
