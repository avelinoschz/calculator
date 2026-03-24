import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { CalculatorForm } from './CalculatorForm';

function setup(onSubmit = vi.fn()) {
  render(<CalculatorForm onSubmit={onSubmit} loading={false} />);
  return { onSubmit };
}

describe('CalculatorForm', () => {
  it('calls onSubmit with parsed numbers when inputs are valid', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.clear(screen.getByLabelText('First operand'));
    await user.type(screen.getByLabelText('First operand'), '10');
    await user.clear(screen.getByLabelText('Second operand'));
    await user.type(screen.getByLabelText('Second operand'), '5');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(onSubmit).toHaveBeenCalledWith({ op: 'add', a: 10, b: 5 });
  });

  it('shows validation error when first operand is empty', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.clear(screen.getByLabelText('Second operand'));
    await user.type(screen.getByLabelText('Second operand'), '5');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Both operands are required.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('shows validation error when second operand is empty', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.type(screen.getByLabelText('First operand'), '10');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Both operands are required.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('shows validation error for non-numeric input', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.type(screen.getByLabelText('First operand'), 'abc');
    await user.type(screen.getByLabelText('Second operand'), '5');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Operands must be valid numbers.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('shows validation error for partial numeric garbage', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.type(screen.getByLabelText('First operand'), '12abc');
    await user.type(screen.getByLabelText('Second operand'), '5');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Operands must be valid numbers.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('shows validation error for whitespace-only input', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.type(screen.getByLabelText('First operand'), '   ');
    await user.type(screen.getByLabelText('Second operand'), '5');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Both operands are required.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('shows validation error for non-finite input', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.type(screen.getByLabelText('First operand'), 'Infinity');
    await user.type(screen.getByLabelText('Second operand'), '5');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Operands must be valid numbers.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('disables the submit button and shows loading text while loading', () => {
    const onSubmit = vi.fn();
    render(<CalculatorForm onSubmit={onSubmit} loading={true} />);
    const button = screen.getByRole('button');
    expect(button).toBeDisabled();
    expect(button).toHaveTextContent('Calculating…');
  });

  it('calls onSubmit with selected operation', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.type(screen.getByLabelText('First operand'), '3');
    await user.type(screen.getByLabelText('Second operand'), '4');
    await user.selectOptions(screen.getByLabelText('Operation'), 'multiply');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(onSubmit).toHaveBeenCalledWith({ op: 'multiply', a: 3, b: 4 });
  });

  it('shows advanced operations in the selector', () => {
    setup();

    expect(
      screen.getByRole('option', { name: 'Power (x^y)' }),
    ).toBeInTheDocument();
    expect(
      screen.getByRole('option', { name: 'Square root (sqrt)' }),
    ).toBeInTheDocument();
    expect(
      screen.getByRole('option', { name: 'Percentage (% of)' }),
    ).toBeInTheDocument();
  });

  it('submits sqrt with only one operand', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.selectOptions(screen.getByLabelText('Operation'), 'sqrt');
    expect(screen.queryByLabelText('Second operand')).not.toBeInTheDocument();

    await user.type(screen.getByLabelText('Operand'), '9');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(onSubmit).toHaveBeenCalledWith({ op: 'sqrt', a: 9 });
  });

  it('blocks negative square root before submit', async () => {
    const user = userEvent.setup();
    const { onSubmit } = setup();

    await user.selectOptions(screen.getByLabelText('Operation'), 'sqrt');
    await user.type(screen.getByLabelText('Operand'), '-9');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Square root requires a non-negative number.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('restores second operand when switching back to a binary operation', async () => {
    const user = userEvent.setup();
    setup();

    await user.selectOptions(screen.getByLabelText('Operation'), 'sqrt');
    expect(screen.queryByLabelText('Second operand')).not.toBeInTheDocument();

    await user.selectOptions(screen.getByLabelText('Operation'), 'power');
    expect(screen.getByLabelText('Exponent')).toBeInTheDocument();
  });
});

describe('CalculatorForm operand limits', () => {
  beforeEach(() => {
    vi.stubEnv('VITE_CALC_MIN', '-100');
    vi.stubEnv('VITE_CALC_MAX', '100');
  });

  afterEach(() => {
    vi.unstubAllEnvs();
  });

  it('shows error when first operand is below min', async () => {
    const user = userEvent.setup();
    const onSubmit = vi.fn();
    render(<CalculatorForm onSubmit={onSubmit} loading={false} />);

    await user.type(screen.getByLabelText('First operand'), '-200');
    await user.type(screen.getByLabelText('Second operand'), '5');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Operands must be at least -100.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('shows error when second operand exceeds max', async () => {
    const user = userEvent.setup();
    const onSubmit = vi.fn();
    render(<CalculatorForm onSubmit={onSubmit} loading={false} />);

    await user.type(screen.getByLabelText('First operand'), '10');
    await user.type(screen.getByLabelText('Second operand'), '200');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Operands must be at most 100.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });

  it('submits when operands are at exact boundaries', async () => {
    const user = userEvent.setup();
    const onSubmit = vi.fn();
    render(<CalculatorForm onSubmit={onSubmit} loading={false} />);

    await user.type(screen.getByLabelText('First operand'), '-100');
    await user.type(screen.getByLabelText('Second operand'), '100');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(onSubmit).toHaveBeenCalledWith({ op: 'add', a: -100, b: 100 });
  });

  it('applies min and max limits only to the active operand for sqrt', async () => {
    const user = userEvent.setup();
    const onSubmit = vi.fn();
    render(<CalculatorForm onSubmit={onSubmit} loading={false} />);

    await user.selectOptions(screen.getByLabelText('Operation'), 'sqrt');
    await user.type(screen.getByLabelText('Operand'), '200');
    await user.click(screen.getByRole('button', { name: 'Calculate' }));

    expect(screen.getByRole('alert')).toHaveTextContent(
      'Operand must be at most 100.',
    );
    expect(onSubmit).not.toHaveBeenCalled();
  });
});
