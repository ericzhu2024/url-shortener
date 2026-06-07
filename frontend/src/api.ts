const BASE = 'http://3.134.106.201:8080'

export async function createLink(original: string) {
  const res = await fetch(`${BASE}/api/links`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ original }),
  })
  const data = await res.json()
  if (!res.ok) throw new Error(data.error)
  return data
}

export async function getStats(code: string) {
  const res = await fetch(`${BASE}/api/links/${code}/stats`)
  const data = await res.json()
  if (!res.ok) throw new Error(data.error)
  return data
}