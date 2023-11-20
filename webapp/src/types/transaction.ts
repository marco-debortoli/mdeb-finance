type TransactionCategory = {
    _id: string;
    name: string;
}

export type Transaction = {
    _id: string;
    amount: number;
    date: string;
    name: string;
    category: TransactionCategory;
}