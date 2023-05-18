<template>
  <el-table :data="tableData" style="width: 100%" height="700" >
    <el-table-column
      fixed
      prop="ID"
      label="充电桩ID"
      sortable
      width="180"
      column-key="ID"
    />
    <el-table-column prop="waitNum" label="排队车辆数量" width="120" />
    <el-table-column prop="chargeType" label="充电类型" width="120" />
    <el-table-column
      prop="state"
      label="工作状态"
      width="100"
      :filters="[
        { text: '工作中', value: '工作中' },
        { text: '待机中', value: '待机中' },
      ]"
      :filter-method="filterTag"
      filter-placement="bottom-end"
    >
      <template #default="scope">
        <el-tag
          :type="scope.row.state === '工作中' ? '' : 'success'"
          disable-transitions
          >{{ scope.row.state }}
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="power" label="充电功率" width="120"></el-table-column>
    <el-table-column prop="location" label="充电桩位置" ></el-table-column>
  </el-table>
</template>

<script lang="ts" setup>
import type {TableColumnCtx} from 'element-plus'

interface ChargePile {
  ID: number
  waitNum: number
  chargeType: string
  state: string
  power: string
  location: string
  algoType: string
}

const workState = {
  work:'工作中',
  wait:'待机中'
}

const chargeType={
  fast:'快充',
  slow:'慢充'
}

const filterTag = (value: string, row: ChargePile) => {
  return row.state === value
}

const filterHandler = (
  value: string,
  row: ChargePile,
  column: TableColumnCtx<ChargePile>
) => {
  const property = column['property']
  return row[property] === value
}

const tableData: ChargePile[] = [
  {
    ID: 1,
    waitNum: 21,
    chargeType: chargeType.fast,
    state: workState.work,
    power: '1231w',
    location:'教三后',
    algoType:'先来先服务'
  },
  {
    ID: 2,
    waitNum: 21,
    chargeType: chargeType.fast,
    state: workState.work,
    power: '1231w',
    location:'教三后',
    algoType:'先来先服务'
  },
  {
    ID: 3,
    waitNum: 21,
    chargeType: chargeType.fast,
    state: workState.work,
    power: '1231w',
    location:'教三后',
    algoType:'先来先服务'
  },
  {
    ID: 4,
    waitNum: 21,
    chargeType: chargeType.fast,
    state: workState.work,
    power: '1231w',
    location:'教三后',
    algoType:'先来先服务'
  },
  {
    ID: 5,
    waitNum: 21,
    chargeType: chargeType.fast,
    state: workState.work,
    power: '1231w',
    location:'教三后',
    algoType:'先来先服务'
  },
  {
    ID: 6,
    waitNum: 21,
    chargeType: chargeType.fast,
    state: workState.work,
    power: '1231w',
    location:'教三后',
    algoType:'先来先服务'
  },
]
</script>
