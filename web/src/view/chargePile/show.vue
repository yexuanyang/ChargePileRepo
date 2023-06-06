<template>
    <div class="page">
        <div class="gva-card-box">
            <div class="gva-card" v-if="active===0">
                <div class="card-header">
                    <span>提交订单</span>
                </div>
                <div class="gva-card gva-order-card">
                    <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
                        <el-form-item label="充电的车牌号:" prop="carId">
                            <el-input v-model="formData.carId" style="width:100%" :clearable="true"
                                      placeholder="请输入"/>
                        </el-form-item>
                        <el-form-item label="充电类型:" prop="chargeType">
                            <el-select
                                v-model="formData.chargeType"
                                placeholder="请选择"
                                style="width:60%"
                                :clearable="true"
                            >
                                <el-option
                                    v-for="item in ['快充', '慢充']"
                                    :key="item"
                                    :label="item"
                                    :value="item"
                                />
                            </el-select>
                        </el-form-item>
                        <el-form-item label="充电度数:" prop="kwh">
                            <el-input-number v-model="formData.kwh" style="width:100%" :precision="2"
                                             :clearable="true"/>
                        </el-form-item>
                        <el-form-item label="充电时长:" prop="time">
                            <el-input-number v-model="formData.time" style="width:100%" :precision="2"
                                             :clearable="true"/>
                        </el-form-item>
                        <el-form-item label="结束充电时间:" prop="stopAt" style="width: 100%">
                            <el-col :span="11">
                                <el-date-picker
                                    v-model="formData.date1"
                                    type="date"
                                    placeholder="选择日期"
                                    style="width: 100%;"
                                />
                            </el-col>
                            <el-col class="line" :span="1">---</el-col>
                            <el-col :span="11">
                                <el-time-picker
                                    v-model="formData.date2"
                                    placeholder="选择时间"
                                    style="width: 100%;"
                                />
                            </el-col>
                        </el-form-item>
                    </el-form>
                </div>
            </div>
            <div class="gva-card" v-if="active===1">
                <div class="card-header">
                    <span>确认订单</span>
                </div>
                <div class="gva-card gva-order-card">
                </div>
            </div>
            <div class="gva-card" v-if="active===2">
                <div class="card-header">
                    <span>开始充电</span>
                </div>
                <div class="gva-card gva-order-card">
                </div>
            </div>
            <div class="gva-card gva-steps">
                <el-steps :active="active" finish-status="success">
                    <el-step title="提交订单"/>
                    <el-step title="确认订单"/>
                    <el-step title="开始充电"/>
                </el-steps>
                <div class="gva-button">
                    <el-button v-if="active!==0" style="margin-top: 12px;margin-right: 0px" @click="lastStep">上一步</el-button>
                    <el-button v-if="active!==2" style="margin-top: 12px;" @click="enterOrder">下一步</el-button>
                    <el-button v-if="active===2" disabled style="margin-top: 12px;">已完成</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import {reactive, ref} from 'vue'

const active = ref(0)
const nextButtonVisable = ref(true)
const lastButtonVisable = ref(true)
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

const elFormRef = ref()
// 验证规则
const rule = reactive({
    carId: [{
        required: true,
        message: '请输入格式正确的车牌号',
        trigger: ['input', 'blur'],
    }],
    chargeType: [{
        required: true,
        message: '请选择充电类型',
        trigger: ['input', 'blur'],
    }],
    kwh: [{
        required: true,
        message: '',
        trigger: ['input', 'blur'],
    }],
    startedAt: [{
        required: true,
        message: '',
        trigger: ['input', 'blur'],
    }],
})

// 返回上一步
function lastStep() {
    active.value--
}

// 提交订单
async function enterOrder() {
    if (active.value === 0) {
        elFormRef.value?.validate(async (valid) => {
            if (!valid) return
            else {
                active.value++
            }
        })
        return
    }
    active.value++
}
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
    justify-content: center;
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
