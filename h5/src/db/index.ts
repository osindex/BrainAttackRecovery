import Dexie, { type Table } from 'dexie'
import type { LocalTrainingRecord } from '@/types/training'

export class RehabDatabase extends Dexie {
  records!: Table<LocalTrainingRecord, number>

  constructor() {
    super('brain-rehab-h5')
    this.version(1).stores({
      records: '++id, clientRecordId, trainingType, occurredOn, syncStatus, createdAt',
    })
  }
}

export const db = new RehabDatabase()
