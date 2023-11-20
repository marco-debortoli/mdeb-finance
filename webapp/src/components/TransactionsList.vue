<script setup lang="ts">
import type { Transaction } from '@/types/transaction';

defineProps<{
  transactions: Transaction[]
  loading: boolean
}>();

</script>

<template>
  <div class="block py-6 px-10 w-full">
    <div class="flex justify-between">
      <div class="block">
        <h2 class="text-xl font-bold">Transactions</h2>
        <small class="text-xs">Recorded financial transactions from DATE to DATE</small>
      </div>
      <div class="block">
        BUTTON
      </div>
    </div>

    <table class="table-auto w-full mt-6" v-if="!loading">
      <thead class="uppercase border-b-2 border-black font-bold">
        <tr>
          <th>Date</th>
          <th>Category</th>
          <th>Amount</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        <template v-for="transaction in transactions" v-bind:key="transaction._id">
          <tr>
            <td>{{ transaction.date }}</td>
            <td>{{ transaction.category.name }}</td>
            <td>{{ transaction.amount }}</td>
            <td>{{ transaction.name }}</td>
          </tr>
        </template>
      </tbody>
    </table>

    <div v-else>
      LOADING
    </div>

  </div>
</template>