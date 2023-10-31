<script setup lang="ts">

import { ref, computed, PropType } from 'vue'
import Image from '../types/Image';
import Button from 'primevue/button';

const props = defineProps({
    images: {
        type: Array as PropType<Image[]>,
        required: true
    }
})

const $emit = defineEmits(['delete'])

const selectedId = ref<string | null>(null)


const imagesSorted = computed(() => {
    return props.images.sort((a, b) => {
        if (a.created_at > b.created_at) {
            return -1
        }
        if (a.created_at < b.created_at) {
            return 1
        }
        return 0
    })
})

const image_count = computed(() => {
    return props.images.length
})

const deleteImage = (image: Image) => {
    $emit('delete', image)
}
</script>


<template>
    <div class="reel">
        <div class="reel-header">
            <span>{{ image_count }} Bilder</span>
        </div>
        <div class="reel-image" v-for="(image) in imagesSorted" :key="image.Id"
            @click="selectedId = selectedId == image.Id ? '' : image.Id" :class="{ selected: image.Id == selectedId }">
            <img :src="image.data" />
            <div class="reel-image-toolbox">
                <Button icon="pi pi-trash" label="Bild Löschen" @click="deleteImage(image)" />
            </div>
        </div>
    </div>
</template>


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
