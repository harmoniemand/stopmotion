<template>
    <div class="reel">
        <div class="reel-header">
            <span>{{ image_count }} Bilder</span>
        </div>
        <div class="reel-image" v-for="(image) in images_reversed" :key="image.Id"
            @click="selectedId = selectedId == image.Id ? '' : image.Id" :class="{ selected: image.Id == selectedId }">
            <img :src="image.BASE64" />
            <div class="reel-image-toolbox">
                <Button icon="pi pi-trash" label="Bild Löschen" @click="deleteImage(image)" />
            </div>
        </div>
    </div>
</template>



<script setup lang="ts">

import { ref, computed } from 'vue'
import Image from '../types/Image';

import { useImageStore } from '../stores/image-store';

import Button from 'primevue/button';


const imageStore = useImageStore()

const selectedId = ref<string | null>(null)


const deleteImage = (image: Image) => {
    imageStore.removeImage(image)
}

const images_reversed = computed(() => {
    return imageStore.reversedImages
})

const image_count = computed(() => {
    return imageStore.images.length
})

</script>



<style scoped>
.reel {
    width: 100%;
    height: 100%;
    overflow: scroll;
    margin-left: 10px;
    margin-right: 10px;
}

.reel .reel-image {
    position: relative;
    margin: 0px;
    margin-bottom: 20px;
    width: calc(100% - 20px);
}

.reel .reel-image img {
    width: 100%;
}

.reel .reel-image .reel-image-toolbox {
    position: absolute;
    bottom: 0;
    left: 0;
    padding: 10px;
    display: none;
    width: 100%;

    background-color: rgba(0, 0, 0, 0.5);
}

.reel .reel-image.selected .reel-image-toolbox {
    display: block;
}

.reel .reel-image .reel-image-toolbox .btn {
    border: solid 1px hsla(160, 100%, 37%, 1);
    height: 42px;
    padding: 0px 45px;
    font-size: 20px;
}
</style>
