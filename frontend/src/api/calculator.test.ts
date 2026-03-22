import { describe, it, expect, vi, afterEach } from 'vitest'
import { calculate } from './calculator'

afterEach(() => {
  vi.unstubAllGlobals()
})

function mockFetch(status: number, body: unknown) {
  vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
    ok: status >= 200 && status < 300,
    headers: { get: () => 'application/json' },
    json: () => Promise.resolve(body),
  }))
}

describe('calculate', () => {
  it('returns result on successful response', async () => {
    mockFetch(200, { result: 15 })
    const response = await calculate({ op: 'add', a: 10, b: 5 })
    expect(response.result).toBe(15)
  })

  it('sends correct request body', async () => {
    const fetchSpy = vi.fn().mockResolvedValue({
      ok: true,
      headers: { get: () => 'application/json' },
      json: () => Promise.resolve({ result: 6 }),
    })
    vi.stubGlobal('fetch', fetchSpy)

    await calculate({ op: 'multiply', a: 2, b: 3 })

    expect(fetchSpy).toHaveBeenCalledWith('/api/v1/calculations', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ op: 'multiply', a: 2, b: 3 }),
    })
  })

  it('throws with backend error message on non-200 response', async () => {
    mockFetch(422, { error: { code: 'DIVISION_BY_ZERO', message: 'division by zero is not allowed' } })
    await expect(calculate({ op: 'divide', a: 10, b: 0 })).rejects.toThrow('division by zero is not allowed')
  })

  it('throws with backend message on 400 response', async () => {
    mockFetch(400, { error: { code: 'INVALID_OPERATION', message: 'operation must be one of add, subtract, multiply, divide' } })
    await expect(calculate({ op: 'add', a: 1, b: 2 })).rejects.toThrow('operation must be one of add, subtract, multiply, divide')
  })

  it('throws a user-friendly message when the server is unreachable', async () => {
    vi.stubGlobal('fetch', vi.fn().mockRejectedValue(new TypeError('Failed to fetch')))
    await expect(calculate({ op: 'add', a: 1, b: 2 })).rejects.toThrow('Unable to reach the server. Make sure the backend is running.')
  })

  it('throws a user-friendly message when the proxy returns a non-JSON error response', async () => {
    vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
      ok: false,
      headers: { get: () => null },
    }))
    await expect(calculate({ op: 'add', a: 1, b: 2 })).rejects.toThrow('Unable to reach the server. Make sure the backend is running.')
  })
})
