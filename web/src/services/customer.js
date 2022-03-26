import {CUSTOMERS} from '@/services/api'
import {request, METHOD} from '@/utils/request'


export async function list() {
  return request(CUSTOMERS, METHOD.GET)
}


export default {
  list
}
