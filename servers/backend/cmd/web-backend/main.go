package main

import (
	"log"

	"auto-monitoring/internal/adapter/bigcache"
	"auto-monitoring/internal/adapter/casbin"
	ginhttp "auto-monitoring/internal/adapter/gin-http"
	"auto-monitoring/internal/adapter/gin-http/controller"
	jsonresponse "auto-monitoring/internal/adapter/gin-http/controller/response/json_response"
	"auto-monitoring/internal/adapter/gin-http/middleware"
	captchaMid "auto-monitoring/internal/adapter/gin-http/middleware/captcha"
	casbinMid "auto-monitoring/internal/adapter/gin-http/middleware/casbin"
	"auto-monitoring/internal/adapter/gin-http/middleware/cors"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
	loggerMid "auto-monitoring/internal/adapter/gin-http/middleware/logger"
	redisRateMid "auto-monitoring/internal/adapter/gin-http/middleware/redis_rate"
	secureHeader "auto-monitoring/internal/adapter/gin-http/middleware/secure_header"
	"auto-monitoring/internal/adapter/gin-http/router"
	"auto-monitoring/internal/adapter/gocron"
	"auto-monitoring/internal/adapter/gocron/job"
	"auto-monitoring/internal/adapter/gorm"
	linenotify "auto-monitoring/internal/adapter/line-notify"
	"auto-monitoring/internal/adapter/redispool"
	"auto-monitoring/internal/adapter/repository"
	"auto-monitoring/internal/application/convert"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/pkg/config"
	"auto-monitoring/pkg/logger"
)

