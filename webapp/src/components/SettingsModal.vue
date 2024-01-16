<script setup lang="ts">
import { ref, defineEmits } from 'vue';

const apiUrl = ref(localStorage.getItem('apiURL'));

const emit = defineEmits(['cancel', 'save']);

function save() {
  if (apiUrl.value) {
    localStorage.setItem('apiURL', apiUrl.value);
  }

  emit('save');
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
        relative m-4 w-4/5 xl:w-1/4 rounded-lg bg-offwhite border-black/30 border
      "
    >
      <div class="flex shrink-0 items-center p-4 text-2xl font-bold uppercase">
        <h2># Settings</h2>
      </div>
      <div class="relative border-t border-b border-t-black/30 border-b-black/30 p-4 text-base font-normal">

        <!-- This is our form -->
        <div class="grid grid-cols-1 gap-y-4">

          <!-- API URL -->
          <div>
            <label for="api" class="block text-sm font-normal">API URL</label>
            <div class="mt-2">
              <input
                type="text" name="api" id="api"
                tabindex="1"
                v-model="apiUrl"
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
          tabindex="5"
          @click="() => { save() }"
        >
          Save
        </button>
      </div>
    </div>
  </div>
</template>
