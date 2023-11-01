
  
<script setup lang="ts">

import { ref, PropType, computed } from 'vue'
import Image from '../types/Image';

const props = defineProps({
  images: {
    type: Array as PropType<Image[]>,
    required: true
  }
})


const fps = ref(30);
const currentImage = ref(0)

setInterval(() => {
  currentImage.value++;
  currentImage.value = currentImage.value % props.images.length;
}, 1000 / fps.value)




const imagesSorted = computed(() => {
  
  const imgs = JSON.parse(JSON.stringify(props.images))
  
  return imgs.sort((a: Image, b: Image) => {
    if (a.CreatedAt > b.CreatedAt) {
      return 1
    }
    if (a.CreatedAt < b.CreatedAt) {
      return -1
    }
    return 0
  })
})

</script>


<template>
  <div class="player">
    <div class="player-video">
      <div class="player-container" style="position: relative;">
        <img class="player-img" v-for="(i, index) in imagesSorted" :src="i.Url"
          :class="{ visible: index == currentImage }" />
      </div>
    </div>
  </div>
</template>
  

<style scoped>
.player {
  width: 100%;
  height: 100%;
  background-color: #212121;

  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  position: relative;
}

.player .player-error-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;

  display: flex;
  align-items: center;
  justify-content: center;

}

.player .player-error-container .player-error {

  width: 250px;
  height: 250px;

  border: solid 1px hsla(17, 100%, 37%, 1);
  background-color: hsla(17, 100%, 37%, 0.2);

  display: flex;
  align-items: center;
  justify-content: center;
}

.player .player-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.player .player-container .player-img {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  visibility: hidden;
}

.player .player-container .player-img.visible {
  visibility: visible;
}
</style>
  