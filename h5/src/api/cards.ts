import { ensureGuestLogin, http } from './client'
import type { CardItem } from '@/types/training'

interface CardListEnvelope {
  code: number
  data?: {
    list?: CardItem[]
  }
}

export async function fetchCards(): Promise<CardItem[]> {
  await ensureGuestLogin()
  const response = await http.get<CardListEnvelope>('/rehab/card', { params: { pageNum: 1, pageSize: 50, status: 1 } })
  return response.data.data?.list ?? []
}
