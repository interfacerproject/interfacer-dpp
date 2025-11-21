package model

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type TransformedValue struct {
	Type  string `json:"type" bson:"type"`
	Value any    `json:"value" bson:"value"`
	Units string `json:"units,omitempty" bson:"units,omitempty"`
}

type Attachment struct {
	ID          string    `bson:"id" json:"id"`
	FileName    string    `bson:"file_name" json:"fileName"`
	ContentType string    `bson:"content_type" json:"contentType"`
	URL         string    `bson:"url" json:"url"`
	Size        int64     `bson:"size" json:"size"`
	Checksum    string    `bson:"checksum" json:"checksum"`
	UploadedAt  time.Time `bson:"uploaded_at" json:"uploadedAt"`
}

type DigitalProductPassport struct {
	ID                       ulid.ULID                 `json:"id" bson:"_id"`
	ProductOverview          *ProductOverview          `json:"productOverview,omitempty" bson:"productOverview,omitempty"`
	Reparability             *Reparability             `json:"reparability,omitempty" bson:"reparability,omitempty"`
	EnvironmentalImpact      *EnvironmentalImpact      `json:"environmentalImpact,omitempty" bson:"environmentalImpact,omitempty"`
	ComplianceAndStandards   *ComplianceAndStandards   `json:"complianceAndStandards,omitempty" bson:"complianceAndStandards,omitempty"`
	Certificates             *Certificates             `json:"certificates,omitempty" bson:"certificates,omitempty"`
	Recyclability            *Recyclability            `json:"recyclability,omitempty" bson:"recyclability,omitempty"`
	EnergyUseAndEfficiency   *EnergyUseAndEfficiency   `json:"energyUseAndEfficiency,omitempty" bson:"energyUseAndEfficiency,omitempty"`
	Components               []ComponentInformation    `json:"components,omitempty" bson:"components,omitempty"`
	EconomicOperator         *EconomicOperator         `json:"economicOperator,omitempty" bson:"economicOperator,omitempty"`
	RepairInformation        *RepairInformation        `json:"repairInformation,omitempty" bson:"repairInformation,omitempty"`
	RefurbishmentInformation *RefurbishmentInformation `json:"refurbishmentInformation,omitempty" bson:"refurbishmentInformation,omitempty"`
	RecyclingInformation     *RecyclingInformation     `json:"recyclingInformation,omitempty" bson:"recyclingInformation,omitempty"`
	ConsumerInformation      *ConsumerInformation      `json:"consumerInformation,omitempty" bson:"consumerInformation,omitempty"`
	DosageInstructions       *DosageInstructions       `json:"dosageInstructions,omitempty" bson:"dosageInstructions,omitempty"`
	Ingredients              *Ingredients              `json:"ingredients,omitempty" bson:"ingredients,omitempty"`
	Packaging                *Packaging                `json:"packaging,omitempty" bson:"packaging,omitempty"`
}

