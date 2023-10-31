<template>
  <div class="camera">
    <div class="camera-video">
      <div style="position: relative;">
        <video :style="{ width: imgWidth + 'px', height: imgHeight + 'px' }" ref="camera" autoplay></video>
        <img v-if="lastImage" :src="lastImage.data" style="opacity: 0.4; position: absolute; top: 0; left: 0;" />
      </div>
    </div>

    <div class="camera-error-container" v-if="error">
      <div class="camera-error">
        <p>Wir benötigen Kamerazugriff.</p>
        <button @click="startCameraStream">Kamerazugriff anfordern</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

import { ref, onMounted, computed } from 'vue'
import { useCameraStore } from '../stores/camera-store'
import Image from '../types/Image';


defineProps({
  lastImage: {
    type: Image,
    // required: true
  }
})


const cameraStore = useCameraStore()


const imgWidth = 800
const imgHeight = 600

const camera = ref<HTMLVideoElement | null>(null)


// const lastImage = ref<Image | null>(imageStore.images.length > 0 ? imageStore.images[imageStore.images.length - 1] : null)

onMounted(() => {
  startCameraStream()
})

const startCameraStream = () => {
  if (camera && camera.value) {
    cameraStore.StartCameraStream(camera.value)
  }
}

const error = computed(() => {
  return cameraStore.HasError
})

</script>

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
