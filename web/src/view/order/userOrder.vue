<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="订单类型">
          <el-select v-model="searchInfo">
            <el-option
              v-for="item in orderType"
              :key="item.key"
              :value="item.value"
              :label="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column align="center" label="创建日期" width="260" fixed>
          <template #default="scope">{{ scope.row.create_time }}</template>
        </el-table-column>
        <el-table-column align="center" label="订单编号" prop="id" width="120" />
        <el-table-column align="center" label="充电的车牌号" prop="car_id" width="120" />
        <el-table-column align="center" label="充电类型" prop="mode" width="120" />
        <el-table-column align="center" label="订单申请充电度数" prop="apply_kwh" width="160" />
        <el-table-column align="center" label="充电费用" prop="charge_price" width="120" />
        <el-table-column align="center" label="充电度数" prop="charge_kwh" width="120" />
        <el-table-column align="center" label="充电站id" prop="stationId" width="120" />
        <el-table-column align="center" label="充电桩id" prop="charge_id" width="120" />
        <el-table-column align="center" label="服务费" prop="service_price" width="120" />
        <el-table-column align="center" label="前车数量" prop="front_cars" width="120" fixed="right" />
        <el-table-column align="center" label="订单状态" prop="state" width="120" fixed="right" />
        <el-table-column align="center" label="开始充电时间" width="260">
          <template #default="scope">{{ scope.row.start_time }}</template>
        </el-table-column>
        <el-table-column align="center" label="结束充电时间" width="260" prop="stopAt">
          <template #default="scope">{{ scope.row.finish_time }}</template>
        </el-table-column>
        <el-table-column align="center" label="总花费" prop="fee" width="120" />
        <el-table-column align="center" label="按钮组" fixed="right" width="110">
          <template #default="scope">
            <el-button
              v-show="isFinished(scope.row.finish_time) && scope.row.state === 'WAITING'"
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateOrderFunc(scope.row)"
            >变更
            </el-button>
            <el-button v-show="isFinished(scope.row.finish_time)" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elFormRef" :model="type === 'create' ? requestForm : updateForm" label-position="right" :rules="type === 'create' ? rule_create : rule_update" label-width="200px">
        <el-form-item v-show="type === 'create'" label="充电的车牌号:" prop="car_id">
          <el-input v-model="requestForm.car_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="充电类型:" prop="mode">
          <el-select v-show="type === 'create'" v-model="requestForm.mode" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="item in modeOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
          <el-select v-show="type === 'update'" v-model="updateForm.mode" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="item in modeOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="充电度数:" prop="apply_kwh">
          <el-input-number v-show="type === 'create'" v-model="requestForm.apply_kwh" style="width:100%" :precision="2" :clearable="true" />
          <el-input-number v-show="type === 'update'" v-model="updateForm.apply_kwh" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderListByUserId, createOrder2, getOrderListByUserId2, updateOrder2, deleteOrder2
} from '@/api/order'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { now } from '@vueuse/shared'
import { ElMessage, ElMessageBox } from 'element-plus'
import { toRef } from 'vue'
import { ref, reactive, computed } from 'vue'
import { getChargeStationList } from '@/api/chargeStation'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  car_id: '',
  chargeCost: 0,
  kwh: 0,
  time: 0,
  pileId: 0,
  mode: '',
  serviceCost: 0,
  startedAt: new Date(),
  stopAt: new Date(Date.now() + 3600 * 1000 * 24),
  totalCost: 0,
  stationId: null,
  state: '',
})

const orderType = ref([
  {
    key: '当前订单（未完成的）',
    value: 'CURRENT'
  },
  {
    key: '历史订单（已完成）',
    value: 'HISTORY'
  }
])

const requestForm = ref({
  car_id: '',
  mode: '',
  apply_kwh: 0,
  stationId: 1,
})
// 验证规则

const validateCarId = () => {
  const regex = /^[\u4e00-\u9fa5]{1}[A-Z]{1}[A-Z_0-9]{5}$/
  const value = toRef(requestForm, 'car_id').value.car_id
  return regex.test(value)
}

const rule_create = reactive({
  car_id: [{
    required: true,
    message: '车辆id不能为空',
    trigger: ['input', 'blur'],
  }],
  mode: [{
    required: true,
    message: '请选择充电类型',
    trigger: ['input', 'blue'],
  }],
  apply_kwh: [{
    required: true,
    message: '请输入充电度数',
    trigger: ['input', 'blue'],
  }],
  stationId: [{
    required: true,
    message: '请输入充电站Id',
    trigger: ['input', 'blue'],
  }],
})

const rule_update = reactive({
  mode: [{
    required: true,
    message: '请选择充电类型',
    trigger: ['input', 'blue'],
  }],
  apply_kwh: [{
    required: true,
    message: '请输入充电度数',
    trigger: ['input', 'blue'],
  }]
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

// 从数据库获取站点选项
// const stationOptions = ref([])
// const getStationOptions = async() => {
//   const res = await getChargeStationList()
//   if (res.code === 0) {
//     res.data.list.forEach((item) => {
//       stationOptions.value.push({ value: item.ID, label: item.position })
//     })
//   }
// }
// getStationOptions()

const elFormRef = ref()

// 是否可以更改
const isFinished = (stopAt) => {
  const now = new Date().getTime()
  const stopAtTime = new Date(stopAt).getTime() - 14 * 3600 * 1000
  return !(stopAtTime < now)
}

// =========== 表格控制部分 ===========
const tableData = ref([])
const searchInfo = ref('CURRENT')

// 重置
const onReset = () => {
  searchInfo.value = 'CURRENT'
  getTableData()
}

// 搜索
const onSubmit = () => {
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getOrderListByUserId2({ mode: searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteOrderFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

const updateForm = ref({
  id: '',
  mode: '',
  apply_kwh: 0,
})

// 更新行
const updateOrderFunc = async(row) => {
  // const res = await findOrder({ID: row.ID})
  type.value = 'update'
  dialogFormVisible.value = true
  updateForm.value.id = row.id
  updateForm.value.mode = row.mode
  updateForm.value.apply_kwh = row.apply_kwh
  // if (res.code === 0) {
  //   formData.value = res.data.reorder
  //   dialogFormVisible.value = true
  // }
}

// 删除行
const deleteOrderFunc = async(row) => {
  const res = await deleteOrder2({ ID: row.id })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    await getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    carId: '',
    chargeCost: 0,
    kwh: 0,
    time: 0,
    pileId: 0,
    serviceCost: 0,
    startedAt: new Date(),
    stopAt: new Date(Date.now() + 3600 * 1000 * 24),
    totalCost: 0,
    stationId: null,
    state: '',
  }
  requestForm.value = {
    car_id: '',
    mode: '',
    apply_kwh: 0,
    stationId: 1,
  }
  updateForm.value = {
    id: '',
    mode: '',
    apply_kwh: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) { // 校验失败
      return
    }
    let res
    switch (type.value) {
      case 'create':
        res = await createOrder2(requestForm.value)
        break
      case 'update':
        res = await updateOrder2(updateForm.value)
        break
      default:
        res = await createOrder2(requestForm.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style></style>
