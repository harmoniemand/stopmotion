<template>
  <ConfirmDialog></ConfirmDialog>


  <div class="app-container" @dblclick="stopDblClick">
    <div class="header">
      <Navigation @onPlayButtonPressed="onButtonPlayPressed" @onTakeImageButtonPressed="onButtonTakeImagePressed"
        @onNewButtonPressed="onNewButtonPressed" />
    </div>
    <div class="content">
      <div class="left" style="max-height: calc(100% - 84px) !important;">
        <Camera v-if="showCamera && projectStore.Project.value" :lastImage="projectStore.Project.value.images[0]" />
        <Player v-if="showPlayer && projectStore.Project.value" :images="projectStore.Project.value.images" />
      </div>
      <div class="right" style="max-height: calc(100% - 84px) !important;">
        <Reel v-if="projectStore.Project.value" :images="projectStore.Project.value.images" @delete="onButtonPressedDeleteImage" />
      </div>
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
import Image from './types/Image';

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

const onButtonPlayPressed = () => {
  console.log('Play')
  showCamera.value = !showCamera.value
  showPlayer.value = !showPlayer.value
}

const onButtonTakeImagePressed = () => {
  console.debug("button take image pressed")
  const imageBase64 = cameraStore.TakePhoto()
  projectStore.AddImage(imageBase64)
}

const onNewButtonPressed = () => {
  console.debug("button new pressed")
  confirm1()
}

const onButtonPressedDeleteImage = (image: Image) => {
  console.debug("button delete pressed")
  projectStore.RemoveImage(image)
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

.content {
  width: 100%;
  height: 100%;

  display: flex;
}


.content .left {
  width: 100%;
  height: 100%;

  display: flex;
  flex-direction: column;
}

.content .right {
  width: 300px;
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
