import { asyncRouterHandle } from '@/utils/asyncRouter'
import { emitter } from '@/utils/bus.js'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { asyncMenu } from '@/api/menu'

const notLayoutRouterArr = []
const keepAliveRoutersArr = []
const nameMap = {}

const formatRouter = (routes, routeMap) => {
  routes && routes.forEach(item => {
    item.meta.btns = item.btns
    item.meta.hidden = item.hidden
    if (item.meta.defaultMenu === true) {
      notLayoutRouterArr.push({
        ...item,
        path: `/${item.path}`,
      })
    } else {
      routeMap[item.name] = item
      if (item.children && item.children.length > 0) {
        formatRouter(item.children, routeMap)
      }
    }
  })
}

const KeepAliveFilter = (routes) => {
  routes && routes.forEach(item => {
    // 子菜单中有 keep-alive 的，父菜单也必须 keep-alive，否则无效。这里将子菜单中有 keep-alive 的父菜单也加入。
    if ((item.children && item.children.some(ch => ch.meta.keepAlive) || item.meta.keepAlive)) {
      item.component && item.component().then(val => {
        keepAliveRoutersArr.push(val.default.name)
        nameMap[item.name] = val.default.name
      })
    }
    if (item.children && item.children.length > 0) {
      KeepAliveFilter(item.children)
    }
  })
}

