import type { Transaction } from '@/types/transaction';
import { defineStore } from 'pinia'

export const useTransactionStore = defineStore(
  'transaction',
  {
    state: () => (
      {
        transactions: [] as Transaction[],
        loading: false
      }
    ),
    getters: {
      sortedTransactions(state) {
        const sorted = [...state.transactions].sort((a, b) => {
          if (a.date < b.date) {
            return -1;
          }
          return 1;
        });
        return sorted;
      },
    },
    actions: {
      async retrieve() {
        this.loading = true;

        const response = await fetch(
          "http://localhost:8080/api/v1/transactions"
        );

        this.loading = false;

        if (response.status != 200) {
          console.log("Failed to fetch transactions")
          return response;
        } else {
          this.transactions = await response.json();
          return response;
        }
      }
    },
  }
);