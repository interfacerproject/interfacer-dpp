-- This script takes a string dictionary from entry, recognizes known
-- DPP fields and turns them into more descriptives data entries.

data_types   = JSON.decode(KEYS)
DPP_examples = JSON.decode(DATA)

function parse(KNOWN, IN)

  local function recon(v, k)
    for kk,vv in pairs(data_types) do
      if array_contains(vv, k) then
        return({type = kk, value = v})
      end
    end
    return({ type = 'text', value = v })
  end -- recon()

  local OUT = { }
  for k,v in pairs(IN) do
    if k == 'components' and isarray(v) then
      -- L1 array
      OUT.components = { }
      for _,v in ipairs(v) do
        -- L2 dictionary
        table.insert(OUT.components, deepmap(recon, v))
      end
    else
      -- L1 dictionary
      OUT[k] = deepmap(recon, v)
    end
  end -- for each category
  return(OUT)
end -- parse()

write(JSON.encode(parse(data_types, DPP_examples)))


-- {
--   "_id": "68e4d76a936dbbc5783e52b5",
--   "certificates": {
--     "nameOfCertificate": "TUV Rheinland GS Mark"
--   },
--   "complianceAndStandards": {
--     "ceMarking": "Yes",
--     "rohsCompliance": "Yes"
--   },
--   "components": [
--     {
--       "componentDescription": "18V 2.0Ah Battery Pack",
--       "componentGTIN": "98765432109876",
--       "linkToDPP": "http://example.com/dpp/battery"
--     }
--   ],
--   "economicOperator": {
--     "addressLine1": "123 Tool Street",
--     "addressLine2": "Toolsville, 12345",
--     "companyName": "Global Tools Inc.",
--     "contactInformation": "support@globaltools.com",
--     "eoriNumber": "DE1234567",
--     "gln": "1122334455667"
--   },
--   "energyUseAndEfficiency": {
--     "batteryChargingTime": "60 minutes",
--     "batteryLife": "500 cycles",
--     "batteryType": "Lithium-Ion",
--     "chargerType": "Fast Charger Model C-200",
--     "dcVoltage": "18V",
--     "maximumCurrent": "25A",
--     "maximumElectricalPower": "500W",
--     "maximumVoltage": "20V",
--     "powerRating": "400W"
--   },
--   "environmentalImpact": {
--     "chemicalConsumptionPerUnit": "0.1 g",
--     "co2eEmissionsPerUnit": "5 kg CO2e",
--     "energyConsumptionPerUnit": "0.5 kWh",
--     "waterConsumptionPerUnit": "0 L"
--   },
--   "productOverview": {
--     "brandName": "PowerDrill Pro",
--     "color": "Blue/Black",
--     "conditionOfTheProduct": "New",
--     "countryOfOrigin": "China",
--     "countryOfSale": "Germany",
--     "dimensions": "20x15x5 cm",
--     "globalProductClassificationCode": "12345678",
--     "gtin": "01234567890123",
--     "maximumDrillingDiameter": "13mm (steel), 30mm (wood)",
--     "modelName": "X-1000",
--     "netContent": "1 unit",
--     "netWeight": "1.5 kg",
--     "nominalMaximumRPM": "2000 RPM",
--     "numberOfGears": "2",
--     "productDescription": "A high-performance cordless drill.",
--     "productImage": "http://example.com/drill.jpg",
--     "productName": "Cordless Drill X-1000",
--     "safetyInstructions": "Always wear safety glasses.",
--     "taricCode": "8467211000",
--     "torque": "50 Nm",
--     "warrantyDuration": "2 years"
--   },
--   "recyclability": {
--     "materialComposition": "Plastic: 40%, Steel: 50%, Copper: 10%",
--     "recyclingInstructions": "Dispose of at an e-waste facility.",
--     "substancesOfConcern": "Lithium-Ion Battery"
--   },
--   "recyclingInformation": {
--     "dateOfRecycling": "2025-10-07T00:00:00Z",
--     "performedAction": "N/A"
--   },
--   "refurbishmentInformation": {
--     "dateOfRefurbishment": "2025-10-07T00:00:00Z",
--     "materialsUsed": "N/A",
--     "performedAction": "N/A"
--   },
--   "repairInformation": {
--     "dateOfRepair": "2025-10-07T00:00:00Z",
--     "materialsUsed": "N/A",
--     "performedAction": "N/A",
--     "reasonForRepair": "N/A"
--   },
--   "reparability": {
--     "availabilityOfSpareParts": "5 years from date of purchase",
--     "serviceAndRepairInstructions": "See manual at http://example.com/manual"
--   }
-- }
