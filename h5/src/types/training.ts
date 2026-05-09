export type TrainingType = 'walk' | 'fist_raise' | 'eye_gaze' | 'card_game'

export type SyncStatus = 'pending' | 'syncing' | 'synced' | 'failed'

export interface LocalTrainingRecord {
  id?: number
  clientRecordId: string
  trainingType: TrainingType
  occurredOn: string
  durationSeconds: number
  setsCount: number
  repsCount: number
  gazeCount: number
  cardId: number
  isCorrect: number
  reactionMs: number
  payload: string
  remark: string
  syncStatus: SyncStatus
  remoteId?: number
  createdAt: string
  updatedAt: string
}

export interface CardItem {
  id: number
  categoryId: number
  categoryName: string
  title: string
  label: string
  imageUrl: string
  difficulty: number
}
