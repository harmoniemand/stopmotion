import { defineStore } from "pinia"
import Image from "../types/Image"
import { ref } from "vue"

export const useImageStore = defineStore('images', {
    state: () => ({
        images: ref<Image[]>([]),
    }),
    getters: {
        reversedImages(state) {
            return state.images.slice().reverse()
        },
        lastImage(state) {
            return state.images.slice().reverse()[0]
        }
    },
    actions: {
        addImage(image: Image) {
            this.images.push(image)
        },
        removeImage(image: Image) {
            this.images = this.images.filter(i => i != image)
        },
    },
})
