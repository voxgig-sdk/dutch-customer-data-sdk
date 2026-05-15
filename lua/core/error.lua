-- DutchCustomerData SDK error

local DutchCustomerDataError = {}
DutchCustomerDataError.__index = DutchCustomerDataError


function DutchCustomerDataError.new(code, msg, ctx)
  local self = setmetatable({}, DutchCustomerDataError)
  self.is_sdk_error = true
  self.sdk = "DutchCustomerData"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function DutchCustomerDataError:error()
  return self.msg
end


function DutchCustomerDataError:__tostring()
  return self.msg
end


return DutchCustomerDataError
