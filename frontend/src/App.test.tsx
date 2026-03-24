import { describe, it, expect, vi, afterEach } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import App from './App';

function mockFetch(status: number, body: unknown) {
  vi.stubGlobal(
    'fetch',
    vi.fn().mockResolvedValue({
      ok: status >= 200 && status < 300,
      headers: { get: () => 'application/json' },
      json: () => Promise.resolve(body),
    }),
  );
}

afterEach(() => {
  vi.unstubAllGlobals();
});

async function fillAndSubmit(a: string, b: string) {
  const user = userEvent.setup();
  await user.type(screen.getByLabelText('First operand'), a);
  await user.type(screen.getByLabelText('Second operand'), b);
  await user.click(screen.getByRole('button', { name: 'Calculate' }));
}

async function fillSqrtAndSubmit(a: string) {
  const user = userEvent.setup();
  await user.selectOptions(screen.getByLabelText('Operation'), 'sqrt');
  await user.type(screen.getByLabelText('Operand'), a);
  await user.click(screen.getByRole('button', { name: 'Calculate' }));
}

async function fillPowerAndSubmit(a: string, b: string) {
  const user = userEvent.setup();
  await user.selectOptions(screen.getByLabelText('Operation'), 'power');
  await user.type(screen.getByLabelText('Base'), a);
  await user.type(screen.getByLabelText('Exponent'), b);
  await user.click(screen.getByRole('button', { name: 'Calculate' }));
}

describe('App', () => {
  it('displays result after successful calculation', async () => {
    mockFetch(200, { result: 15 });
    render(<App />);
    await fillAndSubmit('10', '5');
    await waitFor(() =>
      expect(screen.getByRole('status')).toHaveTextContent('15'),
    );
  });

  it('displays error message when API returns an error', async () => {
    mockFetch(422, {
      error: {
        code: 'DIVISION_BY_ZERO',
        message: 'division by zero is not allowed',
      },
    });
    render(<App />);
    await fillAndSubmit('10', '0');
    await waitFor(() =>
      expect(screen.getByRole('alert')).toHaveTextContent(
        'division by zero is not allowed',
      ),
    );
  });

  it('displays result after successful square root calculation', async () => {
    mockFetch(200, { result: 3 });
    render(<App />);
    await fillSqrtAndSubmit('9');
    await waitFor(() =>
      expect(screen.getByRole('status')).toHaveTextContent('3'),
    );
  });

  it('displays advanced operation errors from the API', async () => {
    mockFetch(422, {
      error: {
        code: 'NON_FINITE_RESULT',
        message: 'calculation result is not a finite real number',
      },
    });
    render(<App />);
    await fillPowerAndSubmit('1e308', '2');
    await waitFor(() =>
      expect(screen.getByRole('alert')).toHaveTextContent(
        'calculation result is not a finite real number',
      ),
    );
  });

  it('clears previous result when a new submission is made', async () => {
    mockFetch(200, { result: 15 });
    render(<App />);
    await fillAndSubmit('10', '5');
    await waitFor(() =>
      expect(screen.getByRole('status')).toHaveTextContent('15'),
    );

    mockFetch(422, {
      error: {
        code: 'DIVISION_BY_ZERO',
        message: 'division by zero is not allowed',
      },
    });
    await userEvent.click(screen.getByRole('button', { name: 'Calculate' }));

    await waitFor(() => {
      expect(screen.queryByRole('status')).toBeNull();
      expect(screen.getByRole('alert')).toHaveTextContent(
        'division by zero is not allowed',
      );
    });
  });
});
