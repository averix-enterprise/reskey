<script lang="ts" setup>
import { WindowSetSize } from "../wailsjs/runtime/runtime";
import Row from "./components/HotKey/Row.vue"
import { ref, watch } from "vue";

const rowAmount = ref(1)

const addRow = () => {
  rowAmount.value++
}

watch(rowAmount, (newValue) => {
  if (newValue >= 4) {
    WindowSetSize(600, 500)
  }
})
</script>

<template>
  <div class="p-4">
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-1">
        <h1 class="text-lg font-bold text-center">
          Res<span class="text-purple-800">Key</span>
        </h1>
        <button @click="addRow">+</button>
      </div>
      <div>
        <button>Refresh</button>
      </div>
    </div>
    <hr class="my-3" />
    <div class="flex flex-col gap-2">
      <Row v-for="row in rowAmount" />
    </div>
  </div>
</template>


