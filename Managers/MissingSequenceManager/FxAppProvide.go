package MissingSequenceManager

//func FxAppProvide() fx.Option {
//	type ReturnType struct {
//		fx.Out
//		FullMarketDepthManager Managers.IMitchDataProcessor `name:"MissingSequenceManager"`
//	}
//	return fx.Provide(
//		func(publisher PubSub.IPublisher, logger *log.Logger) (ReturnType, error) {
//			nextHandler, err := NewManager(logger, publisher)
//			if err != nil {
//				return ReturnType{}, err
//			}
//			handler:= Managers.NewMitchDataProcessorChannelWrapper(nextHandler, publisher)
//			return ReturnType{
//				FullMarketDepthManager: handler,
//			}, nil
//		})
//}
