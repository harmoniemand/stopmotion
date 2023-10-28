<template>
    <Menubar>
        <template #start>
            <img alt="logo" src="./../assets/Logo.png" height="35" class="mr-2" />
            <Button icon="pi pi-plus" label="Neues Projekt" severity="danger" outlined @click="onButtonNewPressed" />

        </template>
        <!-- <template #item="{ label, item, props, root, hasSubmenu }">
            
        </template> -->
        <template #end>
            <Button icon="pi pi-play" v-if="!showPlayer" label="Video Starten" :disabled="!hasImages" outlined @click="onButtonPlayPressed" />
            <Button icon="pi pi-stop" v-if="showPlayer" label="Video Stoppen" outlined @click="onButtonPlayPressed" />

            <Button icon="pi pi-plus" label="Bild aufnehmen" @click="onButtonTakeImagePressed" style="margin-left: 1rem;" />


        </template>
    </Menubar>
</template>


<script setup lang="ts">

import Menubar from 'primevue/menubar';
import { ref, defineEmits, computed } from 'vue'
import { useImageStore } from '../stores/image-store';


const imageStore = useImageStore()

const emit = defineEmits(['onPlayButtonPressed', 'onTakeImageButtonPressed', 'onNewButtonPressed'])

const showPlayer = ref(false)

const hasImages = computed(() => {
    return imageStore.images.length > 0
})

const onButtonNewPressed = () => {
    console.debug("button new pressed")
    emit('onNewButtonPressed')
}

const onButtonPlayPressed = () => {
    showPlayer.value = !showPlayer.value
    console.debug("button play pressed")
    emit('onPlayButtonPressed')
}

const onButtonTakeImagePressed = () => {
    console.debug("button take image pressed")
    emit('onTakeImageButtonPressed')
}

</script>


<style scoped>
.toolbar {
    width: 100%;
    background-color: #222222;
    height: 46px;
    /* border: solid 1px blue; */
    display: flex;
}

.toolbar_item {
    color: hsla(160, 100%, 37%, 1);
    background-color: rgba(0, 0, 0, 0);
    border: solid 1px red;
    cursor: pointer;
    padding: 0px 20px;
}

.logo {
    border: none;
    background-color: rgba(0, 0, 0, 0);
    padding: 5px;
    margin-right: 25px;
}

.logo img {
    height: 36px;
    background-color: rgba(0, 0, 0, 0);
}
</style>
