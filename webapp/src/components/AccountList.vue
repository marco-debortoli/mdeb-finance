<script setup lang="ts">
import { useAccountStore } from '@/stores/account';
import AccountCard from './AccountCard.vue';

defineProps<{
  date: Date
}>();

const accountStore = useAccountStore()

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
        <AccountCard
          :date="date"
          :account="account"
        />
      </template>
    </div>
  </div>
</template>