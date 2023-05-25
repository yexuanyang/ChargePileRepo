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
            <el-table :data="tableData" border show-summary :summary-method="getSummaries"
                style="width: 100%;margin-top: 50px;display:inline-block;"
                :default-sort="{ prop: 'userId', order: 'descending' }">
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
import { onMounted, ref } from "vue"
import { dateEquals, type TableColumnCtx } from 'element-plus'
import { getOrderListByUserId } from '../../api/order'
import { formatDate } from '../../utils/format'
import { computed } from '@vue/reactivity'
import { getDurationTotalCharge, getDurationTotalPrice } from '../../api/report'
import { async } from 'q'

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
    dayChargeRate.value = todayChargeKwh.value - yesterdayChargeKwh.value
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
    monthChargeRate.value = monthChargeKwh.value - LastMonthChargeKwh.value
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
    dayChargePriceRate.value = todayChargePrice.value - yesterdayChargePrice.value
}


//========获取要展示的比较数据==========
const initData = async () => {
    await getDayChargeRate()
    await getMonthChargeRate()
    await getDayChargePriceRate()
}

initData()




// =======图表部分========
const chart = ref();

onMounted(() => {
    init()
})

const init = () => {
    const myChart = echarts.init(chart.value);

    // 此处粘贴图表代码
    let option = {
        legend: {},
        tooltip: {
            trigger: 'axis',
            showContent: false
        },
        dataset: {
            source: [
                ['product', '2012', '2013', '2014', '2015', '2016', '2017'],
                ['Milk Tea', 56.5, 82.1, 88.7, 70.1, 53.4, 85.1],
                ['Matcha Latte', 51.1, 51.4, 55.1, 53.3, 73.8, 68.7],
                ['Cheese Cocoa', 40.1, 62.2, 69.5, 36.4, 45.2, 32.5],
                ['Walnut Brownie', 25.2, 37.1, 41.2, 18, 33.9, 49.1]
            ]
        },
        xAxis: { type: 'category' },
        yAxis: { gridIndex: 0 },
        grid: { top: '55%' },
        series: [
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
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'pie',
                id: 'pie',
                radius: '30%',
                center: ['50%', '25%'],
                emphasis: {
                    focus: 'self'
                },
                label: {
                    formatter: '{b}: {@2012} ({d}%)'
                },
                encode: {
                    itemName: 'product',
                    value: '2012',
                    tooltip: '2012'
                }
            }
        ]
    }

    myChart.setOption(option)
}


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
                        return prev + curr
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
const searchInfo = ref({})

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

getTableData()

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
  