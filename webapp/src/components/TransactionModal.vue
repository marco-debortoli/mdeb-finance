<script setup lang="ts">
import { useCategoryStore } from '@/stores/category';
import { computed, ref } from 'vue';
import { useTransactionStore } from '@/stores/transaction';
import type { Transaction } from '@/types/transaction';

const emit = defineEmits(
  ['cancel', 'create', 'update', 'delete']
);

const props = defineProps<{
  startDate: Date
  editTransaction?: Transaction
}>();

const saving = ref(false);
const error = ref("");

const categoryStore = useCategoryStore();
const transactionStore = useTransactionStore();

let form = ref(
  {
    date: props.editTransaction ?
      new Date(props.editTransaction.date).toISOString().split("T")[0]
      : new Date(props.startDate).toISOString().split('T')[0],
    category: props.editTransaction?.category._id || "",
    amount: props.editTransaction?.amount || 0.0,
    name: props.editTransaction?.name || ""
  }
);

const maxDate = computed(() => {
  const copyDate = new Date(props.startDate);
  copyDate.setMonth(copyDate.getMonth() + 1);
  copyDate.setDate(copyDate.getDate() - 1);

  return copyDate
});

function addOrEditTransaction() {
  // Validation Required
  error.value = "";
  let valid = true;

  if (form.value.date == undefined) {
    error.value = "Date must be selected";
    valid = false;
  }

  if (form.value.category == "") {
    error.value = "Category must be selected";
    valid = false;
  }

  if (form.value.amount <= 0) {
    error.value = "Amount must be greater than or equal to zero";
    valid = false;
  }

  if (form.value.name == undefined || form.value.name == "") {
    valid = false;
    error.value = "Name must be defined and not a blank string";
  }

  if (!valid) return;

  saving.value = true;

  if (props.editTransaction) {
    transactionStore.update(
      props.editTransaction._id,
      form.value.date,
      form.value.amount,
      form.value.category,
      form.value.name      
    ).then(() => {
      saving.value = false;
      emit('update');
    });

  } else {
    transactionStore.create(
      form.value.date,
      form.value.amount,
      form.value.category,
      form.value.name
    ).then(() => {
      saving.value = false;
      emit('create');
    });
  }
}

function deleteTransaction() {
  if (props.editTransaction) {
    transactionStore.delete(
      props.editTransaction?._id
    ).then(() => {
      saving.value = false;
      emit('delete');
    });
  }
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
        relative m-4 w-4/5 xl:w-2/5 rounded-lg bg-offwhite border-black/30 border
      "
    >
      <div class="flex shrink-0 items-center p-4 text-2xl font-bold uppercase">
        <h2># {{ props.editTransaction ? 'Edit' : 'Add' }} Transaction</h2>
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
                v-model="form.date"
                tabindex="1"
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
                tabindex="2"
                v-model="form.amount"
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
                v-model="form.category"
                tabindex="3"
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
                v-model="form.name"
                tabindex="4"
                class="block w-full rounded-md border border-black/30 py-1 px-4"
                required
              >
            </div>
          </div>

          <div v-if="error">
            <h3 class="text-red-500 font-normal">Error: {{ error }}</h3>
          </div>
        </div>
      </div>
      <div :class="{ 'flex justify-between': props.editTransaction }">
        <div class="flex items-center text-red-500 ml-4 font-medium" v-if="editTransaction">
          <button
            class="text-xs uppercase hover:underline"
            @click="() => deleteTransaction()"
          >
            Delete
          </button>
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
            @click="() => addOrEditTransaction()"
            tabindex="5"
            :disabled="saving"
          >
            {{ props.editTransaction ? 'Update' : 'Create' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
