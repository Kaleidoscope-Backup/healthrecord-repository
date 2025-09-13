package resolver_functional_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	c "github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/codegangsta/negroni"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/resolver"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/schema"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/service"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/Kaleidoscope-Backup/microservice-utilities/auth0"
	h "github.com/Kaleidoscope-Backup/microservice-utilities/handler"
	"github.com/Kaleidoscope-Backup/microservice-utilities/loader"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo"
	graphqlGo "github.com/graph-gophers/graphql-go"
)

// GLOBALS
var client *graphql.Client
var ctx context.Context

func TestResolverFunctionalTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResolverFunctional Suite")
}

var _ = BeforeSuite(func() {
	ctx = context.Background()
	testMain("karte_health_records_test")

	t := auth0.GetAuthToken()

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: t.AccessToken},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	client = graphql.NewClient("http://localhost:5000/query", httpClient)
})

var _ = AfterSuite(func() {
	//Drop Database

	//Start with Karte's Default Config and change as needed
	config := util.SetupConfig()

	var err error
	var mongoSession *mgo.Session
	//only delete database if tests are run locally...gitlab clears out old docker containers anyway
	if config.Local == "DEV" {
		mongoURI := config.DatabaseHost + ":" + config.DatabasePort
		mongoSession, err = mgo.Dial(mongoURI)
		if err != nil {
			panic(err)
		}
		//drop database
		err = mongoSession.DB("karte_health_records_test").DropDatabase()
		if err != nil {
			panic(err)
		}
		defer mongoSession.Close()
	}

})

