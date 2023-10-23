<template>
  <div class="app-container">
    <Navigation />

    <div class="content-container" v-if="state?.SessionId != ''">
      <div class="content-container-item left">
        <Camera v-if="showCamera" />
        <Player v-if="showPlayer" />

        <div class="camera-controls">
          <button class="btn" @click="onButtonTakeImagePressed">Bild aufnehmen</button>
          <button class="btn" @click="onButtonPlayPressed">
            <span v-if="!showPlayer">Video Abspielen</span>
            <span v-if="showPlayer">Video Stoppen</span>

            </button>
        </div>
      </div>
      <div class="content-container-item right">
        <Reel :images="images" />
      </div>
    </div>
  </div>

  <ModalNewProjekt v-if="state?.SessionId == ''" @ready="NewProjectReady" />
</template>


<script setup lang="ts">

import AppState from './types/AppState';
import Image from './types/Image';

import CameraService from './services/camera.service'
import ImageService from './services/image.service'

import { ref, watch } from 'vue'

import Navigation from './components/Navigation.vue'
import Camera from './components/Camera.vue'
import Reel from './components/Reel.vue'
import ModalNewProjekt from './components/ModalNewProjekt.vue';
import Player from './components/Player.vue';

const images = ref<Image[]>([])
const showCamera = ref(true)
const showPlayer = ref(false)

const cameraService: CameraService = CameraService.getInstance()
const imageService: ImageService = ImageService.getInstance()

const state = ref<AppState>()
const storageState = localStorage.getItem('vue_app_state')
if (storageState) {
  state.value = JSON.parse(storageState)
} else {
  state.value = new AppState(null)
}

watch(state, (newState) => {
  console.log('state changed', newState)
  localStorage.setItem('vue_app_state', JSON.stringify(newState))
})

watch(images, (newImages) => {
  console.log('images changed', newImages)
})




const NewProjectReady = (e: any) => {
  // state.value.SessionId = '123'
  console.log('NewProjectReady', e)

  state.value = new AppState({
    sessionId: '123',
    name: e.name,
    description: e.description,
    authors: e.authors
  })
}

const onButtonPlayPressed = () => {
  console.log('Play')
  showCamera.value = !showCamera.value
  showPlayer.value = !showPlayer.value
}

const onButtonTakeImagePressed = () => {
  console.debug("button take image pressed")
  cameraService.TakePhoto()
}




</script>



<style scoped>
.app-container {
  width: 100vw;
  height: 100vh;

  display: flex;
  flex-direction: column;
  /* touch-action: none; */
}

.content-container {
  width: 100%;
  height: 100%;

  display: flex;
}

.content-container .content-container-item {
  height: 100%;
  width: 100%;
}

.content-container .content-container-item.left {
  width: 80%;
  height: 100%;

  display: flex;
  flex-direction: column;
}

.content-container .content-container-item.right {
  width: 20%;
  height: 100%;
}


.camera-controls {
  padding: 1rem;
  display: flex;
  justify-content: center;
  align-items: center;
}

.camera-controls .btn {
    border: solid 1px hsla(160, 100%, 37%, 1);
  /* border: solid 1px hsla(17, 100%, 37%, 1); */
  background-color: hsla(160, 100%, 37%, 0.2);
  color: hsla(160, 100%, 37%, 1);

  height: 42px;
  padding: 0px 45px;
  font-size: 20px;
  margin: 0px 10px;
}


.overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;

  background-color: rgba(0, 0, 0, 0.5);

  display: flex;
  justify-content: center;
  align-items: center;
}

.overlay .overlay-inner {
  height: 50vh;
  width: 50vw;

  background-color: white;
}
</style>
