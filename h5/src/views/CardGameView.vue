<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { showFailToast, showSuccessToast, showToast } from 'vant'
import { fetchCards } from '@/api/cards'
import { useTrainingStore } from '@/stores/training'
import { speak } from '@/utils/speech'
import type { CardItem } from '@/types/training'

type GameMode = 'image_to_name' | 'name_to_image' | 'mixed'

interface GameOption {
  key: string
  label: string
  card: CardItem
}

interface Question {
  mode: 'image_to_name' | 'name_to_image'
  target: CardItem
  options: GameOption[]
  optionCount: number
}

interface RoundSummary {
  level: number
  total: number
  correct: number
  averageReactionMs: number
  optionCount: number
}

const QUESTIONS_PER_ROUND = 6
const MIN_OPTIONS = 3
const MAX_OPTIONS = 6

function proxyCard(label: string, file: string, category: string, difficulty: number, _sort: number): CardItem {
  return {
    id: 0,
    categoryId: 0,
    categoryName: category,
    title: label,
    label,
    imageUrl: `/api/v1/rehab/card/image/wiki?file=${encodeURIComponent(file)}`,
    difficulty,
  }
}

const fallbackCards: CardItem[] = [
  // food
  proxyCard('苹果', 'Red_Apple.jpg', '食物', 1, 10),
  proxyCard('香蕉', 'Bananavarieties.jpg', '食物', 1, 11),
  proxyCard('橙子', 'Orange_fruit.jpg', '食物', 1, 12),
  proxyCard('西瓜', 'Watermelon_cross_BNC.jpg', '食物', 1, 13),
  proxyCard('草莓', 'Strawberry_Single1.jpg', '食物', 2, 14),
  proxyCard('梨', 'Yellow_pear_on_a_black_background.jpg', '食物', 2, 15),
  // daily
  proxyCard('杯子', 'White_cup_and_saucer.jpg', '日用品', 1, 20),
  proxyCard('勺子', 'Spoon_silver.jpg', '日用品', 2, 21),
  proxyCard('筷子', 'Chopsticks-candc.jpg', '日用品', 2, 22),
  proxyCard('水壶', 'WWII_Allied_Canteen.jpg', '日用品', 3, 23),
  // animals
  proxyCard('小狗', 'Golden_Retriever_Carlos_(10581910556).jpg', '动物', 1, 30),
  proxyCard('小猫', 'Felis_catus-cat_on_snow.jpg', '动物', 1, 31),
  proxyCard('兔子', 'Oryctolagus_cuniculus_Rcdo.jpg', '动物', 2, 32),
  proxyCard('马', 'Hauspferd.JPG', '动物', 2, 33),
  proxyCard('鸡', 'Gallus_gallus_domesticus_Brown_Leghorn.jpg', '动物', 2, 34),
  proxyCard('牛', 'Cow_female_black_white.jpg', '动物', 2, 35),
  proxyCard('羊', 'Sheep_in_field.JPG', '动物', 2, 36),
  proxyCard('金鱼', 'Goldfish3.jpg', '动物', 3, 37),
  proxyCard('蝴蝶', 'Monarch_In_May.jpg', '动物', 3, 38),
  proxyCard('蚱蜢', 'Grasshopper_1.jpg', '动物', 3, 39),
  proxyCard('熊猫', 'Panda_Cub_from_Wolong,_Sichuan,_China.JPG', '动物', 2, 40),
  // vehicles
  proxyCard('汽车', '2019_Toyota_Corolla_Icon_Tech_VVT-i_Hybrid_1.8.jpg', '交通工具', 1, 50),
  proxyCard('自行车', 'Bicycle_Abbey_Sprotbrough.jpg', '交通工具', 2, 51),
  proxyCard('公交车', 'Bus_in_Hong_Kong.jpg', '交通工具', 2, 52),
  proxyCard('飞机', 'Airbus_A380_blue_sky.jpg', '交通工具', 2, 53),
  proxyCard('救护车', 'Ambulance_Berlin.jpg', '交通工具', 3, 54),
  // clothing
  proxyCard('鞋子', 'Shoes_-_Nike_Air_Jordan_1_Retro_Banned_2016_-_sneakers.jpg', '衣物', 2, 60),
  proxyCard('裤子', 'Blue_Denim_Jeans.jpg', '衣物', 2, 61),
  proxyCard('围巾', 'Red_scarf.jpg', '衣物', 2, 62),
  proxyCard('帽子', 'Winter_hat.jpg', '衣物', 2, 63),
  proxyCard('手套', 'Glove.jpg', '衣物', 3, 64),
  // body parts
  proxyCard('手', 'Hand_(sculpture).jpg', '身体部位', 1, 70),
  proxyCard('眼睛', 'Human_eye.jpg', '身体部位', 1, 71),
  proxyCard('脚', 'Pair_of_feet.jpg', '身体部位', 2, 72),
  proxyCard('脸', 'Face_(Unsplash).jpg', '身体部位', 2, 73),
]

