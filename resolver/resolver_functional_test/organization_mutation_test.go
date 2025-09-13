package resolver_functional_test

// import (
// 	"fmt"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	"github.com/shurcooL/graphql"
// 	"github.com/karte/healthrecord-repository/model"
// )

// var _ = Describe("OrganizationMutation", func() {

// 	/*==========================================================================
// 	Organization Service Tests
// 	==========================================================================*/
// 	Describe("Validating creating an Organization in our MongoDB", func() {
// 		Context("With all fields populated in Organization", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateOrganization struct {
// 					Id           graphql.String
// 					Name         graphql.String
// 					SourceID     graphql.String
// 					SourceIDType graphql.String
// 					Type         model.OrganizationType
// 				} `graphql:"createOrganization(organization:$organization)"`
// 			}
// 			//Mutation to create Organization
// 			It("Should create Organization without error and return an ID", func() {
// 				//type must be strongly type with graphql library (ie. graphql.String() )

// 				sourceID := "890-890"
// 				SourceIDType := "TAX ID"
// 				name := "KAISER PARMANENTE"
// 				variables := map[string]interface{}{
// 					"organization": model.OrganizationCreate{
// 						Name:         name,
// 						Type:         model.OrganizationType("HOSPITAL"),
// 						SourceID:     &sourceID,
// 						SourceIDType: &SourceIDType,
// 					},
// 				}

// 				//run mutation
// 				err := client.Mutate(ctx, &m, variables)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Mutation Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				fmt.Printf("Created an organization with name %v\n", m.CreateOrganization.Name)
// 				Expect(m.CreateOrganization.Id).NotTo(Equal(""))
// 				Expect(m.CreateOrganization.Name).To(Equal(graphql.String("KAISER PARMANENTE")))
// 				Expect(m.CreateOrganization.SourceID).To(Equal(graphql.String("890-890")))
// 				Expect(m.CreateOrganization.SourceIDType).To(Equal(graphql.String("TAX ID")))
// 			})
// 			//Organization query
// 			It("Should read the new Organization without error and all fields populated", func() {
// 				var q struct {
// 					OrganizationQuery struct {
// 						Name         graphql.String
// 						SourceID     graphql.String
// 						SourceIDType graphql.String
// 					} `graphql:"organization(id: $id)"`
// 				}
// 				variables2 := map[string]interface{}{
// 					"id": graphql.String(m.CreateOrganization.Id),
// 				}

// 				// run query and capture the response
// 				err := client.Query(ctx, &q, variables2)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Query String: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				Expect(q.OrganizationQuery.Name).Should(Equal(m.CreateOrganization.Name))
// 				Expect(q.OrganizationQuery.SourceID).Should(Equal(m.CreateOrganization.SourceID))
// 				Expect(q.OrganizationQuery.SourceIDType).Should(Equal(m.CreateOrganization.SourceIDType))
// 			})
// 		})

// 		//When required input is not provided
// 		Context("With MISSING required fields in Organization", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateOrganization struct {
// 					Name graphql.String
// 					Id   graphql.String
// 				} `graphql:"createOrganization()"`
// 			}
// 			It("Should FAIL to create a Organization", func() {
// 				//run mutation
// 				err := client.Mutate(ctx, &m, nil)
// 				if err != nil {
// 					// Handle error.
// 					fmt.Println(err)
// 					Expect(err.Error()).To(Equal("Field \"createOrganization\" argument \"organization\" of type \"model.OrganizationCreate!\" is required but not provided."))
// 				}
// 				Expect(m.CreateOrganization.Id).To(Equal(graphql.String("")))
// 			})
// 		})

// 	})
// })
