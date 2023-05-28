<template>
    <el-row :gutter="20">
        <el-col :span="8">
            <div class="statistic-card">
                <el-statistic :value="todayChargeKwh">
                    <template #title>
                        <div style="display: inline-flex; align-items: center">
                            今日充电度数
                            <el-tooltip effect="dark" content="截止到今天24：00，你所有车辆充电的总度数" placement="top">
                                <el-icon style="margin-left: 4px" :size="12">
                                    <Warning />
                                </el-icon>
                            </el-tooltip>
                        </div>
                    </template>
                </el-statistic>
                <div class="statistic-footer">
                    <div class="footer-item">
                        <span>相较昨天</span>
                        <span :class="dayChargeClassName">
                            {{ dayChargeRate }} 度
                            <el-icon>
                                <CaretTop v-show="dayChargeClassName == 'green'" />
                                <CaretBottom v-show="dayChargeClassName == 'red'" />
                            </el-icon>
                        </span>
                    </div>
                </div>
            </div>
        </el-col>
        <el-col :span="8">
            <div class="statistic-card">
                <el-statistic :value="monthChargeKwh">
                    <template #title>
                        <div style="display: inline-flex; align-items: center">
                            本月充电度数
                            <el-tooltip effect="dark" content="截止到本月最后一天的24：00，你所有车辆充电的总度数" placement="top">
                                <el-icon style="margin-left: 4px" :size="12">
                                    <Warning />
                                </el-icon>
                            </el-tooltip>
                        </div>
                    </template>
                </el-statistic>
                <div class="statistic-footer">
                    <div class="footer-item">
                        <span>相较上月</span>
                        <span :class="monthChargeClassName">
                            {{ monthChargeRate }} 度
                            <el-icon>
                                <CaretTop v-show="monthChargeClassName == 'green'" />
                                <CaretBottom v-show="monthChargeClassName == 'red'" />
                            </el-icon>
                        </span>
                    </div>
                </div>
            </div>
        </el-col>
        <el-col :span="8">
            <div class="statistic-card">
                <el-statistic :value="todayChargePrice" title="New transactions today">
                    <template #title>
                        <div style="display: inline-flex; align-items: center">
                            今日充电使用的总金额
                        </div>
                    </template>
                </el-statistic>
                <div class="statistic-footer">
                    <div class="footer-item">
                        <span>相比昨日</span>
                        <span :class="dayChargePriceClassName">
                            {{ dayChargePriceRate }} ￥
                            <el-icon>
                                <CaretTop v-show="dayChargePriceClassName == 'green'" />
                                <CaretBottom v-show="dayChargePriceClassName == 'red'" />
                            </el-icon>
                        </span>
                    </div>
                    <div class="footer-item">
                        <el-row>
                            <el-button type="success" :icon="ArrowRight" circle />
                        </el-row>
                    </div>
                </div>
            </div>
        </el-col>
    </el-row>
    <el-row :gutter="10">
        <el-col :span="12">
            <div ref="chart" style="width: 100%;height: 800px;margin-top: 50px;display:inline-block;"></div>
        </el-col>
        <el-col :span="12">
            <el-form :inline="true" class="demo-form-inline" v-model="searchInfo" @keyup.enter="onSubmit"
                style="margin-top: 50px;">
                <el-form-item label="报表时间">
                    <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="选择报表开始时间"
                        :shortcuts="shortcuts" />
                    ——
                    <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="选择报表结束时间"
                        :shortcuts="shortcuts" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
            <el-table :data="tableData" border show-summary :summary-method="getSummaries"
                style="width: 100%;display:inline-block;" :default-sort="{ prop: 'userId', order: 'descending' }">
                <el-table-column align="left" label="日期" width="180" fixed prop="CreatedAt" sortable>
                    <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="充电的车牌号" prop="carId" width="120" />
                <el-table-column align="left" label="充电类型" prop="chargeType" width="120" />
                <el-table-column align="left" label="充电费用" prop="chargeCost" width="120" />
                <el-table-column align="left" label="充电度数" prop="kwh" width="120" fixed sortable />
                <el-table-column align="left" label="充电时长" width="180">
                    <template #default="scope">{{ scope.row.time }}</template>
                </el-table-column>
                <el-table-column align="left" label="充电桩id" prop="pileId" width="120" />
                <el-table-column align="left" label="服务费" prop="serviceCost" width="120" />
                <el-table-column align="left" label="开始充电时间" width="180">
                    <template #default="scope">{{ formatDate(scope.row.startedAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="结束充电时间" width="180">
                    <template #default="scope">{{ formatDate(scope.row.stopAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="总花费" prop="totalCost" width="120" fixed="right" sortable />
            </el-table>

            <div class="demo-pagination-block">
                <el-pagination v-model:current-page="page" v-model:page-size="pageSize" :page-sizes="[10, 20, 30, 40]"
                    :small="small" :disabled="disabled" :background="background"
                    layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
                    @current-change="handleCurrentChange" />
            </div>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import {
    ArrowRight,
    CaretBottom,
    CaretTop,
    Warning,
} from '@element-plus/icons-vue'
import * as echarts from 'echarts'
import { onMounted, ref } from 'vue'
import { type TableColumnCtx } from 'element-plus'
import { getOrderListByUserId } from '../../api/order'
import { formatDate } from '../../utils/format'
import { computed } from '@vue/reactivity'
import { getDurationTotalCharge, getDurationTotalPrice, getDurationReportInfo } from '../../api/report'
import { last, template } from 'lodash'
import { nextTick } from 'process'

