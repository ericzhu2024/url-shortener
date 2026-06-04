import { useState } from 'react'
import { createLink, getStats } from './api'
import type { Link, Stats } from './types'

export default function App() {
  const [url, setUrl] = useState('')
  const [result, setResult] = useState<Link | null>(null)
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const [statsCode, setStatsCode] = useState('')
  const [stats, setStats] = useState<Stats | null>(null)
  const [statsError, setStatsError] = useState('')

  async function handleCreate() {
    setError('')
    setResult(null)
    setLoading(true)
    try {
      const data = await createLink(url)
      setResult(data)
      setUrl('')
    } catch (e: any) {
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  async function handleStats() {
    setStatsError('')
    setStats(null)
    try {
      const data = await getStats(statsCode)
      setStats(data)
    } catch (e: any) {
      setStatsError(e.message)
    }
  }

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col items-center py-20 px-4">
      <h1 className="text-4xl font-bold text-gray-800 mb-2">URL Shortener</h1>
      <p className="text-gray-400 mb-12">Turn long URLs into short ones</p>

      {/* Create short link */}
      <div className="w-full max-w-xl bg-white rounded-2xl shadow p-8 mb-8">
        <h2 className="text-lg font-semibold text-gray-700 mb-4">Create Short Link</h2>
        <div className="flex gap-2">
          <input
            className="flex-1 border border-gray-200 rounded-lg px-4 py-2 text-sm outline-none focus:border-blue-400"
            placeholder="https://example.com"
            value={url}
            onChange={e => setUrl(e.target.value)}
            onKeyDown={e => e.key === 'Enter' && handleCreate()}
          />
          <button
            className="bg-blue-500 hover:bg-blue-600 text-white px-5 py-2 rounded-lg text-sm font-medium disabled:opacity-50"
            onClick={handleCreate}
            disabled={loading}
          >
            {loading ? 'Generating...' : 'Generate'}
          </button>
        </div>

        {error && <p className="text-red-400 text-sm mt-3">{error}</p>}

        {result && (
          <div className="mt-5 p-4 bg-blue-50 rounded-lg">
            <p className="text-xs text-gray-400 mb-1">Your short link</p>

              <a href={result.short_url}
              target="_blank"
              className="text-blue-500 font-medium text-sm hover:underline break-all"
            >
              {result.short_url}
            </a>
            <p className="text-xs text-gray-400 mt-2 break-all">Original: {result.original}</p>
          </div>
        )}
      </div>

      {/* View stats */}
      <div className="w-full max-w-xl bg-white rounded-2xl shadow p-8">
        <h2 className="text-lg font-semibold text-gray-700 mb-4">View Stats</h2>
        <div className="flex gap-2">
          <input
            className="flex-1 border border-gray-200 rounded-lg px-4 py-2 text-sm outline-none focus:border-blue-400"
            placeholder="Enter short code, e.g. opU2Y2"
            value={statsCode}
            onChange={e => setStatsCode(e.target.value)}
            onKeyDown={e => e.key === 'Enter' && handleStats()}
          />
          <button
            className="bg-gray-700 hover:bg-gray-800 text-white px-5 py-2 rounded-lg text-sm font-medium"
            onClick={handleStats}
          >
            Look Up
          </button>
        </div>

        {statsError && <p className="text-red-400 text-sm mt-3">{statsError}</p>}

        {stats && (
          <div className="mt-5 p-4 bg-gray-50 rounded-lg space-y-2">
            <div className="flex justify-between text-sm">
              <span className="text-gray-400">Short Code</span>
              <span className="font-medium">{stats.code}</span>
            </div>
            <div className="flex justify-between text-sm">
              <span className="text-gray-400">Clicks</span>
              <span className="font-medium text-blue-500">{stats.clicks}</span>
            </div>
            <div className="flex justify-between text-sm">
              <span className="text-gray-400">Created At</span>
              <span className="font-medium">{new Date(stats.created_at).toLocaleString()}</span>
            </div>
            <div className="text-sm">
              <span className="text-gray-400">Original URL</span>

                <a href={stats.original}
                target="_blank"
                className="block text-blue-500 hover:underline break-all mt-1"
              >
                {stats.original}
              </a>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}