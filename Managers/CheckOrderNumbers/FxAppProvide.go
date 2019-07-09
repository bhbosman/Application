package CheckOrderNumbers

//func FxAppProvide() fx.Option {
//	type ReturnType struct {
//		fx.Out
//		FullMarketDepthManager Managers.IMitchDataProcessor `name:"CheckOrderNumbers"`
//	}
//	return fx.Provide(
//		func(logger *log.Logger, publisher PubSub.IPublisher) (ReturnType, error) {
//			nextHandler, err := NewCheckOrderNumbers(logger)
//			if err != nil {
//				return ReturnType{}, err
//			}
//			handler := Managers.NewMitchDataProcessorChannelWrapper(nextHandler, publisher)
//			return ReturnType{
//				FullMarketDepthManager: handler,
//			}, nil
//		})
//}
