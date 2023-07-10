local person = {
  name = "joe",
  age = 19,
}

local someone = {
  name = "siduck",
}

local result = vim.tbl_deep_extend("force", person, someone)

print(result)
