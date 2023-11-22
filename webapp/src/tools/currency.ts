export function formatCurrency(amount: number) {
  const formatter = new Intl.NumberFormat(
    'en-US', {
      style: 'currency',
      currency: 'CAD',
      currencyDisplay: 'narrowSymbol'
    }
  );

  return formatter.format(amount);
}
