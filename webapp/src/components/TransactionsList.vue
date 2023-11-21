<script setup lang="ts">
import { computed, ref } from 'vue';
import TransactionModal from '@/components/TransactionModal.vue';
import type { Transaction } from '@/types/transaction';
import type { Category } from '@/types/category';

const props = defineProps<{
  transactions: Transaction[]
  categories: Category[]
  loading: boolean
  date: Date
}>();

const sortedTransactions = computed(() => {
  const sorted = [...props.transactions].sort((a, b) => {
    if (a.date < b.date) {
      return -1;
    }
    return 1;
  });
  return sorted;
});

const startDate = computed(() => {
  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }
  return new Date(props.date).toLocaleDateString("en-US", options);
});

const endDate = computed(() => {
  const copyDate = new Date(props.date);
  copyDate.setMonth(copyDate.getMonth() + 1);
  copyDate.setDate(copyDate.getDate() - 1);

  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }
  return copyDate.toLocaleDateString("en-US", options);

});

function formatTransactionDate(isoDate: string) {
  const formattedDate = new Date(isoDate);
  const options: Intl.DateTimeFormatOptions = {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  };
  return formattedDate.toLocaleDateString("en-US", options);
}

function formatCurrency(amount: number) {
  const formatter = new Intl.NumberFormat(
    'en-US', {
      style: 'currency',
      currency: 'CAD',
      currencyDisplay: 'narrowSymbol'
    }
  );

  return formatter.format(amount);
}

// Add and edit
const openModal = ref(false);
const editModal = ref(false);

</script>

<template>
  <div class="block py-6 px-10 w-full h-full relative">
    <div class="flex justify-between">
      <div class="block">
        <h2 class="text-2xl font-bold uppercase"># Transactions</h2>
        <small class="text-xs hidden xl:block">
          > Recorded financial transactions from {{ startDate }} to {{ endDate }}
        </small>
      </div>
      <div class="block">
        <button
          class="border px-4 py-1 rounded-lg border-black/50 hover:border-black align-middle"
          @click="() => { editModal = false; openModal = true }"
        >
          <i class="ti ti-plus hover:cursor-pointer"/>
          Add
        </button>
      </div>
    </div>

    <table class="table-auto xl:table-fixed w-full mt-6" v-if="!loading">
      <thead class="uppercase border-b-2 border-black font-bold text-left text-lg">
        <tr>
          <th>Date</th>
          <th>Category</th>
          <th class="w-36">Amount</th>
          <th>Description</th>
          <th class="w-12"></th>
        </tr>
      </thead>
      <tbody class="text-sm">
        <template v-for="transaction in sortedTransactions" v-bind:key="transaction._id">
          <tr class="border-b border-black/25">
            <td class="font-semibold">{{ formatTransactionDate(transaction.date) }}</td>
            <td>{{ transaction.category.name }}</td>
            <td>{{ formatCurrency(transaction.amount) }}</td>
            <td>{{ transaction.name }}</td>
            <td
              class="text-center hover:cursor-pointer hover:underline"
              @click="() => { editModal = true; openModal = true }"
            >
              <small>Edit</small>
            </td>
          </tr>
        </template>
      </tbody>
    </table>

    <div v-else>
      LOADING
    </div>

    <!-- Add and Edit Modal -->
    <div v-if="openModal">
      <TransactionModal
        v-on:cancel="() => openModal = false"
        v-on:create="() => openModal = false"
        :categories="categories"
        :edit="editModal"
        :start-date="date"
      />
    </div>

  </div>
</template>