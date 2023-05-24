<template>
    <el-select v-model="stationInfo" placeholder="请选择" style="margin-bottom: 10px">
        <el-option
            v-for="item in stationOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        />
    </el-select>
    <el-button icon="search" type="primary" style="margin-left: 10px; margin-bottom: 10px" @click="GeoSearch">查询
    </el-button>
    <el-button icon="search" type="primary" style="margin-left: 10px; margin-bottom: 10px" @click="Geolocation">定位
    </el-button>
    <div id="container"/>
</template>

<script setup>
import AMapLoader from '@amap/amap-jsapi-loader'
import {ref, shallowRef} from 'vue'

let map = shallowRef(null)
let stationCenter = ref([])
let geolocation = ref('')
const stationInfo = ref(1)
const stationOptions = [
    {
        value: 1,
        label: '北京邮电大学充电站'
    },
    {
        value: 2,
        label: '北京市东城区长安街充电站'
    }
]

const initMap = (stationCenter) => {
    console.log(stationCenter)
    AMapLoader.load({
        key: 'befd6edaafb970016694fb14bc220719',
        version: '2.0',
        plugins: ['AMap.ToolBar', 'AMap.Scale', 'AMap.Driving', 'AMap.AutoComplete',
            'AMap.PlaceSearch', 'AMap.Marker', 'AMap.Icon', 'AMap.Geolocation'],
    }).then((AMap) => {
        map = new AMap.Map('container', {
            viewMode: '2D', // 是否为3D地图模式
            zoom: 16, // 初始化地图级别
            center: stationCenter, // 初始化地图中心点位置
        })
        // 点标记显示内容，HTML要素字符串
        const icon = new AMap.Icon({
            size: new AMap.Size(300, 600), // 图标尺寸
            image: 'src/assets/img_1.png',
            imageOffset: new AMap.Pixel(0, 0), // 图像相对展示区域的偏移量x
            imageSize: new AMap.Size(23, 30), // 图标所用图片大小
        })
        const marker1 = new AMap.Marker({
            position: [116.361124, 39.959828],
            offset: new AMap.Pixel(0, 0),
            icon: icon,
            title: '北京邮电大学充电站',
            defaultCursor: 'pointer'
        })
        const marker2 = new AMap.Marker({
            position: [116.399343, 39.908011],
            offset: new AMap.Pixel(-10, -10),
            icon: icon,
            title: '北京市东城区长安街充电站',
            defaultCursor: 'pointer'
        })
        const toolbar = new AMap.ToolBar()
        const scale = new AMap.Scale()
        geolocation = new AMap.Geolocation
        map.add(marker1)
        map.add(marker2)
        map.addControl(toolbar)
        map.addControl(scale)
    }).catch(e => {
        console.log(e)
    })
}
initMap([116.361124, 39.959828])

function GeoSearch() {
    if (stationInfo.value === 1) {
        stationCenter = [116.361124, 39.959828]
    } else {
        stationCenter = [116.399343, 39.908011]
    }
    initMap(stationCenter)
}

function Geolocation() {
    geolocation.getCurrentPosition((status, result) => {
        if (result && result.position) {
            const lng = result.position.lng
            const lat = result.position.lat
            stationCenter = [lng, lat]
            self.loaded = true
        }
        initMap(stationCenter)
    })
}

</script>

<style scoped lang="scss">
#container {
    padding: 0px;
    margin: 0px;
    width: 100%;
    height: 400px;
}

/* 隐藏高德logo  */
.amap-logo {
    display: none !important;
}

/* 隐藏高德版权  */
.amap-copyright {
    display: none !important;
}

.custom-content-marker {
    position: relative;
    width: 25px;
    height: 34px;

    img {
        width: 50px;
        height: 50px;
    }
}
</style>
