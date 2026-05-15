
import { Context } from './Context'


class DutchCustomerDataError extends Error {

  isDutchCustomerDataError = true

  sdk = 'DutchCustomerData'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  DutchCustomerDataError
}

