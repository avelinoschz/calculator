import { useState } from 'react'
import type { Operation } from '../api/calculator'

interface Props {
  onSubmit: (op: Operation, a: number, b: number) => void
  loading: boolean
}

const OPERATIONS: { value: Operation; label: string }[] = [
  { value: 'add', label: 'Add (+)' },
  { value: 'subtract', label: 'Subtract (−)' },
  { value: 'multiply', label: 'Multiply (×)' },
  { value: 'divide', label: 'Divide (÷)' },
]

function parseOperand(value: string): number | null {
  const trimmed = value.trim()
  if (trimmed === '') {
    return null
  }

  const parsed = Number(trimmed)
  return Number.isFinite(parsed) ? parsed : null
}

function parseLimit(value: string | undefined): number | null {
  if (value === undefined) {
    return null
  }

  const trimmed = value.trim()
  if (trimmed === '') {
    return null
  }

  const parsed = Number(trimmed)
  return Number.isFinite(parsed) ? parsed : null
}

export function CalculatorForm({ onSubmit, loading }: Props) {
  const [a, setA] = useState('')
  const [b, setB] = useState('')
  const [op, setOp] = useState<Operation>('add')
  const [validationError, setValidationError] = useState<string | null>(null)

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault()

    if (a.trim() === '' || b.trim() === '') {
      setValidationError('Both operands are required.')
      return
    }

    const numA = parseOperand(a)
    const numB = parseOperand(b)

    if (numA === null || numB === null) {
      setValidationError('Operands must be valid numbers.')
      return
    }

    const calcMin = parseLimit(import.meta.env.VITE_CALC_MIN)
    const calcMax = parseLimit(import.meta.env.VITE_CALC_MAX)

    if (calcMin !== null && (numA < calcMin || numB < calcMin)) {
      setValidationError(`Operands must be at least ${calcMin}.`)
      return
    }

    if (calcMax !== null && (numA > calcMax || numB > calcMax)) {
      setValidationError(`Operands must be at most ${calcMax}.`)
      return
    }

    setValidationError(null)
    onSubmit(op, numA, numB)
  }

  return (
    <form onSubmit={handleSubmit} noValidate>
      <div className="field">
        <label htmlFor="operand-a">First operand</label>
        <input
          id="operand-a"
          type="text"
          inputMode="decimal"
          value={a}
          onChange={(e) => setA(e.target.value)}
          placeholder="e.g. 10"
          aria-label="First operand"
        />
      </div>

      <div className="field">
        <label htmlFor="operation">Operation</label>
        <select
          id="operation"
          value={op}
          onChange={(e) => setOp(e.target.value as Operation)}
          aria-label="Operation"
        >
          {OPERATIONS.map(({ value, label }) => (
            <option key={value} value={value}>{label}</option>
          ))}
        </select>
      </div>

      <div className="field">
        <label htmlFor="operand-b">Second operand</label>
        <input
          id="operand-b"
          type="text"
          inputMode="decimal"
          value={b}
          onChange={(e) => setB(e.target.value)}
          placeholder="e.g. 5"
          aria-label="Second operand"
        />
      </div>

      {validationError && (
        <p className="error" role="alert">{validationError}</p>
      )}

      <button type="submit" disabled={loading}>
        {loading ? 'Calculating…' : 'Calculate'}
      </button>
    </form>
  )
}
