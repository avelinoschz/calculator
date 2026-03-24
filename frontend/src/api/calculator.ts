export type BinaryOperation =
  | 'add'
  | 'subtract'
  | 'multiply'
  | 'divide'
  | 'power'
  | 'percentage';
export type UnaryOperation = 'sqrt';
export type Operation = BinaryOperation | UnaryOperation;

export interface UnaryCalculateRequest {
  op: UnaryOperation;
  a: number;
}

export interface BinaryCalculateRequest {
  op: BinaryOperation;
  a: number;
  b: number;
}

export type CalculateRequest = UnaryCalculateRequest | BinaryCalculateRequest;

export interface CalculateResponse {
  result: number;
}

export interface ApiError {
  code: string;
  message: string;
}

const SERVER_UNAVAILABLE =
  'Unable to reach the server. Make sure the backend is running.';

export async function calculate(
  req: CalculateRequest,
): Promise<CalculateResponse> {
  let response: Response;

  try {
    response = await fetch('/api/v1/calculations', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    });
  } catch {
    throw new Error(SERVER_UNAVAILABLE);
  }

  if (!response.ok) {
    const contentType = response.headers.get('content-type');
    if (contentType?.includes('application/json')) {
      const body = (await response.json()) as { error: ApiError };
      throw new Error(body.error.message);
    }
    throw new Error(SERVER_UNAVAILABLE);
  }

  return response.json() as Promise<CalculateResponse>;
}
