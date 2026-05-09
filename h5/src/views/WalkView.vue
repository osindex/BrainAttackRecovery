<script setup lang="ts">
import { computed, onBeforeUnmount, ref } from 'vue'
import { showSuccessToast } from 'vant'
import { useTrainingStore } from '@/stores/training'
import { speak } from '@/utils/speech'

const trainingStore = useTrainingStore()
const startedAt = ref<number | null>(null)
const elapsedSeconds = ref(0)
let timer: number | undefined

const displayTime = computed(() => `${Math.floor(elapsedSeconds.value / 60)} 分 ${elapsedSeconds.value % 60} 秒`)

function start(): void {
  startedAt.value = Date.now()
  elapsedSeconds.value = 0
  timer = window.setInterval(() => {
    if (startedAt.value !== null) elapsedSeconds.value = Math.floor((Date.now() - startedAt.value) / 1000)
  }, 1000)
  speak('慢走训练开始')
}

async function finish(): Promise<void> {
  if (timer !== undefined) window.clearInterval(timer)
  timer = undefined
  await trainingStore.addRecord({ trainingType: 'walk', durationSeconds: Math.max(elapsedSeconds.value, 1), remark: '慢走训练' })
  speak('训练已保存')
  showSuccessToast('已保存')
  startedAt.value = null
}

onBeforeUnmount(() => { if (timer !== undefined) window.clearInterval(timer) })
</script>

<template>
  <main class="page">
    <van-nav-bar title="慢走训练" left-text="返回" left-arrow @click-left="$router.back()" />
    <section class="card">
      <p class="subtitle">按开始后慢走，结束时记录本次步行时长。</p>
      <div class="big-number">{{ displayTime }}</div>
      <van-button v-if="startedAt === null" type="primary" block @click="start">开始</van-button>
      <van-button v-else type="success" block @click="finish">结束并保存</van-button>
    </section>
  </main>
</template>
