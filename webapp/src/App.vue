<template>
  <ConfirmDialog></ConfirmDialog>


  <div class="grid nested-grid h-screen" @dblclick="stopDblClick">
    <div class="col-12">
      <Navigation @onPlayButtonPressed="onButtonPlayPressed" @onTakeImageButtonPressed="onButtonTakeImagePressed"  @onNewButtonPressed="onNewButtonPressed"/>
    </div>
    <div class="col-9 flex h-full" style="max-height: calc(100% - 84px) !important;">
          <Camera v-if="showCamera" />
          <Player v-if="showPlayer" />
    </div>
    <div class="col-3 h-full" style="max-height: calc(100% - 84px) !important;">
      <Reel :images="images" />
    </div>
  </div>

  <ModalNewProjekt v-if="state?.SessionId == ''" @ready="NewProjectReady"  @dblclick="stopDblClick" />
</template>


<script setup lang="ts">

import AppState from './types/AppState';
import Image from './types/Image';

import { useCameraStore } from './stores/camera-store';
import { useImageStore } from './stores/image-store';

import { ref, watch } from 'vue'

import ConfirmDialog from 'primevue/confirmdialog';
import { useConfirm } from "primevue/useconfirm";

import Navigation from './components/Navigation.vue'
import Camera from './components/Camera.vue'
import Reel from './components/Reel.vue'
import ModalNewProjekt from './components/ModalNewProjekt.vue';
import Player from './components/Player.vue';

const images = ref<Image[]>([])
const showCamera = ref(true)
const showPlayer = ref(false)

const cameraStore = useCameraStore()
const imageStore = useImageStore()


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
  const image = cameraStore.TakePhoto()
  imageStore.addImage(image)

}

const onNewButtonPressed = () => {
  console.debug("button new pressed")

  confirm1()
  // state.value = new AppState(null)
  // showCamera.value = true
  // showPlayer.value = false
}

const confirm = useConfirm();

const confirm1 = () => {
    confirm.require({
        message: 'Wenn du ein neues Projekt anlegst, wird dein aktuelles Projekt gelöscht. Bist du sicher?',
        header: 'Neues Projekt',
        icon: 'pi pi-exclamation-triangle',
        accept: () => {
          console.log('accept')
          localStorage.clear()
          window.location.reload()          
        },
        reject: () => {
          console.log('reject')
        }
    });
};

const stopDblClick = (e: any) => {
  e.preventDefault()
  e.stopPropagation()
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

.camera-controls Button {
  margin: 0 10px;
}
</style>
