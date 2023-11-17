<script setup lang="ts">
import MonthNavigation from '@/components/MonthNavigation.vue';
import { ref, onMounted } from 'vue';

const currentDate = ref(new Date());

// Functions to deal with the date / month that is currently being displayed

function decDate() {
  const copyDate = new Date(currentDate.value);
  copyDate.setMonth(copyDate.getMonth() - 1);
  currentDate.value = copyDate;
}

function incDate() {
  const copyDate = new Date(currentDate.value);
  copyDate.setMonth(copyDate.getMonth() + 1);
  currentDate.value = copyDate;
}

onMounted(() => {
  currentDate.value = new Date(
    currentDate.value.getFullYear(), currentDate.value.getMonth(), 1
  );
});

</script>

<template>
  <div class="flex flex-col w-full h-full">

    <!--Top header which includes the month and net worth component-->
    <header
      class="flex justify-center items-center h-12 mb-2 border rounded-md border-black/30"
    >
      <MonthNavigation
        :date="currentDate"
        @inc="() => incDate()"
        @dec="() => decDate()"
      />
    </header>

    <section class="grid grid-cols-1 xl:grid-cols-4 h-screen gap-2">
      <div class="xl:col-span-3 border rounded-md border-black/30">{{ currentDate }}</div>
      <div class="xl:col-span-1 border rounded-md border-black/30">ACCOUNTS</div>
    </section>
  </div>
</template>