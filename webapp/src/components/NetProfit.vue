<script setup lang="ts">
import { useTransactionStore } from '@/stores/transaction';
import { formatCurrency } from '@/tools/currency';
import { computed } from 'vue'; 


defineProps({
  date: Date
});

const transactionStore = useTransactionStore();

const monthProfit = computed(() => {
  let profit = 0;

  transactionStore.transactions.forEach((tr) => {
    if (tr.category.type == "debit") {
      profit += tr.amount;
    } else {
      profit -= tr.amount;
    }
  });

  return profit
});

const color = computed(() => {
  if (monthProfit.value > 0) return 'text-green-700';
  if (monthProfit.value == 0) return 'text-current';
  return 'text-red-700';
})

</script>

<template>
  <div class="flex flex-grow items-center border rounded-md border-black/30">
    <div class="flex justify-center px-4 border-r border-black/30">
      MONTH
    </div>

    <div class="flex justify-center flex-grow">
      <span class="font-black" :class="color">
        {{ formatCurrency(monthProfit, true) }}
      </span>
    </div>
  </div>
</template>