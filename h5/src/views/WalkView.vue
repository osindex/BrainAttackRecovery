<script setup lang="ts">
import { computed, onBeforeUnmount, ref } from 'vue'
import { showSuccessToast } from 'vant'
import TrainingInstruction from '@/components/TrainingInstruction.vue'
import walkImage from '@/assets/training/walk.svg'
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
    <TrainingInstruction
      title="慢走训练说明"
      :image-src="walkImage"
      image-alt="患者扶稳慢走训练示意图"
      :steps="['穿好防滑鞋，站稳后再开始。', '可以扶栏杆、扶墙或让家属在旁边陪同。', '用舒服的速度慢慢走，不追求速度。', '结束后点击保存，记录本次步行时长。']"
      :safety-tips="['头晕、胸闷、腿软时立即停止。', '地面要干燥，避免拖鞋和湿滑地面。', '首次训练建议有人陪同。']"
    />
    <section class="card">
      <p class="subtitle">按开始后慢走，结束时记录本次步行时长。</p>
      <div class="big-number">{{ displayTime }}</div>
      <van-button v-if="startedAt === null" type="primary" block @click="start">开始</van-button>
      <van-button v-else type="success" block @click="finish">结束并保存</van-button>
    </section>
  </main>
</template>
