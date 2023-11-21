<script setup lang="ts">

defineEmits(
  ['cancel']
);

const props = defineProps<{
  edit: boolean
}>();

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
      <div class="relative border-t border-b border-t-blue-gray-100 border-b-blue-gray-100 p-4 text-base font-normal">
        CONTENT
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
        >
          Confirm
        </button>
      </div>
    </div>
  </div>
</template>