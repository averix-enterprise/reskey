<script lang="ts" setup>
import { useQueryClient } from '@tanstack/vue-query'
import { AddHotKey } from "../wailsjs/go/backend/App";
import useQuery from './composables/useQuery'
import Row from "./components/HotKey/HotKeyRow.vue"
import {useKeyModifier, useMagicKeys} from "@vueuse/core";
import {watch} from "vue";
import {LogInfo} from "../wailsjs/runtime";

const queryClient = useQueryClient()

const { hotKeys: hotKeysQuery } = useQuery()
const { data: hotKeys, isLoading: areHotKeysLoading } = hotKeysQuery.all()

const addHotKey = async () => {
  await AddHotKey()
  queryClient.invalidateQueries({ queryKey: ['hotKeys'] })
}

const { current } = useMagicKeys()

watch(current, (e) => {
  LogInfo(Array.from(e.keys()).join(" "))
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
        Averix.Tech
      </div>
    </div>
    <hr class="my-3" />
    <div v-if="!areHotKeysLoading" class="flex flex-col gap-2">
      <Row v-for="hotkey in hotKeys" :key="hotkey.id" :hotkey="hotkey" :refresh-hot-keys="() => {}" />
    </div>
  </div>
</template>
