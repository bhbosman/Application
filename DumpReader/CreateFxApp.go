package main

import (
	"github.com/bhbosman/Application/Managers/TimeService"
	"github.com/bhbosman/Application/PubSub"
	"github.com/bhbosman/Application/UniqueNumber"
	"go.uber.org/fx"
	"io"
	"log"
	"path"
)

type ICurrentOpenFile io.Reader

func CreateFxApplication(applicationLogger *log.Logger) (*fx.App, IApplicationContext, error) {
	var applicationContext IApplicationContext

	return fx.New(
		fx.StartTimeout(fx.DefaultTimeout),
		fx.StopTimeout(fx.DefaultTimeout),
		UniqueNumber.FxAppProvide(),
		UniqueNumber.FxAppInvoke(),
		PubSub.FxAppProvide(),
		PubSub.FxAppInvoke(),
		TimeService.FxAppProvide(),
		//MissingSequenceManager.FxAppProvide(),
		//MissingSequenceManager.FxAppInvoke(),
		//CheckOrderNumbers.FxAppProvide(),
		//CheckOrderNumbers.FxAppInvoke(),
		TimeService.FxAppInvoke(),

		FxAppProvideMissingSequencesManager(),
		FxAppProvideApplicationContext(),
		FxAppProvideApplicationLogger(applicationLogger),
		FxAppProvideFxAppOverrideLogger(applicationLogger),
		FxAppProviderUserHomeDir(""),
		FxAppProvidePlayBackFileFromUserFolder(path.Join("Data", "MitchData", "20190516173819519_21651.new")),
		FxAppProvideCounters(),
		FxAppInvokeCreatePrometheusService(),
		FxAppProvideMitchFeedProcessor(),
		FxAppProvideCurrentOpenFile(),
		FxAppInvokeMitchFeedProcessor(),
		fx.Populate(&applicationContext),
		//fx.Invoke(
		//	func(lc fx.Lifecycle, appCtx context.Context, appCancel context.CancelFunc) error {
		//		lc.Append(fx.Hook{
		//			OnStart: func(startContext context.Context) error {
		//				go func() {
		//					time.Sleep(time.Second * 10)
		//					appCancel()
		//				}()
		//				return nil
		//			},
		//			OnStop: nil,
		//		})
		//		return nil
		//	}),
	), applicationContext, nil
}
