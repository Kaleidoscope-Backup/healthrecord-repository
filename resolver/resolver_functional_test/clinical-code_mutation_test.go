package resolver_functional_test

// import (
// 	"fmt"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	"github.com/shurcooL/graphql"
// )

// var _ = Describe("ClinicalCodeTypeMutation", func() {

// 	/*==========================================================================
// 	Clinical Code Type Service Tests
// 	==========================================================================*/
// 	Describe("Validating creating a Clinical Code Type in our MongoDB", func() {
// 		Context("With all fields populated in Clinical Code Type", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateClinicalCodeType struct {
// 					Type        graphql.String
// 					Description graphql.String
// 					Id          graphql.String
// 				} `graphql:"createClinicalCodeType(type:$type, description:$description)"`
// 			}
// 			It("Should create a Clinical Code Type without error and return an ID", func() {
// 				//type must be strongly type with graphql library (ie. graphql.String() )
// 				variables := map[string]interface{}{
// 					"type":        graphql.String("TestType"),
// 					"description": graphql.String("Test Description"),
// 				}

// 				//run mutation
// 				err := client.Mutate(ctx, &m, variables)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Mutation Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				// fmt.Printf("Created a %v clinical code type\nwith description: %v\nand id:%v\n", m.CreateClinicalCodeType.Type, m.CreateClinicalCodeType.Description, m.CreateClinicalCodeType.Id)
// 				Expect(m.CreateClinicalCodeType.Id).NotTo(Equal(""))
// 				Expect(m.CreateClinicalCodeType.Id).NotTo(Equal(nil))
// 				Expect(m.CreateClinicalCodeType.Type).To(Equal(graphql.String("TestType")))
// 				Expect(m.CreateClinicalCodeType.Description).To(Equal(graphql.String("Test Description")))
// 			})
// 			It("Should read the new Clinical Code Type without error and all fields populated", func() {
// 				var q struct {
// 					ClinicalCodeTypeQuery struct {
// 						Type        graphql.String
// 						Description graphql.String
// 					} `graphql:"clinicalCodeType(id: $id)"`
// 				}
// 				variables2 := map[string]interface{}{
// 					"id": graphql.String(m.CreateClinicalCodeType.Id),
// 				}

// 				// run query and capture the response
// 				err := client.Query(ctx, &q, variables2)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Query String: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				Expect(q.ClinicalCodeTypeQuery.Type).Should(Equal(m.CreateClinicalCodeType.Type))
// 				Expect(q.ClinicalCodeTypeQuery.Description).Should(Equal(m.CreateClinicalCodeType.Description))
// 			})
// 		})

// 		Context("With only required fields populated in Clinical Code Type", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateClinicalCodeType struct {
// 					Type graphql.String
// 					Id   graphql.String
// 				} `graphql:"createClinicalCodeType(type:$type)"`
// 			}
// 			It("Should create a Clinical Code Type without error and return an ID", func() {
// 				//type must be strongly type with graphql library (ie. graphql.String() )
// 				variables := map[string]interface{}{
// 					"type": graphql.String("TestType"),
// 				}

// 				//run mutation
// 				err := client.Mutate(ctx, &m, variables)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Mutation Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				// fmt.Printf("Created a %v clinical code type\nwith description: %v\nand id:%v\n", m.CreateClinicalCodeType.Type, m.CreateClinicalCodeType.Description, m.CreateClinicalCodeType.Id)
// 				Expect(m.CreateClinicalCodeType.Id).NotTo(Equal(""))
// 				Expect(m.CreateClinicalCodeType.Id).NotTo(Equal(nil))
// 				Expect(m.CreateClinicalCodeType.Type).To(Equal(graphql.String("TestType")))
// 			})
// 			It("Should read the new Clinical Code Type without error and only required fields populated", func() {
// 				var q struct {
// 					ClinicalCodeTypeQuery struct {
// 						Type        graphql.String
// 						Description graphql.String
// 					} `graphql:"clinicalCodeType(id: $id)"`
// 				}
// 				variables2 := map[string]interface{}{
// 					"id": graphql.String(m.CreateClinicalCodeType.Id),
// 				}

// 				// run query and capture the response
// 				err := client.Query(ctx, &q, variables2)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Query Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				Expect(q.ClinicalCodeTypeQuery.Type).Should(Equal(m.CreateClinicalCodeType.Type))
// 				Expect(q.ClinicalCodeTypeQuery.Description).Should(Equal(graphql.String("")))
// 			})
// 		})

// 		Context("With MISSING required fields in Clinical Code Type", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateClinicalCodeType struct {
// 					Id graphql.String
// 				} `graphql:"createClinicalCodeType()"`
// 			}
// 			It("Should FAIL to create a Clinical Code Type", func() {
// 				//run mutation
// 				err := client.Mutate(ctx, &m, nil)
// 				if err != nil {
// 					// Handle error.
// 					Expect(err.Error()).To(Equal("Field \"createClinicalCodeType\" argument \"type\" of type \"String!\" is required but not provided."))
// 				}
// 				Expect(m.CreateClinicalCodeType.Id).To(Equal(graphql.String("")))
// 			})
// 		})
// 	})
// })

