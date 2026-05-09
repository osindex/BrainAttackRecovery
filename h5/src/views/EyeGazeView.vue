<script setup lang="ts">
import { ref } from 'vue'
import { showSuccessToast } from 'vant'
import TrainingInstruction from '@/components/TrainingInstruction.vue'
import eyeGazeImage from '@/assets/training/eye-gaze.svg'
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
    <TrainingInstruction
      title="眼睛凝视训练说明"
      :image-src="eyeGazeImage"
      image-alt="眼睛从左到右凝视训练示意图"
      :steps="['坐稳，头保持不动。', '先看左侧目标点。', '只移动眼睛，看向右侧目标点。', '从左看到右算一次，然后点击计数。']"
      :safety-tips="['眼睛酸胀、头晕时立即休息。', '不要快速甩头，训练时头部尽量不动。', '每组时间不宜太长，可以少量多次。']"
    />
    <section class="card">
      <p class="subtitle">眼睛从左看向右，完成一次后点击按钮。</p>
      <div class="big-number">{{ count }}</div>
      <van-button type="primary" block @click="count += 1">左到右 1 次</van-button>
      <van-button style="margin-top: 12px" block @click="count = Math.max(0, count - 1)">撤销 1 次</van-button>
      <van-button style="margin-top: 12px" type="success" block @click="save">保存记录</van-button>
    </section>
  </main>
</template>
