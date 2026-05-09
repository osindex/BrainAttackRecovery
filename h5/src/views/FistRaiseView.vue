<script setup lang="ts">
import { ref } from 'vue'
import { showSuccessToast } from 'vant'
import TrainingInstruction from '@/components/TrainingInstruction.vue'
import fistRaiseImage from '@/assets/training/fist-raise.svg'
import { useTrainingStore } from '@/stores/training'
import { speak } from '@/utils/speech'

const trainingStore = useTrainingStore()
const setsCount = ref(1)
const repsCount = ref(0)

async function save(): Promise<void> {
  await trainingStore.addRecord({ trainingType: 'fist_raise', setsCount: setsCount.value, repsCount: repsCount.value, payload: JSON.stringify({ sets: setsCount.value, reps: repsCount.value }), remark: '双手握拳平举' })
  speak('握拳平举记录已保存')
  showSuccessToast('已保存')
  repsCount.value = 0
}
</script>

<template>
  <main class="page">
    <van-nav-bar title="握拳平举" left-text="返回" left-arrow @click-left="$router.back()" />
    <TrainingInstruction
      title="双手握拳平举说明"
      :image-src="fistRaiseImage"
      image-alt="双手握拳平举训练示意图"
      :steps="['坐稳或站稳，双脚踩实地面。', '双手轻轻握拳，不要太用力。', '双臂慢慢向前抬到肩膀高度。', '慢慢放下，完成一次后点击计数。']"
      :safety-tips="['不要憋气，动作要慢。', '肩膀疼痛时降低高度或停止。', '身体不要后仰，必要时坐着完成。']"
    />
    <section class="card">
      <p class="subtitle">记录双手握拳平举的组数和次数。</p>
      <van-stepper v-model="setsCount" min="1" integer button-size="44" />
      <div class="big-number">{{ repsCount }}</div>
      <van-button type="primary" block @click="repsCount += 1">完成 1 次</van-button>
      <van-button style="margin-top: 12px" block @click="repsCount = Math.max(0, repsCount - 1)">撤销 1 次</van-button>
      <van-button style="margin-top: 12px" type="success" block @click="save">保存本组</van-button>
    </section>
  </main>
</template>
