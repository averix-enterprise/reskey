<script lang="ts" setup>
import { WindowSetSize } from "../wailsjs/runtime/runtime";
import Row from "./components/HotKey/HotKeyRow.vue"
import { ref, watch } from "vue";

const rowAmount = ref(1)

const addRow = () => {
  rowAmount.value++
}

watch(rowAmount, (newValue) => {
  const baseHeight = 250
  const windowHeight = baseHeight + (rowAmount.value * 50 )
  if (windowHeight <= 800) {
      WindowSetSize(600, windowHeight)
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
        <button @click="rowAmount = 1">Refresh</button>
      </div>
    </div>
    <hr class="my-3" />
    <div class="flex flex-col gap-2">
      <Row v-for="row in rowAmount" />
    </div>
  </div>
</template>
