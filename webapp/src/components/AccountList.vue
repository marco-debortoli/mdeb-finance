<script setup lang="ts">
import { useAccountStore } from '@/stores/account';
import { formatCurrency } from '@/tools/currency';
import type { Account } from '@/types/account';
import dayjs from 'dayjs';

const props = defineProps({
  date: Date
});

const accountStore = useAccountStore()

function formatDate(date: string | undefined) {
  return dayjs(date).local().format("YYYY-MM-DD");
}

function accountValue(account: Account) {
  const currentDate = dayjs(props.date);
  const fVal = account.values.find((v) => currentDate.isSame(dayjs(v.date)));

  if (fVal) return fVal.value;
  return 0;
}

</script>

<template>
  <div class="flex flex-col p-4 border rounded-md border-black/30 w-full h-full">
    <div class="block mb-6">
      <h2 class="text-2xl font-bold uppercase"># Accounts</h2>
      <small class="text-xs hidden xl:block">
        > Latest account values for this month
      </small>
    </div>

    <div class="flex flex-col flex-grow gap-y-2 overflow-auto">
      <div class="captitalize underline flex justify-center">
        DEBIT
      </div>
      <template v-for="account in accountStore.accounts" :key="account._id">
        <div class="flex border border-black rounded-lg py-2 px-4 shadow-md justify-between">
          <div class="flex flex-col flex-grow justify-center">
            <span class="font-bold">{{ account.name }}</span>
            <span class="text-xs">As of {{ formatDate(account.current_value?.date) }}</span>
          </div>

          <div class="flex flex-col">
            <span>{{ formatCurrency(accountValue(account)) }}</span>
            <span>+100%</span>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>