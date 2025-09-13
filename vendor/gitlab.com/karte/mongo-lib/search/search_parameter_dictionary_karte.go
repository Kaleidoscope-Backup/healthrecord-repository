package search

//SearchParameterDictionaryKarte ...
var SearchParameterDictionaryKarte = map[string]map[string]SearchParamInfo{

	"Account": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Account",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"userName": SearchParamInfo{
			Resource: "Account",
			Name:     "userName",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "userName", Type: "string"},
			},
		},
	},

	"Application": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Application",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"name": SearchParamInfo{
			Resource: "Application",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "name", Type: "string"},
			},
		},
	},

	"Organization": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Organization",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"name": SearchParamInfo{
			Resource: "Organization",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "name", Type: "string"},
			},
		},
		"email": SearchParamInfo{
			Resource: "Organization",
			Name:     "email",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "email", Type: "string"},
			},
		},
		"type": SearchParamInfo{
			Resource: "Organization",
			Name:     "type",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "type", Type: "string"},
			},
		},
	},

	"Consumer": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Consumer",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "Consumer",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"firstName": SearchParamInfo{
			Resource: "Consumer",
			Name:     "firstName",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "actor.firstName", Type: "string"},
			},
		},
		"lastName": SearchParamInfo{
			Resource: "Consumer",
			Name:     "lastName",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "actor.lastName", Type: "string"},
			},
		},
		"email": SearchParamInfo{
			Resource: "Consumer",
			Name:     "email",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "actor.email", Type: "string"},
			},
		},
	},

	"Practitioner": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Practitioner",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "Practitioner",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"firstName": SearchParamInfo{
			Resource: "Practitioner",
			Name:     "firstName",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "actor.firstName", Type: "string"},
			},
		},
		"lastName": SearchParamInfo{
			Resource: "Practitioner",
			Name:     "lastName",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "actor.lastName", Type: "string"},
			},
		},
		"email": SearchParamInfo{
			Resource: "Practitioner",
			Name:     "email",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "actor.email", Type: "string"},
			},
		},
	},

	"HealthRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "HealthRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "HealthRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "HealthRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "consumerID", Type: "string"},
			},
		},
	},

	"MedicationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "MedicationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "MedicationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "MedicationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"prescribedBy": SearchParamInfo{
			Resource: "MedicationRecord",
			Name:     "prescribedBy",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "prescribedBy", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "MedicationRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"ImmunizationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ImmunizationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "ImmunizationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "ImmunizationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "ImmunizationRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"ConditionRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ConditionRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "ConditionRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "ConditionRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "ConditionRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"AllergyRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "AllergyRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "AllergyRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "AllergyRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "AllergyRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"VitalObservationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "VitalObservationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "VitalObservationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "VitalObservationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "VitalObservationRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"SocialHistoryObservationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "SocialHistoryObservationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "SocialHistoryObservationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "SocialHistoryObservationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"type": SearchParamInfo{
			Resource: "SocialHistoryObservationRecord",
			Name:     "type",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "type", Type: "string"},
			},
		},
		"status": SearchParamInfo{
			Resource: "SocialHistoryObservationRecord",
			Name:     "status",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "status", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "SocialHistoryObservationRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"ClinicalAssesmentObservationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ClinicalAssesmentObservationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "ClinicalAssesmentObservationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "ClinicalAssesmentObservationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "ClinicalAssesmentObservationRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"EncounterRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "EncounterRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "EncounterRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "EncounterRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "EncounterRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"FamilyMemberHistoryRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "FamilyMemberHistoryRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "FamilyMemberHistoryRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "FamilyMemberHistoryRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "FamilyMemberHistoryRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"ImagingResultObservationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ImagingResultObservationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "ImagingResultObservationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "ImagingResultObservationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "ImagingResultObservationRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
	},

	"LabResultObservationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "LabResultObservationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "LabResultObservationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "LabResultObservationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "LabResultObservationRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
		"category": SearchParamInfo{
			Resource: "LabResultObservationRecord",
			Name:     "category",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "category", Type: "string"},
			},
		},
	},

	"PersonalCharacteristicsObservationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "PersonalCharacteristicsObservationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "PersonalCharacteristicsObservationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "PersonalCharacteristicsObservationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
	},

	"ProcedureRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ProcedureRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "ProcedureRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "ProcedureRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
	},
	"SleepRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "SleepRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "SleepRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "SleepRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
	},

	"MealRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "MealRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "MealRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "MealRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
	},

	"GoalRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "GoalRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "GoalRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "GoalRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
	},

	"ActivityRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ActivityRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "ActivityRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "ActivityRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"activityType": SearchParamInfo{
			Resource: "ActivityRecord",
			Name:     "activityType",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "activityType", Type: "string"},
			},
		},
	},

	"AdverseEventRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "AdverseEventRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "AdverseEventRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "AdverseEventRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
	},
	"ObservationRecord": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ObservationRecord",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "ObservationRecord",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "ObservationRecord",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.consumerID", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "ObservationRecord",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.name", Type: "string"},
			},
		},
		"source": SearchParamInfo{
			Resource: "ObservationRecord",
			Name:     "source",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "healthrecord.source", Type: "string"},
			},
		},
	},
	"Notification": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Notification",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "Notification",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "consumerID", Type: "string"},
			},
		},
		"status": SearchParamInfo{
			Resource: "Notification",
			Name:     "status",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "status", Type: "string"},
			},
		},
	},
	"Acknowledgement": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Acknowledgement",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "Acknowledgement",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "consumerID", Type: "string"},
			},
		},
		"status": SearchParamInfo{
			Resource: "Acknowledgement",
			Name:     "refrenceNotification",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "refrenceNotification", Type: "string"},
			},
		},
	},
	"Product": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Product",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"supplier": SearchParamInfo{
			Resource: "Product",
			Name:     "supplier",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "supplier", Type: "string"},
			},
		},
		"vendor": SearchParamInfo{
			Resource: "Product",
			Name:     "vendor",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "vendor", Type: "string"},
			},
		},
		"category": SearchParamInfo{
			Resource: "Product",
			Name:     "category",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "category", Type: "string"},
			},
		},
		"language": SearchParamInfo{
			Resource: "Product",
			Name:     "language",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "language", Type: "string"},
			},
		},
		"currency": SearchParamInfo{
			Resource: "Product",
			Name:     "currency",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "currency", Type: "string"},
			},
		},
		"label": SearchParamInfo{
			Resource: "Product",
			Name:     "label",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "label", Type: "string"},
			},
		},
	},
	"Order": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Order",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"supplier": SearchParamInfo{
			Resource: "Order",
			Name:     "supplier",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "supplier", Type: "string"},
			},
		},
		"orderedItem": SearchParamInfo{
			Resource: "Order",
			Name:     "orderedItem",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "orderedItem", Type: "string"},
			},
		},
		"fromID": SearchParamInfo{
			Resource: "Order",
			Name:     "actorID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "from.actorID", Type: "string"},
			},
		},
		"toID": SearchParamInfo{
			Resource: "Order",
			Name:     "actorID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "to.actorID", Type: "string"},
			},
		},
		"status": SearchParamInfo{
			Resource: "Order",
			Name:     "status",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "status", Type: "string"},
			},
		},
	},
	"OrderEvent": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "OrderEvent",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"orderID": SearchParamInfo{
			Resource: "OrderEvent",
			Name:     "orderID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "orderID", Type: "string"},
			},
		},
		"externalID": SearchParamInfo{
			Resource: "OrderEvent",
			Name:     "externalID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "externalID", Type: "string"},
			},
		},
		"type": SearchParamInfo{
			Resource: "OrderEvent",
			Name:     "type",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "type", Type: "string"},
			},
		},
		"fromID": SearchParamInfo{
			Resource: "OrderEvent",
			Name:     "actorID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "from.actorID", Type: "string"},
			},
		},
		"toID": SearchParamInfo{
			Resource: "OrderEvent",
			Name:     "actorID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "to.actorID", Type: "string"},
			},
		},
	},
	"ReferenceActor": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ReferenceActor",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"active": SearchParamInfo{
			Resource: "ReferenceActor",
			Name:     "type",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "type", Type: "string"},
			},
		},
		"actorID": SearchParamInfo{
			Resource: "Relationship",
			Name:     "actorID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "actorID", Type: "string"},
			},
		},
	},
	"Relationship": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Relationship",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"_lastUpdated": SearchParamInfo{
			Resource: "Relationship",
			Name:     "_lastUpdated",
			Type:     "date",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "meta.lastUpdated", Type: "instant"},
			},
		},
		"fromID": SearchParamInfo{
			Resource: "Relationship",
			Name:     "actorID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "from.actorID", Type: "string"},
			},
		},
		"toID": SearchParamInfo{
			Resource: "Relationship",
			Name:     "actorID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "to.actorID", Type: "string"},
			},
		},
		"fromType": SearchParamInfo{
			Resource: "Relationship",
			Name:     "fromType",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "from.actorType", Type: "string"},
			},
		},
		"toType": SearchParamInfo{
			Resource: "Relationship",
			Name:     "toType",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "to.actorType", Type: "string"},
			},
		},
		"active": SearchParamInfo{
			Resource: "Relationship",
			Name:     "active",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "active", Type: "boolean"},
			},
		},
		"label": SearchParamInfo{
			Resource: "Relationship",
			Name:     "label",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "label", Type: "string"},
			},
		},
		"type": SearchParamInfo{
			Resource: "Relationship",
			Name:     "type",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "type", Type: "string"},
			},
		},
	},
	"Consent": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Consent",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"name": SearchParamInfo{
			Resource: "Consent",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "name", Type: "string"},
			},
		},
		"category": SearchParamInfo{
			Resource: "Consent",
			Name:     "category",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "category", Type: "string"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "Consent",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "consumerID", Type: "string"},
			},
		},
	},

	"Questionnaire": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Questionnaire",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"purpose": SearchParamInfo{
			Resource: "Questionnaire",
			Name:     "purpose",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "purpose", Type: "string"},
			},
		},
		"publisher": SearchParamInfo{
			Resource: "Questionnaire",
			Name:     "publisher",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "publisher", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "Questionnaire",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "name", Type: "string"},
			},
		},
		"language": SearchParamInfo{
			Resource: "Questionnaire",
			Name:     "language",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "language", Type: "string"},
			},
		},
	},

	"QuestionnaireResponse": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "QuestionnaireResponse",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"consumerID": SearchParamInfo{
			Resource: "QuestionnaireResponse",
			Name:     "consumerID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "consumerID", Type: "string"},
			},
		},
	},

	"ObservationDefinitionCollection": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ObservationDefinitionCollection",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"publisher": SearchParamInfo{
			Resource: "ObservationDefinitionCollection",
			Name:     "publisher",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "publisher", Type: "string"},
			},
		},
		"name": SearchParamInfo{
			Resource: "ObservationDefinitionCollection",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "name", Type: "string"},
			},
		},
		"purpose": SearchParamInfo{
			Resource: "ObservationDefinitionCollection",
			Name:     "purpose",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "purpose", Type: "string"},
			},
		},
		"language": SearchParamInfo{
			Resource: "ObservationDefinitionCollection",
			Name:     "language",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "language", Type: "string"},
			},
		},
	},
	"Message": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "Message",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"fromID": SearchParamInfo{
			Resource: "Message",
			Name:     "actorID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "from.actorID", Type: "string"},
			},
		},
		"toID": SearchParamInfo{
			Resource: "Message",
			Name:     "toID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "to.actorID", Type: "string"},
			},
		},
	},
	"DocumentReference": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "DocumentReference",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"custodian": SearchParamInfo{
			Resource: "DocumentReference",
			Name:     "custodian",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "custodian.actorID", Type: "string"},
			},
		},
		"class": SearchParamInfo{
			Resource: "DocumentReference",
			Name:     "class",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "class", Type: "string"},
			},
		},
		"type": SearchParamInfo{
			Resource: "DocumentReference",
			Name:     "type",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "type", Type: "string"},
			},
		},
		"language": SearchParamInfo{
			Resource: "DocumentReference",
			Name:     "language",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "language", Type: "string"},
			},
		},
	},
	"ConceptClass": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "ConceptClass",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"name": SearchParamInfo{
			Resource: "ConceptClass",
			Name:     "name",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "name", Type: "string"},
			},
		},
		"externalID": SearchParamInfo{
			Resource: "ConceptClass",
			Name:     "externalID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "externalID", Type: "string"},
			},
		},
	},

	"CustomerFeedback": map[string]SearchParamInfo{
		"_id": SearchParamInfo{
			Resource: "CustomerFeedback",
			Name:     "_id",
			Type:     "token",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "_id", Type: "id"},
			},
		},
		"applicationID": SearchParamInfo{
			Resource: "CustomerFeedback",
			Name:     "applicationID",
			Type:     "string",
			Paths: []SearchParamPath{
				SearchParamPath{Path: "application.entityID", Type: "string"},
			},
		},
	},
}
