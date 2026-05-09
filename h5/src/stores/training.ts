import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { db } from '@/db'
import { uploadTrainingRecord } from '@/api/records'
import { makeClientRecordId, nowIso, todayString } from '@/utils/date'
import type { LocalTrainingRecord, TrainingType } from '@/types/training'

interface AddRecordInput {
  trainingType: TrainingType
  durationSeconds?: number
  setsCount?: number
  repsCount?: number
  gazeCount?: number
  cardId?: number
  isCorrect?: number
  reactionMs?: number
  payload?: string
  remark?: string
}

export const useTrainingStore = defineStore('training', () => {
  const records = ref<LocalTrainingRecord[]>([])
  const syncing = ref(false)

  const pendingCount = computed(() => records.value.filter((item) => item.syncStatus !== 'synced').length)

  async function loadRecords(): Promise<void> {
    records.value = await db.records.orderBy('createdAt').reverse().toArray()
  }

  async function addRecord(input: AddRecordInput): Promise<LocalTrainingRecord> {
    const now = nowIso()
    const record: LocalTrainingRecord = {
      clientRecordId: makeClientRecordId(input.trainingType),
      trainingType: input.trainingType,
      occurredOn: todayString(),
      durationSeconds: input.durationSeconds ?? 0,
      setsCount: input.setsCount ?? 0,
      repsCount: input.repsCount ?? 0,
      gazeCount: input.gazeCount ?? 0,
      cardId: input.cardId ?? 0,
      isCorrect: input.isCorrect ?? 0,
      reactionMs: input.reactionMs ?? 0,
      payload: input.payload ?? '{}',
      remark: input.remark ?? '',
      syncStatus: 'pending',
      createdAt: now,
      updatedAt: now,
    }
    record.id = await db.records.add(record)
    await loadRecords()
    void syncPending()
    return record
  }

  async function syncPending(): Promise<void> {
    if (syncing.value) return
    syncing.value = true
    try {
      const pending = await db.records.where('syncStatus').anyOf('pending', 'failed').toArray()
      for (const record of pending) {
        if (record.id === undefined) continue
        await db.records.update(record.id, { syncStatus: 'syncing', updatedAt: nowIso() })
        try {
          const remoteId = await uploadTrainingRecord(record)
          await db.records.update(record.id, { syncStatus: 'synced', remoteId, updatedAt: nowIso() })
        } catch {
          await db.records.update(record.id, { syncStatus: 'failed', updatedAt: nowIso() })
        }
      }
    } finally {
      syncing.value = false
      await loadRecords()
    }
  }

  return { records, syncing, pendingCount, loadRecords, addRecord, syncPending }
})
