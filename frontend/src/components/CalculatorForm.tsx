import { useState } from 'react';
import type { CalculateRequest, Operation } from '../api/calculator';

interface Props {
  onSubmit: (request: CalculateRequest) => void;
  loading: boolean;
}

type OperationConfig = {
  value: Operation;
  label: string;
  arity: 'unary' | 'binary';
  firstLabel: string;
  firstPlaceholder: string;
  secondLabel?: string;
  secondPlaceholder?: string;
  helperText?: string;
};

const OPERATIONS: OperationConfig[] = [
  {
    value: 'add',
    label: 'Add (+)',
    arity: 'binary',
    firstLabel: 'First operand',
    firstPlaceholder: 'e.g. 10',
    secondLabel: 'Second operand',
    secondPlaceholder: 'e.g. 5',
  },
  {
    value: 'subtract',
    label: 'Subtract (−)',
    arity: 'binary',
    firstLabel: 'First operand',
    firstPlaceholder: 'e.g. 10',
    secondLabel: 'Second operand',
    secondPlaceholder: 'e.g. 3',
  },
  {
    value: 'multiply',
    label: 'Multiply (×)',
    arity: 'binary',
    firstLabel: 'First operand',
    firstPlaceholder: 'e.g. 4',
    secondLabel: 'Second operand',
    secondPlaceholder: 'e.g. 5',
  },
  {
    value: 'divide',
    label: 'Divide (÷)',
    arity: 'binary',
    firstLabel: 'First operand',
    firstPlaceholder: 'e.g. 20',
    secondLabel: 'Second operand',
    secondPlaceholder: 'e.g. 4',
  },
  {
    value: 'power',
    label: 'Power (x^y)',
    arity: 'binary',
    firstLabel: 'Base',
    firstPlaceholder: 'e.g. 2',
    secondLabel: 'Exponent',
    secondPlaceholder: 'e.g. 3',
  },
  {
    value: 'sqrt',
    label: 'Square root (sqrt)',
    arity: 'unary',
    firstLabel: 'Operand',
    firstPlaceholder: 'e.g. 9',
    helperText: 'Square root uses only one operand.',
  },
  {
    value: 'percentage',
    label: 'Percentage (% of)',
    arity: 'binary',
    firstLabel: 'Percentage',
    firstPlaceholder: 'e.g. 10',
    secondLabel: 'Of value',
    secondPlaceholder: 'e.g. 200',
    helperText: 'Calculates a% of b.',
  },
];

function parseOperand(value: string): number | null {
  const trimmed = value.trim();
  if (trimmed === '') {
    return null;
  }

  const parsed = Number(trimmed);
  return Number.isFinite(parsed) ? parsed : null;
}

function parseLimit(value: string | undefined): number | null {
  if (value === undefined) {
    return null;
  }

  const trimmed = value.trim();
  if (trimmed === '') {
    return null;
  }

  const parsed = Number(trimmed);
  return Number.isFinite(parsed) ? parsed : null;
}

export function CalculatorForm({ onSubmit, loading }: Props) {
  const [a, setA] = useState('');
  const [b, setB] = useState('');
  const [op, setOp] = useState<Operation>('add');
  const [validationError, setValidationError] = useState<string | null>(null);
  const selectedOperation =
    OPERATIONS.find((operation) => operation.value === op) ?? OPERATIONS[0];
  const requiresSecondOperand = selectedOperation.arity === 'binary';

  function updateFirstOperand(value: string) {
    setA(value);
    setValidationError(null);
  }

  function updateSecondOperand(value: string) {
    setB(value);
    setValidationError(null);
  }

  function updateOperation(value: Operation) {
    setOp(value);
    setValidationError(null);
  }

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();

    if (a.trim() === '') {
      setValidationError(
        requiresSecondOperand
          ? 'Both operands are required.'
          : 'Operand is required.',
      );
      return;
    }

    const numA = parseOperand(a);

    if (numA === null) {
      setValidationError(
        requiresSecondOperand
          ? 'Operands must be valid numbers.'
          : 'Operand must be a valid number.',
      );
      return;
    }

    const calcMin = parseLimit(import.meta.env.VITE_CALC_MIN);
    const calcMax = parseLimit(import.meta.env.VITE_CALC_MAX);

    if (requiresSecondOperand && b.trim() === '') {
      setValidationError('Both operands are required.');
      return;
    }

    let numB: number | null = null;
    if (requiresSecondOperand) {
      numB = parseOperand(b);
      if (numB === null) {
        setValidationError('Operands must be valid numbers.');
        return;
      }
    }

    if (selectedOperation.value === 'sqrt' && numA < 0) {
      setValidationError('Square root requires a non-negative number.');
      return;
    }

    if (calcMin !== null && numA < calcMin) {
      setValidationError(
        requiresSecondOperand
          ? `Operands must be at least ${calcMin}.`
          : `Operand must be at least ${calcMin}.`,
      );
      return;
    }

    if (calcMax !== null && numA > calcMax) {
      setValidationError(
        requiresSecondOperand
          ? `Operands must be at most ${calcMax}.`
          : `Operand must be at most ${calcMax}.`,
      );
      return;
    }

    if (requiresSecondOperand && calcMin !== null && numB < calcMin) {
      setValidationError(`Operands must be at least ${calcMin}.`);
      return;
    }

    if (requiresSecondOperand && calcMax !== null && numB > calcMax) {
      setValidationError(`Operands must be at most ${calcMax}.`);
      return;
    }

    setValidationError(null);
    if (selectedOperation.value === 'sqrt') {
      onSubmit({ op: selectedOperation.value, a: numA });
      return;
    }

    onSubmit({ op: selectedOperation.value, a: numA, b: numB });
  }

  return (
    <form onSubmit={handleSubmit} noValidate>
      <div className="field">
        <label htmlFor="operand-a">{selectedOperation.firstLabel}</label>
        <input
          id="operand-a"
          type="text"
          inputMode="decimal"
          value={a}
          onChange={(e) => updateFirstOperand(e.target.value)}
          placeholder={selectedOperation.firstPlaceholder}
          aria-label={selectedOperation.firstLabel}
        />
      </div>

      <div className="field">
        <label htmlFor="operation">Operation</label>
        <select
          id="operation"
          value={op}
          onChange={(e) => updateOperation(e.target.value as Operation)}
          aria-label="Operation"
        >
          {OPERATIONS.map(({ value, label }) => (
            <option key={value} value={value}>
              {label}
            </option>
          ))}
        </select>
      </div>

      {selectedOperation.helperText && (
        <p className="helper-text">{selectedOperation.helperText}</p>
      )}

      {requiresSecondOperand && (
        <div className="field">
          <label htmlFor="operand-b">{selectedOperation.secondLabel}</label>
          <input
            id="operand-b"
            type="text"
            inputMode="decimal"
            value={b}
            onChange={(e) => updateSecondOperand(e.target.value)}
            placeholder={selectedOperation.secondPlaceholder}
            aria-label={selectedOperation.secondLabel}
          />
        </div>
      )}

      {validationError && (
        <p className="error" role="alert">
          {validationError}
        </p>
      )}

      <button type="submit" disabled={loading}>
        {loading ? 'Calculating…' : 'Calculate'}
      </button>
    </form>
  );
}
