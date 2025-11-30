export const formatMoney = (number: number): string => {
    if (isNaN(number)) {
        return 'Valor no válido';
    }

    const formatter = new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
        minimumFractionDigits: 2,
        maximumFractionDigits: 2,
    });

    return formatter.format(number);
}