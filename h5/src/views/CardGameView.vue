<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { showFailToast, showSuccessToast, showToast } from 'vant'
import { fetchCards } from '@/api/cards'
import { useTrainingStore } from '@/stores/training'
import { speak } from '@/utils/speech'
import type { CardItem } from '@/types/training'

type GameMode = 'name_to_image' | 'image_to_name'

interface GameOption {
  key: string
  label: string
  card: CardItem
}

const fallbackCards: CardItem[] = [
  { id: 0, categoryId: 0, categoryName: '实物图', title: '苹果', label: '苹果', imageUrl: 'https://commons.wikimedia.org/wiki/Special:FilePath/Red_Apple.jpg', difficulty: 1 },
  { id: 0, categoryId: 0, categoryName: '实物图', title: '杯子', label: '杯子', imageUrl: 'https://commons.wikimedia.org/wiki/Special:FilePath/White_cup_and_saucer.jpg', difficulty: 1 },
  { id: 0, categoryId: 0, categoryName: '实物图', title: '小狗', label: '小狗', imageUrl: 'https://commons.wikimedia.org/wiki/Special:FilePath/Golden_Retriever_Carlos_(10581910556).jpg', difficulty: 1 },
  { id: 0, categoryId: 0, categoryName: '实物图', title: '汽车', label: '汽车', imageUrl: 'https://commons.wikimedia.org/wiki/Special:FilePath/2019_Toyota_Corolla_Icon_Tech_VVT-i_Hybrid_1.8.jpg', difficulty: 1 },
]

const trainingStore = useTrainingStore()
const cards = ref<CardItem[]>([])
const currentCard = ref<CardItem | null>(null)
const options = ref<GameOption[]>([])
const mode = ref<GameMode>('image_to_name')
const startedAt = ref(Date.now())

const modeText = computed(() => (mode.value === 'image_to_name' ? '看图片，选名称' : '看名称，选图片'))

function shuffle<T>(items: T[]): T[] {
  return [...items].sort(() => Math.random() - 0.5)
}

function makeOptions(target: CardItem): GameOption[] {
  const distractors = shuffle(cards.value.filter((item) => item.label !== target.label)).slice(0, 3)
  return shuffle([target, ...distractors]).map((card, index) => ({ key: `${card.label}-${index}`, label: card.label || card.title, card }))
}

function nextQuestion(nextMode?: GameMode): void {
  if (nextMode) mode.value = nextMode
  const pool = cards.value.length > 0 ? cards.value : fallbackCards
  const target = shuffle(pool)[0]
  currentCard.value = target
  options.value = makeOptions(target)
  startedAt.value = Date.now()
  speak(modeText.value)
}

onMounted(async () => {
  try {
    const remoteCards = await fetchCards()
    cards.value = remoteCards.length >= 2 ? remoteCards : fallbackCards
  } catch {
    cards.value = fallbackCards
  }
  nextQuestion('image_to_name')
})

async function choose(option: GameOption): Promise<void> {
  const target = currentCard.value
  if (!target) {
    showToast('暂无卡片')
    return
  }
  const isCorrect = option.card.label === target.label
  const reactionMs = Date.now() - startedAt.value
  await trainingStore.addRecord({
    trainingType: 'card_game',
    cardId: target.id,
    isCorrect: isCorrect ? 1 : 0,
    reactionMs,
    payload: JSON.stringify({ mode: mode.value, target: target.label, selected: option.label }),
    remark: modeText.value,
  })
  if (isCorrect) {
    speak('回答正确')
    showSuccessToast('答对了')
  } else {
    speak(`不对，正确答案是${target.label}`)
    showFailToast(`正确答案：${target.label}`)
  }
  window.setTimeout(() => nextQuestion(), 800)
}
</script>

<template>
  <main class="page">
    <van-nav-bar title="图片卡辨物" left-text="返回" left-arrow @click-left="$router.back()" />

    <section class="card">
      <p class="subtitle">选择玩法</p>
      <van-button
        :type="mode === 'image_to_name' ? 'primary' : 'default'"
        block
        @click="nextQuestion('image_to_name')"
      >
        看图片，选名称
      </van-button>
      <van-button
        style="margin-top: 12px"
        :type="mode === 'name_to_image' ? 'primary' : 'default'"
        block
        @click="nextQuestion('name_to_image')"
      >
        看名称，选图片
      </van-button>
    </section>

    <section class="card" v-if="currentCard">
      <p class="subtitle">{{ modeText }}</p>

      <template v-if="mode === 'image_to_name'">
        <img :src="currentCard.imageUrl" :alt="currentCard.title" class="question-image" />
        <div class="option-grid">
          <van-button v-for="option in options" :key="option.key" block type="primary" plain @click="choose(option)">
            {{ option.label }}
          </van-button>
        </div>
      </template>

      <template v-else>
        <div class="target-name">{{ currentCard.label }}</div>
        <div class="image-grid">
          <button v-for="option in options" :key="option.key" class="image-option" @click="choose(option)">
            <img :src="option.card.imageUrl" :alt="option.label" />
          </button>
        </div>
      </template>
    </section>
  </main>
</template>

<style scoped>
.question-image {
  width: 100%;
  border-radius: 16px;
  margin: 16px 0;
  background: #f5f7fb;
}

.option-grid {
  display: grid;
  gap: 12px;
}

.target-name {
  margin: 16px 0;
  padding: 24px;
  border-radius: 18px;
  background: #eef6ff;
  text-align: center;
  color: #0958d9;
  font-size: 42px;
  font-weight: 900;
}

.image-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 12px;
}

.image-option {
  padding: 8px;
  border: 3px solid #d6e4ff;
  border-radius: 18px;
  background: #fff;
}

.image-option img {
  width: 100%;
  display: block;
  border-radius: 12px;
}
</style>
