<script setup lang="ts">
import { useCategoryStore } from '@/stores/category';
import { computed, ref } from 'vue';
import { useTransactionStore } from '@/stores/transaction';

const emit = defineEmits(
  ['cancel', 'create']
);

const props = defineProps<{
  edit: boolean
  startDate: Date
}>();

const adding = ref(false);

const categoryStore = useCategoryStore();
const transactionStore = useTransactionStore();

let createForm = ref(
  {
    date: new Date(props.startDate).toISOString().split('T')[0],
    category: "",
    amount: 0.00,
    name: ""
  }
);

const maxDate = computed(() => {
  const copyDate = new Date(props.startDate);
  copyDate.setMonth(copyDate.getMonth() + 1);
  copyDate.setDate(copyDate.getDate() - 1);

  return copyDate
});

function addTransaction() {
  adding.value = true;
  
  transactionStore.create(
    createForm.value.date,
    createForm.value.amount,
    createForm.value.category,
    createForm.value.name
  ).then(() => {
    adding.value = false;
    emit('create');
  })
}

</script>

<template>
  <div
    class="
      absolute inset-0 z-[999] grid h-full w-full
      place-items-center backdrop-blur-lg border-0 rounded-md
      transition-all duration-500
    "
  >
    <div
      data-dialog="dialog"
      class="
        relative m-4 w-2/5 min-w-[40%] max-w-[40%] rounded-lg bg-offwhite border-black/30 border
      "
    >
      <div class="flex shrink-0 items-center p-4 text-2xl font-bold uppercase">
        <h2># {{ props.edit ? 'Edit' : 'Add' }} Transaction</h2>
      </div>
      <div class="relative border-t border-b border-t-black/30 border-b-black/30 p-4 text-base font-normal">

        <!-- This is our form -->
        <div class="grid grid-cols-1 gap-y-4">

          <!-- Transaction Date -->
          <div>
            <label for="date" class="block text-sm font-normal">Transaction Date</label>
            <div class="mt-2">
              <input
                type="date" name="date" id="date"
                :min="startDate.toISOString().split('T')[0]"
                :max="maxDate.toISOString().split('T')[0]"
                v-model="createForm.date"
                class="block w-full rounded-md border border-black/30 py-1 px-4"
                required
              >
            </div>
          </div>

          <!-- Amount -->
          <div>
            <label for="amount" class="block text-sm font-normal">Amount</label>
            <div class="mt-2">
              <input
                type="number" name="amount" id="amount"
                :min="0"
                v-model="createForm.amount"
                class="block w-full rounded-md border border-black/30 py-1 px-4"
                required
              >
            </div>
          </div>

          <!-- Category -->
          <div>
            <label for="category" class="block text-sm font-normal">Category</label>
            <div class="mt-2">
              <select
                id="category" name="category"
                v-model="createForm.category"
                class="block w-full rounded-md border border-black/30 py-1 px-4 bg-white"
              >
                <template v-for="category in categoryStore.categories" :key="category._id">
                  <option :value="category._id">{{ category.name }}</option>
                </template>
              </select>
            </div>
          </div>

          <!-- Short Description / Name -->
          <div>
            <label for="name" class="block text-sm font-normal">Description</label>
            <div class="mt-2">
              <input
                type="text" name="name" id="name"
                placeholder="Short Description"
                v-model="createForm.name"
                class="block w-full rounded-md border border-black/30 py-1 px-4"
                required
              >
            </div>
          </div>
        </div>
      </div>
      <div class="flex shrink-0 flex-wrap items-center justify-end p-4 text-blue-gray-500">
        <button
          class="mr-6 text-xs uppercase hover:underline"
          @click="() => $emit('cancel')"
        >
          Cancel
        </button>
        <button
          data-ripple-light="true"
          data-dialog-close="true"
          class="uppercase border px-4 py-1 rounded-lg border-black/50 hover:border-black align-middle"
          @click="() => addTransaction()"
          :disabled="adding"
        >
          {{ props.edit ? 'Update' : 'Create' }}
        </button>
      </div>
    </div>
  </div>
</template>