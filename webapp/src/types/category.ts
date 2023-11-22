export type Category = {
    _id: string;
    name: string;
    type: 'debit' | 'credit' | 'transfer';
    active: boolean;
}