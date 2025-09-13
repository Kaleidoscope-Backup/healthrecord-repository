package service_test

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/mserver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	logging "github.com/op/go-logging"
)

var testDAL mserver.DataAccessLayer
var testLog *logging.Logger

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = BeforeSuite(func() {
	//Start with Karte's Default Config and change as needed
	config := util.SetupConfig()
	config.DatabaseName = "karte_health_records_test"
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

	//Create DAL that is used for Data Interactions with Mongo directly
	testDAL = mserver.NewMongoDataAccessLayer(masterSession, s.Interceptors, s.Config)

	//This is to set up logging
	testLog = service.NewLogger(&config)
})

var _ = AfterSuite(func() {
	//Drop Database
	//Start with Karte's Default Config and change as needed
	config := util.SetupConfig()
	config.DatabaseName = "karte_health_records_test"
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
