<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
            <div class="gva-top-card-left-title">你好 智能充电桩管理员</div>
            <div class="gva-top-card-left-rows">
                <div class="card-overview">
                    <img src="../../assets/chargePile.png">
                    <i class="fa fa-map" @click="dialogMapVisible = true"></i>
                    <el-button type="text" @click="dialogMapVisible = true">查看详细地址</el-button>
                </div>
                <el-dialog v-model="dialogMapVisible" title="充电站详细地址">
                    <MapContainer></MapContainer>
                    <span slot="footer" class="dialog-footer" style="display: flex; justify-content: right">
                      <el-button type="primary" style="margin-top: 10px"
                                 @click="dialogMapVisible = false">确 定</el-button>
                  </span>
                </el-dialog>
            </div>
        </div>
        <img src="@/assets/dashboard.png" class="gva-top-card-right" alt>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>数据统计</span>
        </div>
        <div class="gva-data-statistics">
          <el-row :gutter="20">
            <el-col :span="6" class="gva-data-row">
              <div>
                <el-statistic title="站点数目" group-separator="," :value="dataStatistic.stationNum" />
              </div>
            </el-col>
            <el-col :span="6" class="gva-data-row">
              <div>
                <el-statistic title="用户总数" group-separator="," :value="dataStatistic.userNum" />
              </div>
            </el-col>
            <el-col :span="6" class="gva-data-row">
              <div>
                  <el-statistic title="充电桩数目" group-separator="," :value="dataStatistic.pileNum"/>
              </div>
            </el-col>
            <el-col :span="6" class="gva-data-row">
              <div>
                  <el-statistic title="充电站内汽车数量" group-separator="," :value="dataStatistic.unFinishedNum"/>
              </div>
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="bottom-items">
          <div class="right">
            <el-row :gutter="20">
              <el-col
                v-for="(card, key) in toolCards"
                :key="key"
                :span="6"
                :xs="8"
                class="quick-entrance-items"
                @click="toTarget(card.name)"
              >
                <div class="quick-entrance-item">
                  <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                    <el-icon>
                      <component :is="card.icon" :style="{ color: card.color }" />
                    </el-icon>
                  </div>
                  <p>{{ card.label }}</p>
                </div>
              </el-col>
            </el-row>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup>
import {getUserList} from '@/api/user'
import {reactive, ref} from 'vue'
import {useRouter} from 'vue-router'
import {getChargeStationList} from '@/api/chargeStation'
import MapContainer from "@/view/chargePile/MapContainer.vue";
import {getChargePileList} from "@/api/chargePile";
import {getUnFinishedOrderNumber} from "@/api/order";
import Icon from "@/view/superAdmin/menu/icon.vue";

const dialogMapVisible = ref(false)

const total = ref(0)
const toolCards = ref([
    {
        label: '用户管理',
        icon: 'monitor',
        name: 'user',
        color: '#ff9c6e',
        bg: 'rgba(255, 156, 110,.3)'
    },
    {
        label: '角色管理',
        icon: 'setting',
        name: 'authority',
        color: '#69c0ff',
        bg: 'rgba(105, 192, 255,.3)'
    },
    {
        label: '菜单管理',
        icon: 'menu',
        name: 'menu',
        color: '#b37feb',
        bg: 'rgba(179, 127, 235,.3)'
    },
    {
        label: '充电桩管理',
        icon: 'cpu',
        name: 'chargePile',
        color: '#ffd666',
        bg: 'rgba(255, 214, 102,.3)'
    },
])

const dataStatistic = reactive({
    stationNum: 0,
    userNum: 0,
    pileNum: 0,
    unFinishedNum: 0,
})
const getTableData = async () => {
    const res = await getUserList({page: 1, pageSize: 10})
    if (res.code === 0) {
        dataStatistic.userNum = res.data.total
    }
}
const getStationData = async () => {
    const res = await getChargeStationList()
    if (res.code === 0) {
        dataStatistic.stationNum = res.data.total
    }
}
const getChargePileData = async () => {
    const res = await getChargePileList()
    if (res.code === 0) {
        dataStatistic.pileNum = res.data.total
    }
}
const getOrderData = async () => {
    const res = await getUnFinishedOrderNumber()
    if (res.code === 0) {
        dataStatistic.unFinishedNum = res.data
    }
}

