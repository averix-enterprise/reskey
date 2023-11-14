<script setup lang="ts">
import {backend} from "../../../wailsjs/go/models";
import HotKey = backend.HotKey;
import {DeleteHotKey} from "../../../wailsjs/go/backend/App";

const props =defineProps<{hotkey: HotKey, refreshHotKeys: () => void}>()

const deleteHotKey = () => {
  DeleteHotKey(props.hotkey.id)
  props.refreshHotKeys()
}
</script>
<template>
    <div class="bg-gray-200 p-2 rounded">
        <div class="flex items-center gap-2 justify-between">
          Hotkey: {{ hotkey.key }}
          <input class="p-1 text-sm rounded" placeholder="Alt + SHIFT" />
        </div>
        <div>
          <input class="bg-transparent text-2xl" placeholder="1920 x 1080" />
          <button class="text-white bg-red-700 rounded-2xl px-2 py-1" @click="deleteHotKey">Delete</button>
        </div>
    </div>
</template>
