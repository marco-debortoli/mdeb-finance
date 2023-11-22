<script setup lang="ts">
import { computed, ref } from 'vue';
import TransactionModal from '@/components/TransactionModal.vue';
import { useTransactionStore } from '@/stores/transaction';
import { formatCurrency } from '@/tools/currency';
import type { Transaction } from '@/types/transaction';

const props = defineProps<{
  date: Date
}>();

const transactionStore = useTransactionStore();

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

// Add and edit
const openModal = ref(false);
const editTransaction = ref<Transaction | undefined>(undefined);

function closeTransactionModal(refresh: boolean = false) {
  openModal.value = false;
  editTransaction.value = undefined;

  if (refresh) transactionStore.refresh();
}

</script>

<template>
  <div class="flex flex-col py-6 px-10 w-full h-full relative">
    <div class="flex justify-between mb-6">
      <div class="block">
        <h2 class="text-2xl font-bold uppercase"># Transactions</h2>
        <small class="text-xs hidden xl:block">
          > Recorded financial transactions from {{ startDate }} to {{ endDate }}
        </small>
      </div>
      <div class="block">
        <button
          class="border px-4 py-1 rounded-lg border-black/50 hover:border-black align-middle"
          @click="() => { openModal = true }"
        >
          <i class="ti ti-plus hover:cursor-pointer"/>
          Add
        </button>
      </div>
    </div>

    <div class="overflow-auto">
      <table class="table-auto xl:table-fixed w-full">
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
          <template v-for="transaction in transactionStore.sortedTransactions" v-bind:key="transaction._id">
            <tr class="border-b border-black/30 hover:bg-black/30">
              <td class="font-semibold">{{ formatTransactionDate(transaction.date) }}</td>
              <td>{{ transaction.category.name }}</td>
              <td>{{ formatCurrency(transaction.amount) }}</td>
              <td>{{ transaction.name }}</td>
              <td
                class="text-center hover:cursor-pointer hover:underline"
                @click="() => { editTransaction = transaction; openModal = true }"
              >
                <small>Edit</small>
              </td>
            </tr>
          </template>
        </tbody>
      </table>
    </div>

    <div
      class="
        flex flex-col mt-2 border border-dashed border-black/30
        rounded-md h-60 text-center justify-center text-black/30
      "
      v-if="
        transactionStore.sortedTransactions.length == 0
        && !transactionStore.loading
      "
    >
      <i class="ti ti-file-x text-5xl"></i>
      <span class="uppercase">NO TRANSACTIONS</span>
    </div>

    <!-- Add and Edit Modal -->
    <div v-if="openModal">
      <TransactionModal
        v-on:cancel="() => closeTransactionModal(false)"
        v-on:create="() => closeTransactionModal(true)"
        v-on:update="() => closeTransactionModal(true)"
        v-on:delete="() => closeTransactionModal(true)"
        :edit-transaction="editTransaction"
        :start-date="date"
      />
    </div>

  </div>
</template>