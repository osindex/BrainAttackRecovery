<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { showSuccessToast, showToast } from 'vant'
import { fetchCards } from '@/api/cards'
import { useTrainingStore } from '@/stores/training'
import { speak } from '@/utils/speech'
import type { CardItem } from '@/types/training'

const trainingStore = useTrainingStore()
const cards = ref<CardItem[]>([])
const index = ref(0)
const startedAt = ref(Date.now())

const currentCard = computed(() => cards.value[index.value])

onMounted(async () => {
  try {
    cards.value = await fetchCards()
  } catch {
    cards.value = [{ id: 0, categoryId: 0, categoryName: '占位', title: '苹果', label: '苹果', imageUrl: 'https://dummyimage.com/512x512/f5f5f5/333333&text=%E8%8B%B9%E6%9E%9C', difficulty: 1 }]
  }
  startedAt.value = Date.now()
})

async function answer(isCorrect: boolean): Promise<void> {
  const card = currentCard.value
  if (!card) {
    showToast('暂无卡片')
    return
  }
  const reactionMs = Date.now() - startedAt.value
  await trainingStore.addRecord({ trainingType: 'card_game', cardId: card.id, isCorrect: isCorrect ? 1 : 0, reactionMs, payload: JSON.stringify({ title: card.title, label: card.label }), remark: '图片卡辨物' })
  speak(isCorrect ? '回答正确' : '已记录')
  showSuccessToast('已保存')
  index.value = (index.value + 1) % Math.max(cards.value.length, 1)
  startedAt.value = Date.now()
}
</script>

<template>
  <main class="page">
    <van-nav-bar title="图片卡辨物" left-text="返回" left-arrow @click-left="$router.back()" />
    <section class="card" v-if="currentCard">
      <p class="subtitle">看图片，说出它是什么，然后记录结果。</p>
      <img :src="currentCard.imageUrl" :alt="currentCard.title" style="width: 100%; border-radius: 16px; margin: 16px 0" />
      <h2 class="title">{{ currentCard.title }}</h2>
      <van-button type="success" block @click="answer(true)">答对了</van-button>
      <van-button style="margin-top: 12px" type="danger" block @click="answer(false)">没答对</van-button>
    </section>
  </main>
</template>
