import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { CalculatorForm } from './CalculatorForm'

function setup(onSubmit = vi.fn()) {
  render(<CalculatorForm onSubmit={onSubmit} loading={false} />)
  return { onSubmit }
}

describe('CalculatorForm', () => {
  it('calls onSubmit with parsed numbers when inputs are valid', async () => {
    const user = userEvent.setup()
    const { onSubmit } = setup()

    await user.clear(screen.getByLabelText('First operand'))
    await user.type(screen.getByLabelText('First operand'), '10')
    await user.clear(screen.getByLabelText('Second operand'))
    await user.type(screen.getByLabelText('Second operand'), '5')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(onSubmit).toHaveBeenCalledWith('add', 10, 5)
  })

  it('shows validation error when first operand is empty', async () => {
    const user = userEvent.setup()
    const { onSubmit } = setup()

    await user.clear(screen.getByLabelText('Second operand'))
    await user.type(screen.getByLabelText('Second operand'), '5')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(screen.getByRole('alert')).toHaveTextContent('Both operands are required.')
    expect(onSubmit).not.toHaveBeenCalled()
  })

  it('shows validation error when second operand is empty', async () => {
    const user = userEvent.setup()
    const { onSubmit } = setup()

    await user.type(screen.getByLabelText('First operand'), '10')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(screen.getByRole('alert')).toHaveTextContent('Both operands are required.')
    expect(onSubmit).not.toHaveBeenCalled()
  })

  it('shows validation error for non-numeric input', async () => {
    const user = userEvent.setup()
    const { onSubmit } = setup()

    await user.type(screen.getByLabelText('First operand'), 'abc')
    await user.type(screen.getByLabelText('Second operand'), '5')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(screen.getByRole('alert')).toHaveTextContent('Operands must be valid numbers.')
    expect(onSubmit).not.toHaveBeenCalled()
  })

  it('disables the submit button while loading', () => {
    const onSubmit = vi.fn()
    render(<CalculatorForm onSubmit={onSubmit} loading={true} />)
    expect(screen.getByRole('button')).toBeDisabled()
  })
})

describe('CalculatorForm operand limits', () => {
  beforeEach(() => {
    vi.stubEnv('VITE_CALC_MIN', '-100')
    vi.stubEnv('VITE_CALC_MAX', '100')
  })

  afterEach(() => {
    vi.unstubAllEnvs()
  })

  it('shows error when first operand is below min', async () => {
    const user = userEvent.setup()
    const onSubmit = vi.fn()
    render(<CalculatorForm onSubmit={onSubmit} loading={false} />)

    await user.type(screen.getByLabelText('First operand'), '-200')
    await user.type(screen.getByLabelText('Second operand'), '5')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(screen.getByRole('alert')).toHaveTextContent('Operands must be at least -100.')
    expect(onSubmit).not.toHaveBeenCalled()
  })

  it('shows error when second operand exceeds max', async () => {
    const user = userEvent.setup()
    const onSubmit = vi.fn()
    render(<CalculatorForm onSubmit={onSubmit} loading={false} />)

    await user.type(screen.getByLabelText('First operand'), '10')
    await user.type(screen.getByLabelText('Second operand'), '200')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(screen.getByRole('alert')).toHaveTextContent('Operands must be at most 100.')
    expect(onSubmit).not.toHaveBeenCalled()
  })

  it('submits when operands are at exact boundaries', async () => {
    const user = userEvent.setup()
    const onSubmit = vi.fn()
    render(<CalculatorForm onSubmit={onSubmit} loading={false} />)

    await user.type(screen.getByLabelText('First operand'), '-100')
    await user.type(screen.getByLabelText('Second operand'), '100')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(onSubmit).toHaveBeenCalledWith('add', -100, 100)
  })
})
