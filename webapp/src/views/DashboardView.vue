<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

const selectedDate = ref(new Date());

const displayDate = computed(() => {
  const monthName = Intl.DateTimeFormat('en', {month: 'long'}).format(selectedDate.value);
  const year = selectedDate.value.getFullYear();

  return monthName + " " + year;
})

onMounted(() => {
  selectedDate.value = new Date(
    selectedDate.value.getFullYear(), selectedDate.value.getMonth(), 1
  );
})

</script>

<template>
  <div class="flex flex-col w-full h-full">

    <!--Top header which includes the month and net worth component-->
    <header class="flex justify-center items-center h-12 mb-2 bg-slate-400">
      <div class="text-3xl uppercase font-bold">
        {{ displayDate }}
      </div>
    </header>

    <section class="grid grid-cols-4 h-screen gap-2">
      <div class="col-span-3 bg-green-300">TRANSACTIONS</div>
      <div class="col-span-1 bg-blue-300">ACCOUNTS</div>
    </section>
  </div>
</template>