type ProductOverview struct {
	BrandName                       *TransformedValue `json:"brandName,omitempty" bson:"brandName,omitempty"`
	ProductImage                    *TransformedValue `json:"productImage,omitempty" bson:"productImage,omitempty"`
	GlobalProductClassificationCode *TransformedValue `json:"globalProductClassificationCode,omitempty" bson:"globalProductClassificationCode,omitempty"`
	CountryOfSale                   *TransformedValue `json:"countryOfSale,omitempty" bson:"countryOfSale,omitempty"`
	ProductDescription              *TransformedValue `json:"productDescription,omitempty" bson:"productDescription,omitempty"`
	ProductName                     *TransformedValue `json:"productName,omitempty" bson:"productName,omitempty"`
	NetWeight                       *TransformedValue `json:"netWeight,omitempty" bson:"netWeight,omitempty"`
	Gtin                            *TransformedValue `json:"gtin,omitempty" bson:"gtin,omitempty"`
	Color                           *TransformedValue `json:"color,omitempty" bson:"color,omitempty"`
	CountryOfOrigin                 *TransformedValue `json:"countryOfOrigin,omitempty" bson:"countryOfOrigin,omitempty"`
	Dimensions                      *TransformedValue `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
	ModelName                       *TransformedValue `json:"modelName,omitempty" bson:"modelName,omitempty"`
	TaricCode                       *TransformedValue `json:"taricCode,omitempty" bson:"taricCode,omitempty"`
	ConditionOfTheProduct           *TransformedValue `json:"conditionOfTheProduct,omitempty" bson:"conditionOfTheProduct,omitempty"`
	NetContent                      *TransformedValue `json:"netContent,omitempty" bson:"netContent,omitempty"`
	NominalMaximumRPM               *TransformedValue `json:"nominalMaximumRPM,omitempty" bson:"nominalMaximumRPM,omitempty"`
	MaximumDrillingDiameter         *TransformedValue `json:"maximumDrillingDiameter,omitempty" bson:"maximumDrillingDiameter,omitempty"`
	NumberOfGears                   *TransformedValue `json:"numberOfGears,omitempty" bson:"numberOfGears,omitempty"`
	Torque                          *TransformedValue `json:"torque,omitempty" bson:"torque,omitempty"`
	WarrantyDuration                *TransformedValue `json:"warrantyDuration,omitempty" bson:"warrantyDuration,omitempty"`
	SafetyInstructions              *TransformedValue `json:"safetyInstructions,omitempty" bson:"safetyInstructions,omitempty"`
	ConsumerUnit                    *TransformedValue `json:"consumerUnit,omitempty" bson:"consumerUnit,omitempty"`
	NetContentAndUnitOfMeasure      *TransformedValue `json:"netContentAndUnitOfMeasure,omitempty" bson:"netContentAndUnitOfMeasure,omitempty"`
	YearOfSale                      *TransformedValue `json:"yearOfSale,omitempty" bson:"yearOfSale,omitempty"`
}

type Reparability struct {
	ServiceAndRepairInstructions *TransformedValue `json:"serviceAndRepairInstructions,omitempty" bson:"serviceAndRepairInstructions,omitempty"`
	AvailabilityOfSpareParts     *TransformedValue `json:"availabilityOfSpareParts,omitempty" bson:"availabilityOfSpareParts,omitempty"`
}

type EnvironmentalImpact struct {
	WaterConsumptionPerUnit                                 *TransformedValue `json:"waterConsumptionPerUnit,omitempty" bson:"waterConsumptionPerUnit,omitempty"`
	ChemicalConsumptionPerUnit                              *TransformedValue `json:"chemicalConsumptionPerUnit,omitempty" bson:"chemicalConsumptionPerUnit,omitempty"`
	Co2eEmissionsPerUnit                                    *TransformedValue `json:"co2eEmissionsPerUnit,omitempty" bson:"co2eEmissionsPerUnit,omitempty"`
	EnergyConsumptionPerUnit                                *TransformedValue `json:"energyConsumptionPerUnit,omitempty" bson:"energyConsumptionPerUnit,omitempty"`
	CleaningPerformanceAtLowTemperature                     *TransformedValue `json:"cleaningPerformanceAtLowTemperature,omitempty" bson:"cleaningPerformanceAtLowTemperature,omitempty"`
	MinimumContentOfMaterialWithSustainabilityCertification *TransformedValue `json:"minimumContentOfMaterialWithSustainabilityCertification,omitempty" bson:"minimumContentOfMaterialWithSustainabilityCertification,omitempty"`
}

type ComplianceAndStandards struct {
	CeMarking      *TransformedValue `json:"ceMarking,omitempty" bson:"ceMarking,omitempty"`
	RohsCompliance *TransformedValue `json:"rohsCompliance,omitempty" bson:"rohsCompliance,omitempty"`
}

type Certificates struct {
	NameOfCertificate *TransformedValue `json:"nameOfCertificate,omitempty" bson:"nameOfCertificate,omitempty"`
}

type Recyclability struct {
	RecyclingInstructions *TransformedValue `json:"recyclingInstructions,omitempty" bson:"recyclingInstructions,omitempty"`
	MaterialComposition   *TransformedValue `json:"materialComposition,omitempty" bson:"materialComposition,omitempty"`
	SubstancesOfConcern   *TransformedValue `json:"substancesOfConcern,omitempty" bson:"substancesOfConcern,omitempty"`
}

type EnergyUseAndEfficiency struct {
	BatteryType            *TransformedValue `json:"batteryType,omitempty" bson:"batteryType,omitempty"`
	BatteryChargingTime    *TransformedValue `json:"batteryChargingTime,omitempty" bson:"batteryChargingTime,omitempty"`
	BatteryLife            *TransformedValue `json:"batteryLife,omitempty" bson:"batteryLife,omitempty"`
	ChargerType            *TransformedValue `json:"chargerType,omitempty" bson:"chargerType,omitempty"`
	MaximumElectricalPower *TransformedValue `json:"maximumElectricalPower,omitempty" bson:"maximumElectricalPower,omitempty"`
	MaximumVoltage         *TransformedValue `json:"maximumVoltage,omitempty" bson:"maximumVoltage,omitempty"`
	MaximumCurrent         *TransformedValue `json:"maximumCurrent,omitempty" bson:"maximumCurrent,omitempty"`
	PowerRating            *TransformedValue `json:"powerRating,omitempty" bson:"powerRating,omitempty"`
	DcVoltage              *TransformedValue `json:"dcVoltage,omitempty" bson:"dcVoltage,omitempty"`
}

type ComponentInformation struct {
	ComponentDescription *TransformedValue `json:"componentDescription,omitempty" bson:"componentDescription,omitempty"`
	ComponentGTIN        *TransformedValue `json:"componentGTIN,omitempty" bson:"componentGTIN,omitempty"`
	LinkToDPP            *TransformedValue `json:"linkToDPP,omitempty" bson:"linkToDPP,omitempty"`
}

type EconomicOperator struct {
	CompanyName        *TransformedValue `json:"companyName,omitempty" bson:"companyName,omitempty"`
	Gln                *TransformedValue `json:"gln,omitempty" bson:"gln,omitempty"`
	EoriNumber         *TransformedValue `json:"eoriNumber,omitempty" bson:"eoriNumber,omitempty"`
	AddressLine1       *TransformedValue `json:"addressLine1,omitempty" bson:"addressLine1,omitempty"`
	AddressLine2       *TransformedValue `json:"addressLine2,omitempty" bson:"addressLine2,omitempty"`
	ContactInformation *TransformedValue `json:"contactInformation,omitempty" bson:"contactInformation,omitempty"`
}

type RepairInformation struct {
	ReasonForRepair *TransformedValue `json:"reasonForRepair,omitempty" bson:"reasonForRepair,omitempty"`
	PerformedAction *TransformedValue `json:"performedAction,omitempty" bson:"performedAction,omitempty"`
	MaterialsUsed   *TransformedValue `json:"materialsUsed,omitempty" bson:"materialsUsed,omitempty"`
	DateOfRepair    *TransformedValue `json:"dateOfRepair,omitempty" bson:"dateOfRepair,omitempty"`
}

type RefurbishmentInformation struct {
	PerformedAction     *TransformedValue `json:"performedAction,omitempty" bson:"performedAction,omitempty"`
	MaterialsUsed       *TransformedValue `json:"materialsUsed,omitempty" bson:"materialsUsed,omitempty"`
	DateOfRefurbishment *TransformedValue `json:"dateOfRefurbishment,omitempty" bson:"dateOfRefurbishment,omitempty"`
}

type RecyclingInformation struct {
	PerformedAction *TransformedValue `json:"performedAction,omitempty" bson:"performedAction,omitempty"`
	DateOfRecycling *TransformedValue `json:"dateOfRecycling,omitempty" bson:"dateOfRecycling,omitempty"`
}

type ConsumerInformation struct {
	MarketingClaim *TransformedValue `json:"marketingClaim,omitempty" bson:"marketingClaim,omitempty"`
}

type DosageInstructions struct {
	UsageAndDisposalInfo *TransformedValue `json:"usageAndDisposalInfo,omitempty" bson:"usageAndDisposalInfo,omitempty"`
}

type Ingredients struct {
	IngredientList                          *TransformedValue `json:"ingredientList,omitempty" bson:"ingredientList,omitempty"`
	MinimumContentOfBiodegradableSubstances *TransformedValue `json:"minimumContentOfBiodegradableSubstances,omitempty" bson:"minimumContentOfBiodegradableSubstances,omitempty"`
	PresenceOfNonBiodegradableMicroplastics *TransformedValue `json:"presenceOfNonBiodegradableMicroplastics,omitempty" bson:"presenceOfNonBiodegradableMicroplastics,omitempty"`
}

type ChemicalConsumption struct {
	Amount     *TransformedValue `json:"amount,omitempty" bson:"amount,omitempty"`
	Ingredient *TransformedValue `json:"ingredient,omitempty" bson:"ingredient,omitempty"`
}

type Packaging struct {
	ChemicalConsumption    *ChemicalConsumption `json:"chemicalConsumption,omitempty" bson:"chemicalConsumption,omitempty"`
	DisposalInstructions   *TransformedValue    `json:"disposalInstructions,omitempty" bson:"disposalInstructions,omitempty"`
	MinimumRecycledContent *TransformedValue    `json:"minimumRecycledContent,omitempty" bson:"minimumRecycledContent,omitempty"`
	RecyclablePackaging    *TransformedValue    `json:"recyclablePackaging,omitempty" bson:"recyclablePackaging,omitempty"`
}