func testMain(dbName string) {
	//Start with Karte's Default Config and change as needed
	config := util.SetupConfig()
	config.DatabaseName = dbName
	mconfig := config.MongoConfig()

	//This creates a new Mongo Server
	s := mserver.NewServer(mconfig)
	var err error
	var mongoSession *mgo.Session

	// Establish initial connection to mongo
	if config.Local == "DEV" {
		mongoURI := config.DatabaseHost + ":" + config.DatabasePort
		mongoSession, err = mgo.Dial(mongoURI)
		if err != nil {
			panic(err)
		}
	} else {
		mongoURI := "mongodb://" + config.DatabaseUsername + ":" + config.DatabasePassword + "@" + config.DatabaseHost
		dialInfo, err := mgo.ParseURL(mongoURI)
		if err != nil {
			fmt.Println(err)
		}

		//Below part is similar to above.
		tlsConfig := &tls.Config{}
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}

		mongoSession, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			panic(err)
		}
	}

	mongoSession.SetSocketTimeout(s.Config.DatabaseSocketTimeout)

	//establish mongo session and DB object
	s.Database = mongoSession.DB(s.Config.DatabaseName)

	// Establish karte database masterSession
	masterSession := mserver.NewMasterSession(mongoSession, s.Config.DatabaseName)
	log.Println("MongoDB: Connected")

	// Ensure all indexes in Mongo (builds indexes in NoSQL)
	mserver.NewIndexer(s.Config).ConfigureIndexes(masterSession)

	// Establish admin mongoSession
	masterAdminSession := mserver.NewMasterSession(mongoSession, "admin")

	// Kick off the database op monitoring routine. This periodically checks db.currentOp() and
	// kills client-initiated operations exceeding the configurable timeout. Do this AFTER the index
	// build to ensure no index build processes are killed unintentionally.
	ticker := time.NewTicker(s.Config.DatabaseKillOpPeriod)
	go mserver.KillLongRunningOps(ticker, masterAdminSession, s.Config)

	//Create DAL that is used for Data Interactions with Mongo directly
	dal := mserver.NewMongoDataAccessLayer(masterSession, s.Interceptors, s.Config)

	//This is to set up graphql
	ctx := context.Background()
	log := service.NewLogger(&config)

	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "log", log)

	//add singleton services for Organization
	organizationService := service.NewOrganizationService(dal, log)
	ctx = context.WithValue(ctx, c.OrganizationService, organizationService)

	//add singleton service for Practitioner
	practitionerService := service.NewPractitionerService(dal, log)
	ctx = context.WithValue(ctx, c.PractitionerService, practitionerService)

	//add singleton service for Consumer
	consumerService := service.NewConsumerService(dal, log)
	ctx = context.WithValue(ctx, c.ConsumerService, consumerService)

	heartRateService := service.NewHeartRateService(dal, log)
	ctx = context.WithValue(ctx, c.HeartRateService, heartRateService)

	//db admin services
	clinicalCodeService := service.NewClinicalCodeService(dal, log)
	ctx = context.WithValue(ctx, c.ClinicalCodeService, clinicalCodeService)

	dosageService := service.NewDosageService(dal, log)
	ctx = context.WithValue(ctx, c.DosageService, dosageService)

	strengthService := service.NewStrengthService(dal, log)
	ctx = context.WithValue(ctx, c.StrengthService, strengthService)

	medicationService := service.NewMedicationService(dal, log)
	ctx = context.WithValue(ctx, c.MedicationService, medicationService)

	contactPointService := service.NewContactPointService(dal, log)
	ctx = context.WithValue(ctx, c.ContactPointService, contactPointService)

	sourceRecordIDService := service.NewSourceRecordIDService(dal, log)
	ctx = context.WithValue(ctx, c.SourceRecordIDService, sourceRecordIDService)

	sourceConsumerIDService := service.NewSourceConsumerIDService(dal, log)
	ctx = context.WithValue(ctx, c.SourceConsumerIDService, sourceConsumerIDService)

	sourceOrganizationIDService := service.NewSourceOrganizationIDService(dal, log)
	ctx = context.WithValue(ctx, c.SourceOrganizationIDService, sourceOrganizationIDService)

	accountService := service.NewAccountService(dal, log)
	ctx = context.WithValue(ctx, c.AccountService, accountService)

	socialHistoryObservationRecordService := service.NewSocialHistoryObservationRecordService(dal, log)
	ctx = context.WithValue(ctx, c.SocialHistoryObservationRecordService, socialHistoryObservationRecordService)

	medicationRecordService := service.NewMedicationRecordService(dal, log)
	ctx = context.WithValue(ctx, c.MedicationRecordService, medicationRecordService)

	allergyRecordService := service.NewAllergyRecordService(dal, log)
	ctx = context.WithValue(ctx, c.AllergyRecordService, allergyRecordService)

	encounterRecordService := service.NewEncounterRecordService(dal, log)
	ctx = context.WithValue(ctx, c.EncounterRecordService, encounterRecordService)

	procedureRecordService := service.NewProcedureRecordService(dal, log)
	ctx = context.WithValue(ctx, c.ProcedureRecordService, procedureRecordService)

	familyMemberHistoryRecordService := service.NewFamilyMemberHistoryRecordService(dal, log)
	ctx = context.WithValue(ctx, c.FamilyMemberHistoryRecordService, familyMemberHistoryRecordService)

	clinicalAssesmentObservationRecordService := service.NewClinicalAssesmentObservationRecordService(dal, log)
	ctx = context.WithValue(ctx, c.ClinicalAssesmentObservationRecordService, clinicalAssesmentObservationRecordService)

	personalCharacteristicsObservationRecordService := service.NewPersonalCharacteristicsObservationRecordService(dal, log)
	ctx = context.WithValue(ctx, c.PersonalCharacteristicsObservationRecordService, personalCharacteristicsObservationRecordService)

	imagingResultObservationRecordService := service.NewImagingResultObservationRecordService(dal, log)
	ctx = context.WithValue(ctx, c.ImagingResultObservationRecordService, imagingResultObservationRecordService)

	labResultObservationRecordService := service.NewLabResultObservationRecordService(dal, log)
	ctx = context.WithValue(ctx, c.LabResultObservationRecordService, labResultObservationRecordService)

	vitalObservationRecordService := service.NewVitalObservationRecordService(dal, log)
	ctx = context.WithValue(ctx, c.VitalObservationRecordService, vitalObservationRecordService)

	conditionRecordService := service.NewConditionRecordService(dal, log)
	ctx = context.WithValue(ctx, c.ConditionRecordService, conditionRecordService)

	addressService := service.NewAddressService(dal, log)
	ctx = context.WithValue(ctx, c.AddressService, addressService)

	immunizationRecordService := service.NewImmunizationRecordService(dal, log)
	ctx = context.WithValue(ctx, c.ImmunizationRecordService, immunizationRecordService)

	activityRecordService := service.NewActivityRecordService(dal, log)
	ctx = context.WithValue(ctx, c.ActivityRecordService, activityRecordService)

	mealRecordService := service.NewMealRecordService(dal, log)
	ctx = context.WithValue(ctx, c.MealRecordService, mealRecordService)

	goalRecordService := service.NewGoalRecordService(dal, log)
	ctx = context.WithValue(ctx, c.GoalRecordService, goalRecordService)

	sleepRecordService := service.NewSleepRecordService(dal, log)
	ctx = context.WithValue(ctx, c.SleepRecordService, sleepRecordService)

	healthRecordService := service.NewHealthRecordService(dal, log)
	ctx = context.WithValue(ctx, c.HealthRecordService, healthRecordService)

	graphqlSchema := graphqlGo.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	//Establish HTTP Handlers
	loggerHandler := &h.LoggerHandler{config.Debug}
	jwtMiddleware := auth0.NewJWTMiddleware("https://healthrecord-repository.karte.io", "https://karte-dev.auth0.com/")
	graphqlHandler := &h.GraphQL{Schema: graphqlSchema, Loaders: loader.NewLoaderCollection()}

	n := negroni.New()
	n.Use(negroni.HandlerFunc(jwtMiddleware.HandlerWithNext))
	n.Use(negroni.HandlerFunc(h.AuthScopeHandler))
	n.Use(negroni.Wrap(graphqlHandler))

	http.Handle("/query", h.AddContext(ctx, loggerHandler.Logging(n)))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	go http.ListenAndServe(":"+config.AppPort, nil)
}