const trainingStore = useTrainingStore()
const cards = ref<CardItem[]>([])
const mode = ref<GameMode>('mixed')
const currentQuestion = ref<Question | null>(null)
const startedAt = ref(Date.now())
const questionIndex = ref(0)
const correctCount = ref(0)
const reactionTotalMs = ref(0)
const level = ref(1)
const cleared = ref(0)
const roundSummary = ref<RoundSummary | null>(null)

const modeText = computed(() => {
  if (!currentQuestion.value) return ''
  return currentQuestion.value.mode === 'image_to_name' ? '看图片，选名称' : '看名称，选图片'
})

const progressText = computed(() => `第 ${level.value} 关 · 第 ${Math.min(questionIndex.value + 1, QUESTIONS_PER_ROUND)} / ${QUESTIONS_PER_ROUND} 题`)

const optionCountForLevel = computed(() => {
  const target = MIN_OPTIONS + Math.floor((level.value - 1) / 2)
  return Math.min(MAX_OPTIONS, Math.max(MIN_OPTIONS, target))
})

const allowedDifficulty = computed(() => {
  if (level.value <= 2) return [1]
  if (level.value <= 4) return [1, 2]
  return [1, 2, 3]
})

function shuffle<T>(items: T[]): T[] {
  return [...items].sort(() => Math.random() - 0.5)
}

function pickQuestionMode(): 'image_to_name' | 'name_to_image' {
  if (mode.value === 'mixed') return Math.random() < 0.5 ? 'image_to_name' : 'name_to_image'
  return mode.value
}

function pickPool(): CardItem[] {
  const source = cards.value.length >= MIN_OPTIONS + 1 ? cards.value : fallbackCards
  const allowed = allowedDifficulty.value
  const filtered = source.filter((card) => allowed.includes(card.difficulty || 1))
  return filtered.length >= MIN_OPTIONS + 1 ? filtered : source
}

function buildOptions(target: CardItem, pool: CardItem[], optionCount: number): GameOption[] {
  const distractorPoolByLabel = new Map<string, CardItem>()
  for (const card of pool) {
    if (card.label !== target.label && !distractorPoolByLabel.has(card.label)) {
      distractorPoolByLabel.set(card.label, card)
    }
  }
  const distractors = shuffle(Array.from(distractorPoolByLabel.values())).slice(0, optionCount - 1)
  return shuffle([target, ...distractors]).map((card, index) => ({
    key: `${card.label}-${index}`,
    label: card.label || card.title,
    card,
  }))
}

function pickQuestion(): void {
  const pool = pickPool()
  const target = shuffle(pool)[0]
  const optionCount = Math.min(optionCountForLevel.value, pool.length)
  currentQuestion.value = {
    mode: pickQuestionMode(),
    target,
    options: buildOptions(target, pool, optionCount),
    optionCount,
  }
  startedAt.value = Date.now()
  speak(currentQuestion.value.mode === 'image_to_name' ? '看图片，选名称' : '看名称，选图片')
}

function startRound(nextMode?: GameMode, resetLevel = true): void {
  if (nextMode) mode.value = nextMode
  if (resetLevel) {
    level.value = 1
    cleared.value = 0
  }
  questionIndex.value = 0
  correctCount.value = 0
  reactionTotalMs.value = 0
  roundSummary.value = null
  pickQuestion()
}

function goToNextLevel(): void {
  level.value += 1
  startRound(undefined, false)
}

function retryCurrentLevel(): void {
  startRound(undefined, false)
}

onMounted(async () => {
  try {
    const remoteCards = await fetchCards()
    cards.value = remoteCards.length >= MIN_OPTIONS + 1 ? remoteCards : fallbackCards
  } catch {
    cards.value = fallbackCards
  }
  startRound()
})

