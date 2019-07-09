package MissingSequenceManager

//func FxAppInvoke() fx.Option {
//	type InputType struct {
//		fx.In
//		FullMarketDepthManager Managers.IMitchDataProcessor `name:"MissingSequenceManager"`
//	}
//	return fx.Invoke(
//		func(lifecycle fx.Lifecycle, logger *log.Logger, input InputType) error {
//			lifecycle.Append(fx.Hook{
//				OnStart: nil,
//				OnStop: func(stopContext context.Context) error {
//					logger.Println("CheckOrderNumbers")
//					return input.FullMarketDepthManager.Close()
//				},
//			})
//			return nil
//		})
//}
