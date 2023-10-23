<template>
    <div class="reel">
        <div class="reel-header">
            <span>{{ image_count }} Bilder</span>
        </div>
        <div class="reel-image" v-for="(image) in images_reversed" :key="image.Id" @click="selectedId = selectedId == image.Id ? '' : image.Id" :class="{ selected: image.Id == selectedId}">
            <img :src="image.BASE64" />
            <div class="reel-image-toolbox">
                <button class="btn" @click="deleteImage(image)">Löschen</button>
            </div>
        </div>
    </div>
</template>



<script setup lang="ts">

import { ref, computed, watch } from 'vue'
import Image from '../types/Image';

import ImageService from '../services/image.service';

const imageService: ImageService = ImageService.getInstance()

const images = imageService.Images

// const props = defineProps({ images: { type: Array as () => Image[], required: true } })

const selectedId = ref<string | null>(null)


const deleteImage = (image: Image) => {
    imageService.Delete(image)
}


watch(images, (newImages) => {
    console.log('images changed', newImages)
})

const images_reversed = computed(() => {
    return images.value.slice().reverse();
})

const image_count = computed(() => {
    return images.value.length;
})

</script>



<style scoped>
.reel {
    width: 100%;
    height: calc(100vh - 46px);
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

    background-color: rgba(0,0,0,0.5);
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