// var _ = Describe("ClinicalCodeMutation", func() {

// 	/*==========================================================================
// 	Clinical Code Service Tests
// 	==========================================================================*/
// 	Describe("Validating creating a Clinical Code in our MongoDB", func() {
// 		Context("With all fields populated in Clinical Code", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m1 struct {
// 				CreateClinicalCodeType struct {
// 					Type graphql.String
// 					Id   graphql.String
// 				} `graphql:"createClinicalCodeType(type:$type, description:$description)"`
// 			}

// 			var m2 struct {
// 				CreateClinicalCode struct {
// 					Id             graphql.String
// 					Code           graphql.String
// 					Name           graphql.String
// 					CodeSystem     graphql.String
// 					CodeSystemName graphql.String
// 				} `graphql:"createClinicalCode(code:$code, name:$name, codeSystem: $codeSystem, codeSystemName: $codeSystemName, elementTypeID: $elementTypeId)"`
// 			}
// 			It("Should create a Clinical Code without error and return an ID", func() {
// 				variablesType := map[string]interface{}{
// 					"type":        graphql.String("TestType"),
// 					"description": graphql.String("Test Description"),
// 				}

// 				//run mutation
// 				err := client.Mutate(ctx, &m1, variablesType)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Mutation Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}

// 				//type must be strongly type with graphql library (ie. graphql.String() )
// 				variables := map[string]interface{}{
// 					"code":           graphql.String("TestCode"),
// 					"name":           graphql.String("Test Clinical Code"),
// 					"codeSystem":     graphql.String("TestCodeSystem"),
// 					"codeSystemName": graphql.String("TestCodeSystemName"),
// 					"elementTypeId":  graphql.String(m1.CreateClinicalCodeType.Id),
// 				}

// 				//run mutation
// 				err = client.Mutate(ctx, &m2, variables)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Mutation Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				// fmt.Printf("Created a %v clinical code type\nwith description: %v\nand id:%v\n", m.CreateClinicalCodeType.Type, m.CreateClinicalCodeType.Description, m.CreateClinicalCodeType.Id)
// 				Expect(m2.CreateClinicalCode.Id).NotTo(Equal(""))
// 				Expect(m2.CreateClinicalCode.Id).NotTo(Equal(nil))
// 			})
// 			It("Should read the new Clinical Code without error and all fields populated", func() {
// 				var q struct {
// 					ClinicalCodeQuery struct {
// 						Id             graphql.String
// 						Code           graphql.String
// 						Name           graphql.String
// 						CodeSystem     graphql.String
// 						CodeSystemName graphql.String
// 						ElementType    struct {
// 							Type        graphql.String
// 							Description graphql.String
// 						}
// 					} `graphql:"clinicalCode(id: $id)"`
// 				}
// 				variables2 := map[string]interface{}{
// 					"id": graphql.String(m2.CreateClinicalCode.Id),
// 				}

// 				// run query and capture the response
// 				err := client.Query(ctx, &q, variables2)
// 				if err != nil {
// 					// Handle error.
// 					failureStr := fmt.Sprintf("Query Failed: %s", err.Error())
// 					Fail(failureStr)
// 				}
// 				Expect(q.ClinicalCodeQuery.Id).NotTo(Equal(""))
// 				Expect(q.ClinicalCodeQuery.Id).NotTo(Equal(nil))
// 				Expect(q.ClinicalCodeQuery.Code).To(Equal(graphql.String("TestCode")))
// 				Expect(q.ClinicalCodeQuery.Name).To(Equal(graphql.String("Test Clinical Code")))
// 				Expect(q.ClinicalCodeQuery.CodeSystem).To(Equal(graphql.String("TestCodeSystem")))
// 				Expect(q.ClinicalCodeQuery.CodeSystemName).To(Equal(graphql.String("TestCodeSystemName")))
// 				Expect(q.ClinicalCodeQuery.ElementType.Type).To(Equal(graphql.String("TestType")))
// 				Expect(q.ClinicalCodeQuery.ElementType.Description).To(Equal(graphql.String("Test Description")))
// 			})
// 		})

// 		Context("With MISSING required fields in Clinical Code", func() {
// 			//Create Mutation Object to Send to GraphQL /query endpoint
// 			var m struct {
// 				CreateClinicalCode struct {
// 					Id             graphql.String
// 					Code           graphql.String
// 					Name           graphql.String
// 					CodeSystem     graphql.String
// 					CodeSystemName graphql.String
// 				} `graphql:"createClinicalCode(name:$name, codeSystem: $codeSystem, codeSystemName: $codeSystemName, elementTypeID: $elementTypeId)"`
// 			}
// 			It("Should FAIL to create a Clinical Code Type", func() {
// 				//run mutation
// 				err := client.Mutate(ctx, &m, nil)
// 				if err != nil {
// 					//Handle error
// 					Expect(err.Error()).To(Equal("Field \"createClinicalCode\" argument \"code\" of type \"String!\" is required but not provided."))
// 				}
// 				Expect(m.CreateClinicalCode.Id).To(Equal(graphql.String("")))
// 			})
// 		})
// 	})
// })
