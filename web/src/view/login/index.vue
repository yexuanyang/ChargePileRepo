<template>
  <div id="userLayout">
    <div class="login_panel">
      <div class="login_panel_form">
        <div class="login_panel_form_title">
          <img
            class="login_panel_form_title_logo"
            :src="$GIN_VUE_ADMIN.appLogo"
            alt
          >
          <p class="login_panel_form_title_p">{{ $GIN_VUE_ADMIN.appName }}</p>
        </div>
        <el-form
          ref="logInform"
          :model="logInformData"
          :rules="rules"
          :validate-on-rule-change="false"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username">
            <el-input
              v-model="logInformData.account"
              size="large"
              placeholder="请输入用户名"
              suffix-icon="user"
            />
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="logInformData.password"
              show-password
              size="large"
              type="password"
              placeholder="请输入密码"
            />
          </el-form-item>
          <el-form-item v-if="logInformData.openCaptcha" prop="captcha">
            <div class="vPicBox">
              <el-input
                v-model="logInformData.captcha"
                placeholder="请输入验证码"
                size="large"
                style="flex:1;padding-right: 20px;"
              />
              <div class="vPic">
                <img
                  v-if="picPath"
                  :src="picPath"
                  alt="请输入验证码"
                  @click="loginVerify()"
                >
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
                       style="width: 30%"
                       size="large"
                       @click="addUser">注册</el-button>
            <el-button
              type="primary"
              size="large"
              style="width: 30%; "
              @click="submitForm('user')"
            >用 户 登 录</el-button>
            <el-button
                type="primary"
                size="large"
                style="width: 30%; "
                @click="submitForm('admin')"
            >管 理 员 登 录</el-button>
          </el-form-item>
        </el-form>
        <el-dialog
            v-model="addUserDialog"
            custom-class="user-dialog"
            title="用户"
            :show-close="false"
            :close-on-press-escape="false"
            :close-on-click-modal="false"
        >
          <div style="height:60vh;overflow:auto;padding:0 12px;">
            <el-form ref="userForm" :rules="rules1" :model="userInfo" label-width="80px">
              <el-form-item label="用户名" prop="account">
                <el-input v-model="userInfo.account" />
              </el-form-item>
              <el-form-item label="密码" prop="password">
                <el-input v-model="userInfo.password" />
              </el-form-item>
              <el-form-item label="昵称" prop="name">
                <el-input v-model="userInfo.name" />
              </el-form-item>
            </el-form>
          </div>
          <template #footer>
            <div class="dialog-footer">
              <el-button @click="closeAddUserDialog">取 消</el-button>
              <el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
            </div>
          </template>
        </el-dialog>
      </div>
      <div class="login_panel_right" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
}
</script>

<script setup>
import {localRegister, register} from '@/api/user'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
const addUserDialog = ref(false)
const userInfo = ref({
  account: '',
  password: '',
  name: '',
})

const closeAddUserDialog = () => {
  userForm.value.resetFields()
  addUserDialog.value = false
}
const userForm = ref(null)
const enterAddUserDialog = async() => {
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      await localRegister(req)
      const res = await register(req)
      if (res.code === 0) {
        ElMessage({ type: 'success', message: '创建成功' })
        await getTableData()
        closeAddUserDialog()
      }
    }
  })
}

const rules1 = ref({
  account: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur'}
  ],
  password: [
    { required: true, message: '请输入用户密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' }
  ]
})

// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

// 登录相关操作
const logInform = ref(null)
const picPath = ref('')
const logInformData = reactive({
  account: 'admin',
  password: '111111',
})

const rules = reactive({
  account: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
})

const userStore = useUserStore()
const login = async(type) => {
  return await userStore.LoginIn(logInformData, type)
}
const submitForm = (type) => {
  logInform.value.validate(async(v) => {
    if (v) {
      const flag = await login(type)
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true,
      })
      return false
    }
  })
}

const addUser = () => {
  addUserDialog.value = true
}

</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";
</style>
