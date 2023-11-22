<script setup lang="ts">
import { computed, ref } from 'vue';

const leftHover = ref(false);
const rightHover = ref(false);

defineEmits(['dec', 'inc']);

const props = defineProps<{
  date: Date
}>();

const displayDate = computed(() => {
  const monthName = Intl.DateTimeFormat('en', {month: 'long'}).format(props.date);
  const year = props.date.getFullYear();
  return monthName + " " + year;
});

</script>

<template>
  <div class="flex flex-grow align-middle justify-center items-center">
    <i
      class="ti text-2xl mr-2 hover:cursor-pointer"
      :class="{'ti-arrow-big-left': !leftHover, 'ti-arrow-big-left-filled': leftHover}"
      v-on:mouseenter="() => leftHover=true"
      v-on:mouseleave="() => leftHover=false"
      v-on:click="() => $emit('dec')"
    ></i>
    <h1 class="text-3xl uppercase font-bold">
      {{ displayDate }}
    </h1>
    <i
      class="ti text-2xl ml-2 hover:cursor-pointer"
      :class="{'ti-arrow-big-right': !rightHover, 'ti-arrow-big-right-filled': rightHover}"
      v-on:mouseenter="() => rightHover=true"
      v-on:mouseleave="() => rightHover=false"
      v-on:click="() => $emit('inc')"
    ></i>
  </div>
</template>