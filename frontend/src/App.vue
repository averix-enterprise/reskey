<script lang="ts" setup>
import {LogInfo, WindowSetSize} from "../wailsjs/runtime";
import Row from "./components/HotKey/HotKeyRow.vue"
import {onMounted, ref, watch} from "vue";
import {backend} from "../wailsjs/go/models";
import {AddHotKey, GetAllHotKeys} from "../wailsjs/go/backend/App";

const rowAmount = ref(1)

const hotKeys = ref<Array<backend.HotKey>>([]);

onMounted(async () => {
  hotKeys.value = await GetAllHotKeys()
})

const addHotKey = async () => {
  await AddHotKey()
  await refreshHotKeys()
}

const refreshHotKeys = async () => {
  hotKeys.value = await GetAllHotKeys()
}

watch(hotKeys, (newValue) => {
  const baseHeight = 250
  const windowHeight = baseHeight + (hotKeys.value.length * 50 )
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
        <button @click="addHotKey">+</button>
      </div>
      <div>
        <button @click="refreshHotKeys">Refresh</button>
      </div>
    </div>
    <hr class="my-3" />
    <div class="flex flex-col gap-2">
      <Row v-for="hotkey in hotKeys" :key="hotkey.id" :hotkey="hotkey" :refresh-hot-keys="refreshHotKeys" />
    </div>
  </div>
</template>
