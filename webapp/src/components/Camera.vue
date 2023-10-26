<template>
  <div class="camera">
    <div class="camera-video">
      <div style="position: relative;">
        <video :style="{ width: imgWidth + 'px', height: imgHeight + 'px' }" ref="camera" autoplay></video>
        <img v-if="lastImage" :src="lastImage.BASE64" style="opacity: 0.4; position: absolute; top: 0; left: 0;" />
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

import Image from '../types/Image';
import { ref, onMounted } from 'vue'
import CameraService from './../services/camera.service'
import ImageService from './../services/image.service'
import CameraEvent from '../Events/CameraEvent';


const cameraService = CameraService.getInstance()
const imageService = ImageService.getInstance()

const imgWidth = 800
const imgHeight = 600

const error = ref(false)
const camera = ref<HTMLVideoElement | null>(null)


const lastImage = ref<Image | null>(imageService.Images.value.length > 0 ? imageService.Images.value[imageService.Images.value.length - 1] : null)

onMounted(() => {
  startCameraStream()
})

const startCameraStream = () => {
  if (camera && camera.value) {
    cameraService.StartCameraStream(camera.value)
  }
}

const onCameraImage = (e: any) => {
  console.log('onCameraImage', e)
  const img = (e as CameraEvent).Image;
  lastImage.value = img
}
cameraService.addEventListener('onNewImage', onCameraImage)

const onCameraError = () => {
  error.value = cameraService.HasError
}

cameraService.addEventListener('onCameraError', onCameraError)


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
