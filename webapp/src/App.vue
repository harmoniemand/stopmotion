<template>
  <ConfirmDialog></ConfirmDialog>


  <div class="grid nested-grid h-screen" @dblclick="stopDblClick">
    <div class="col-12">
      <Navigation @onPlayButtonPressed="onButtonPlayPressed" @onTakeImageButtonPressed="onButtonTakeImagePressed"
        @onNewButtonPressed="onNewButtonPressed" />
    </div>
    <div class="col-9 flex h-full" style="max-height: calc(100% - 84px) !important;">
      <Camera v-if="showCamera && projectStore.Project.value" :lastImage="projectStore.Project.value.images[0]" />
      <Player v-if="showPlayer && projectStore.Project.value" :images="projectStore.Project.value.images" />
    </div>
    <div class="col-3 h-full" style="max-height: calc(100% - 84px) !important;">
      <Reel v-if="projectStore.Project.value" :images="projectStore.Project.value.images" />
    </div>
  </div>

  <ModalNewProjekt v-if="projectStore.Project.value == null" @ready="NewProjectReady" @dblclick="stopDblClick" />
</template>


<script setup lang="ts">

import Project from './types/Project';

import { useCameraStore } from './stores/camera-store';
import ProjectStore from './stores/project-store';

import { ref } from 'vue'

import ConfirmDialog from 'primevue/confirmdialog';
import { useConfirm } from "primevue/useconfirm";

import Navigation from './components/Navigation.vue'
import Camera from './components/Camera.vue'
import Reel from './components/Reel.vue'
import ModalNewProjekt from './components/ModalNewProjekt.vue';
import Player from './components/Player.vue';
import { Guid } from 'guid-typescript';

const showCamera = ref(true)
const showPlayer = ref(false)

const cameraStore = useCameraStore()
const projectStore = ProjectStore.getInstance()

projectStore.LoadProject()


const NewProjectReady = (e: any) => {
  // state.value.SessionId = '123'
  console.log('NewProjectReady', e)

  projectStore.CreateProject(new Project({
    id: Guid.create().toString(),
    name: e.name,
    description: e.description,
    authors: e.authors,
    images: [],
  }))
}

const leadingZero = (num: number, size: number) => {
  let s = num + "";
  while (s.length < size) s = "0" + s;
  return s;
}

const getDateString = () => {
  let d = new Date()
  let dStr = d.getFullYear() + '-'
  dStr += leadingZero((d.getMonth() + 1), 2) + '-'
  dStr += leadingZero(d.getDate(), 2) + ' '
  dStr += leadingZero(d.getHours(), 2) + ':'
  dStr += leadingZero(d.getMinutes(), 2) + ':'
  dStr += leadingZero(d.getSeconds(), 2)

  return dStr
}

const onButtonPlayPressed = () => {
  console.log('Play')
  showCamera.value = !showCamera.value
  showPlayer.value = !showPlayer.value
}

const onButtonTakeImagePressed = () => {
  console.debug("button take image pressed")
  const image = cameraStore.TakePhoto()
  image.created_at = getDateString()
  projectStore.AddImage(image)

}

const onNewButtonPressed = () => {
  console.debug("button new pressed")
  confirm1()
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
