package resolver_functional_test

// import (
// 	"fmt"
// 	"time"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	"github.com/shurcooL/graphql"
// 	"github.com/karte/healthrecord-repository/model"
// 	"github.com/karte/healthrecord-repository/util"
// )

// var _ = Describe("MedicationMutation", func() {

// 	/*==========================================================================
// 	Medication Service Tests
// 	==========================================================================*/
// 	Describe("Validating creating an MEDICATION in our MongoDB", func() {
// 		Context("With all fields populated in Medication", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateMedication struct {
// 					Id               graphql.String
// 					MedicationStatus graphql.String
// 					ProductName      graphql.String
// 					IsOverTheCounter graphql.String
// 					Route            model.OrganizationType
// 				} `graphql:"createMedication(medication:$medication)"`
// 			}
// 			//Mutation to create Medication
// 			It("Should create Medication without error and return an ID", func() {
// 				//type must be strongly type with graphql library (ie. graphql.String() )
// 				productName := "Advil"
// 				isOverTheCounter := true
// 				route := model.ORAL_ADMINISTRATION
// 				instructions := "With Plenty of Water"
// 				var dosageValue int32 = 10
// 				dosageFrequency := "1 tablet per day"
// 				dosageUnit := "10 mg"

// 				var refillRemainingVar int32 = 10
// 				refillsRemaining := &refillRemainingVar
// 				var refillsTotalVar int32 = 50
// 				refillsTotal := &refillsTotalVar
// 				var strengthNumber int32 = 10
// 				strengthUnit := "mg"

// 				var t = util.Time{time.Now()}
// 				start := t

// 				variables := map[string]interface{}{
// 					"medication": model.MedicationCreate{
// 						ProductName:      productName,
// 						IsOverTheCounter: isOverTheCounter,
// 						Route:            route,
// 						Instructions:     instructions,
// 						DosageValue:      dosageValue,
// 						DosageFrequency:  dosageFrequency,
// 						DosageUnit:       dosageUnit,
// 						RefillsRemaining: refillsRemaining,
// 						RefillsTotal:     refillsTotal,
// 						StrengthNumber:   strengthNumber,
// 						StrengthUnit:     strengthUnit,
// 						Start:            start,
// 					},
// 				}

// 				//run mutation
// 				err := client.Mutate(ctx, &m, variables)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Mutation Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				fmt.Printf("Created an organization with name %v\n", m.CreateMedication.ProductName)
// 				Expect(m.CreateMedication.Id).NotTo(Equal(""))
// 				Expect(m.CreateMedication.IsOverTheCounter).To(Equal(true))
// 				Expect(m.CreateMedication.Route).To(Equal(model.ORAL_ADMINISTRATION))
// 			})
// 			//Medication query
// 			It("Should read the new Medication without error and all fields populated", func() {
// 				var q struct {
// 					MedicationQuery struct {
// 						Id               graphql.String
// 						ProductName      graphql.String
// 						Route            graphql.String
// 						DosageValue      graphql.String
// 						IsOverTheCounter graphql.Boolean
// 					} `graphql:"medication(id: $id)"`
// 				}
// 				variables2 := map[string]interface{}{
// 					"id": graphql.String(m.CreateMedication.Id),
// 				}

// 				// run query and capture the response
// 				err := client.Query(ctx, &q, variables2)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Query String: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				Expect(q.MedicationQuery.ProductName).Should(Equal(m.CreateMedication.ProductName))
// 				Expect(q.MedicationQuery.Route).Should(Equal(m.CreateMedication.Route))
// 				Expect(q.MedicationQuery.IsOverTheCounter).Should(Equal(m.CreateMedication.IsOverTheCounter))
// 			})
// 		})

// 		//When required input is not provided
// 		Context("With MISSING required fields in Organization", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateMedication struct {
// 					ProductName graphql.String
// 					Id          graphql.String
// 				} `graphql:"createOrganization()"`
// 			}
// 			It("Should FAIL to create a Medication", func() {
// 				//run mutation
// 				err := client.Mutate(ctx, &m, nil)
// 				if err != nil {
// 					// Handle error.
// 					Expect(err.Error()).To(Equal("Field \"createMedication\" argument \"medication\" of type \"model.MedicationCreate!\" is required but not provided."))
// 				}
// 				Expect(m.CreateMedication.Id).To(Equal(graphql.String("")))
// 			})
// 		})

// 	})
// })
