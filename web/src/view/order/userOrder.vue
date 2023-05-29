<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
        </el-form-item>
        <el-form-item label="充电的车牌号">
          <el-input v-model="searchInfo.carId" placeholder="搜索条件"/>

        </el-form-item>
        <el-form-item label="充电类型">
          <el-input v-model="searchInfo.chargeType" placeholder="搜索条件"/>

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
          @selection-change="handleSelectionChange">
        <el-table-column align="left" label="日期" width="180" fixed>
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="充电的车牌号" prop="carId" width="120"/>
        <el-table-column align="left" label="充电类型" prop="chargeType" width="120"/>
        <el-table-column align="left" label="充电费用" prop="chargeCost" width="120"/>
        <el-table-column align="left" label="充电度数" prop="kwh" width="120"/>
        <el-table-column align="left" label="充电时长" width="180">
          <template #default="scope">{{ scope.row.time }} min</template>
        </el-table-column>
        <el-table-column align="left" label="充电桩id" prop="pileId" width="120"/>
        <el-table-column align="left" label="服务费" prop="serviceCost" width="120"/>
        <el-table-column align="left" label="订单状态" prop="state" width="120"/>
        <el-table-column align="left" label="开始充电时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.startedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="结束充电时间" width="180" prop="stopAt">
          <template #default="scope">{{ formatDate(scope.row.stopAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="总花费" prop="totalCost" width="120"/>
        <el-table-column align="left" label="按钮组" fixed="right" width="110">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" v-show="isFinished(scope.row.stopAt)"
                       @click="updateOrderFunc(scope.row)">变更
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
                       :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
                       @size-change="handleSizeChange"/>
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="200px">
        <el-form-item label="充电的车牌号:" prop="carId">
          <el-input v-model="formData.carId" :clearable="true" placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="充电类型:" prop="chargeType">
          <el-select v-model="formData.chargeType" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="item in ['快充', '慢充', '其他']" :key="item" :label="item" :value="item"/>
          </el-select>
        </el-form-item>
        <el-form-item label="充电度数:" prop="kwh">
          <el-input-number v-model="formData.kwh" style="width:100%" :precision="2" :clearable="true"/>
        </el-form-item>
        <el-form-item label="充电站ID:" prop="stationId">
          <el-input-number v-model="formData.stationId" style="width:100%" :step="1" :min="0" :clearable="true"/>
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
  getOrderListByUserId
} from '@/api/order'

// 全量引入格式化工具 请按需保留
import {getDictFunc, formatDate, formatBoolean, filterDict} from '@/utils/format'
import {now} from '@vueuse/shared';
import {ElMessage, ElMessageBox} from 'element-plus'
import {toRef} from 'vue';
import {ref, reactive, computed} from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  carId: '',
  chargeCost: 0,
  kwh: 0,
  time: 0,
  pileId: 0,
  serviceCost: 0,
  startedAt: new Date(),
  stopAt: new Date().setDate(Date.now() + 3600 * 1000 * 24),
  totalCost: 0,
  stationId: 0,
})

// 验证规则

const validateCarId = () => {
  const regex = /^[\u4e00-\u9fa5]{1}[A-Z]{1}[A-Z_0-9]{5}$/
  const value = toRef(formData, 'carId').value.carId
  return regex.test(value)
}

const rule = reactive({
  carId: [{
    required: true,
    validator: validateCarId,
    message: '请输入格式正确的车牌号',
    trigger: ['input', 'blur'],
  }],
  chargeType: [{
    required: true,
    message: '请选择充电类型',
    trigger: ['input', 'blue'],
  }],
  kwh: [{
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

const elFormRef = ref()

// 是否可以更改
const isFinished = (stopAt) => {
  const now = new Date().getTime()
  const stopAtTime = new Date(stopAt).getTime()
  return stopAtTime < now ? false : true
}

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getOrderListByUserId({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
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

// 更新行
const updateOrderFunc = async (row) => {
  const res = await findOrder({ID: row.ID})
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reorder
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteOrderFunc = async (row) => {
  const res = await deleteOrder({ID: row.ID})
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
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
    time: new Date(),
    pileId: 0,
    serviceCost: 0,
    startedAt: new Date(),
    stopAt: new Date(),
    totalCost: 0,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) {//校验失败
      return
    }
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
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style></style>
  