const shortcuts = [
    {
        text: '今天',
        value: new Date()
    },
    {
        text: '昨天',
        value: () => {
            const date = new Date()
            date.setTime(date.getTime() - 3600 * 1000 * 24)
            return date
        },
    },
    {
        text: '一周前',
        value: () => {
            const date = new Date()
            date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
            return date
        },
    },
]

//=======数据对比部分========

//=======每日充电量对比========
const todayChargeKwh = ref(0)
const yesterdayChargeKwh = ref(0)
const dayChargeRate = ref(0)

const dayChargeClassName = computed(() => {
    return dayChargeRate.value >= 0 ? 'green' : 'red'
})

const getTodayChargeKwh = async () => {
    const date = new Date()
    date.setHours(0, 0, 0, 0)
    const endDate = new Date()
    endDate.setHours(24, 0, 0)
    const total = await getDurationTotalCharge({ date: date, endDate: endDate })
    if (total.code == 0) {
        todayChargeKwh.value = total.data.total
    }
}

const getYesterDayChargeKwh = async () => {
    const date = new Date()
    date.setHours(-24, 0, 0)
    const endDate = new Date()
    const total = await getDurationTotalCharge({ date: date, endDate: endDate })
    if (total.code == 0) {
        yesterdayChargeKwh.value = total.data.total
    }
}

const getDayChargeRate = async () => {
    await getYesterDayChargeKwh()
    await getTodayChargeKwh()
    dayChargeRate.value = Number((todayChargeKwh.value - yesterdayChargeKwh.value).toFixed(2))
}


//=======每月充电量对比========
const monthChargeKwh = ref(0)
const monthChargeRate = ref(0)
const LastMonthChargeKwh = ref(0)

const monthChargeClassName = computed(() => {
    return monthChargeRate.value >= 0 ? 'green' : 'red'
})

const getMonthChargeKwh = async () => {
    const date = new Date()
    date.setMonth(date.getMonth(), 0)
    date.setHours(24, 0, 0)
    const endDate = new Date()
    endDate.setMonth(endDate.getMonth() + 1, 0)
    endDate.setHours(24, 0, 0)
    const total = await getDurationTotalCharge({ date: date, endDate: endDate })
    if (total.code == 0) {
        monthChargeKwh.value = total.data.total
    }
}

const getLastMonthChargeKwh = async () => {
    const date = new Date()
    date.setMonth(date.getMonth() - 1, 0)
    date.setHours(24, 0, 0)
    const endDate = new Date()
    endDate.setMonth(endDate.getMonth(), 0)
    endDate.setHours(24, 0, 0)
    const total = await getDurationTotalCharge({ date: date, endDate: endDate })
    if (total.code == 0) {
        LastMonthChargeKwh.value = total.data.total
    }
}

const getMonthChargeRate = async () => {
    await getLastMonthChargeKwh()
    await getMonthChargeKwh()
    monthChargeRate.value = Number((monthChargeKwh.value - LastMonthChargeKwh.value).toFixed(2))
}

//=======每日充电总金额对比========
const todayChargePrice = ref(0)
const dayChargePriceRate = ref(0)
const yesterdayChargePrice = ref(0)
const dayChargePriceClassName = computed(() => {
    return dayChargePriceRate.value >= 0 ? 'green' : 'red'
})

const getTodayChargePrice = async () => {
    const date = new Date()
    date.setMonth(date.getMonth(), 0)
    date.setHours(24, 0, 0)
    const endDate = new Date()
    endDate.setMonth(endDate.getMonth() + 1, 0)
    endDate.setHours(24, 0, 0)
    const total = await getDurationTotalPrice({ date: date, endDate: endDate })
    if (total.code == 0) {
        todayChargePrice.value = total.data.total
    }
}

const getYesterdayChargePrice = async () => {
    const date = new Date()
    date.setMonth(date.getMonth() - 1, 0)
    date.setHours(24, 0, 0)
    const endDate = new Date()
    endDate.setMonth(endDate.getMonth(), 0)
    endDate.setHours(24, 0, 0)
    const total = await getDurationTotalCharge({ date: date, endDate: endDate })
    if (total.code == 0) {
        yesterdayChargePrice.value = total.data.total
    }
}

const getDayChargePriceRate = async () => {
    await getYesterdayChargePrice()
    await getTodayChargePrice()
    dayChargePriceRate.value = Number((todayChargePrice.value - yesterdayChargePrice.value).toFixed(2))
}


