export function todayString(): string {
  const now = new Date()
  const year = now.getFullYear()
  const month = `${now.getMonth() + 1}`.padStart(2, '0')
  const day = `${now.getDate()}`.padStart(2, '0')
  return `${year}-${month}-${day}`
}

export function nowIso(): string {
  return new Date().toISOString()
}

export function makeClientRecordId(prefix: string): string {
  const random = crypto.randomUUID()
  return `${prefix}-${Date.now()}-${random}`
}
