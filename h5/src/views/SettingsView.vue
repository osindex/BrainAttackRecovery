<script setup lang="ts">
import { ref } from 'vue'
import { showSuccessToast } from 'vant'
import { savePairingCredentials } from '@/api/client'

const username = ref('')
const password = ref('')

function save(): void {
  savePairingCredentials(username.value.trim(), password.value)
  showSuccessToast('已保存本机配对信息')
}
</script>

<template>
  <main class="page">
    <van-nav-bar title="本机配对" left-text="返回" left-arrow @click-left="$router.back()" />
    <section class="card">
      <h1 class="title">本机配对</h1>
      <p class="subtitle">请输入后台为患者设备创建的低权限账号。不要使用管理员账号。</p>
      <van-field v-model="username" label="账号" placeholder="例如 patient-device-1" />
      <van-field v-model="password" label="密码" type="password" placeholder="请输入配对密码" />
      <van-button type="primary" block style="margin-top: 16px" :disabled="!username || !password" @click="save">保存配对</van-button>
    </section>
  </main>
</template>
