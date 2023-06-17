<template>
  <div class="page">
    <div class="gva-card-box">
      <div v-if="active.active===0" class="gva-card">
        <div class="card-header">
          <span>提交订单</span>
        </div>
        <div class="gva-card gva-order-card">
          <el-form
            ref="elFormRef"
            :model="requestForm"
            label-position="right"
            :rules="rule"
            label-width="120px"
          >
            <el-form-item label="充电的车牌号:" prop="car_id">
              <el-input
                v-model="requestForm.car_id"
                style="width:100%"
                :clearable="true"
                placeholder="请输入"
              />
            </el-form-item>
            <el-form-item label="充电类型:" prop="mode">
              <el-select
                v-model="requestForm.mode"
                placeholder="请选择"
                style="width:60%"
                :clearable="true"
              >
                <el-option
                  v-for="item in modeOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="充电度数:" prop="apply_kwh">
              <el-input-number
                v-model="requestForm.apply_kwh"
                style="width:100%"
                :precision="2"
                :clearable="true"
              />
            </el-form-item>
            <div>
              <div class="card-overview">
                <i class="fa fa-map" @click="dialogMapVisible = true" />
                <el-button type="text" @click="dialogMapVisible = true">查看详细地址</el-button>
              </div>
              <el-dialog v-model="dialogMapVisible" title="充电站详细地址">
                <MapContainer />
                <span slot="footer" class="dialog-footer" style="display: flex; justify-content: right">
                  <el-button
                    type="primary"
                    style="margin-top: 10px"
                    @click="dialogMapVisible = false"
                  >确 定</el-button>
                </span>
              </el-dialog>
            </div>
          </el-form>
        </div>
      </div>
      <div v-if="active.active===1" class="gva-card">
        <div class="card-header">
          <span>确认订单</span>
        </div>
        <div class="gva-card gva-order-card">
          <el-form
            label-position="right"
            label-width="120px"
          >
            <el-form-item label="充电的车牌号:">
              <el-input
                style="width:100%"
                :disabled="true"
                :placeholder="chargeStore.chargeRequest.car_id"
              />
            </el-form-item>
            <el-form-item label="充电类型:">
              <el-input
                style="width:100%"
                :disabled="true"
                :placeholder="chargeStore.chargeRequest.mode"
              />
            </el-form-item>
            <el-form-item label="充电度数:">
              <el-input
                style="width:100%"
                :disabled="true"
                :placeholder="chargeStore.chargeRequest.apply_kwh"
              />
            </el-form-item>
            <el-form-item label="预计充电时间:">
              <el-input
                style="width:100%"
                :disabled="true"
                :placeholder="chargeStore.chargeRequest.apply_time"
              />
            </el-form-item>
          </el-form>
        </div>
      </div>
      <div v-if="active.active===2" class="gva-card">
        <div class="card-header">
          <span>开始充电</span>
        </div>
        <div class="gva-card gva-order-card">
          <div class="gva-charging">已进入充电队列...</div>
          <bubble />
        </div>
      </div>
      <div class="gva-card gva-steps">
        <el-steps :active="active.active" finish-status="success">
          <el-step title="提交订单" />
          <el-step title="确认订单" />
          <el-step title="开始充电" />
        </el-steps>
        <div class="gva-button">
          <el-button v-if="active.active===2" style="margin-top: 12px;" @click="active.backHome">开启新订单</el-button>
          <el-button
            v-if="active.active===2"
            style="margin-top: 12px;margin-right: 0px"
            @click="active.lastStep"
          >上一步
          </el-button>
          <el-button v-if="active.active===1" style="margin-top: 12px;" @click="enterOrder">下一步</el-button>
          <el-button v-if="active.active===0" style="margin-top: 12px;" @click="enterOrder">提交</el-button>
          <el-button v-if="active.active===2" disabled style="margin-top: 12px;">已完成</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import MapContainer from '@/view/chargePile/MapContainer.vue'
import {createOrder2, getOrderListByUserId2} from '@/api/order'
import { ActiveStore } from '@/pinia/modules/active'
import { ChargeStore } from '@/pinia/modules/chargeRequest'
import bubble from '@/view/order/bubble.vue'

const active = ActiveStore()
active.loadActive()
const chargeStore = ChargeStore()
chargeStore.loadChargeInfo()
const dialogMapVisible = ref(false)
const searchInfo = ref('CURRENT')
const formData = ref({
  userId: 0,
  carId: '',
  chargeCost: 0,
  kwh: 0,
  time: 0,
  pileId: 0,
  serviceCost: 0,
  startedAt: new Date(),
  stopAt: new Date(),
  totalCost: 0,
  date1: '',
  date2: '',
})
const requestForm = ref({
  car_id: '',
  mode: '',
  apply_kwh: 0,
  stationId: 1,
})
const tmpForm = reactive({
  car_id: '',
  mode: '',
  apply_kwh: 0,
  apply_time: 0,
  stationId: 1,
})
const modeOptions = ref([
  {
    value: 'T',
    label: '慢充',
  },
  {
    value: 'F',
    label: '快充',
  }
])
const elFormRef = ref()
// 验证规则
const rule = reactive({
  car_id: [{
    required: true,
    message: '请输入格式正确的车牌号',
    trigger: ['input', 'blur'],
  }],
  mode: [{
    required: true,
    message: '请选择充电类型',
    trigger: ['input', 'blur'],
  }],
  apply_kwh: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

// 提交订单
async function enterOrder() {
  if (active.active === 0) {
        elFormRef.value?.validate(async(valid) => {
          if (!valid) return
          else {
            active.nextStep()
          }
        })
        tmpForm.car_id = requestForm.value.car_id
        tmpForm.mode = requestForm.value.mode
        tmpForm.apply_kwh = requestForm.value.apply_kwh
        if (tmpForm.mode === 'T') {
          tmpForm.apply_time = tmpForm.apply_kwh / 7
        } else {
          tmpForm.apply_time = tmpForm.apply_kwh / 30
        }
        const res = await createOrder2(requestForm.value)
        console.log(res)
        chargeStore.setChargeInfo(tmpForm)
        return
  }
  active.nextStep()
}

const getTableData = async() => {
  const res = await getOrderListByUserId2({ mode: searchInfo.value })
  console.log(res)
}
getTableData()

</script>

<style lang="scss">
.gva-card-box {
  height: 80vh;
  position: relative;
}

.card-header {
  padding-bottom: 20px;
  border-bottom: 1px solid #e8e8e8;
  font-size: 20px;
}

.gva-card {
  box-sizing: border-box;
  background-color: #fff;
  border-radius: 2px;
  height: auto;
  padding: 26px 30px;
  overflow: hidden;
  box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
}

.gva-order-card {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;

  .gva-charging {
    margin: 10px;
    font-size: 20px;
  }
}

.card-overview {
  display: flex;
  align-items: center;
  justify-content: center;

  i {
    font-size: 20px;
    margin-right: 10px;
    cursor: pointer;
  }
}

.gva-steps {
  width: 100%;
  position: absolute;
  bottom: 5px;

  .gva-button {
    display: flex;
    justify-content: right;
  }
}

</style>
