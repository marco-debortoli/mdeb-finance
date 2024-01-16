<script setup lang="ts">
import MonthNavigation from '@/components/MonthNavigation.vue';
import TransactionsList from '@/components/TransactionsList.vue';
import CategoryTotals from '@/components/CategoryTotals.vue';
import NetProfit from '@/components/NetProfit.vue';
import AccountList from '@/components/AccountList.vue';
import NetWorth from '@/components/NetWorth.vue';
import SettingsModal from '@/components/SettingsModal.vue';

import { ref, onMounted } from 'vue';
import { useTransactionStore } from '@/stores/transaction';
import { useCategoryStore } from '@/stores/category';
import { useAccountStore } from '@/stores/account';

const currentDate = ref(new Date());
const showSettingsModal = ref(false);

// Functions to deal with the date / month that is currently being displayed

function decDate() {
  const copyDate = new Date(currentDate.value);
  copyDate.setMonth(copyDate.getMonth() - 1);
  currentDate.value = copyDate;

  transactionStore.retrieve(currentDate.value);
}

function incDate() {
  const copyDate = new Date(currentDate.value);
  copyDate.setMonth(copyDate.getMonth() + 1);
  currentDate.value = copyDate;

  transactionStore.retrieve(currentDate.value);
}

const transactionStore = useTransactionStore();
const categoryStore = useCategoryStore();
const accountStore = useAccountStore();

// Save settings
function handleSave() {
  showSettingsModal.value = false;

  transactionStore.retrieve(currentDate.value);
  accountStore.retrieve();
  categoryStore.retrieve();
}

// On page mount
onMounted(() => {
  currentDate.value = new Date(
    currentDate.value.getFullYear(), currentDate.value.getMonth(), 1
  );

  transactionStore.retrieve(currentDate.value);
  accountStore.retrieve();
  categoryStore.retrieve();
})

</script>

<template>
  <div class="flex flex-col w-full p-2 h-screen">

    <!-- First row which includes the month selector, monthly review and net worth component-->
    <div class="flex flex-grow-0 mb-2">
      <div class="flex flex-col flex-grow gap-2 xl:flex-row">
        <div class="flex xl:w-1/2 gap-2">
          <div class="flex w-10 border rounded-md border-black/30 justify-center items-center">
            <i
              class="ti ti-settings text-xl text-black/30 hover:text-black hover:cursor-pointer"
              @click="() => { showSettingsModal = true }"
            ></i>
          </div>
          <div class="flex flex-grow border rounded-md border-black/30">
            <MonthNavigation
              :date="currentDate"
              @inc="() => incDate()"
              @dec="() => decDate()"
            />
          </div>
        </div>
        <div class="flex xl:w-1/4">
          <NetProfit
            :date="currentDate"
          />
        </div>
        <div class="flex xl:w-1/4">
          <NetWorth
            :date="currentDate"
          />
        </div>        
      </div>
    </div>

    <!-- Second row containing the transaction list, categories and accounts -->
    <div class="flex grow overflow-auto">
      <div class="flex flex-col xl:flex-row gap-2 w-full">
        <div class="xl:w-3/4 border rounded-md border-black/30">
          <TransactionsList
            :date="currentDate"
          />
        </div>
        <div class="flex flex-col gap-2 xl:w-1/4">
          <div class="flex h-1/2">
            <AccountList
              :date="currentDate"
            />
          </div>
          <div class="flex h-1/2">
            <CategoryTotals
              :date="currentDate"
            />
          </div>
        </div>
      </div>
    </div>

    <div v-if="showSettingsModal">
      <SettingsModal
        v-on:cancel="() => { showSettingsModal = false }"
        v-on:save="() => { handleSave() }"
      />
    </div>

  </div>
</template>