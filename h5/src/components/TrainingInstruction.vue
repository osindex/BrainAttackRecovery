<script setup lang="ts">
import { speak } from '@/utils/speech'

const props = defineProps<{
  title: string
  imageSrc: string
  imageAlt: string
  steps: string[]
  safetyTips: string[]
}>()

function playInstruction(): void {
  const content = [...props.steps, ...props.safetyTips].join('。')
  speak(`${props.title}。${content}`)
}
</script>

<template>
  <section class="card instruction-card">
    <img class="instruction-image" :src="imageSrc" :alt="imageAlt" />
    <h2 class="instruction-title">{{ title }}</h2>
    <div class="instruction-section">
      <strong>动作步骤</strong>
      <ol>
        <li v-for="step in steps" :key="step">{{ step }}</li>
      </ol>
    </div>
    <div class="instruction-section safety">
      <strong>安全提醒</strong>
      <ul>
        <li v-for="tip in safetyTips" :key="tip">{{ tip }}</li>
      </ul>
    </div>
    <van-button block plain type="primary" @click="playInstruction">播放语音说明</van-button>
  </section>
</template>

<style scoped>
.instruction-card {
  overflow: hidden;
}

.instruction-image {
  width: 100%;
  display: block;
  border-radius: 16px;
  margin-bottom: 14px;
  background: #f5f7fb;
}

.instruction-title {
  margin: 0 0 12px;
  font-size: 22px;
  font-weight: 800;
}

.instruction-section {
  margin-bottom: 12px;
  font-size: 18px;
  line-height: 1.7;
}

.instruction-section ol,
.instruction-section ul {
  margin: 8px 0 0;
  padding-left: 24px;
}

.safety {
  color: #ad4e00;
}
</style>
