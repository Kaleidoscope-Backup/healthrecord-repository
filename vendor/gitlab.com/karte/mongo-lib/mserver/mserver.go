package mserver

import (
	"github.com/globalsign/mgo"
)

type MServer struct {
	Config       Config
	Database     *mgo.Database
	Interceptors map[string]InterceptorList
}

// AddInterceptor adds a new interceptor for a particular database operation and Karte resource.
// For example:
// AddInterceptor("Create", "HealthRecord", healthRecordInterceptorHandler) would register the
// healthRecordInterceptorHandler methods to be run against a HealthRecord resource when it is created.
//
// To run a handler against ALL resources pass "*" as the resourceType.
//
// Supported database operations are: "Create", "Update", "Delete"
// func (s *MServer) AddInterceptor(op, resourceType string, handler InterceptorHandler) error {

// 	if op == "Create" || op == "Update" || op == "Delete" {
// 		s.Interceptors[op] = append(s.Interceptors[op], Interceptor{ResourceType: resourceType, Handler: handler})
// 		return nil
// 	}
// 	return fmt.Errorf("AddInterceptor: unsupported database operation %s", op)
// }

func NewServer(config Config) *MServer {
	server := &MServer{
		Config:       config,
		Interceptors: make(map[string]InterceptorList),
	}

	return server
}

// Connect to the Karte Mongo Server, which manages the database connection, interceptors, and
// data handling for Mongo transactions
// func (s *MServer) Connect() *mgo.Session {
// 	var err error

// 	// Establish initial connection to mongo
// 	mongoSession, err := mgo.Dial(s.Config.DatabaseHost)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer mongoSession.Close()

// 	mongoSession.SetSocketTimeout(s.Config.DatabaseSocketTimeout)

// 	//establish mongo session and DB object
// 	s.Database = mongoSession.DB(s.Config.DatabaseName)

// 	// Establish karte database masterSession
// 	masterSession := NewMasterSession(mongoSession, s.Config.DatabaseName)
// 	log.Println("MongoDB: Connected")

// 	// Ensure all indexes
// 	NewIndexer(s.Config).ConfigureIndexes(masterSession)

// 	// Establish admin mongoSession
// 	// masterAdminSession := NewMasterSession(mongoSession, "admin")

// 	// Kick off the database op monitoring routine. This periodically checks db.currentOp() and
// 	// kills client-initiated operations exceeding the configurable timeout. Do this AFTER the index
// 	// build to ensure no index build processes are killed unintentionally.
// 	// ticker := time.NewTicker(s.Config.DatabaseKillOpPeriod)
// 	// go killLongRunningOps(ticker, masterAdminSession, s.Config)

// 	// Register all API routes, including GraphQL endpoint
// 	// RegisterRoutes(s.Engine, s.MiddlewareConfig, NewMongoDataAccessLayer(masterSession, s.Interceptors, s.Config), s.Config)

// 	// If not in -readonly mode, clear the count cache
// 	// if !s.Config.ReadOnly {
// 	// 	worker := masterSession.GetWorkerSession()
// 	// 	defer worker.Close()
// 	// 	err = worker.DB().C("countcache").DropCollection()
// 	// 	if err != nil {
// 	// 		log.Println("Server: Failed to clear cache, or cache was empty")
// 	// 	}
// 	// } else {
// 	// 	log.Println("Server: Running in read-only mode")
// 	// }
// 	return mongoSession
// }
