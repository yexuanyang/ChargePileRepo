import service from '@/utils/request'

export const getChargePileList = () => {
  return service({
    url: '/chargePile/getChargePileList',
    method: 'POST',
  })
}
