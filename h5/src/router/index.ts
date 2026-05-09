import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import WalkView from '@/views/WalkView.vue'
import FistRaiseView from '@/views/FistRaiseView.vue'
import EyeGazeView from '@/views/EyeGazeView.vue'
import CardGameView from '@/views/CardGameView.vue'
import HistoryView from '@/views/HistoryView.vue'
import SettingsView from '@/views/SettingsView.vue'

const routes: RouteRecordRaw[] = [
  { path: '/', name: 'home', component: HomeView },
  { path: '/walk', name: 'walk', component: WalkView },
  { path: '/fist-raise', name: 'fistRaise', component: FistRaiseView },
  { path: '/eye-gaze', name: 'eyeGaze', component: EyeGazeView },
  { path: '/card-game', name: 'cardGame', component: CardGameView },
  { path: '/history', name: 'history', component: HistoryView },
  { path: '/settings', name: 'settings', component: SettingsView },
]

export const router = createRouter({ history: createWebHistory(), routes })
