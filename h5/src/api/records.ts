import { ensureGuestLogin, http } from './client'
import type { LocalTrainingRecord } from '@/types/training'

interface CreateRecordEnvelope {
  code: number
  data?: { id?: number }
}

export async function uploadTrainingRecord(record: LocalTrainingRecord): Promise<number | undefined> {
  await ensureGuestLogin()
  const response = await http.post<CreateRecordEnvelope>('/rehab/record', {
    clientRecordId: record.clientRecordId,
    trainingType: record.trainingType,
    occurredOn: record.occurredOn,
    durationSeconds: record.durationSeconds,
    setsCount: record.setsCount,
    repsCount: record.repsCount,
    gazeCount: record.gazeCount,
    cardId: record.cardId,
    isCorrect: record.isCorrect,
    reactionMs: record.reactionMs,
    payload: record.payload,
    remark: record.remark,
  })
  return response.data.data?.id
}
