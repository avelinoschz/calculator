import { describe, it, expect, vi, afterEach } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import App from './App'

function mockFetch(status: number, body: unknown) {
  vi.stubGlobal('fetch', vi.fn().mockResolvedValue({
    ok: status >= 200 && status < 300,
    json: () => Promise.resolve(body),
  }))
}

afterEach(() => {
  vi.unstubAllGlobals()
})

async function fillAndSubmit(a: string, b: string) {
  const user = userEvent.setup()
  await user.type(screen.getByLabelText('First operand'), a)
  await user.type(screen.getByLabelText('Second operand'), b)
  await user.click(screen.getByRole('button', { name: 'Calculate' }))
}

describe('App', () => {
  it('displays result after successful calculation', async () => {
    mockFetch(200, { result: 15 })
    render(<App />)
    await fillAndSubmit('10', '5')
    await waitFor(() => expect(screen.getByRole('status')).toHaveTextContent('15'))
  })

  it('displays error message when API returns an error', async () => {
    mockFetch(422, { error: { code: 'DIVISION_BY_ZERO', message: 'division by zero is not allowed' } })
    render(<App />)
    await fillAndSubmit('10', '0')
    await waitFor(() => expect(screen.getByRole('alert')).toHaveTextContent('division by zero is not allowed'))
  })

  it('clears previous result when a new submission is made', async () => {
    mockFetch(200, { result: 15 })
    render(<App />)
    await fillAndSubmit('10', '5')
    await waitFor(() => expect(screen.getByRole('status')).toHaveTextContent('15'))

    mockFetch(422, { error: { code: 'DIVISION_BY_ZERO', message: 'division by zero is not allowed' } })
    await userEvent.click(screen.getByRole('button', { name: 'Calculate' }))

    await waitFor(() => {
      expect(screen.queryByRole('status')).toBeNull()
      expect(screen.getByRole('alert')).toHaveTextContent('division by zero is not allowed')
    })
  })
})
