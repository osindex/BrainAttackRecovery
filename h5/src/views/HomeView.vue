<script setup lang="ts">
import { onMounted } from 'vue'
import { showToast } from 'vant'
import { useTrainingStore } from '@/stores/training'
import { hasPairingCredentials } from '@/api/client'

const trainingStore = useTrainingStore()

onMounted(async () => {
  await trainingStore.loadRecords()
  if (navigator.onLine && hasPairingCredentials()) void trainingStore.syncPending()
})

function notifyComingSoon(): void {
  showToast('请先完成一项训练')
}
</script>

<template>
  <main class="page safe-bottom">
    <section class="card">
      <h1 class="title">卒中康复训练</h1>
      <p class="subtitle">大按钮、少步骤，先记录训练，再自动同步到本地 linapro 后台。</p>
    </section>

    <section class="grid">
      <van-button type="primary" block to="/walk">慢走训练</van-button>
      <van-button type="success" block to="/fist-raise">握拳平举</van-button>
      <van-button type="warning" block to="/eye-gaze">眼睛凝视</van-button>
      <van-button type="danger" block to="/card-game">图片辨物</van-button>
    </section>

    <section class="card" style="margin-top: 16px">
      <p class="subtitle">待同步记录：{{ trainingStore.pendingCount }} 条</p>
      <van-button block plain type="primary" to="/history">查看历史图表</van-button>
      <van-button block style="margin-top: 12px" @click="notifyComingSoon">今日建议</van-button>
      <van-button block style="margin-top: 12px" to="/settings">本机配对设置</van-button>
    </section>
  </main>
</template>
