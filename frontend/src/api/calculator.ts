export type Operation = 'add' | 'subtract' | 'multiply' | 'divide'

export interface CalculateRequest {
  op: Operation
  a: number
  b: number
}

export interface CalculateResponse {
  result: number
}

export interface ApiError {
  code: string
  message: string
}

export async function calculate(req: CalculateRequest): Promise<CalculateResponse> {
  const response = await fetch('/api/v1/calculations', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(req),
  })

  if (!response.ok) {
    const body = await response.json() as { error: ApiError }
    throw new Error(body.error.message)
  }

  return response.json() as Promise<CalculateResponse>
}
