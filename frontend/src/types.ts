export interface Link {
  code: string
  short_url: string
  original: string
}

export interface Stats {
  id: number
  code: string
  original: string
  clicks: number
  created_at: string
  expires_at: string | null
}