import type { Transaction } from '@/types/transaction';
import { defineStore } from 'pinia'
import dayjs from 'dayjs';
import { apiDelete, apiGet, apiPost, apiPut } from '@/tools/api';

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

        const response = await apiGet(
          `/api/v1/transactions?start_date=${startDateFilter}&end_date=${endDateFilter}`
        );

        this.loading = false;

        if (response === undefined || response.status != 200) {
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

        const response = await apiPost(
          '/api/v1/transactions',
          createData
        );

        return response;
      },

      async update(
        id: string, date: string, amount: number, category: string, name: string
      ) {
        // First we should convert our date - which is a YYYY-MM-DD string to the proper format
        const updateDate = dayjs(date).startOf('day');

        const updateData = {
          date: updateDate.local().format(),
          amount: amount,
          category: category,
          name: name
        };

        const response = await apiPut(
          `/api/v1/transactions/${id}`,
          updateData
        );

        return response;
      },

      async delete(id: string) {
        return await apiDelete(`/api/v1/transactions/${id}`);
      }
    },
  }
);