export const useRouterStore = defineStore('router', () => {
  const keepAliveRouters = ref([])
  const asyncRouterFlag = ref(0)
  const setKeepAliveRouters = (history) => {
    const keepArrTemp = []
    history.forEach(item => {
      if (nameMap[item.name]) {
        keepArrTemp.push(nameMap[item.name])
      }
    })
    keepAliveRouters.value = Array.from(new Set(keepArrTemp))
  }
  emitter.on('setKeepAlive', setKeepAliveRouters)

  const asyncRouters = ref([])
  const routeMap = ({})
  // 获取动态路由
  const SetAsyncRouter = async(type) => {
    asyncRouterFlag.value++
    const baseRouter = [{
      path: '/layout',
      name: 'layout',
      component: 'view/layout/index.vue',
      meta: {
        title: '底层layout'
      },
      children: []
    }]

    const asyncRouterRes = await asyncMenu()
    const asyncRouter = asyncRouterRes.data.menus

    // let asyncRouter = [
    //   {
    //     'ID': 43,
    //     'CreatedAt': '2023-05-23T16:32:53.942+08:00',
    //     'UpdatedAt': '2023-06-02T18:13:03.147+08:00',
    //     'parentId': '0',
    //     'path': 'chargestationTest',
    //     'name': 'chargestationTest',
    //     'hidden': false,
    //     'component': 'view/chargePile/StationBoard.vue',
    //     'sort': 0,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '仪表盘',
    //       'icon': 'compass',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '43',
    //     'children': null,
    //     'parameters': [],
    //     'btns': null
    //   },
    //   {
    //     'ID': 3,
    //     'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //     'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //     'parentId': '0',
    //     'path': 'admin',
    //     'name': 'superAdmin',
    //     'hidden': false,
    //     'component': 'view/superAdmin/index.vue',
    //     'sort': 3,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '超级管理员',
    //       'icon': 'user',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '3',
    //     'children': [
    //       {
    //         'ID': 4,
    //         'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'parentId': '3',
    //         'path': 'authority',
    //         'name': 'authority',
    //         'hidden': false,
    //         'component': 'view/superAdmin/authority/authority.vue',
    //         'sort': 1,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '角色管理',
    //           'icon': 'avatar',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '4',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 9,
    //         'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'parentId': '3',
    //         'path': 'dictionaryDetail/:id',
    //         'name': 'dictionaryDetail',
    //         'hidden': true,
    //         'component': 'view/superAdmin/dictionary/sysDictionaryDetail.vue',
    //         'sort': 1,
    //         'meta': {
    //           'activeName': 'dictionary',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '字典详情-${id}',
    //           'icon': 'list',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '9',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 5,
    //         'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'parentId': '3',
    //         'path': 'menu',
    //         'name': 'menu',
    //         'hidden': false,
    //         'component': 'view/superAdmin/menu/menu.vue',
    //         'sort': 2,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': true,
    //           'defaultMenu': false,
    //           'title': '菜单管理',
    //           'icon': 'tickets',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '5',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 6,
    //         'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'parentId': '3',
    //         'path': 'api',
    //         'name': 'api',
    //         'hidden': false,
    //         'component': 'view/superAdmin/api/api.vue',
    //         'sort': 3,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': true,
    //           'defaultMenu': false,
    //           'title': 'api管理',
    //           'icon': 'platform',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '6',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 7,
    //         'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'parentId': '3',
    //         'path': 'user',
    //         'name': 'user',
    //         'hidden': false,
    //         'component': 'view/superAdmin/user/user.vue',
    //         'sort': 4,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '用户管理',
    //           'icon': 'coordinate',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '7',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 8,
    //         'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'parentId': '3',
    //         'path': 'dictionary',
    //         'name': 'dictionary',
    //         'hidden': false,
    //         'component': 'view/superAdmin/dictionary/sysDictionary.vue',
    //         'sort': 5,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '字典管理',
    //           'icon': 'notebook',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '8',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 10,
    //         'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //         'parentId': '3',
    //         'path': 'operation',
    //         'name': 'operation',
    //         'hidden': false,
    //         'component': 'view/superAdmin/operation/sysOperationRecord.vue',
    //         'sort': 6,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '操作历史',
    //           'icon': 'pie-chart',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '10',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       }
    //     ],
    //     'parameters': [],
    //     'btns': null
    //   },
    //   {
    //     'ID': 11,
    //     'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //     'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //     'parentId': '0',
    //     'path': 'person',
    //     'name': 'person',
    //     'hidden': true,
    //     'component': 'view/person/person.vue',
    //     'sort': 4,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '个人信息',
    //       'icon': 'message',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '11',
    //     'children': null,
    //     'parameters': [],
    //     'btns': null
    //   },
    //   {
    //     'ID': 41,
    //     'CreatedAt': '2023-05-22T16:13:09.598+08:00',
    //     'UpdatedAt': '2023-06-02T18:15:15.458+08:00',
    //     'parentId': '0',
    //     'path': 'carRoot',
    //     'name': 'carRoot',
    //     'hidden': false,
    //     'component': 'view/routerHolder.vue',
    //     'sort': 4,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '车辆',
    //       'icon': 'aim',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '41',
    //     'children': [
    //       {
    //         'ID': 39,
    //         'CreatedAt': '2023-05-22T16:03:30.189+08:00',
    //         'UpdatedAt': '2023-05-22T16:32:04.056+08:00',
    //         'parentId': '41',
    //         'path': 'car',
    //         'name': 'car',
    //         'hidden': false,
    //         'component': 'view/car/car.vue',
    //         'sort': 99,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '汽车列表',
    //           'icon': 'aim',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '39',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       }
    //     ],
    //     'parameters': [],
    //     'btns': null
    //   },
    //   {
    //     'ID': 33,
    //     'CreatedAt': '2023-05-20T15:46:04.108+08:00',
    //     'UpdatedAt': '2023-05-21T21:14:25.629+08:00',
    //     'parentId': '0',
    //     'path': 'PileManage',
    //     'name': 'PileManage',
    //     'hidden': false,
    //     'component': 'view/routerHolder.vue',
    //     'sort': 10,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '充电桩',
    //       'icon': 'lightning',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '33',
    //     'children': [
    //       {
    //         'ID': 47,
    //         'CreatedAt': '2023-06-02T18:40:52.227+08:00',
    //         'UpdatedAt': '2023-06-02T18:40:52.227+08:00',
    //         'parentId': '33',
    //         'path': 'CSposition',
    //         'name': 'CSposition',
    //         'hidden': false,
    //         'component': 'view/chargeStation/userChargeStation.vue',
    //         'sort': 10,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '充电站位置',
    //           'icon': 'collection-tag',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '47',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 45,
    //         'CreatedAt': '2023-05-24T21:14:13.745+08:00',
    //         'UpdatedAt': '2023-05-24T21:14:28.641+08:00',
    //         'parentId': '33',
    //         'path': 'chargePileReport',
    //         'name': 'chargePileReport',
    //         'hidden': false,
    //         'component': 'view/chargePile/report.vue',
    //         'sort': 10,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '充电桩报表',
    //           'icon': 'briefcase',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '45',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 31,
    //         'CreatedAt': '2023-05-18T21:54:37.021+08:00',
    //         'UpdatedAt': '2023-05-21T21:22:40.181+08:00',
    //         'parentId': '33',
    //         'path': 'chargePile',
    //         'name': 'chargePile',
    //         'hidden': false,
    //         'component': 'view/chargePile/chargePile.vue',
    //         'sort': 10,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '充电桩详细信息',
    //           'icon': 'document',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '31',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 38,
    //         'CreatedAt': '2023-05-21T22:00:07.543+08:00',
    //         'UpdatedAt': '2023-06-02T18:37:21.829+08:00',
    //         'parentId': '33',
    //         'path': 'ChargeStation',
    //         'name': 'ChargeStation',
    //         'hidden': false,
    //         'component': 'view/chargeStation/chargeStation.vue',
    //         'sort': 11,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '充电站位置详细信息',
    //           'icon': 'add-location',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '38',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       }
    //     ],
    //     'parameters': [],
    //     'btns': null
    //   },
    //   {
    //     'ID': 35,
    //     'CreatedAt': '2023-05-21T14:27:27.42+08:00',
    //     'UpdatedAt': '2023-05-21T14:27:27.42+08:00',
    //     'parentId': '0',
    //     'path': 'orderRoot',
    //     'name': 'orderRoot',
    //     'hidden': false,
    //     'component': 'view/routerHolder.vue',
    //     'sort': 20,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '订单',
    //       'icon': 'box',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '35',
    //     'children': [
    //       {
    //         'ID': 42,
    //         'CreatedAt': '2023-05-23T16:09:30.272+08:00',
    //         'UpdatedAt': '2023-05-23T16:09:40.158+08:00',
    //         'parentId': '35',
    //         'path': 'report',
    //         'name': 'report',
    //         'hidden': false,
    //         'component': 'view/report/report.vue',
    //         'sort': 20,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '订单报表',
    //           'icon': 'arrow-down-bold',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '42',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 34,
    //         'CreatedAt': '2023-05-21T13:45:50.855+08:00',
    //         'UpdatedAt': '2023-05-21T14:27:34.117+08:00',
    //         'parentId': '35',
    //         'path': 'order',
    //         'name': 'order',
    //         'hidden': false,
    //         'component': 'view/order/order.vue',
    //         'sort': 20,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '订单详细信息',
    //           'icon': 'document-checked',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '34',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       }
    //     ],
    //     'parameters': [],
    //     'btns': null
    //   }
    // ]
    // const userRouter = [
    //   {
    //     'ID': 11,
    //     'CreatedAt': '2023-05-18T21:18:16.934+08:00',
    //     'UpdatedAt': '2023-05-18T21:18:16.934+08:00',
    //     'parentId': '0',
    //     'path': 'person',
    //     'name': 'person',
    //     'hidden': true,
    //     'component': 'view/person/person.vue',
    //     'sort': 4,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '个人信息',
    //       'icon': 'message',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '11',
    //     'children': null,
    //     'parameters': [],
    //     'btns': null
    //   },
    //   {
    //     'ID': 41,
    //     'CreatedAt': '2023-05-22T16:13:09.598+08:00',
    //     'UpdatedAt': '2023-06-02T18:15:15.458+08:00',
    //     'parentId': '0',
    //     'path': 'carRoot',
    //     'name': 'carRoot',
    //     'hidden': false,
    //     'component': 'view/routerHolder.vue',
    //     'sort': 4,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '车辆',
    //       'icon': 'aim',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '41',
    //     'children': [
    //       {
    //         'ID': 40,
    //         'CreatedAt': '2023-05-22T16:11:23.235+08:00',
    //         'UpdatedAt': '2023-05-22T16:32:19.173+08:00',
    //         'parentId': '41',
    //         'path': 'userCarList',
    //         'name': 'userCarList',
    //         'hidden': false,
    //         'component': 'view/car/car_list.vue',
    //         'sort': 99,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '我的汽车列表',
    //           'icon': 'aim',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '40',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       }
    //     ],
    //     'parameters': [],
    //     'btns': null
    //   },
    //   {
    //     'ID': 33,
    //     'CreatedAt': '2023-05-20T15:46:04.108+08:00',
    //     'UpdatedAt': '2023-05-21T21:14:25.629+08:00',
    //     'parentId': '0',
    //     'path': 'PileManage',
    //     'name': 'PileManage',
    //     'hidden': false,
    //     'component': 'view/routerHolder.vue',
    //     'sort': 10,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '充电桩',
    //       'icon': 'lightning',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '33',
    //     'children': [
    //       {
    //         'ID': 46,
    //         'CreatedAt': '2023-05-27T14:03:04.086+08:00',
    //         'UpdatedAt': '2023-06-02T18:10:32.612+08:00',
    //         'parentId': '33',
    //         'path': 'showTest',
    //         'name': 'showTest',
    //         'hidden': false,
    //         'component': 'view/chargePile/show.vue',
    //         'sort': 0,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '充电桩列表',
    //           'icon': 'goods-filled',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '46',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 38,
    //         'CreatedAt': '2023-05-21T22:00:07.543+08:00',
    //         'UpdatedAt': '2023-06-02T18:37:21.829+08:00',
    //         'parentId': '33',
    //         'path': 'ChargeStation',
    //         'name': 'ChargeStation',
    //         'hidden': false,
    //         'component': 'view/chargeStation/chargeStation.vue',
    //         'sort': 11,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '充电站位置详细信息',
    //           'icon': 'add-location',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '38',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       }
    //     ],
    //     'parameters': [],
    //     'btns': null
    //   },
    //   {
    //     'ID': 35,
    //     'CreatedAt': '2023-05-21T14:27:27.42+08:00',
    //     'UpdatedAt': '2023-05-21T14:27:27.42+08:00',
    //     'parentId': '0',
    //     'path': 'orderRoot',
    //     'name': 'orderRoot',
    //     'hidden': false,
    //     'component': 'view/routerHolder.vue',
    //     'sort': 20,
    //     'meta': {
    //       'activeName': '',
    //       'keepAlive': false,
    //       'defaultMenu': false,
    //       'title': '订单',
    //       'icon': 'box',
    //       'closeTab': false
    //     },
    //     'authoritys': null,
    //     'menuBtn': null,
    //     'menuId': '35',
    //     'children': [
    //       {
    //         'ID': 36,
    //         'CreatedAt': '2023-05-21T14:50:34.343+08:00',
    //         'UpdatedAt': '2023-05-21T14:52:22.416+08:00',
    //         'parentId': '35',
    //         'path': 'userOrder',
    //         'name': 'userOrder',
    //         'hidden': false,
    //         'component': 'view/order/userOrder.vue',
    //         'sort': 20,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '我的订单',
    //           'icon': 'chat-line-round',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '36',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       },
    //       {
    //         'ID': 42,
    //         'CreatedAt': '2023-05-23T16:09:30.272+08:00',
    //         'UpdatedAt': '2023-05-23T16:09:40.158+08:00',
    //         'parentId': '35',
    //         'path': 'report',
    //         'name': 'report',
    //         'hidden': false,
    //         'component': 'view/report/report.vue',
    //         'sort': 20,
    //         'meta': {
    //           'activeName': '',
    //           'keepAlive': false,
    //           'defaultMenu': false,
    //           'title': '订单报表',
    //           'icon': 'arrow-down-bold',
    //           'closeTab': false
    //         },
    //         'authoritys': null,
    //         'menuBtn': null,
    //         'menuId': '42',
    //         'children': null,
    //         'parameters': [],
    //         'btns': null
    //       }
    //     ],
    //     'parameters': [],
    //     'btns': null
    //   }
    // ]
    // if (type === undefined) {
    //   type = window.localStorage.getItem('type')
    // }
    // if (type === 'user') {
    //   asyncRouter = userRouter
    // }
    asyncRouter && asyncRouter.push({
      path: 'reload',
      name: 'Reload',
      hidden: true,
      meta: {
        title: '',
        closeTab: true,
      },
      component: 'view/error/reload.vue'
    })
    formatRouter(asyncRouter, routeMap)
    baseRouter[0].children = asyncRouter
    if (notLayoutRouterArr.length !== 0) {
      baseRouter.push(...notLayoutRouterArr)
    }
    asyncRouterHandle(baseRouter)
    KeepAliveFilter(asyncRouter)
    asyncRouters.value = baseRouter
    return true
  }

  return {
    asyncRouters,
    keepAliveRouters,
    asyncRouterFlag,
    SetAsyncRouter,
    routeMap
  }
})

