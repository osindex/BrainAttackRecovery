<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useTrainingStore } from '@/stores/training'

const trainingStore = useTrainingStore()

const todayRecords = computed(() => trainingStore.records.filter((item) => item.occurredOn === new Date().toISOString().slice(0, 10)))
const walkMinutes = computed(() => Math.round(todayRecords.value.filter((item) => item.trainingType === 'walk').reduce((sum, item) => sum + item.durationSeconds, 0) / 60))
const fistReps = computed(() => todayRecords.value.filter((item) => item.trainingType === 'fist_raise').reduce((sum, item) => sum + item.repsCount, 0))
const gazeCount = computed(() => todayRecords.value.filter((item) => item.trainingType === 'eye_gaze').reduce((sum, item) => sum + item.gazeCount, 0))
const cardCorrect = computed(() => todayRecords.value.filter((item) => item.trainingType === 'card_game' && item.isCorrect === 1).length)

onMounted(async () => { await trainingStore.loadRecords() })
</script>

<template>
  <main class="page safe-bottom">
    <van-nav-bar title="历史图表" left-text="返回" left-arrow @click-left="$router.back()" />
    <section class="grid" style="margin-top: 16px">
      <div class="card"><p class="subtitle">今日慢走</p><div class="big-number">{{ walkMinutes }}</div><p class="subtitle">分钟</p></div>
      <div class="card"><p class="subtitle">握拳平举</p><div class="big-number">{{ fistReps }}</div><p class="subtitle">次</p></div>
      <div class="card"><p class="subtitle">眼睛凝视</p><div class="big-number">{{ gazeCount }}</div><p class="subtitle">次</p></div>
      <div class="card"><p class="subtitle">图卡答对</p><div class="big-number">{{ cardCorrect }}</div><p class="subtitle">张</p></div>
    </section>
    <section class="card">
      <p class="subtitle">本地记录共 {{ trainingStore.records.length }} 条，待同步 {{ trainingStore.pendingCount }} 条。</p>
      <van-button type="primary" block @click="trainingStore.syncPending">立即同步</van-button>
    </section>
  </main>
</template>
