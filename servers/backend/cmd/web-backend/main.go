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
	deviceAlarmRecord := repository.NewDeviceAlarmRecordRepository(mysql)
	deviceAlarmSetting := repository.NewDeviceAlarmSettingRepository(mysql, linenotify)
	device := repository.NewDeviceRepository(mysql)
	deviceStation := repository.NewDeviceStationRepository(mysql)
	last := repository.NewLastRepository(mysql)
	login := repository.NewLoginRepository(mysql)
	physicalQuantity := repository.NewPhysicalQuantityRepository(mysql)
	physicalQuantityCatchDetail := repository.NewPhysicalQuantityCatchDetailRepository(mysql)
	record := repository.NewRecordRepository(mysql)
	signalInputMapping := repository.NewSignalInputMappingRepository(mysql, bigcache)
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
	alarmSettingUsecase := usecase.NewAlarmSettingUsecase(alarmSetting, device)
	alarmUsecase := usecase.NewAlarmUsecase(alarmSetting, alarmRecord)
	calibrationUsecase := usecase.NewCalibrationUsecase(calibration)
	captchaUsecase := usecase.NewCaptchaUsecase(captcha)
	catchInputUsecase := usecase.NewCatchInputUsecase(physicalQuantity, physicalQuantityCatchDetail, device, station, record, *alarmUsecase)
	deviceAlarmSettingUsecase := usecase.NewDeviceAlarmSettingUsecase(deviceAlarmSetting, deviceAlarmRecord, device)
	deviceUsecase := usecase.NewDeviceUsecase(device)
	deviceStationUsecase := usecase.NewDeviceStationUsecase(deviceStation)
	lastUsecase := usecase.NewLastUsecase(last)
	loginUsecase := usecase.NewLoginUsecase(login, *tokenUsecase)
	physicalQuantityUsecase := usecase.NewPhysicalQuantityUsecase(physicalQuantity)
	recordUsecase := usecase.NewRecordUsecase(record, device, station, physicalQuantity)
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

	alarmRecordController := controller.NewAlarmRecordController(jsonResponse, alarmRecordUsecase)
	alarmSettingController := controller.NewAlarmSettingController(jsonResponse, alarmSettingUsecase)
	calibrationController := controller.NewCalibrationController(jsonResponse, calibrationUsecase, stationUsecase, deviceUsecase)
	captchaController := controller.NewCaptchaController(jsonResponse, captchaUsecase)
	catchInputController := controller.NewCatchInputController(jsonResponse, catchInputUsecase, signalInputToInput)
	deviceController := controller.NewDeviceController(jsonResponse, deviceUsecase)
	deviceStationController := controller.NewDeviceStationController(jsonResponse, deviceStationUsecase)
	lastController := controller.NewLastController(jsonResponse, lastUsecase, alarmRecordUsecase)
	loginController := controller.NewLoginController(jsonResponse, loginUsecase)
	physicalQuantityController := controller.NewPhysicalQuantityController(jsonResponse, physicalQuantityUsecase)
	recordController := controller.NewRecordController(jsonResponse, recordUsecase)
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

	alarmRecordRouter := router.NewAlarmRecordRouter(alarmRecordController)
	alarmSettingRouter := router.NewAlarmSettingRouter(alarmSettingController)
	calibrationRouter := router.NewCalibrationRouter(calibrationController)
	captchaRouter := router.NewCaptchaRouter(captchaController)
	catchInputRouter := router.NewCatchInputRouter(catchInputController)
	deviceRouter := router.NewDeviceRouter(deviceController)
	deviceStationRouter := router.NewDeviceStationRouter(deviceStationController)
	lastRouter := router.NewLastRouter(lastController)
	loginRouter := router.NewLoginRouter(loginController, captchaMid, jwtMid)
	physicalQuantityRouter := router.NewPhysicalQuantityRouter(physicalQuantityController)
	recordRouter := router.NewRecordRouter(recordController, loggerMid)
	stationRouter := router.NewStationRouter(stationController)
	timeSeriesRouter := router.NewTimeSeriesRouter(timeSeriesController)
	userRouter := router.NewUserRouter(userController, jwtMid, casbinMid)

	app := ginhttp.NewApp(
		[]router.IRoute{
			alarmRecordRouter,
			alarmSettingRouter,
			calibrationRouter,
			captchaRouter,
			catchInputRouter,
			deviceRouter,
			deviceStationRouter,
			lastRouter,
			loginRouter,
			physicalQuantityRouter,
			recordRouter,
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

	log.Println(conf.GetString("log.application_name"), "v0.11.0")

	ginhttp.Run(app, ginhttp.NewOption(conf))
}
