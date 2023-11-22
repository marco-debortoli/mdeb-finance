import type { Transaction } from '@/types/transaction';
import { defineStore } from 'pinia'
import dayjs from 'dayjs';

export const useTransactionStore = defineStore(
  'transaction',
  {
    state: () => (
      {
        transactions: [] as Transaction[],
        loading: false,
        date: undefined as Date | undefined
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
      async retrieve(date: Date) {
        // We need to get the upper and lower bounds of transaction dates
        // The provided date should always be the start of the month
        this.date = date;

        const startDate = dayjs(date).startOf('month');
        const endDate = dayjs(date).endOf('month');

        const startDateFilter = startDate.local().format();
        const endDateFilter = endDate.local().format()

        this.loading = true;
        this.transactions = [];

        const response = await fetch(
          `http://localhost:8080/api/v1/transactions?start_date=${startDateFilter}&end_date=${endDateFilter}`
        );

        this.loading = false;

        if (response.status != 200) {
          console.log("Failed to fetch transactions")
          return response;
        } else {
          this.transactions = await response.json();
          return response;
        }
      },

      async refresh() {
        if (this.date) {
          return this.retrieve(this.date)
        }
        return null;
      },

      async create(date: string, amount: number, category: string, name: string) {
        // First we should convert our date - which is a YYYY-MM-DD string to the proper format
        const createDate = dayjs(date).startOf('day');

        const createData = {
          date: createDate.local().format(),
          amount: amount,
          category: category,
          name: name
        };

        const response = await fetch(
          "http://localhost:8080/api/v1/transactions",
          {
          method: "post",
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },

          //make sure to serialize your JSON body
          body: JSON.stringify(createData)
        });

        return response;
      }
    },
  }
);