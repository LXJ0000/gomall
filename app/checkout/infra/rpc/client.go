package rpc

// var (
// 	UserClient userservice.Client
// 	once       sync.Once
// 	err        error
// )

// func Init() {
// 	once.Do(func() {
// 		initUserClient()
// 	})
// }

// func initUserClient() {
// 	opts := []client.Option{
// 		client.WithSuite(&clientsuite.CommonClientSuite{
// 			CurrentServiceName: conf.GetConf().Kitex.Service,
// 			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
// 		}),
// 	}
// 	UserClient, err = userservice.NewClient("user", opts...)
// 	utils.MustHandleError(err)
// }
