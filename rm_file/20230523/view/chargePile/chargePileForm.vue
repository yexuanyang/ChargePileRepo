<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="left" :rules="rule" label-width="200px">
        <el-form-item label="是否开启:" prop="isOpen">
          <el-switch
            v-model="formData.isOpen"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item label="充电桩类型:" prop="pileType">
          <el-select v-model="formData.pileType" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="item in ['快充','慢充','其他']" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="充电桩累计充电次数:" prop="chargeCount">
          <el-input v-model.number="formData.chargeCount" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="充电站的ID:" prop="stationId">
          <el-input v-model.number="formData.stationId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="充电桩累计充电电量:" prop="electricity">
          <el-input-number v-model="formData.electricity" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="充电时间:" prop="chargeTime">
          <el-date-picker
            v-model="formData.chargeTime"
            type="date"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ChargePile'
}
</script>

<script setup>
import {
  createChargePile,
  updateChargePile,
  findChargePile
} from '@/api/chargePile'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  isOpen: false,
  chargeCount: 0,
  stationId: 0,
  electricity: 0,
  chargeTime: new Date(),
})
// 验证规则
const rule = reactive({
  isOpen: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  pileType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  stationId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findChargePile({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.rechargePile
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
    elFormRef.value?.validate(async(valid) => {
      if (!valid) return
      let res
      switch (type.value) {
        case 'create':
          res = await createChargePile(formData.value)
          break
        case 'update':
          res = await updateChargePile(formData.value)
          break
        default:
          res = await createChargePile(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
