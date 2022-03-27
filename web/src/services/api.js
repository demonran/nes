//跨域代理前缀
const BASE_URL = process.env.VUE_APP_API_BASE_URL
module.exports = {
  LOGIN: `${BASE_URL}/login`,
  ROUTES: `${BASE_URL}/routes`,
  CUSTOMERS: `${BASE_URL}/api/customers`
}
