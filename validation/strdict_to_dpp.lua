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

RES = { }
for _,v in pairs(DPP_examples) do
  RES[v.productOverview.brandName] = parse(data_types, v)
end

print(JSON.encode(RES))
