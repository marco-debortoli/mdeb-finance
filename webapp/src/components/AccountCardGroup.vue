<script setup lang="ts">
import { useAccountStore } from '@/stores/account';
import { getAccountsTotalForDate, getActiveAccountsForDate, absoluteAccountValueDiff } from '@/tools/account';
import AccountCard from '@/components/AccountCard.vue';
import { computed } from 'vue';
import { formatCurrency } from '@/tools/currency';
import dayjs from 'dayjs';

const props = defineProps<{
  date: Date,
  accType: 'debit' | 'credit' | 'investment'
}>();

const accountStore = useAccountStore();

const lastMonth = computed(() => {
  return dayjs(props.date).startOf('month').subtract(1, 'month').toDate();
})

const accounts = computed(() => {
  switch (props.accType) {
    case 'debit':
      return accountStore.debitAccounts;
      
    case 'credit':
      return accountStore.creditAccounts;
  
    default:
      return accountStore.investAccounts;
  }
});

const thisMonthAccounts = computed(() => {
  return getActiveAccountsForDate(accounts.value, props.date);
});

const thisMonthTotal = computed(() => {
  return getAccountsTotalForDate(thisMonthAccounts.value, props.date);
});

const lastMonthAccounts = computed(() => {
  return getActiveAccountsForDate(accounts.value, props.date);
});

const lastMonthTotal = computed(() => {
  return getAccountsTotalForDate(lastMonthAccounts.value, lastMonth.value);
});

const lastMonthDiff = computed(() => {
  return absoluteAccountValueDiff(thisMonthTotal.value, lastMonthTotal.value);
});

function formatColour(amount: number) {
  if (amount > 0) return 'text-green-700';
  if (amount == 0) return 'text-current';
  return 'text-red-700';
}

</script>

<template>
  <div class="flex flex-col gap-y-2" v-if="thisMonthAccounts.length > 0">
    <div class="uppercase underline flex justify-center">
      {{ accType }}
    </div>

    <div class="flex justify-center">
      <span>{{ formatCurrency(thisMonthTotal) }}&nbsp;/&nbsp;</span>
      <span :class="formatColour(lastMonthDiff)">{{ formatCurrency(lastMonthDiff, true) }}</span>
    </div>

    <template v-for="account in thisMonthAccounts" :key="account._id">
      <AccountCard
        :date="date"
        :account="account"
      />
    </template>
  </div>
</template>