package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DigitalProductPassport struct {
	ID                       primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductOverview          ProductOverview          `json:"productOverview" bson:"productOverview"`
	Reparability             Reparability             `json:"reparability" bson:"reparability"`
	EnvironmentalImpact      EnvironmentalImpact      `json:"environmentalImpact" bson:"environmentalImpact"`
	ComplianceAndStandards   ComplianceAndStandards   `json:"complianceAndStandards" bson:"complianceAndStandards"`
	Certificates             Certificates             `json:"certificates" bson:"certificates"`
	Recyclability            Recyclability            `json:"recyclability" bson:"recyclability"`
	EnergyUseAndEfficiency   EnergyUseAndEfficiency   `json:"energyUseAndEfficiency" bson:"energyUseAndEfficiency"`
	Components               []ComponentInformation   `json:"components" bson:"components"`
	EconomicOperator         EconomicOperator         `json:"economicOperator" bson:"economicOperator"`
	RepairInformation        RepairInformation        `json:"repairInformation" bson:"repairInformation"`
	RefurbishmentInformation RefurbishmentInformation `json:"refurbishmentInformation" bson:"refurbishmentInformation"`
	RecyclingInformation     RecyclingInformation     `json:"recyclingInformation" bson:"recyclingInformation"`
}

type ProductOverview struct {
	BrandName                     string `json:"brandName" bson:"brandName"`
	ProductImage                  string `json:"productImage" bson:"productImage"`
	GlobalProductClassificationCode string `json:"globalProductClassificationCode" bson:"globalProductClassificationCode"`
	CountryOfSale                 string `json:"countryOfSale" bson:"countryOfSale"`
	ProductDescription            string `json:"productDescription" bson:"productDescription"`
	ProductName                   string `json:"productName" bson:"productName"`
	NetWeight                     string `json:"netWeight" bson:"netWeight"`
	Gtin                          string `json:"gtin" bson:"gtin"`
	Color                         string `json:"color" bson:"color"`
	CountryOfOrigin               string `json:"countryOfOrigin" bson:"countryOfOrigin"`
	Dimensions                    string `json:"dimensions" bson:"dimensions"`
	ModelName                     string `json:"modelName" bson:"modelName"`
	TaricCode                     string `json:"taricCode" bson:"taricCode"`
	ConditionOfTheProduct         string `json:"conditionOfTheProduct" bson:"conditionOfTheProduct"`
	NetContent                    string `json:"netContent" bson:"netContent"`
	NominalMaximumRPM             string `json:"nominalMaximumRPM" bson:"nominalMaximumRPM"`
	MaximumDrillingDiameter       string `json:"maximumDrillingDiameter" bson:"maximumDrillingDiameter"`
	NumberOfGears                 string `json:"numberOfGears" bson:"numberOfGears"`
	Torque                        string `json:"torque" bson:"torque"`
	WarrantyDuration              string `json:"warrantyDuration" bson:"warrantyDuration"`
	SafetyInstructions            string `json:"safetyInstructions" bson:"safetyInstructions"`
}

type Reparability struct {
	ServiceAndRepairInstructions string `json:"serviceAndRepairInstructions" bson:"serviceAndRepairInstructions"`
	AvailabilityOfSpareParts     string `json:"availabilityOfSpareParts" bson:"availabilityOfSpareParts"`
}

type EnvironmentalImpact struct {
	WaterConsumptionPerUnit  string `json:"waterConsumptionPerUnit" bson:"waterConsumptionPerUnit"`
	ChemicalConsumptionPerUnit string `json:"chemicalConsumptionPerUnit" bson:"chemicalConsumptionPerUnit"`
	Co2eEmissionsPerUnit     string `json:"co2eEmissionsPerUnit" bson:"co2eEmissionsPerUnit"`
	EnergyConsumptionPerUnit string `json:"energyConsumptionPerUnit" bson:"energyConsumptionPerUnit"`
}

type ComplianceAndStandards struct {
	CeMarking      string `json:"ceMarking" bson:"ceMarking"`
	RohsCompliance string `json:"rohsCompliance" bson:"rohsCompliance"`
}

type Certificates struct {
	NameOfCertificate string `json:"nameOfCertificate" bson:"nameOfCertificate"`
}

type Recyclability struct {
	RecyclingInstructions string `json:"recyclingInstructions" bson:"recyclingInstructions"`
	MaterialComposition   string `json:"materialComposition" bson:"materialComposition"`
	SubstancesOfConcern   string `json:"substancesOfConcern" bson:"substancesOfConcern"`
}

type EnergyUseAndEfficiency struct {
	BatteryType           string `json:"batteryType" bson:"batteryType"`
	BatteryChargingTime   string `json:"batteryChargingTime" bson:"batteryChargingTime"`
	BatteryLife           string `json:"batteryLife" bson:"batteryLife"`
	ChargerType           string `json:"chargerType" bson:"chargerType"`
	MaximumElectricalPower string `json:"maximumElectricalPower" bson:"maximumElectricalPower"`
	MaximumVoltage        string `json:"maximumVoltage" bson:"maximumVoltage"`
	MaximumCurrent        string `json:"maximumCurrent" bson:"maximumCurrent"`
	PowerRating           string `json:"powerRating" bson:"powerRating"`
	DcVoltage             string `json:"dcVoltage" bson:"dcVoltage"`
}

type ComponentInformation struct {
	ComponentDescription string `json:"componentDescription" bson:"componentDescription"`
	ComponentGTIN        string `json:"componentGTIN" bson:"componentGTIN"`
	LinkToDPP            string `json:"linkToDPP" bson:"linkToDPP"`
}

type EconomicOperator struct {
	CompanyName      string `json:"companyName" bson:"companyName"`
	Gln              string `json:"gln" bson:"gln"`
	EoriNumber       string `json:"eoriNumber" bson:"eoriNumber"`
	AddressLine1     string `json:"addressLine1" bson:"addressLine1"`
	AddressLine2     string `json:"addressLine2" bson:"addressLine2"`
	ContactInformation string `json:"contactInformation" bson:"contactInformation"`
}

type RepairInformation struct {
	ReasonForRepair string `json:"reasonForRepair" bson:"reasonForRepair"`
	PerformedAction string `json:"performedAction" bson:"performedAction"`
	MaterialsUsed   string `json:"materialsUsed" bson:"materialsUsed"`
	DateOfRepair    string `json:"dateOfRepair" bson:"dateOfRepair"`
}

type RefurbishmentInformation struct {
	PerformedAction     string `json:"performedAction" bson:"performedAction"`
	MaterialsUsed       string `json:"materialsUsed" bson:"materialsUsed"`
	DateOfRefurbishment string `json:"dateOfRefurbishment" bson:"dateOfRefurbishment"`
}

type RecyclingInformation struct {
	PerformedAction string `json:"performedAction" bson:"performedAction"`
	DateOfRecycling string `json:"dateOfRecycling" bson:"dateOfRecycling"`
}
