<template>
    <el-row :gutter="20">
        <el-col :span="8">
            <div class="statistic-card">
                <el-statistic :value="98500">
                    <template #title>
                        <div style="display: inline-flex; align-items: center">
                            Daily active users
                            <el-tooltip effect="dark" content="Number of users who logged into the product in one day"
                                placement="top">
                                <el-icon style="margin-left: 4px" :size="12">
                                    <Warning />
                                </el-icon>
                            </el-tooltip>
                        </div>
                    </template>
                </el-statistic>
                <div class="statistic-footer">
                    <div class="footer-item">
                        <span>than yesterday</span>
                        <span class="green">
                            24%
                            <el-icon>
                                <CaretTop />
                            </el-icon>
                        </span>
                    </div>
                </div>
            </div>
        </el-col>
        <el-col :span="8">
            <div class="statistic-card">
                <el-statistic :value="693700">
                    <template #title>
                        <div style="display: inline-flex; align-items: center">
                            Monthly Active Users
                            <el-tooltip effect="dark" content="Number of users who logged into the product in one month"
                                placement="top">
                                <el-icon style="margin-left: 4px" :size="12">
                                    <Warning />
                                </el-icon>
                            </el-tooltip>
                        </div>
                    </template>
                </el-statistic>
                <div class="statistic-footer">
                    <div class="footer-item">
                        <span>month on month</span>
                        <span class="red">
                            12%
                            <el-icon>
                                <CaretBottom />
                            </el-icon>
                        </span>
                    </div>
                </div>
            </div>
        </el-col>
        <el-col :span="8">
            <div class="statistic-card">
                <el-statistic :value="72000" title="New transactions today">
                    <template #title>
                        <div style="display: inline-flex; align-items: center">
                            New transactions today
                        </div>
                    </template>
                </el-statistic>
                <div class="statistic-footer">
                    <div class="footer-item">
                        <span>than yesterday</span>
                        <span class="green">
                            16%
                            <el-icon>
                                <CaretTop />
                            </el-icon>
                        </span>
                    </div>
                    <div class="footer-item">
                        <el-icon :size="14">
                            <ArrowRight />
                        </el-icon>
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
            <el-table :data="tableData" border show-summary style="width: 100%;margin-top: 50px;display:inline-block;"
                :default-sort="{ prop: 'userId', order: 'descending' }">
                <el-table-column align="left" label="日期" width="180" fixed prop="CreatedAt" sortable>
                    <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="充电的车牌号" prop="carId" width="120" />
                <el-table-column align="left" label="充电类型" prop="chargeType" width="120" />
                <el-table-column align="left" label="充电费用" prop="chargeCost" width="120" />
                <el-table-column align="left" label="充电度数" prop="kwh" width="120" />
                <el-table-column align="left" label="充电时长" width="180">
                    <template #default="scope">{{ formatDate(scope.row.time) }}</template>
                </el-table-column>
                <el-table-column align="left" label="充电桩id" prop="pileId" width="120" />
                <el-table-column align="left" label="服务费" prop="serviceCost" width="120" />
                <el-table-column align="left" label="开始充电时间" width="180">
                    <template #default="scope">{{ formatDate(scope.row.startedAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="结束充电时间" width="180">
                    <template #default="scope">{{ formatDate(scope.row.stopAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="总花费" prop="totalCost" width="120" />
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
import type { TableColumnCtx } from 'element-plus'
import { getOrderListByUserId } from '../../api/order'
import { formatDate } from '../../utils/format'

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

interface Product {
    id: string
    name: string
    amount1: string
    amount2: string
    amount3: number
}

interface SummaryMethodProps<T = Product> {
    columns: TableColumnCtx<T>[]
    data: T[]
}

const getSummaries = (param: SummaryMethodProps) => {
    const { columns, data } = param
    const sums: string[] = []
    columns.forEach((column, index) => {
        if (index === 0) {
            sums[index] = 'Total Cost'
            return
        }
        const values = data.map((item) => Number(item[column.property]))
        if (!values.every((value) => Number.isNaN(value))) {
            sums[index] = `$ ${values.reduce((prev, curr) => {
                const value = Number(curr)
                if (!Number.isNaN(value)) {
                    return prev + curr
                } else {
                    return prev
                }
            }, 0)}`
        } else {
            sums[index] = 'N/A'
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
  