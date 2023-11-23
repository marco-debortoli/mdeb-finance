<script setup lang="ts">
import AccountCardGroup from '@/components/AccountCardGroup.vue'

defineProps<{
  date: Date
}>();

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
      <AccountCardGroup
        :date="date"
        :acc-type="'debit'"
      />

      <AccountCardGroup
        :date="date"
        :acc-type="'credit'"
      />

      <AccountCardGroup
        :date="date"
        :acc-type="'investment'"
      />

    </div>
  </div>
</template>