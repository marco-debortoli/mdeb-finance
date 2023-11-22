<script setup lang="ts">
import MonthNavigation from '@/components/MonthNavigation.vue';
import { ref, onMounted } from 'vue';
import TransactionsList from '@/components/TransactionsList.vue';
import { useTransactionStore } from '@/stores/transaction';
import { useCategoryStore } from '@/stores/category';

const currentDate = ref(new Date());

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

// Transactions

const transactionStore = useTransactionStore();

// Categories

const categoryStore = useCategoryStore();

// On page mount
onMounted(() => {
  currentDate.value = new Date(
    currentDate.value.getFullYear(), currentDate.value.getMonth(), 1
  );

  transactionStore.retrieve(currentDate.value);
  categoryStore.retrieve();
})

</script>

<template>
  <div class="flex flex-col w-full p-2 h-screen">

    <!--Top header which includes the month and net worth component-->
    <header
      class="flex flex-grow-0 justify-center items-center mb-2 border rounded-md border-black/30"
    >
      <MonthNavigation
        :date="currentDate"
        @inc="() => incDate()"
        @dec="() => decDate()"
      />
    </header>

    <div class="flex flex-grow overflow-auto">
      <section class="flex flex-col xl:flex-row gap-2 w-full">
        <div class="xl:w-3/4 border rounded-md border-black/30">
          <TransactionsList
            :date="currentDate"
          />
        </div>
        <div class="xl:w-1/4 border rounded-md border-black/30">ACCOUNTS</div>
      </section>
    </div>
  </div>
</template>