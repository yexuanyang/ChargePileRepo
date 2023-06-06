import {defineStore} from 'pinia'

export const ActiveStore = defineStore('active', {
  state: () => {
    return {
      active: 0
    }
  },
  getters: {},
  actions: {
    // 返回上一步
    lastStep() {
      this.active--
      localStorage.setItem('active', this.active)
    },
    // 前往下一步
    nextStep() {
      this.active++
      localStorage.setItem('active', this.active)
    },
    // 重新提交订单
    backHome() {
      this.active = 0
      localStorage.setItem('active', this.active)
    },
    // 加载active
    loadActive() {
      const activeValue = localStorage.getItem('active')
      if (activeValue !== null) {
        this.active = parseInt(activeValue)
      }
    }
  }
})