//========获取要展示的比较数据==========
const initData = async () => {
    await getDayChargeRate()
    await getMonthChargeRate()
    await getDayChargePriceRate()
}

initData()


// =======表格部分========

interface Order {
    CreatedAt: Date
    userId: number
    carId: string
    chargeType: string
    chargeCost: number
    kwh: number
    serviceCost: number
    time: number
    startedAt: Date
    stopAt: Date
    pileId: number
    totalCost: number
}


const SummaryDataKey = ["chargeCost", "kwh", "serviceCost", "totalCost"]

interface SummaryMethodProps<T = Order> {
    columns: TableColumnCtx<T>[]
    data: T[]
}

const getSummaries = (param: SummaryMethodProps) => {
    const { columns, data } = param
    const sums: string[] = []
    columns.forEach((column, index) => {
        if (index === 0) {
            sums[index] = '合计'
            return
        }
        const values = data.map((item) => Number(item[column.property]))
        if (!values.every((value) => Number.isNaN(value))) {
            if (SummaryDataKey.includes(column.property))
                sums[index] = `${values.reduce((prev, curr) => {
                    const value = Number(curr)
                    if (!Number.isNaN(value)) {
                        return Number((prev + curr).toFixed(2))
                    } else {
                        return prev
                    }
                }, 0)}`
        } else {
            sums[index] = ''
        }
    })
    return sums
}

const page = ref(1)
const pageSize = ref(10)
const small = ref(false)
const disabled = ref(false)
const background = ref(false)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({ startCreatedAt: null, endCreatedAt: null })

const handleSizeChange = (val: number) => {
    pageSize.value = val
    getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}


const getTableData = async () => {
    const table = await getOrderListByUserId({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

const onSubmit = () => {
    page.value = 1
    pageSize.value = 10
    getTableData()
    getPictureData()
}

const onReset = () => {
    searchInfo.value = { startCreatedAt: null, endCreatedAt: null }
    getTableData()
    getPictureData()
}

getTableData()


//=======图形部分========

// 定义后端返回的json的key和value
type resType = {
    date: string
    totalCost: number
    totalKwh: number
    serviceCost: number
    chargeCost: number
}
const chart = ref();

// 初始化echarts图例
var myChart: echarts.ECharts
nextTick(() => { myChart = echarts.init(chart.value) })
// 初始化echarts要用到的dataset
const dataSet = ref<(number | string)[][]>([])
// 存储返回回来的json列表
const pictureData = ref<resType[]>([])


const setOption = (dataset) => {
    let option = {
        legend: {},
        tooltip: {
            trigger: 'axis',
            showContent: true
        },
        dataset: {
            dimensions: ['日期', '充电度数', '充电总金额', '服务费用', '充电费用'],
            source: dataset
        },
        xAxis: { type: 'category' },
        yAxis: { gridIndex: 0 },
        dataZoom: [
            {
                type: 'inside',
            },
        ],
        series: [
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' },
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
        ]
    };
    myChart.setOption(option)
}
// 从后端获取数据（默认是一个星期时间段），将数据设置到图形中
const getPictureData = async () => {
    const now = new Date()
    const lastWeek = new Date(now.getTime() - 3600 * 1000 * 24 * 7)
    // 用户选择了起始日期和结束日期就指定，否则默认一周
    const start = searchInfo.value.startCreatedAt == null ? lastWeek : searchInfo.value.startCreatedAt
    const end = searchInfo.value.endCreatedAt == null ? now : searchInfo.value.endCreatedAt
    console.log(start)
    console.log(end)
    const resData = await getDurationReportInfo({ date: start, endDate: end })
    if (resData.code == 0) {
        pictureData.value = resData.data
    }
    // 清空dataSet的旧值(后续调用这个函数的时候可能要更换dataSet数据，所以清空旧值)
    dataSet.value = []
    //将json列表转换成列表,dataSet存储最后的列表
    let keyList = Object.keys(pictureData.value[0])
    let len = keyList.length
    for (let i = 0; i < len; i++) {
        let tempList: (number | string)[] = []
        for (let j = 0; j < pictureData.value.length; j++) {
            tempList.push(pictureData.value[j][keyList[i]])
        }
        dataSet.value.push(tempList)
    }
    console.log(dataSet.value)
    setOption(dataSet.value)
}

getPictureData()

</script>  

<style scoped>
:global(h2#card-usage ~ .example .example-showcase) {
    background-color: var(--el-fill-color) !important;
}

.el-statistic {
    --el-statistic-content-font-size: 28px;
}

.statistic-card {
    height: 100%;
    padding: 20px;
    border-radius: 4px;
    background-color: var(--el-bg-color-overlay);
}

.statistic-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
    font-size: 12px;
    color: var(--el-text-color-regular);
    margin-top: 16px;
}

.statistic-footer .footer-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.statistic-footer .footer-item span:last-child {
    display: inline-flex;
    align-items: center;
    margin-left: 4px;
}

.green {
    color: var(--el-color-success);
}

.red {
    color: var(--el-color-error);
}
</style>
  