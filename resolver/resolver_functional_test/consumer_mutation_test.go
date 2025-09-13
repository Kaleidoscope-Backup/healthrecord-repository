package resolver_functional_test

// import (
// 	"fmt"

// 	"github.com/karte/healthrecord-repository/model"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	"github.com/shurcooL/graphql"
// )

// var _ = Describe("ConsumerMutation", func() {

// 	/*==========================================================================
// 	Consumer Service Tests
// 	==========================================================================*/
// 	Describe("Validating creating an Consumer in our MongoDB", func() {
// 		Context("With all fields populated in Consumer", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateConsumer struct {
// 					Id                 graphql.String
// 					FirstName          graphql.String
// 					LastName           graphql.String
// 					Email              graphql.String
// 					LanguagePreference graphql.String
// 					Gender             model.Gender
// 					MarritalStatus     model.MarritalStatus
// 					Ethnicity          *graphql.String
// 				} `graphql:"createConsumer(consumer:$consumer)"`
// 			}
// 			//Mutation to create Consumer
// 			It("Should create Consumer without error and return an ID", func() {
// 				//type must be strongly type with graphql library (ie. graphql.String() )

// 				firstName := "Suparna"
// 				lastName := "Pal"
// 				email := "suparna2702@gmail.com"
// 				password := "1234"
// 				sourceConsumerIDSystem := "KAISER-6789"
// 				sourceConsumerIDValue := "567-9090"
// 				sourceConsumerIDAssigner := "ORG-67890"
// 				ethnicity := "Asian"
// 				gender := model.MALE
// 				marritalStatus := model.SINGLE
// 				languagePreference := "English"

// 				variables := map[string]interface{}{
// 					"consumer": model.ConsumerCreate{
// 						FirstName: firstName,
// 						LastName:  lastName,
// 						Email:     email,
// 						Password:  password,
// 						// SourceConsumerIDSystem:   &sourceConsumerIDSystem,
// 						// SourceConsumerIDValue:    &sourceConsumerIDValue,
// 						// SourceConsumerIDAssigner: &sourceConsumerIDAssigner,
// 						Ethnicity:          &ethnicity,
// 						Gender:             &gender,
// 						MarritalStatus:     &marritalStatus,
// 						LanguagePreference: &languagePreference,
// 					},
// 				}

// 				//run mutation
// 				err := client.Mutate(ctx, &m, variables)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Mutation Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				fmt.Printf("Created an consumer with name %v\n", m.CreateConsumer.FirstName)
// 				Expect(m.CreateConsumer.Id).NotTo(Equal(""))
// 				Expect(m.CreateConsumer.FirstName).To(Equal(graphql.String("Suparna")))
// 				Expect(m.CreateConsumer.LastName).To(Equal(graphql.String("Pal")))
// 				Expect(m.CreateConsumer.Email).To(Equal(graphql.String("suparna2702@gmail.com")))
// 				Expect(m.CreateConsumer.Gender).To(Equal(model.MALE))
// 				Expect(m.CreateConsumer.MarritalStatus).To(Equal(model.SINGLE))
// 				Expect(m.CreateConsumer.LanguagePreference).To(Equal(graphql.String("English")))
// 			})
// 			//Consumer query
// 			It("Should read the new Consumer without error and all fields populated", func() {
// 				var q struct {
// 					ConsumerQuery struct {
// 						FirstName graphql.String
// 						LastName  graphql.String
// 					} `graphql:"consumer(id: $id)"`
// 				}
// 				variables2 := map[string]interface{}{
// 					"id": graphql.String(m.CreateConsumer.Id),
// 				}

// 				// run query and capture the response
// 				err := client.Query(ctx, &q, variables2)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Query String: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				Expect(q.ConsumerQuery.FirstName).Should(Equal(m.CreateConsumer.FirstName))
// 			})
// 		})

// 		//When required input is not provided
// 		Context("With MISSING required fields in  Consumer", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateConsumer struct {
// 					FirstName graphql.String
// 					Id        graphql.String
// 				} `graphql:"createConsumer()"`
// 			}
// 			It("Should FAIL to create a Consumer", func() {
// 				//run mutation
// 				err := client.Mutate(ctx, &m, nil)
// 				if err != nil {
// 					// Handle error.
// 					Expect(err.Error()).To(Equal("Field \"createConsumer\" argument \"consumer\" of type \"model.ConsumerCreate!\" is required but not provided."))
// 				}
// 				Expect(m.CreateConsumer.Id).To(Equal(graphql.String("")))
// 			})
// 		})

// 	})
// })