async function choose(option: GameOption): Promise<void> {
  const question = currentQuestion.value
  if (!question) {
    showToast('暂无卡片')
    return
  }
  const isCorrect = option.card.label === question.target.label
  const reactionMs = Date.now() - startedAt.value
  reactionTotalMs.value += reactionMs
  if (isCorrect) correctCount.value += 1

  await trainingStore.addRecord({
    trainingType: 'card_game',
    cardId: question.target.id,
    isCorrect: isCorrect ? 1 : 0,
    reactionMs,
    payload: JSON.stringify({
      mode: question.mode,
      target: question.target.label,
      selected: option.label,
      questionIndex: questionIndex.value + 1,
      level: level.value,
      optionCount: question.optionCount,
      allowedDifficulty: allowedDifficulty.value,
    }),
    remark: question.mode === 'image_to_name' ? '看图片，选名称' : '看名称，选图片',
  })

  if (isCorrect) {
    speak('回答正确')
    showSuccessToast('答对了')
  } else {
    speak(`不对，正确答案是${question.target.label}`)
    showFailToast(`正确答案：${question.target.label}`)
  }

  window.setTimeout(() => {
    if (questionIndex.value + 1 >= QUESTIONS_PER_ROUND) {
      finishRound()
    } else {
      questionIndex.value += 1
      pickQuestion()
    }
  }, 900)
}

function finishRound(): void {
  const total = QUESTIONS_PER_ROUND
  const summary: RoundSummary = {
    level: level.value,
    total,
    correct: correctCount.value,
    averageReactionMs: total > 0 ? Math.round(reactionTotalMs.value / total) : 0,
    optionCount: optionCountForLevel.value,
  }
  if (summary.correct === total) cleared.value += 1
  roundSummary.value = summary
  currentQuestion.value = null
  if (summary.correct === total) speak(`第${summary.level}关全对，进入下一关`)
  else if (summary.correct === 0) speak('本关结束，再来一次')
  else speak(`本关答对${summary.correct}题，共${total}题`)
}
</script>

<template>
  <main class="page safe-bottom">
    <van-nav-bar title="图片卡辨物" left-text="返回" left-arrow @click-left="$router.back()" />

    <section class="card">
      <p class="subtitle">玩法（每关 {{ QUESTIONS_PER_ROUND }} 题，关卡递增）</p>
      <van-button :type="mode === 'image_to_name' ? 'primary' : 'default'" block @click="startRound('image_to_name')">看图片，选名称</van-button>
      <van-button style="margin-top: 12px" :type="mode === 'name_to_image' ? 'primary' : 'default'" block @click="startRound('name_to_image')">看名称，选图片</van-button>
      <van-button style="margin-top: 12px" :type="mode === 'mixed' ? 'primary' : 'default'" block @click="startRound('mixed')">混合关卡</van-button>
      <p class="subtitle" style="margin-top: 12px">
        累计通关：{{ cleared }} 关 · 当前选项数：{{ optionCountForLevel }} · 难度池：{{ allowedDifficulty.join('、') }}
      </p>
    </section>

    <section class="card" v-if="currentQuestion">
      <p class="subtitle">{{ progressText }} · {{ modeText }}</p>

      <template v-if="currentQuestion.mode === 'image_to_name'">
        <img :src="currentQuestion.target.imageUrl" :alt="currentQuestion.target.title" class="question-image" />
        <div class="option-grid">
          <van-button v-for="option in currentQuestion.options" :key="option.key" block type="primary" plain @click="choose(option)">
            {{ option.label }}
          </van-button>
        </div>
      </template>

      <template v-else>
        <div class="target-name">{{ currentQuestion.target.label }}</div>
        <div class="image-grid">
          <button v-for="option in currentQuestion.options" :key="option.key" class="image-option" @click="choose(option)">
            <img :src="option.card.imageUrl" :alt="option.label" />
          </button>
        </div>
      </template>
    </section>

    <section class="card" v-else-if="roundSummary">
      <h2 class="title">第 {{ roundSummary.level }} 关结束</h2>
      <p class="subtitle">答对 {{ roundSummary.correct }} / {{ roundSummary.total }} 题（每题 {{ roundSummary.optionCount }} 选 1）</p>
      <p class="subtitle">平均反应：{{ roundSummary.averageReactionMs }} 毫秒</p>
      <van-button v-if="roundSummary.correct === roundSummary.total" type="primary" block style="margin-top: 12px" @click="goToNextLevel">进入第 {{ level + 1 }} 关</van-button>
      <template v-else>
        <van-button type="primary" block style="margin-top: 12px" @click="retryCurrentLevel">本关重来（第 {{ level }} 关）</van-button>
        <van-button block style="margin-top: 12px" @click="goToNextLevel">跳到第 {{ level + 1 }} 关</van-button>
      </template>
      <van-button block plain style="margin-top: 12px" @click="startRound()">从第一关重新开始</van-button>
      <van-button block plain style="margin-top: 12px" @click="$router.back()">返回首页</van-button>
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
