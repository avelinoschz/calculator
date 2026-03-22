import { describe, it, expect, vi, afterEach } from 'vitest'
import { calculate } from './calculator'

afterEach(() => {
  vi.unstubAllGlobals()
})

function mockFetch(status: number, body: unknown) {
  vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
    ok: status >= 200 && status < 300,
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
})
