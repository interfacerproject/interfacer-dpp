package model

import (
	"github.com/oklog/ulid/v2"
)

type TransformedValue struct {
	Type  string      `json:"type" bson:"type"`
	Value any `json:"value" bson:"value"`
	Units string      `json:"units,omitempty" bson:"units,omitempty"`
}

type DigitalProductPassport struct {
	ID                       ulid.ULID                `json:"id" bson:"_id"`
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
	BrandName                     TransformedValue `json:"brandName" bson:"brandName"`
	ProductImage                  TransformedValue `json:"productImage" bson:"productImage"`
	GlobalProductClassificationCode TransformedValue `json:"globalProductClassificationCode" bson:"globalProductClassificationCode"`
	CountryOfSale                 TransformedValue `json:"countryOfSale" bson:"countryOfSale"`
	ProductDescription            TransformedValue `json:"productDescription" bson:"productDescription"`
	ProductName                   TransformedValue `json:"productName" bson:"productName"`
	NetWeight                     TransformedValue `json:"netWeight" bson:"netWeight"`
	Gtin                          TransformedValue `json:"gtin" bson:"gtin"`
	Color                         TransformedValue `json:"color" bson:"color"`
	CountryOfOrigin               TransformedValue `json:"countryOfOrigin" bson:"countryOfOrigin"`
	Dimensions                    TransformedValue `json:"dimensions" bson:"dimensions"`
	ModelName                     TransformedValue `json:"modelName" bson:"modelName"`
	TaricCode                     TransformedValue `json:"taricCode" bson:"taricCode"`
	ConditionOfTheProduct         TransformedValue `json:"conditionOfTheProduct" bson:"conditionOfTheProduct"`
	NetContent                    TransformedValue `json:"netContent" bson:"netContent"`
	NominalMaximumRPM             TransformedValue `json:"nominalMaximumRPM" bson:"nominalMaximumRPM"`
	MaximumDrillingDiameter       TransformedValue `json:"maximumDrillingDiameter" bson:"maximumDrillingDiameter"`
	NumberOfGears                 TransformedValue `json:"numberOfGears" bson:"numberOfGears"`
	Torque                        TransformedValue `json:"torque" bson:"torque"`
	WarrantyDuration              TransformedValue `json:"warrantyDuration" bson:"warrantyDuration"`
	SafetyInstructions            TransformedValue `json:"safetyInstructions" bson:"safetyInstructions"`
}

type Reparability struct {
	ServiceAndRepairInstructions TransformedValue `json:"serviceAndRepairInstructions" bson:"serviceAndRepairInstructions"`
	AvailabilityOfSpareParts     TransformedValue `json:"availabilityOfSpareParts" bson:"availabilityOfSpareParts"`
}

type EnvironmentalImpact struct {
	WaterConsumptionPerUnit  TransformedValue `json:"waterConsumptionPerUnit" bson:"waterConsumptionPerUnit"`
	ChemicalConsumptionPerUnit TransformedValue `json:"chemicalConsumptionPerUnit" bson:"chemicalConsumptionPerUnit"`
	Co2eEmissionsPerUnit     TransformedValue `json:"co2eEmissionsPerUnit" bson:"co2eEmissionsPerUnit"`
	EnergyConsumptionPerUnit TransformedValue `json:"energyConsumptionPerUnit" bson:"energyConsumptionPerUnit"`
}

type ComplianceAndStandards struct {
	CeMarking      TransformedValue `json:"ceMarking" bson:"ceMarking"`
	RohsCompliance TransformedValue `json:"rohsCompliance" bson:"rohsCompliance"`
}

type Certificates struct {
	NameOfCertificate TransformedValue `json:"nameOfCertificate" bson:"nameOfCertificate"`
}

type Recyclability struct {
	RecyclingInstructions TransformedValue `json:"recyclingInstructions" bson:"recyclingInstructions"`
	MaterialComposition   TransformedValue `json:"materialComposition" bson:"materialComposition"`
	SubstancesOfConcern   TransformedValue `json:"substancesOfConcern" bson:"substancesOfConcern"`
}

type EnergyUseAndEfficiency struct {
	BatteryType           TransformedValue `json:"batteryType" bson:"batteryType"`
	BatteryChargingTime   TransformedValue `json:"batteryChargingTime" bson:"batteryChargingTime"`
	BatteryLife           TransformedValue `json:"batteryLife" bson:"batteryLife"`
	ChargerType           TransformedValue `json:"chargerType" bson:"chargerType"`
	MaximumElectricalPower TransformedValue `json:"maximumElectricalPower" bson:"maximumElectricalPower"`
	MaximumVoltage        TransformedValue `json:"maximumVoltage" bson:"maximumVoltage"`
	MaximumCurrent        TransformedValue `json:"maximumCurrent" bson:"maximumCurrent"`
	PowerRating           TransformedValue `json:"powerRating" bson:"powerRating"`
	DcVoltage             TransformedValue `json:"dcVoltage" bson:"dcVoltage"`
}

type ComponentInformation struct {
	ComponentDescription TransformedValue `json:"componentDescription" bson:"componentDescription"`
	ComponentGTIN        TransformedValue `json:"componentGTIN" bson:"componentGTIN"`
	LinkToDPP            TransformedValue `json:"linkToDPP" bson:"linkToDPP"`
}

type EconomicOperator struct {
	CompanyName      TransformedValue `json:"companyName" bson:"companyName"`
	Gln              TransformedValue `json:"gln" bson:"gln"`
	EoriNumber       TransformedValue `json:"eoriNumber" bson:"eoriNumber"`
	AddressLine1     TransformedValue `json:"addressLine1" bson:"addressLine1"`
	AddressLine2     TransformedValue `json:"addressLine2" bson:"addressLine2"`
	ContactInformation TransformedValue `json:"contactInformation" bson:"contactInformation"`
}

type RepairInformation struct {
	ReasonForRepair TransformedValue `json:"reasonForRepair" bson:"reasonForRepair"`
	PerformedAction TransformedValue `json:"performedAction" bson:"performedAction"`
	MaterialsUsed   TransformedValue `json:"materialsUsed" bson:"materialsUsed"`
	DateOfRepair    TransformedValue `json:"dateOfRepair" bson:"dateOfRepair"`
}

type RefurbishmentInformation struct {
	PerformedAction     TransformedValue `json:"performedAction" bson:"performedAction"`
	MaterialsUsed       TransformedValue `json:"materialsUsed" bson:"materialsUsed"`
	DateOfRefurbishment TransformedValue `json:"dateOfRefurbishment" bson:"dateOfRefurbishment"`
}

type RecyclingInformation struct {
	PerformedAction TransformedValue `json:"performedAction" bson:"performedAction"`
	DateOfRecycling TransformedValue `json:"dateOfRecycling" bson:"dateOfRecycling"`
}