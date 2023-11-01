<script setup lang="ts">

import { ref, computed, PropType } from 'vue'
import Image from '../types/Image';
// import Button from 'primevue/button';
import Listbox from 'primevue/listbox';

const props = defineProps({
    images: {
        type: Array as PropType<Image[]>,
        required: true
    }
})

const $emit = defineEmits(['delete'])

const selectedImage = ref<Image | null>(null)


const imagesSorted = computed(() => {
    return props.images.sort((a, b) => {
        if (a.CreatedAt > b.CreatedAt) {
            return -1
        }
        if (a.CreatedAt < b.CreatedAt) {
            return 1
        }
        return 0
    })
})

const image_count = computed(() => {
    return props.images.length
})

const deleteImage = (image: Image) => {
    console.log("deleteImage", image)
    $emit('delete', image)
}
</script>


<template>
    <div class="reel">

        <div class="reel-header">
            Dein Video besteht aus {{ image_count }} Bildern
        </div>

        <div class="reel-body">
            <Listbox v-model="selectedImage" :options="imagesSorted" optionLabel="id" class=""
                listStyle="width: 100%; height:100%" :virtualScrollerOptions="{ itemSize: 150 }" style="height: 100%;">
                <template #option="slotProps">
                    <div class="flex align-items-center" style="position: relative;">
                        <img :alt="slotProps.option.id" :src="slotProps.option.Url" style="width: 200px; height: 150px;" />
                        <div class="actions" v-if="selectedImage?.Id == slotProps.option.Id">
                            <Button icon="pi pi-trash" @click="deleteImage(slotProps.option)" />
                        </div>
                    </div>
                </template>
            </Listbox>
        </div>
    </div>
</template>


<style scoped>
.reel {
    width: 275px;
    height: 100%;
    overflow: scroll;
    margin-left: 10px;
    margin-right: 10px;

    display: flex;
    flex-direction: column;
}

.reel .reel-header {
    height: 3rem;
    line-height: 3rem;
}

.reel .reel-body {
    width: calc(100% - 20px);
    height: 100%;
    /* display: flex; */
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
    /* border: solid 1px red; */
}

.actions {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-left: 10px;
    margin-right: 10px;

    position: absolute;
    top: 0;
    left: 0;
}
</style>
