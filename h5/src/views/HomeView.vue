<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { showDialog } from 'vant'
import { useTrainingStore } from '@/stores/training'
import { hasPairingCredentials } from '@/api/client'
import { todayString } from '@/utils/date'

const trainingStore = useTrainingStore()

const todayRecords = computed(() => trainingStore.records.filter((item) => item.occurredOn === todayString()))
const completedTypes = computed(() => new Set(todayRecords.value.map((item) => item.trainingType)))

onMounted(async () => {
  await trainingStore.loadRecords()
  if (navigator.onLine && hasPairingCredentials()) void trainingStore.syncPending()
})

function showTodayAdvice(): void {
  const done = completedTypes.value
  let title = '今日训练建议'
  let message = '建议先从慢走训练开始，时间不用太长，5 到 10 分钟即可。训练前确认地面不滑，旁边最好有人陪同。'

  if (done.has('walk') && !done.has('fist_raise')) {
    message = '今天已经完成慢走训练。可以休息 5 分钟后，再做一组双手握拳平举，动作要慢，不要憋气。'
  } else if (done.has('fist_raise') && !done.has('eye_gaze')) {
    message = '今天已经完成握拳平举。接下来可以做眼睛凝视训练，头保持不动，眼睛从左看到右，感觉疲劳就停止。'
  } else if (done.has('eye_gaze') && !done.has('card_game')) {
    message = '今天已经完成眼睛凝视训练。可以做几张图片卡辨物练习，重点是慢慢说清楚，不追求速度。'
  } else if (done.size >= 3) {
    title = '今日完成不错'
    message = '今天已经完成多项训练。建议先休息，喝水，观察是否有头晕、乏力或不适。如果状态很好，再少量补充图片卡辨物练习。'
  }

  showDialog({ title, message, confirmButtonText: '知道了' })
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
      <van-button block style="margin-top: 12px" @click="showTodayAdvice">今日建议</van-button>
      <van-button block style="margin-top: 12px" to="/settings">本机配对设置</van-button>
    </section>
  </main>
</template>
