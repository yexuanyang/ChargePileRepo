<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单所属的用户Id:" prop="user_id">
          <el-input v-model.number="formData.user_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="充电的车牌号:" prop="carId">
          <el-input v-model="formData.carId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="充电类型:" prop="chargeType">
        <el-select v-model="formData.chargeType" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['快充','慢充','其他']" :key="item" :label="item" :value="item" />
        </el-select>
        </el-form-item>
        <el-form-item label="充电费用:" prop="chargeCost">
          <el-input-number v-model="formData.chargeCost" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="充电度数:" prop="kwh">
          <el-input-number v-model="formData.kwh" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="充电时长:" prop="time">
          <el-date-picker v-model="formData.time" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="充电桩id:" prop="pileId">
          <el-input v-model.number="formData.pileId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="服务费:" prop="serviceCost">
          <el-input-number v-model="formData.serviceCost" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="开始充电时间:" prop="startedAt">
          <el-date-picker v-model="formData.startedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="结束充电时间:" prop="stopAt">
          <el-date-picker v-model="formData.stopAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="总花费:" prop="totalCost">
          <el-input-number v-model="formData.totalCost" :precision="2" :clearable="true"></el-input-number>
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
  name: 'Order'
}
</script>

<script setup>
import {
  createOrder,
  updateOrder,
  findOrder
} from '@/api/order'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            user_id: 0,
            carId: '',
            chargeCost: 0,
            kwh: 0,
            time: new Date(),
            pileId: 0,
            serviceCost: 0,
            startedAt: new Date(),
            stopAt: new Date(),
            totalCost: 0,
        })
// 验证规则
const rule = reactive({
               user_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               carId : [{
                   required: true,
                   message: '请输入格式正确的车牌号',
                   trigger: ['input','blur'],
               }],
               chargeType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               kwh : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pileId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startedAt : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reorder
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOrder(formData.value)
               break
             case 'update':
               res = await updateOrder(formData.value)
               break
             default:
               res = await createOrder(formData.value)
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
