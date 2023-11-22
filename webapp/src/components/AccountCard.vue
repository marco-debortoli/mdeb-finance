<script setup lang="ts">
import type { Account } from '@/types/account';
import { formatCurrency } from '@/tools/currency';
import dayjs from 'dayjs';
import { computed, ref } from 'vue';
import { useAccountStore } from '@/stores/account';

const props = defineProps<{
  account: Account
  date: Date
}>();

const edit = ref(false);
const newValue = ref(0);

const accountStore = useAccountStore();

// Value of the account this month
const accountValue = computed(() => {
  const currentDate = dayjs(props.date);
  return props.account.values.find((v) => currentDate.isSame(dayjs(v.date)));
});

const accountValueAmount = computed(() => {
  if (accountValue.value) return accountValue.value.value;
  return 0.00;
});

// Value of the account last month
const lastMonthValue = computed(() => {
  const lastMonthDate = dayjs(props.date).subtract(1, 'month');
  return props.account.values.find((v) => lastMonthDate.isSame(dayjs(v.date)));
});

// Difference between this month and last month (absolute)
const lastMonthAbsoluteDiff = computed(() => {
  let diff = 0;

  if (accountValue.value && lastMonthValue.value) {
    diff = accountValue.value.value - lastMonthValue.value.value;
  } else if (accountValue.value && !lastMonthValue.value) {
    diff = accountValue.value.value;
  } else if (!accountValue.value && lastMonthValue.value) {
    diff = -lastMonthValue.value.value;
  } else if (!accountValue.value && !lastMonthValue.value) {
    diff = 0;
  }

  return diff
});

// Date & Colour Formatting
function formatDate(date: string | undefined) {
  return dayjs(date).local().format("MMMM");
};

function formatColour(amount: number) {
  if (amount > 0) return 'text-green-700';
  if (amount == 0) return 'text-current';
  return 'text-red-700';
}

// Saving
function saveValue() {
  if (newValue.value <= 0) return;

  accountStore.addValue(
    props.account._id, props.date, newValue.value
  ).then(() => {
    accountStore.refresh();
    edit.value = false;
    newValue.value = 0;
  });
}

</script>

<template>
  <div class="flex border border-black rounded-lg py-2 px-4 shadow-md justify-between">
    <div class="flex flex-col justify-center">
      <span class="font-bold">{{ account.name }}</span>
      <span class="text-xs">{{ formatDate(accountValue?.date) }}</span>
    </div>

    <div class="flex flex-col text-right">
      <span
        class="hover:cursor-pointer hover:font-bold"
        v-if="!edit"
        @click="() => { newValue = accountValueAmount; edit = true }"
      >
        {{ formatCurrency(accountValueAmount) }}
      </span>

      <div class="flex items-center" v-else>
        <i
          class="ti ti-check mr-0.5 hover:cursor-pointer hover:text-green-700"
          @click="() => saveValue()"
        />

        <i
          class="ti ti-x mr-0.5 hover:cursor-pointer hover:text-red-700"
          @click="() => edit = false"
        />

        <input
          type="number" name="amount" id="amount"
          :min="0"
          :step="0.01"
          tabindex="2"
          v-model="newValue"
          class="rounded-md border border-black/30 w-28 pl-1"
          required
        >
      </div>

      <span :class="formatColour(lastMonthAbsoluteDiff)">{{ formatCurrency(lastMonthAbsoluteDiff, true) }}</span>

    </div>
  </div>
</template>