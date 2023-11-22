import type { Account } from '@/types/account';
import { defineStore } from 'pinia'
import dayjs from 'dayjs';

type Nullable<T> = T | null;

export const useAccountStore = defineStore(
  'account',
  {
    state: () => (
      {
        accounts: [] as Account[],
        loading: false,
        date: null as Nullable<Date>
      }
    ),
    getters: {},
    actions: {
      async retrieve(date: Date | null = null) {
        this.date = date;

        let url = "http://localhost:8080/api/v1/accounts";

        if (date !== null) {
            const startDate = dayjs(date).startOf('month');
            url = `http://localhost:8080/api/v1/accounts?date=${startDate.local().format()}`
        }

        this.loading = true;
        this.accounts = [];

        const response = await fetch(url);

        this.loading = false;

        if (response.status != 200) {
          console.log("Failed to fetch accounts")
          return response;
        } else {
          this.accounts = await response.json();
          return response;
        }
      }
    },
  }
);