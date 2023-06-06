<template>
    <div class="block">
        <el-form :inline="true" class="demo-form-inline" v-model="searchInfo" @keyup.enter="onSubmit">
            <el-form-item label="报表时间">
                <el-date-picker v-model="searchInfo.beginDate" type="datetime" placeholder="选择报表开始时间"
                    :shortcuts="shortcuts" />
                ——
                <el-date-picker v-model="searchInfo.endDate" type="datetime" placeholder="选择报表结束时间"
                    :shortcuts="shortcuts" />
            </el-form-item>
            <el-form-item label="充电桩ID">
                <el-input v-model="pileId" placeholder="选择生成报表的充电桩ID" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                <el-button icon="refresh" @click="onReset">重置</el-button>
            </el-form-item>
        </el-form>
        <el-row>
            <el-descriptions class="margin-top" title="充电桩报表" :column="7" :size="size" border>
                <el-descriptions-item>
                    <template #label>
                        <div class="cell-item">
                            充电桩ID
                        </div>
                    </template>
                    <div class="cell-item">
                        {{ reportInfo.pileId }}
                    </div>
                </el-descriptions-item>
                <el-descriptions-item>
                    <template #label>
                        <div class="cell-item">
                            充电桩累计充电次数
                        </div>
                    </template>
                    <div class="cell-item">
                        {{ reportInfo.chargeCount }}
                    </div>
                </el-descriptions-item>
                <el-descriptions-item>
                    <template #label>
                        <div class="cell-item">
                            充电桩累计充电时间（小时）
                        </div>
                    </template>
                    <div class="cell-item">
                        {{ reportInfo.chargeTime }}
                    </div>
                </el-descriptions-item>
                <el-descriptions-item>
                    <template #label>
                        <div class="cell-item">
                            充电桩累计充电度数(kw*h)
                        </div>
                    </template>
                    <div class="cell-item">
                        {{ reportInfo.chargeElectricity }}
                    </div>
                </el-descriptions-item>
                <el-descriptions-item>
                    <template #label>
                        <div class="cell-item">
                            充电桩累计充电费用(￥)
                        </div>
                    </template>
                    <div class="cell-item">
                        {{ reportInfo.chargeCost }}
                    </div>
                </el-descriptions-item>
                <el-descriptions-item>
                    <template #label>
                        <div class="cell-item">
                            充电桩累计服务费用(￥)
                        </div>
                    </template>
                    <div class="cell-item">
                        {{ reportInfo.serviceCost }}
                    </div>
                </el-descriptions-item>
                <el-descriptions-item>
                    <template #label>
                        <div class="cell-item">
                            充电桩累计总费用(￥)
                        </div>
                    </template>
                    <div class="cell-item">
                        {{ reportInfo.totalCost }}
                    </div>
                </el-descriptions-item>
            </el-descriptions>
        </el-row>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { getDurationChargeInfo } from '../../api/report';
import { ElMessage } from 'element-plus';

const searchInfo = ref({})
const beginDate = ref('')
const endDate = ref('')
const pileId = ref(0)
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
const reportInfo = ref({})

const successMessage = () => {
    ElMessage({
        type: 'success',
        message: '查询成功',
    })
}
const errorMessage = (message) => {
    ElMessage({
        type: 'error',
        message: messgae
    })
}
const getReportInfo = async () => {
    const Info = await getDurationChargeInfo({ pileId: Number(pileId.value), ...searchInfo.value })
    if (Info.code == 0) {
        reportInfo.value = Info.data
        successMessage()
    }
    else {
        errorMessage(Info.msg)
    }
}
const onSubmit = () => {
    getReportInfo()
}
const onReset = () => {
    beginDate.value = ''
    endDate.value = ''
}
</script>

<style>
.el-descriptions {
    margin-top: 20px;
}

.cell-item {
    display: flex;
    align-items: center;
}

.margin-top {
    margin-top: 20px;
}
</style>