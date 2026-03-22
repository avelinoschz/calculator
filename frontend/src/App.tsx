import { useState } from 'react'
import { calculate } from './api/calculator'
import type { Operation } from './api/calculator'
import { CalculatorForm } from './components/CalculatorForm'
import './App.css'

function App() {
  const [result, setResult] = useState<number | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  async function handleSubmit(op: Operation, a: number, b: number) {
    setResult(null)
    setError(null)
    setLoading(true)

    try {
      const response = await calculate({ op, a, b })
      setResult(response.result)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An unexpected error occurred.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="card">
      <h1>Calculator</h1>
      <CalculatorForm onSubmit={handleSubmit} loading={loading} />
      {result !== null && (
        <div className="result" role="status" aria-live="polite">
          <p>Result</p>
          <span>{result}</span>
        </div>
      )}
      {error !== null && (
        <p className="error" role="alert">{error}</p>
      )}
    </div>
  )
}

export default App
