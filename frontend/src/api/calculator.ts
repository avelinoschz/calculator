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

const SERVER_UNAVAILABLE = 'Unable to reach the server. Make sure the backend is running.'

export async function calculate(req: CalculateRequest): Promise<CalculateResponse> {
  let response: Response

  try {
    response = await fetch('/api/v1/calculations', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    })
  } catch {
    throw new Error(SERVER_UNAVAILABLE)
  }

  if (!response.ok) {
    const contentType = response.headers.get('content-type')
    if (contentType?.includes('application/json')) {
      const body = await response.json() as { error: ApiError }
      throw new Error(body.error.message)
    }
    throw new Error(SERVER_UNAVAILABLE)
  }

  return response.json() as Promise<CalculateResponse>
}
