import service from '@/utils/request'


// 传送一个日期date，根据token中的id和日期date获取当日的充电总度数
// 返回一个total float64
export const getDurationTotalCharge = (data) => {
    return service({
        url: '/report/getDurationTotalCharge',
        method: 'post',
        data: data
    })
}

export const getDurationTotalPrice = (data) => {
    return service({
        url: '/report/getDurationTotalPrice',
        method: 'post',
        data: data
    })
}

export const getDurationChargeInfo = (data) => {
    return service({
        url: '/report/getDurationChargeInfo',
        method: 'post',
        data: data
    })
}

export const getDurationReportInfo = (data)=>{
    return service({
        url: '/report/getDurationReportInfo',
        method: 'post',
        data: data
    })
}
