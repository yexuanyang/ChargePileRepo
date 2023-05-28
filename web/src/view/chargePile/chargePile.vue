<template>
    <div>
        <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
                <el-form-item label="创建时间">
                    <el-date-picker
                        v-model="searchInfo.startCreatedAt"
                        placeholder="开始时间"
                        type="datetime"
                    />
                    —
                    <el-date-picker
                        v-model="searchInfo.endCreatedAt"
                        placeholder="结束时间"
                        type="datetime"
                    />
                </el-form-item>
                <el-form-item label="充电桩类型">
                    <el-select v-model="searchInfo.pileType" placeholder="请选择">
                        <el-option
                            v-for="item in pileOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="充电桩站点ID">
                    <el-select v-model="searchInfo.stationId" placeholder="请选择" style="width: 200px">
                        <el-option
                            v-for="item in stationOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button icon="search" type="primary" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button icon="plus" type="primary" @click="openDialog">新增</el-button>
                <el-popover v-model:visible="deleteVisible" placement="top" width="160">
                    <p>确定要删除吗？</p>
                    <div style="text-align: right; margin-top: 8px;">
                        <el-button link type="primary" @click="deleteVisible = false">取消</el-button>
                        <el-button type="primary" @click="onDelete">确定</el-button>
                    </div>
                    <template #reference>
                        <el-button
                            :disabled="!multipleSelection.length"
                            icon="delete"
                            style="margin-left: 10px;"
                            @click="deleteVisible = true"
                        >删除
                        </el-button>
                    </template>
                </el-popover>
                <el-button
                    :disabled="!multipleSelection.length"
                    icon="open"
                    style="margin-left: 10px;"
                    @click="openConfirm"
                >开启充电桩
                </el-button>
                <el-button
                    :disabled="!multipleSelection.length"
                    icon="close"
                    style="margin-left: 10px;"
                    @click="closeConfirm"
                >关闭充电桩
                </el-button>
            </div>
            <el-table
                ref="multipleTable"
                :data="tableData"
                row-key="ID"
                style="width: 100%"
                tooltip-effect="dark"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55"/>
                <el-table-column align="left" label="日期" width="180">
                    <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="状态" prop="isOpen" width="120">
                    <template #default="scope">
                        <i v-if="scope.row.isOpen" class="fa fa-power-off" style="color: #52c41a">
                            <span style=" color: #595959;margin-left: 10px">开机</span>
                        </i>
                        <i v-else="!scope.row.isOpen" class="fa fa-power-off">
                            <span style="margin-left: 10px">关机</span>
                        </i>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="充电桩类型" prop="pileType" width="120"/>
                <el-table-column align="left" label="充电桩累计充电次数" prop="chargeCount" width="120"/>
                <el-table-column align="left" label="充电站的ID" prop="stationId" width="120"/>
                <el-table-column align="left" label="充电桩累计充电电量（度）" prop="electricity" width="120"/>
                <el-table-column align="left" label="充电桩累计充电时间（小时）" prop="chargeTime" width="140"/>
                <el-table-column align="left" label="按钮组">
                    <template #default="scope">
                        <el-button
                            class="table-button"
                            icon="edit"
                            link
                            type="primary"
                            @click="updateChargePileFunc(scope.row)"
                        >变更
                        </el-button>
                        <el-button icon="delete" link type="primary" @click="deleteRow(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination
                    :current-page="page"
                    :page-size="pageSize"
                    :page-sizes="[10, 30, 50, 100]"
                    :total="total"
                    layout="total, sizes, prev, pager, next, jumper"
                    @current-change="handleCurrentChange"
                    @size-change="handleSizeChange"
                />
            </div>
        </div>
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
            <el-form ref="elFormRef" :model="formData" :rules="rule" label-position="right" label-width="200px">
                <el-form-item label="是否开启:" prop="isOpen">
                    <el-switch
                        v-model="formData.isOpen"
                        active-color="#13ce66"
                        active-text="是"
                        clearable
                        inactive-color="#ff4949"
                        inactive-text="否"
                    />
                </el-form-item>
                <el-form-item label="充电桩类型:" prop="pileType">
                    <el-select v-model="formData.pileType" :clearable="true" placeholder="请选择" style="width:100%">
                        <el-option v-for="item in ['快充','慢充']" :key="item" :label="item" :value="item"/>
                    </el-select>
                </el-form-item>
                <el-form-item label="充电桩累计充电次数:" prop="chargeCount">
                    <el-input v-model.number="formData.chargeCount" :clearable="true" placeholder="请输入"/>
                </el-form-item>
                <el-form-item label="充电站的ID:" prop="stationId">
                    <el-input v-model.number="formData.stationId" :clearable="true" placeholder="请输入"/>
                </el-form-item>
                <el-form-item label="充电桩累计充电电量（度）:" prop="electricity">
                    <el-input-number
                        v-model="formData.electricity"
                        :clearable="true"
                        :precision="2"
                        style="width:100%"
                    />
                </el-form-item>
                <el-form-item label="充电桩累计充电时间（小时）:" prop="chargeTime">
                    <el-input-number
                        v-model="formData.chargeTime"
                        :clearable="true"
                        :precision="2"
                        style="width:100%"
                    />
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
    name: 'ChargePile'
}
</script>

