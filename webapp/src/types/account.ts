type AccountValue = {
  value: number;
  date: string;
}

export type Account = {
    _id: string;
    name: string;
    type: 'debit' | 'credit' | 'investment';
    active: boolean;
    include_net_worth: boolean;
    start_date: string;
    end_date: string | null;
    current_value: AccountValue | null;
    values: AccountValue[];
}