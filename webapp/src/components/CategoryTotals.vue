<script setup lang="ts">
import { useCategoryStore } from '@/stores/category';
import { useTransactionStore } from '@/stores/transaction';
import { formatCurrency } from '@/tools/currency';

defineProps({
  date: Date
});

const categoryStore = useCategoryStore();
const transactionStore = useTransactionStore();

function transactionTotal(catId: string) {
  let total = 0;

  transactionStore.transactions.forEach((tr) => {
    if (tr.category._id == catId) total += tr.amount;
  });

  return total;
}

</script>

<template>
  <div class="flex flex-col p-4 border rounded-md border-black/30 w-full h-full">
    <div class="block mb-6">
      <h2 class="text-2xl font-bold uppercase"># Categories</h2>
      <small class="text-xs hidden xl:block">
        > Calculated category totals for this month
      </small>
    </div>

    <div class="overflow-auto">
      <table class="w-full">
        <thead class="uppercase text-left">
          <tr class="border-b-2 border-black/30">
            <th class="w-1/2 border-r border-black/30 pl-2">Category</th>
            <th class="w-1/2 border-l border-black/30 pl-2">Total</th>
          </tr>
        </thead>
        <tbody class="text-sm">
          <tr class="border-b border-black/30 bg-black/10" v-if="categoryStore.debitCategories.length > 0">
            <td class="text-left pl-2 uppercase" colspan="100%">Debit</td>
          </tr>
          <template v-for="category in categoryStore.debitCategories" :key="category._id">
            <tr class="border-b border-black/30 hover:bg-black/30">
              <td class="pl-2 border-r border-black/30">{{ category.name }}</td>
              <td class="pl-2 border-l border-black/30">{{ formatCurrency(transactionTotal(category._id)) }}</td>
            </tr>
          </template>

          <tr class="border-b border-black/30 bg-black/10">
            <td class="text-left pl-2 uppercase" colspan="100%">Credit</td>
          </tr>
          <template v-for="category in categoryStore.creditCategories" :key="category._id">
            <tr class="border-b border-black/30 hover:bg-black/30">
              <td class="pl-2 border-r border-black/30">{{ category.name }}</td>
              <td class="pl-2 border-l border-black/30">{{ formatCurrency(transactionTotal(category._id)) }}</td>
            </tr>
          </template>
        </tbody>
      </table>
    </div>

  </div>
</template>