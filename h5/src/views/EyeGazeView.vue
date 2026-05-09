<script setup lang="ts">
import { ref } from 'vue'
import { showSuccessToast } from 'vant'
import { useTrainingStore } from '@/stores/training'
import { speak } from '@/utils/speech'

const trainingStore = useTrainingStore()
const count = ref(0)

async function save(): Promise<void> {
  await trainingStore.addRecord({ trainingType: 'eye_gaze', gazeCount: count.value, payload: JSON.stringify({ count: count.value }), remark: '眼睛凝视训练' })
  speak('凝视训练记录已保存')
  showSuccessToast('已保存')
  count.value = 0
}
</script>

<template>
  <main class="page">
    <van-nav-bar title="眼睛凝视训练" left-text="返回" left-arrow @click-left="$router.back()" />
    <section class="card">
      <p class="subtitle">眼睛从左看向右，完成一次后点击按钮。</p>
      <div class="big-number">{{ count }}</div>
      <van-button type="primary" block @click="count += 1">左到右 1 次</van-button>
      <van-button style="margin-top: 12px" block @click="count = Math.max(0, count - 1)">撤销 1 次</van-button>
      <van-button style="margin-top: 12px" type="success" block @click="save">保存记录</van-button>
    </section>
  </main>
</template>
