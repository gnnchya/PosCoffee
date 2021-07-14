package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	migrate "github.com/xakep666/mongo-migrate"

	"github.com/gnnchya/PosCoffee/authorize/app"
	"github.com/gnnchya/PosCoffee/authorize/config"
	"github.com/gnnchya/PosCoffee/authorize/middleware"

	"github.com/gnnchya/PosCoffee/authorize/repository/mongodb"
	"github.com/gnnchya/PosCoffee/authorize/service/util"

	serviceValidator "github.com/gnnchya/PosCoffee/authorize/service/validator"

	repoPermissions "github.com/gnnchya/PosCoffee/authorize/repository/permissions"
	permissionsService "github.com/gnnchya/PosCoffee/authorize/service/permissions/implement"

	repoRoles "github.com/gnnchya/PosCoffee/authorize/repository/role"
	rolesService "github.com/gnnchya/PosCoffee/authorize/service/roles/implement"

	grpcRepo "github.com/gnnchya/PosCoffee/authorize/repository/grpc"
	grpcServiceServer "github.com/gnnchya/PosCoffee/authorize/service/grpc/implement"
	grpcService "github.com/gnnchya/PosCoffee/authorize/service/grpcClient/implement"
)

const (
	NETWORK = "tcp"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()

	uuid, err := util.NewUUID()
	panicIfErr(err)

	// repositories
	permissionRepo, err := repoPermissions.New(ctx, configMongodb(appConfig, appConfig.MongoDBPermissionTableName))
	panicIfErr(err)

	roleRepo, err := repoRoles.New(ctx, configMongodb(appConfig, appConfig.MongoDBRoleTableName))
	panicIfErr(err)

	grpcRepo := grpcRepo.New(configGrpc(appConfig))

	// validators
	validator := serviceValidator.New(&serviceValidator.ValidatorRepository{
		Permission: permissionRepo,
		Role:       roleRepo,
	})

	// services
	permission := permissionsService.New(validator, permissionRepo, uuid)
	role := rolesService.New(validator, roleRepo, uuid)
	gService := grpcService.New(grpcRepo)

	// middleware services
	midService := middleware.New(appConfig, gService, role)
	go grpcServiceServer.New(grpcRepo, permission, role)

	return app.New(permission, role, appConfig, midService)
}

func setupJaeger(appConfig *config.Config) io.Closer {
	cfg, err := jaegerConf.FromEnv()
	panicIfErr(err)

	cfg.ServiceName = appConfig.AppName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter = &jaegerConf.ReporterConfig{LogSpans: true}

	jLogger := jaegerLog.NullLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
	)
	panicIfErr(err)
	opentracing.SetGlobalTracer(tracer)

	return closer
}

func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})
	lr.SetLevel(logrus.ErrorLevel)
	return lr
}

func configMongodb(appConfig *config.Config, collName string) *mongodb.Config {
	return &mongodb.Config{
		URI:      appConfig.MongoDBEndpoint,
		DBName:   appConfig.MongoDBName,
		CollName: collName,
	}
}

func MigrateMongoDB(appConfig *config.Config) {
	ctx := context.Background()
	mongoConfig := configMongodb(appConfig, "_migrations")
	repo, err := mongodb.New(ctx, mongoConfig)
	panicIfErr(err)

	migrate.SetDatabase(repo.DB)
	migrate.SetMigrationsCollection("_migrations")

	if len(os.Args) == 1 {
		if err = migrate.Up(migrate.AllAvailable); err != nil {
			log.Panic("Error migration up", err.Error())
			return
		}
		log.Println("Auto migration done.")
		return
	}

	// Command for use migration
	// go run *.go new <filename>
	// go run *.go migrate
	// go run *.go rollback

	option := os.Args[1]
	switch option {
	case "new":
		doNewMigrateFile()
		os.Exit(0)
	case "migrate":
		doMigrate()
		os.Exit(0)
	case "rollback":
		doRollback()
		os.Exit(0)
	default:

	}
}

func doNewMigrateFile() {
	if len(os.Args) != 3 {
		log.Fatal("Invalid command to new migration file")
		return
	}

	folder := "migrations"
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := os.Stat(path + "/" + folder); os.IsNotExist(err) {
		_ = os.Mkdir(path+"/"+folder, 0755)
	}

	file := fmt.Sprintf("%s/%s_%s.go", folder, time.Now().Format("20060102150405"), strings.ReplaceAll(os.Args[2], "-", "_"))

	from, err := os.Open(fmt.Sprintf("%s/template.txt", folder))
	if err != nil {
		log.Fatal("File migration template not found !!")
	}
	defer func() {
		_ = from.Close()
	}()

	to, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panicIfErr(err)
	}
	defer func() {
		_ = to.Close()
	}()

	_, err = io.Copy(to, from)
	if err != nil {
		panicIfErr(err)
	}
	log.Printf(fmt.Sprintf("%s %s\n", "New migration file created =>", file))
}

func doMigrate() {
	log.Println("Beginning migration ...")
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		panicIfErr(err)
	}
	log.Println("Migration done.")
}

func doRollback() {
	log.Println("Beginning rollback ...")
	if err := migrate.Down(migrate.AllAvailable); err != nil {
		panicIfErr(err)
	}
	log.Println("Rollback done.")
}

func configGrpc(appConfig *config.Config) *grpcRepo.Config {
	return &grpcRepo.Config{
		Network: NETWORK,
		Port:    appConfig.GRPCAuthorizeHost,
	}
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
