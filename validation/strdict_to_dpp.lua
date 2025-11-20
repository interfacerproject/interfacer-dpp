-- This script takes a string dictionary from entry, recognizes known
-- DPP fields and turns them into more descriptives data entries.

data_types   = JSON.decode(KEYS)
DPP_examples = JSON.decode(DATA)

-- Find out if the key name is part of a known type (category)
function istype(arr, obj)
    if luatype(arr) ~= 'table' then error("Internal error: array_contains argument is not a table",2) end
    for k, v in pairs(arr) do
        if luatype(k) ~= 'number' then error("Internal error: array_contains argument is not an array", 2) end
        if obj:lower():match(v:lower()) then return true end
    end
    return false
end

-- Parse strings like:
-- "577 Wh", "577Wh", "652.0 kWh", "652,0kWh", "0.65 Watt/m²K"
local function parse_number_unit(s)
   local num, unit = s:match("^(%d+[.,]?%d*)%s*(.+)$")
    if num and unit then
        return { tonumber(num), unit }
    end
    return {s,'parser error'}
end

-- Parse strings like:
-- "20x15x5 cm", "14.1 x 18.8 x 32.9 cm", "1400 x 1200 mm"
local function parse_dimensions(s)
    local dims, unit = s:match("^(.-)%s*([%a²µ]+)$")
    if not dims then return nil end
    local results = {}
    for num in dims:gmatch("(%d+[.,]?%d*)") do
      table.insert(results, tonumber(num))
    end
    return { value = results, units = unit }
end
-- May render back simply with:
-- print(table.concat(value, "x")) --> 20x15x5

-- Parse a value into a boolean.
-- Supports: true/false, "true"/"false", "yes"/"no", "1"/"0", 1/0
local function parse_boolean(v)
    if type(v) == "boolean" then
        return v
    elseif type(v) == "number" then
        return v ~= 0
    elseif type(v) == "string" then
        local s = v:lower():gsub("^%s*(.-)%s*$", "%1") -- trim
        if s == "true" or s == "yes" or s == "1" then
            return true
        elseif s == "false" or s == "no" or s == "0" then
            return false
        end
    end
    return 'parse error'
end


function parse(KNOWN, IN)

  local function recon(v, k)
    for kk,vv in pairs(data_types) do
      if istype(vv, k) then
        if kk == 'quantity' then
          local toks <const> = parse_number_unit(v)
          return({type = kk, value = toks[1], units = toks[2]})
        elseif kk == 'dimensions' then
          local toks <const> = parse_dimensions(v)
          return({type = kk, value = toks.value, units = toks.units})
        elseif kk == 'boolean' then
          return({type = kk, value = parse_boolean(v)})
        end
      end
    end
    if v == 'N/A' then return nil end
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

-- writerr(I.print(DATA))
RES = { }
for _,v in pairs(DPP_examples) do
  RES[v.productOverview.brandName] = parse(data_types, v)
end

print(JSON.encode(RES))


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
