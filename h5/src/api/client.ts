import axios from 'axios'

interface LoginResponseEnvelope {
  code: number
  message?: string
  data?: {
    accessToken?: string
  }
}

const tokenStorageKey = 'brain-rehab-token'
const pairingUserKey = 'brain-rehab-pairing-user'
const pairingPasswordKey = 'brain-rehab-pairing-password'

export const http = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
})

http.interceptors.request.use((config) => {
  const token = localStorage.getItem(tokenStorageKey)
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

export async function ensureGuestLogin(): Promise<void> {
  if (localStorage.getItem(tokenStorageKey)) return
  const username = localStorage.getItem(pairingUserKey)
  const password = localStorage.getItem(pairingPasswordKey)
  if (!username || !password) {
    throw new Error('请先在设置中完成本机配对')
  }
  const response = await http.post<LoginResponseEnvelope>('/auth/login', {
    username,
    password,
  })
  const token = response.data.data?.accessToken
  if (!token) throw new Error('登录失败：后端未返回 token')
  localStorage.setItem(tokenStorageKey, token)
}

export function savePairingCredentials(username: string, password: string): void {
  localStorage.setItem(pairingUserKey, username)
  localStorage.setItem(pairingPasswordKey, password)
  localStorage.removeItem(tokenStorageKey)
}

export function hasPairingCredentials(): boolean {
  return Boolean(localStorage.getItem(pairingUserKey) && localStorage.getItem(pairingPasswordKey))
}
