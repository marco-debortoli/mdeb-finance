<script setup lang="ts">
import { useAccountStore } from '@/stores/account';
import { getActiveAccountsForDate, absoluteAccountValueDiff } from '@/tools/account';
import { formatCurrency } from '@/tools/currency';
import dayjs from 'dayjs';
import { computed } from 'vue'; 

const props = defineProps<{
  date: Date
}>();

const accountStore = useAccountStore();

const lastMonth = computed(() => {
  return dayjs(props.date).startOf('month').subtract(1, 'month').toDate();
});

const thisMonthAccounts = computed(() => {
  return getActiveAccountsForDate(accountStore.netWorthAccounts, props.date);
});

const lastMonthAccounts = computed(() => {
  return getActiveAccountsForDate(accountStore.netWorthAccounts, lastMonth.value);
});

const thisMonthNetWorth = computed(() => {
  let netWorth = 0;
  const currentDate = dayjs(props.date);

  thisMonthAccounts.value.forEach((a) => {
    const av = a.values.find((v) => currentDate.isSame(dayjs(v.date)));
    if (av) {
      if (a.type === "debit" || a.type === "investment") netWorth += av.value;
      else netWorth -= av.value;
    }
  });

  return netWorth;
});

const lastMonthNetWorth = computed(() => {
  let netWorth = 0;
  const currentDate = dayjs(lastMonth.value);

  lastMonthAccounts.value.forEach((a) => {
    const av = a.values.find((v) => currentDate.isSame(dayjs(v.date)));
    if (av) {
      if (a.type === "debit" || a.type === "investment") netWorth += av.value;
      else netWorth -= av.value;
    }
  });

  return netWorth;
});

const netWorthDiff = computed(() => {
  return absoluteAccountValueDiff(thisMonthNetWorth.value, lastMonthNetWorth.value);
});

function formatColour(amount: number) {
  if (amount > 0) return 'text-green-700';
  if (amount == 0) return 'text-current';
  return 'text-red-700';
}

</script>

<template>
  <div class="flex flex-grow items-center border rounded-md border-black/30">
    <div class="flex justify-center px-4 border-r border-black/30">
      NET WORTH
    </div>

    <div class="flex justify-center flex-grow font-black">
      <span>{{ formatCurrency(thisMonthNetWorth) }}&nbsp;/&nbsp;</span> 
      <span :class="formatColour(netWorthDiff)">{{ formatCurrency(netWorthDiff, true) }}</span>
    </div>
  </div>
</template>