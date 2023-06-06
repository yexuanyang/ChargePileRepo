import { defineStore } from 'pinia'

export const ChargeStore = defineStore('chargeRequest', {
  state: () => {
    return {
      chargeRequest: {
        car_id: '',
        mode: '',
        apply_kwh: 0,
        apply_time: 0,
        stationId: 1,
      }
    }
  },
  getters: {},
  actions: {
    setChargeInfo(info) {
      this.$patch({
        chargeRequest: info
      })
      localStorage.setItem('chargeRequest', JSON.stringify(info))
    },
    loadChargeInfo() {
      const chargeRequestString = localStorage.getItem('chargeRequest')
      if (chargeRequestString !== null) {
        const chargeRequest = JSON.parse(chargeRequestString)
        this.chargeRequest = chargeRequest
      }
    },
  }
})
