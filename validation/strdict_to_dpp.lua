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

  local function recon(v, k, allData)
    -- Skip nil values
    if v == nil or v == 'N/A' then return nil end
    
    -- Skip unit fields as they are handled with their corresponding value fields
    if k:match("Unit$") then
      return nil
    end
    
    -- Check if this is a field that has a separate unit field
    local unitFieldName = k .. "Unit"
    if allData[unitFieldName] then
      -- This field has a separate unit field, combine them
      return({type = 'quantity', value = tonumber(v) or v, units = allData[unitFieldName]})
    end
    
    for kk,vv in pairs(data_types) do
      if istype(vv, k) then
        if kk == 'quantity' then
          local toks <const> = parse_number_unit(v)
          return({type = kk, value = toks[1], units = toks[2]})
        elseif kk == 'dimensions' then
          local toks <const> = parse_dimensions(v)
          if toks then
            return({type = kk, value = toks.value, units = toks.units})
          else
            return({type = 'text', value = v})
          end
        elseif kk == 'boolean' then
          return({type = kk, value = parse_boolean(v)})
        elseif kk == 'date' then
          return({type = kk, value = v})
        elseif kk == 'link' then
          return({type = kk, value = v})
        end
      end
    end
    return({ type = 'text', value = v })
  end -- recon()

  -- Map input category names to Go model field names
  local categoryMapping = {
    compliance = 'complianceAndStandards',
    repairability = 'reparability',
    energyUseEfficiency = 'energyUseAndEfficiency'
  }
  
  -- Map field names to Go model field names
  local fieldMapping = {
    -- ProductOverview fields
    gpc = 'globalProductClassificationCode',
    
    -- EnvironmentalImpact fields
    waterConsumption = 'waterConsumptionPerUnit',
    chemicalConsumption = 'chemicalConsumptionPerUnit',
    co2Emissions = 'co2eEmissionsPerUnit',
    energyConsumption = 'energyConsumptionPerUnit',
    
    -- Certificates fields
    certificateName = 'nameOfCertificate',
    
    -- Reparability fields
    spareParts = 'availabilityOfSpareParts'
  }

  local OUT = { }
  for k,v in pairs(IN) do
    -- Map category name if needed
    local outputKey = categoryMapping[k] or k
    
    -- Handle all known DPP top-level categories
    if k == 'componentInformation' and isarray(v) then
      -- componentInformation is an array
      OUT.componentInformation = { }
      for _,item in ipairs(v) do
        local processed = {}
        for fieldKey, fieldValue in pairs(item) do
          local mappedFieldKey = fieldMapping[fieldKey] or fieldKey
          local result = recon(fieldValue, mappedFieldKey, item)
          if result ~= nil then
            processed[mappedFieldKey] = result
          end
        end
        table.insert(OUT.componentInformation, processed)
      end
    elseif k == 'components' and isarray(v) then
      -- Legacy support for 'components' key
      OUT.components = { }
      for _,item in ipairs(v) do
        local processed = {}
        for fieldKey, fieldValue in pairs(item) do
          local mappedFieldKey = fieldMapping[fieldKey] or fieldKey
          local result = recon(fieldValue, mappedFieldKey, item)
          if result ~= nil then
            processed[mappedFieldKey] = result
          end
        end
        table.insert(OUT.components, processed)
      end
    elseif k == 'productOverview' or 
           k == 'repairability' or 
           k == 'environmentalImpact' or 
           k == 'compliance' or 
           k == 'certificates' or 
           k == 'recyclability' or 
           k == 'energyUseEfficiency' or 
           k == 'economicOperator' or 
           k == 'repairInformation' or 
           k == 'refurbishmentInformation' or 
           k == 'recyclingInformation' or
           k == 'consumerInformation' or
           k == 'dosageInstructions' or
           k == 'ingredients' or
           k == 'packaging' then
      -- Known DPP categories (objects)
      OUT[outputKey] = {}
      for fieldKey, fieldValue in pairs(v) do
        local mappedFieldKey = fieldMapping[fieldKey] or fieldKey
        local result = recon(fieldValue, mappedFieldKey, v)
        if result ~= nil then
          OUT[outputKey][mappedFieldKey] = result
        end
      end
    else
      -- Unknown/legacy categories - still process them
      OUT[outputKey] = {}
      for fieldKey, fieldValue in pairs(v) do
        local mappedFieldKey = fieldMapping[fieldKey] or fieldKey
        local result = recon(fieldValue, mappedFieldKey, v)
        if result ~= nil then
          OUT[outputKey][mappedFieldKey] = result
        end
      end
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
