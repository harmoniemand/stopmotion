
  
<script setup lang="ts">

import { ref, PropType, computed } from 'vue'
import Image from '../types/Image';

const props = defineProps({
  images: {
    type: Array as PropType<Image[]>,
    required: true
  }
})


const fps = ref(10);
const currentImage = ref(0)

setInterval(() => {
  currentImage.value++;
  currentImage.value = currentImage.value % props.images.length;
}, 1000 / fps.value)




const imagesSorted = computed(() => {
  return props.images.sort((a, b) => {
    if (a.created_at > b.created_at) {
      return 1
    }
    if (a.created_at < b.created_at) {
      return -1
    }
    return 0
  })
})

</script>

<template>
  <div class="camera">
    <div class="camera-video">
      <div style="position: relative;">
        <img :src="imagesSorted[currentImage].data" />
      </div>
    </div>
  </div>
</template>
  
<style scoped>
.camera {
  width: 100%;
  height: 100%;
  background-color: #212121;

  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  position: relative;
}

.camera .camera-error-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;

  display: flex;
  align-items: center;
  justify-content: center;

}

.camera .camera-error-container .camera-error {

  width: 250px;
  height: 250px;

  border: solid 1px hsla(17, 100%, 37%, 1);
  background-color: hsla(17, 100%, 37%, 0.2);

  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
  