func main() {
	conf := config.NewConfig()

	bigcache := bigcache.NewBigcache(bigcache.NewOption(conf))

	mysql := gorm.NewMysql(gorm.NewOption(conf))

	zapLogger := logger.NewZapLogger(logger.NewOption(conf))

	linenotify := linenotify.NewLineNotify(linenotify.NewOption(conf))

	redispool := redispool.NewRedisPool(redispool.NewOption(conf))

	casbinSyncedEnforcer := casbin.NewCasbin(mysql, casbin.NewOption(conf))

	alarmRecord := repository.NewAlarmRecordRepository(mysql)
	alarmSetting := repository.NewAlarmSettingRepository(mysql, linenotify)
	calibration := repository.NewCalibrationRepository(mysql)
	captcha := repository.NewCaptchaRepository()
	controlSignal := repository.NewControlSignalRepository(mysql)
	deviceAlarmRecord := repository.NewDeviceAlarmRecordRepository(mysql)
	deviceAlarmSetting := repository.NewDeviceAlarmSettingRepository(mysql, linenotify)
	device := repository.NewDeviceRepository(mysql)
	last := repository.NewLastRepository(mysql)
	login := repository.NewLoginRepository(mysql)
	physicalQuantity := repository.NewPhysicalQuantityRepository(mysql)
	physicalQuantityPreset := repository.NewPhysicalQuantityPresetRepository(mysql)
	physicalQuantityCatchDetail := repository.NewPhysicalQuantityCatchDetailRepository(mysql)
	record := repository.NewRecordRepository(mysql)
	signalInputMapping := repository.NewSignalInputMappingRepository(mysql, bigcache)
	signalInputMappingDetail := repository.NewSignalInputMappingDetailRepository(mysql)
	station := repository.NewStationRepository(mysql)
	timeSeries := repository.NewTimeSeriesRepository(mysql)
	token := repository.NewTokenRepository(redispool, conf.GetDuration("jwt.expiration"))
	user := repository.NewUserRepository(mysql)
	userAuth := repository.NewUserAuthRepository(mysql)

	signalInputToInput := convert.NewSignalInputToInput(signalInputMapping)

	userUsecase := usecase.NewUserUsecase(user)
	userAuthUsecase := usecase.NewUserAuthUsecase(userAuth)
	tokenUsecase := usecase.NewTokenUsecase(token, *userAuthUsecase)

	alarmRecordUsecase := usecase.NewAlarmRecordUsecase(alarmRecord, alarmSetting, physicalQuantity, device, station)
	alarmSettingUsecase := usecase.NewAlarmSettingUsecase(alarmSetting, physicalQuantity)
	alarmRecordCollectionUsecase := usecase.NewAlarmRecordCollectionUsecase(alarmRecord, deviceAlarmRecord, deviceAlarmSetting, physicalQuantity, alarmSetting)
	alarmUsecase := usecase.NewAlarmUsecase(alarmSetting, alarmRecord)
	calibrationUsecase := usecase.NewCalibrationUsecase(calibration)
	captchaUsecase := usecase.NewCaptchaUsecase(captcha)
	catchInputUsecase := usecase.NewCatchInputUsecase(physicalQuantity, physicalQuantityCatchDetail, device, station, record, *alarmUsecase)
	controlSignalUsecase := usecase.NewControlSignalUsecase(controlSignal)
	deviceAlarmSettingUsecase := usecase.NewDeviceAlarmSettingUsecase(deviceAlarmSetting, deviceAlarmRecord, device)
	deviceUsecase := usecase.NewDeviceUsecase(device)
	lastUsecase := usecase.NewLastUsecase(last)
	loginUsecase := usecase.NewLoginUsecase(login, *tokenUsecase)
	physicalQuantityUsecase := usecase.NewPhysicalQuantityUsecase(physicalQuantity)
	physicalQuantityPresetUsecase := usecase.NewPhysicalQuantityPresetUsecase(physicalQuantityPreset)
	recordUsecase := usecase.NewRecordUsecase(record, device, station, physicalQuantity)
	signalInputMappingUsecase := usecase.NewSignalInputMappingUsecase(signalInputMapping)
	signalInputMappingDetailUsecase := usecase.NewSignalInputMappingDetailUsecase(signalInputMappingDetail)
	stationUsecase := usecase.NewStationUsecase(station)
	timeSeriesUsecase := usecase.NewTimeSeriesUsecase(physicalQuantity, device, timeSeries)

	deviceAlarmJob := job.NewDeviceAlarmJob(zapLogger, deviceAlarmSettingUsecase, deviceUsecase)

	cronApp := gocron.NewApp(
		[]job.IJob{
			deviceAlarmJob,
		},
	)
	go gocron.Run(cronApp)

	jsonResponse := jsonresponse.NewJSONResponse(zapLogger)

	accountController := controller.NewAccountController(jsonResponse)
	alarmRecordController := controller.NewAlarmRecordController(jsonResponse, alarmRecordUsecase)
	alarmRecordCollectionController := controller.NewAlarmRecordCollectionController(jsonResponse, alarmRecordCollectionUsecase)
	alarmSettingController := controller.NewAlarmSettingController(jsonResponse, alarmSettingUsecase)
	calibrationController := controller.NewCalibrationController(jsonResponse, calibrationUsecase, stationUsecase, deviceUsecase)
	captchaController := controller.NewCaptchaController(jsonResponse, captchaUsecase)
	catchInputController := controller.NewCatchInputController(jsonResponse, catchInputUsecase, signalInputToInput)
	controlSignalController := controller.NewControlSignalController(jsonResponse, controlSignalUsecase)
	deviceController := controller.NewDeviceController(jsonResponse, deviceUsecase)
	lastController := controller.NewLastController(jsonResponse, lastUsecase, alarmRecordUsecase)
	loginController := controller.NewLoginController(jsonResponse, loginUsecase)
	physicalQuantityController := controller.NewPhysicalQuantityController(jsonResponse, physicalQuantityUsecase)
	physicalQuantityPresetController := controller.NewPhysicalQuantityPresetController(jsonResponse, physicalQuantityPresetUsecase)
	recordController := controller.NewRecordController(jsonResponse, recordUsecase)
	signalInputMappingController := controller.NewSignalInputMappingController(jsonResponse, signalInputMappingUsecase)
	signalInputMappingDetailController := controller.NewSignalInputMappingDetailController(jsonResponse, signalInputMappingDetailUsecase)
	stationController := controller.NewStationController(jsonResponse, stationUsecase)
	timeSeriesController := controller.NewTimeSeriesController(jsonResponse, timeSeriesUsecase)
	userController := controller.NewUserController(jsonResponse, userUsecase, loginUsecase)

	corsMid := cors.NewCORS(cors.NewOption(conf))
	redisRate := redisRateMid.NewRateLimiter(redispool, zapLogger, redisRateMid.NewOption(conf))
	secureHeader := secureHeader.NewSecureHeader()

	captchaMid := captchaMid.NewCaptcha(jsonResponse, captchaUsecase)
	casbinMid := casbinMid.NewCasbin(jsonResponse, casbinSyncedEnforcer)
	jwtMid := jwt.NewJWT(jsonResponse, tokenUsecase)
	loggerMid := loggerMid.NewLogger(zapLogger)

	accountRouter := router.NewAccountRouter(accountController, jwtMid)
	alarmRecordRouter := router.NewAlarmRecordRouter(alarmRecordController)
	alarmRecordCollectionRouter := router.NewAlarmRecordCollectionRouter(alarmRecordCollectionController)
	alarmSettingRouter := router.NewAlarmSettingRouter(alarmSettingController)
	calibrationRouter := router.NewCalibrationRouter(calibrationController)
	captchaRouter := router.NewCaptchaRouter(captchaController)
	catchInputRouter := router.NewCatchInputRouter(catchInputController)
	controlSignalRouter := router.NewControlSignalRouter(controlSignalController)
	deviceRouter := router.NewDeviceRouter(deviceController, jwtMid)
	lastRouter := router.NewLastRouter(lastController)
	loginRouter := router.NewLoginRouter(loginController, captchaMid, jwtMid)
	physicalQuantityRouter := router.NewPhysicalQuantityRouter(physicalQuantityController, jwtMid)
	physicalQuantityPresetRouter := router.NewPhysicalQuantityPresetRouter(physicalQuantityPresetController, jwtMid)
	recordRouter := router.NewRecordRouter(recordController, loggerMid)
	signalInputMappingRouter := router.NewSignalInputMappingRouter(signalInputMappingController, jwtMid)
	signalInputMappingDetailRouter := router.NewSignalInputMappingDetailRouter(signalInputMappingDetailController, jwtMid)
	stationRouter := router.NewStationRouter(stationController, jwtMid)
	timeSeriesRouter := router.NewTimeSeriesRouter(timeSeriesController)
	userRouter := router.NewUserRouter(userController, jwtMid, casbinMid)

	app := ginhttp.NewApp(
		[]router.IRoute{
			accountRouter,
			alarmRecordRouter,
			alarmRecordCollectionRouter,
			alarmSettingRouter,
			calibrationRouter,
			captchaRouter,
			catchInputRouter,
			controlSignalRouter,
			deviceRouter,
			lastRouter,
			loginRouter,
			physicalQuantityRouter,
			physicalQuantityPresetRouter,
			recordRouter,
			signalInputMappingRouter,
			signalInputMappingDetailRouter,
			stationRouter,
			timeSeriesRouter,
			userRouter,
		},
		[]middleware.IMiddleware{
			corsMid,
			redisRate,
			secureHeader,
		},
	)

	log.Println(conf.GetString("log.application_name"), "v0.8.0")

	ginhttp.Run(app, ginhttp.NewOption(conf))
}
