<script lang="ts" setup>
import {WindowSetSize} from "../wailsjs/runtime/runtime"

function test() {
  WindowSetSize(500, 250)
}
</script>

<template>
  <button class="bg-blue-900"  @click="test()">Click me bitch</button>
</template>