getTableData()
getStationData()
getChargePileData()
getOrderData()

const router = useRouter()

const toTarget = (name) => {
    router.push({name})
}

</script>
<script>
export default {
  name: 'Dashboard'
}
</script>

<style scoped lang="scss">
@mixin flex-center {
    display: flex;
    align-items: center;
}

.gva-data-row {
    .el-statistic {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;

      .el-statistic__head {
        font-size: 15px;
      }
    }
}

.bottom-items {
  display: flex;
}

.card-overview {
  display: flex;
  align-items: center;

  img {
    height: 100px;
    width: 100px;
  }
  i {
    font-size: 30px;
    margin-left: 20px;
    margin-right: 10px;
    cursor: pointer;
  }
}


.right {
  width: 1000px;

  &::before {
    padding-left: 200px;
  }
}

.page {
    background: #f0f2f5;
    padding: 0;

    .gva-card-box {
        padding: 12px 16px;

        & + .gva-card-box {
            padding-top: 0px;
        }
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

    .gva-top-card {
        height: 260px;
        @include flex-center;
        justify-content: space-between;
        color: #777;

        &-left {
            height: 100%;
            display: flex;
            flex-direction: column;

            &-title {
                font-size: 22px;
                color: #343844;
            }

            &-dot {
                font-size: 16px;
                color: #6B7687;
                margin-top: 24px;
            }

            &-rows {
                // margin-top: 15px;
                margin-top: 18px;
                color: #6B7687;
                width: 600px;
                align-items: center;
            }

            &-item {
                + .gva-top-card-left-item {
                    margin-top: 24px;
                }

                margin-top: 14px;
            }
        }

        &-right {
            height: 600px;
            width: 600px;
            margin-top: 28px;
        }
    }

    ::v-deep(.el-card__header) {
        padding: 0;
        border-bottom: none;
    }

    .card-header {
        padding-bottom: 20px;
        border-bottom: 1px solid #e8e8e8;
        font-size: 20px;
    }

    .quick-entrance-title {
        height: 30px;
        font-size: 22px;
        color: #333;
        width: 100%;
        border-bottom: 1px solid #eee;
    }

    .quick-entrance-items {
        @include flex-center;
        justify-content: center;
        text-align: center;
        color: #333;

        .quick-entrance-item {
            padding: 16px 28px;
            margin-top: -16px;
            margin-bottom: -16px;
            border-radius: 4px;
            transition: all 0.2s;

            &:hover {
                box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
            }

            cursor: pointer;
            height: auto;
            text-align: center;
            // align-items: center;
            &-icon {
                width: 50px;
                height: 50px !important;
                border-radius: 8px;
                @include flex-center;
                justify-content: center;
                margin: 0 auto;

                i {
                    font-size: 24px;
                }
            }

            p {
                margin-top: 10px;
            }
        }
    }

    .echart-box {
        padding: 14px;
    }
}

.dashboard-icon {
    font-size: 20px;
    color: rgb(85, 160, 248);
    width: 30px;
    height: 30px;
    margin-right: 10px;
    @include flex-center;
}

.flex-center {
    @include flex-center;
}

//小屏幕不显示右侧，将登录框居中
@media (max-width: 750px) {
    .gva-card {
        padding: 20px 10px !important;

        .gva-top-card {
            height: auto;

            &-left {
                &-title {
                    font-size: 20px !important;
                }

                &-rows {
                    margin-top: 15px;
                    align-items: center;
                }
            }

            &-right {
                display: none;
            }
        }

        .gva-middle-card {
            &-item {
                line-height: 20px;
            }
        }

        .dashboard-icon {
            font-size: 18px;
        }
    }
}
</style>

