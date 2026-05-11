import { http } from './client'
import type { CardItem } from '@/types/training'

interface CardListEnvelope {
  code: number
  data?: {
    list?: CardItem[]
  }
}

export async function fetchCards(): Promise<CardItem[]> {
  const response = await http.get<CardListEnvelope>('/rehab/card/public', { params: { limit: 200 } })
  return response.data.data?.list ?? []
}