<script setup>
import {
    createChargePile,
    deleteChargePile,
    deleteChargePileByIds,
    updateChargePile,
    findChargePile,
    getChargePileList,
    UpdateChargePileByIds
} from '@/api/chargePile'

// 全量引入格式化工具 请按需保留
import {getDictFunc, formatDate, formatBoolean, filterDict} from '@/utils/format'
import {ElMessage, ElMessageBox} from 'element-plus'
import {ref, reactive} from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
    isOpen: false,
    chargeCount: 0,
    stationId: 0,
    electricity: 0,
    chargeTime: 0,
})
console.log(formData)
// 验证规则
const rule = reactive({
    isOpen: [{
        required: false,
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

// 充电类型选项
const pileOptions = [
    {
        value: '快充',
        label: '快充'
    },
    {
        value: '慢充',
        label: '慢充'
    }
]

// 站点选项
const stationOptions = [
    {
        value: '1',
        label: '站点一'
    },
    {
        value: '2',
        label: '站点二'
    }
]
const elFormRef = ref()

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
    if (searchInfo.value.isOpen === '') {
        searchInfo.value.isOpen = null
    }
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
    const res = await getChargePileList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
    console.log(res)
    if (res.code === 0) {
        tableData.value = res.data.list
        total.value = res.data.total
        page.value = res.data.page
        pageSize.value = res.data.pageSize
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

const openConfirm = () => {
    ElMessageBox.confirm('确定要对以下充电桩进行开机操作吗?', '开启充电桩', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
    }).then(() => {
        OpenPile()
    })
}

const closeConfirm = () => {
    ElMessageBox.confirm('确定要对以下充电桩进行关机操作吗?', '关闭充电桩', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        ClosePile()
    })
}
// 打开开关
async function OpenPile() {
    const ids = []
    multipleSelection.value &&
    multipleSelection.value.map(item => {
        ids.push(item.ID)
    })
    await UpdateChargePileByIds({
        is_open: true,
        ids: ids
    })
    getTableData()
}

// 关上开关
async function ClosePile() {
    const ids = []
    multipleSelection.value &&
    multipleSelection.value.map(item => {
        ids.push(item.ID)
    })
    await UpdateChargePileByIds({
        is_open: false,
        ids: ids
    })
    getTableData()
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        deleteChargePileFunc(row)
    })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
    const ids = []
    if (multipleSelection.value.length === 0) {
        ElMessage({
            type: 'warning',
            message: '请选择要删除的数据'
        })
        return
    }
    multipleSelection.value &&
    multipleSelection.value.map(item => {
        ids.push(item.ID)
    })
    const res = await deleteChargePileByIds({ids})
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
            page.value--
        }
        deleteVisible.value = false
        getTableData()
    }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateChargePileFunc = async (row) => {
    const res = await findChargePile({ID: row.ID})
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rechargePile
        dialogFormVisible.value = true
    }
}

// 删除行
const deleteChargePileFunc = async (row) => {
    const res = await deleteChargePile({ID: row.ID})
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
        isOpen: false,
        chargeCount: 0,
        stationId: 0,
        electricity: 0,
        chargeTime: 0,
    }
}
// 弹窗确定
const enterDialog = async () => {
    elFormRef.value?.validate(async (valid) => {
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
            closeDialog()
            getTableData()
        }
    })
}
</script>
<style>
</style>
