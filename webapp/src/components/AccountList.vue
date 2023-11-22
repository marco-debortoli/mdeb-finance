<script setup lang="ts">
import { useAccountStore } from '@/stores/account';
import { computed } from 'vue';
import AccountCard from './AccountCard.vue';
import dayjs from 'dayjs';

const props = defineProps<{
  date: Date
}>();

const accountStore = useAccountStore()

const activeAccounts = computed(() => {
  const currentDate = dayjs(props.date);
  return accountStore.accounts.filter((acc) => {
    const accStartDate = dayjs(acc.start_date).startOf('month');

    if (accStartDate.isAfter(currentDate)) return false;
    if (!acc.end_date) return true;

    const accEndDate = dayjs(acc.end_date).startOf('month');
    console.log(accEndDate, currentDate);
    return accEndDate.isSame(currentDate) || accEndDate.isAfter(currentDate);
  });
});

const debitAccounts = computed(() => {
  return activeAccounts.value.filter((acc) => acc.type == "debit");
});

const creditAccounts = computed(() => {
  return activeAccounts.value.filter((acc) => acc.type == "credit");
});

const investAccounts = computed(() => {
  return activeAccounts.value.filter((acc) => acc.type == "investment");
});

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
      <!-- Debit Accounts -->
      <template v-if="debitAccounts.length > 0">
        <div class="uppercase underline flex justify-center">
          Debit
        </div>
        <template v-for="account in debitAccounts" :key="account._id">
          <AccountCard
            :date="date"
            :account="account"
          />
        </template>
      </template>

      <!-- Credit Accounts -->
      <template v-if="creditAccounts.length > 0">
        <div class="uppercase underline flex justify-center">
          Credit
        </div>
        <template v-for="account in creditAccounts" :key="account._id">
          <AccountCard
            :date="date"
            :account="account"
          />
        </template>
      </template>

      <!-- Investment Accounts -->
      <template v-if="investAccounts.length > 0">
        <div class="uppercase underline flex justify-center">
          Investment
        </div>
        <template v-for="account in investAccounts" :key="account._id">
          <AccountCard
            :date="date"
            :account="account"
          />
        </template>
      </template>

    </div>
  </div>
</template>