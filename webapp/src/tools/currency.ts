export function formatCurrency(amount: number, alwaysShowSign: boolean = false) {
  const options = {
    style: 'currency',
    currency: 'CAD',
    currencyDisplay: 'narrowSymbol'
  } as Intl.NumberFormatOptions

  if (alwaysShowSign) options.signDisplay = "exceptZero";

  const formatter = new Intl.NumberFormat('en-US', options);
  return formatter.format(amount);
}
