<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="活跃颜色:" prop="activeColor">
          <el-input v-model="formData.activeColor" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="基础颜色:" prop="baseColor">
          <el-input v-model="formData.baseColor" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户是否被冻结 1正常 2冻结:" prop="enable">
          <el-input v-model.number="formData.enable" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户头像:" prop="headerImg">
          <el-input v-model="formData.headerImg" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户昵称:" prop="nickName">
          <el-input v-model="formData.nickName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户登录密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户侧边主题:" prop="sideMode">
          <el-input v-model="formData.sideMode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户登录名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户UUID:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Users'
}
</script>

<script setup>
import {
  createUsers,
  updateUsers,
  findUsers
} from '@/api/users'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            activeColor: '',
            baseColor: '',
            email: '',
            enable: 0,
            headerImg: '',
            nickName: '',
            password: '',
            phone: '',
            sideMode: '',
            username: '',
            uuid: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUsers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reuserInfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createUsers(formData.value)
               break
             case 'update':
               res = await updateUsers(formData.value)
               break
             default:
               res = await createUsers(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
